package main

import (
	"context"
	"encoding/base64"
	"testing"
	"time"

	android10 "github.com/wnxd/microdbg-android/10"
	"github.com/wnxd/microdbg-android/extend"
	gava "github.com/wnxd/microdbg-android/java"
	"github.com/wnxd/microdbg-android/wrapper"
	"github.com/wnxd/microdbg-example/tiger"
	java "github.com/wnxd/microdbg-java"
	unicorn "github.com/wnxd/microdbg-unicorn"
	"github.com/wnxd/microdbg/emulator"
	"github.com/wnxd/microdbg/filesystem"
)

func TestTiger(t *testing.T) {
	emu, err := unicorn.New(emulator.ARCH_ARM64)
	if err != nil {
		t.Fatal(err)
	}
	defer emu.Close()
	env := wrapper.NewFake(tiger.Handler{})
	art, err := android10.NewRuntime(emu, android10.WithApkPath("app/1.apk"), android10.WithRuntimeDir("runtime/10.zip"), android10.WithJNIEnv(env))
	if err != nil {
		t.Fatal(err)
	}
	defer art.Close()

	pkg := art.Package()
	art.LinkFS(pkg.CodePath(), filesystem.SysFileFS("app/1.apk"))

	cf := env.ClassFactory()
	ex := extend.Define(art, cf)

	altt := tiger.SharedPreference{
		"TT_COOKIEID":     gava.FakeString(base64.StdEncoding.EncodeToString(RandomBytes(64))),
		"TT_COOKIEID_NEW": gava.FakeString(base64.StdEncoding.EncodeToString(RandomBytes(64))),
	}
	ex.SharedPreference("altt", altt)

	module, err := pkg.LoadModule(context.TODO(), art, "tiger_tally")
	if err != nil {
		t.Fatal(err)
	}
	defer module.Close()
	t.Log(module.Name())
	t.Log(module.CallOnLoad(context.TODO()))

	cls := cf.GetClass("com.aliyun.TigerTally.t.B")

	method := cls.GetMethod("genericNt1", "(Ljava/util/Map;)I")
	r := method.CallPrimitive(nil, gava.MapOf(gava.Map{
		gava.FakeString("AppKey"):      gava.FakeString("ypWt5wEEQwOEgLM4e12Gl26wHlW6Qj_XOG0-l7p3ju05wOt2jZ0tNkr5he6ei73A2AQQUH2QbJfvfJoKU_rKkdwvHEn75U6xYNgpgVYUVSjxZt1Ks5MdUQoZY_SK-ETAArxOUW1Mhf8uTnvvLUOB9tQMlSNcntBETjvhg8xB2CA="),
		gava.FakeString("CollectType"): gava.FakeString("1"),
	}))
	t.Log(r)

	method = cls.GetMethod("genericNt3", "(I[B)Ljava/lang/String;")
	var data [16]byte

	start := time.Now()
	r = method.CallPrimitive(nil, java.JInt(1), gava.BytesOf(data[:]))
	t.Logf("Result: %s, off=%v", r, time.Since(start))
}

func TestTiger2(t *testing.T) {
	emu, err := unicorn.New(emulator.ARCH_ARM64)
	if err != nil {
		t.Fatal(err)
	}
	defer emu.Close()
	env := wrapper.NewFake(tiger.Handler{})
	art, err := android10.NewRuntime(emu, android10.WithApkPath("app/2.apk"), android10.WithRuntimeDir("runtime/10.zip"), android10.WithJNIEnv(env))
	if err != nil {
		t.Fatal(err)
	}
	defer art.Close()

	pkg := art.Package()
	art.LinkFS(pkg.CodePath(), filesystem.SysFileFS("app/2.apk"))

	cf := env.ClassFactory()
	ex := extend.Define(art, cf)

	altt := tiger.SharedPreference{
		"TT_COOKIEID":     gava.FakeString(base64.StdEncoding.EncodeToString(RandomBytes(64))),
		"TT_COOKIEID_NEW": gava.FakeString(base64.StdEncoding.EncodeToString(RandomBytes(64))),
	}
	ex.SharedPreference("altt", altt)

	module, err := pkg.LoadModule(context.TODO(), art, "tiger_tally")
	if err != nil {
		t.Fatal(err)
	}
	defer module.Close()
	t.Log(module.Name())
	t.Log(module.CallOnLoad(context.TODO()))

	cls := cf.GetClass("com.aliyun.TigerTally.t.B")

	cls.GetMethod("genericNt2", "(ILjava/lang/String;)I").CallPrimitive(nil, java.JInt(1), gava.FakeString("1450767262844751"))
	cls.GetMethod("genericNt1", "(ILjava/lang/String;)I").CallPrimitive(nil, java.JInt(1), gava.FakeString("xPEj7uv0KuziQnXUyPIBNUjnDvvHuW09VOYFuLYBcY-jV6fgqmfy5B1y75_iSuRM5U2zNq7MRoR9N1F-UthTEgv-QBWk68gr95BrAySzWuDzt08FrkeBZWQCGyZ0iAybalYLOJEF7nkKBtmDGLewcw=="))

	method := cls.GetMethod("genericNt3", "(I[B)Ljava/lang/String;")
	var data [16]byte

	start := time.Now()
	r := method.CallPrimitive(nil, java.JInt(1), gava.BytesOf(data[:]))
	t.Logf("Result: %s, off=%v", r, time.Since(start))
}
