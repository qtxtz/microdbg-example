package main

import (
	"context"
	"os"
	"testing"

	android10 "github.com/wnxd/microdbg-android/10"
	gava "github.com/wnxd/microdbg-android/java"
	"github.com/wnxd/microdbg-android/wrapper"
	java "github.com/wnxd/microdbg-java"
	unicorn "github.com/wnxd/microdbg-unicorn"
	"github.com/wnxd/microdbg/emulator"
)

func TestSignUtil(t *testing.T) {
	emu, err := unicorn.New(emulator.ARCH_ARM)
	if err != nil {
		t.Fatal(err)
	}
	defer emu.Close()
	env := wrapper.NewFake(nil)
	art, err := android10.NewRuntime(emu, android10.WithRuntimeDir("runtime/10.zip"), android10.WithJNIEnv(env))
	if err != nil {
		t.Fatal(err)
	}
	defer art.Close()

	file, _ := os.Open("signutil/libsignutil.so")
	defer file.Close()
	module, err := art.LoadModule(context.TODO(), file)
	if err != nil {
		t.Fatal(err)
	}
	defer module.Close()
	t.Log(module.Name())

	cf := env.ClassFactory()

	SignUtil := cf.GetClass("com/anjuke/mobile/sign/SignUtil")

	method, err := module.FindNativeMethod(art.JavaVM(), SignUtil, "getSign0", "(Ljava/lang/String;Ljava/lang/String;Ljava/util/Map;Ljava/lang/String;I)Ljava/lang/String;")
	if err != nil {
		t.Fatal(err)
	}

	data := gava.MapOf(map[java.IObject]java.IObject{
		gava.FakeString("a"): gava.BytesOf([]byte("b")),
		gava.FakeString("b"): gava.BytesOf([]byte("b")),
	})
	r := method(nil, gava.FakeString("aa"), gava.FakeString("bb"), data, gava.FakeString("cc"), java.JInt(10))
	t.Log(r)
}
