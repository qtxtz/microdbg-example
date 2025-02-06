package main

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"testing"
	"time"

	android10 "github.com/wnxd/microdbg-android/10"
	"github.com/wnxd/microdbg-android/extend"
	gava "github.com/wnxd/microdbg-android/java"
	"github.com/wnxd/microdbg-android/wrapper"
	java "github.com/wnxd/microdbg-java"
	unicorn "github.com/wnxd/microdbg-unicorn"
	"github.com/wnxd/microdbg/emulator"
	"github.com/wnxd/microdbg/filesystem"
)

func TestJMEncryptBox(t *testing.T) {
	emu, err := unicorn.New(emulator.ARCH_ARM64)
	if err != nil {
		t.Fatal(err)
	}
	defer emu.Close()
	env := wrapper.NewFake(nil)
	art, err := android10.NewRuntime(emu, android10.WithApkPath("app/2.apk"), android10.WithRuntimeDir("runtime/10.zip"), android10.WithJNIEnv(env))
	if err != nil {
		t.Fatal(err)
	}
	defer art.Close()

	pkg := art.Package()
	art.LinkFS(pkg.CodePath(), filesystem.SysFileFS("app/2.apk"))

	cf := env.ClassFactory()
	extend.Define(art, cf)

	JMEncryptBox := cf.GetClass("com.ijiami.JMEncryptBox")
	JMEncryptBoxByRandom := cf.GetClass("com.ijiami.JMEncryptBoxByRandom")

	JMEncryptBox.DefineMethod("getFinger", "(Ljava/lang/String;[B)Ljava/lang/String;", gava.Modifier_PUBLIC|gava.Modifier_STATIC).BindCall(func(obj java.IObject, args ...any) any {
		algorithm := args[0].(java.IString).String()
		bytes := gava.GetBytes(args[1].(java.IByteArray))
		switch algorithm {
		case "MD5":
			h := md5.Sum(bytes)
			return gava.BytesOf(h[:])
		case "SHA1":
			h := sha1.Sum(bytes)
			return gava.BytesOf(h[:])
		}
		return nil
	})

	module, err := pkg.LoadModule(context.TODO(), art, "JMEncryptBox")
	if err != nil {
		t.Fatal(err)
	}
	defer module.Close()
	t.Log(module.Name())
	t.Log(module.CallOnLoad(context.TODO()))

	setKey, err := module.FindNativeMethod(art.JavaVM(), JMEncryptBox, "setKey", "(Ljava/lang/String;)Ljava/lang/String;")
	if err != nil {
		t.Fatal(err)
	}

	encryptByRandomType1, err := module.FindNativeMethod(art.JavaVM(), JMEncryptBoxByRandom, "encryptByRandomType1", "([B)[B")
	if err != nil {
		t.Fatal(err)
	}

	r := setKey(nil, gava.FakeString("C0612E92E4F99FA8A1C73EB94D3969F2D094476F4DC16620AF850ACCDB9DA2D3024796E1A65AEBE504235D04E85520391E7B694D83C7F58C2C70C3E90E81CB70A97F6855F40243F5852E04D013DBC263984BBF58F8F9EFBBA59C9E51E50AF320E6BD"))
	t.Log(r)

	var data [16]byte
	start := time.Now()
	r = encryptByRandomType1(nil, gava.BytesOf(data[:]))
	t.Logf("Result: %v, off=%v", r, time.Since(start))
}

func TestSm4EncryptBox(t *testing.T) {
	emu, err := unicorn.New(emulator.ARCH_ARM64)
	if err != nil {
		t.Fatal(err)
	}
	defer emu.Close()
	env := wrapper.NewFake(nil)
	art, err := android10.NewRuntime(emu, android10.WithApkPath("app/2.apk"), android10.WithRuntimeDir("runtime/10.zip"), android10.WithJNIEnv(env))
	if err != nil {
		t.Fatal(err)
	}
	defer art.Close()

	pkg := art.Package()
	art.LinkFS(pkg.CodePath(), filesystem.SysFileFS("app/2.apk"))

	cf := env.ClassFactory()
	extend.Define(art, cf)

	module, err := pkg.LoadModule(context.TODO(), art, "ijmWhitebox")
	if err != nil {
		t.Fatal(err)
	}
	defer module.Close()
	t.Log(module.Name())
	t.Log(module.CallOnLoad(context.TODO()))

	Sm4EncryptBox := cf.GetClass("com.ijiami.whitebox.Sm4EncryptBox")

	sm4AlgorithmOperation, err := module.FindNativeMethod(art.JavaVM(), Sm4EncryptBox, "sm4AlgorithmOperation", "([B[B[B[BII)[B")
	if err != nil {
		t.Fatal(err)
	}

	const (
		key = "e58f12c530982087a5791af451815416996b5b628cec2def690f8c30336b07f617554f11d4cae0d9e4942121c2889f0087485ffd758ebf072dec3c98d4ce66ef146c7173571cf4a27c230d84e156b841ffd895f1482091c4e5c3c65793d21a4035590e6ce5e623af1970d3858df05c1835c3e72991780c3053a2613f2d68d786328373a2"
		iv  = "00000000000000000000000000000000"
	)

	var data [16]byte
	start := time.Now()
	r := sm4AlgorithmOperation(nil, gava.BytesOf([]byte(key)), gava.BytesOf([]byte(iv)), gava.BytesOf(data[:]), gava.BytesOf([]byte("cbc")), java.JInt(0), java.JInt(0))
	t.Logf("Result: %v, off=%v", r, time.Since(start))
}
