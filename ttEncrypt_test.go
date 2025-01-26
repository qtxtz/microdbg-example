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

func TestTTEncrypt(t *testing.T) {
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

	file, _ := os.Open("ttEncrypt/libttEncrypt.so")
	defer file.Close()
	module, err := art.LoadModule(context.TODO(), file)
	if err != nil {
		t.Fatal(err)
	}
	defer module.Close()
	t.Log(module.Name())

	cf := env.ClassFactory()

	TTEncryptUtils := cf.GetClass("com/bytedance/frameworks/core/encrypt/TTEncryptUtils")

	method, err := module.FindNativeMethod(art.JavaVM(), TTEncryptUtils, "ttEncrypt", "([BI)[B")
	if err != nil {
		t.Fatal(err)
	}

	var data [16]byte
	r := method(nil, gava.BytesOf(data[:]), java.JInt(len(data)))
	t.Log(r)
}
