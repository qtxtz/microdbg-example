package main

import (
	"context"
	"os"
	"testing"
	"time"

	android10 "github.com/wnxd/microdbg-android/10"
	gava "github.com/wnxd/microdbg-android/java"
	"github.com/wnxd/microdbg-android/wrapper"
	java "github.com/wnxd/microdbg-java"
	unicorn "github.com/wnxd/microdbg-unicorn"
	"github.com/wnxd/microdbg/emulator"
	"golang.org/x/sync/semaphore"
)

func TestKanxue(t *testing.T) {
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

	file, _ := os.Open("kanxue/libnative-lib.so")
	defer file.Close()
	module, err := art.LoadModule(context.TODO(), file)
	if err != nil {
		t.Fatal(err)
	}
	defer module.Close()
	t.Log(module.Name())

	cf := env.ClassFactory()

	cls := cf.GetClass("com.kanxue.test2.MainActivity")

	method, err := module.FindNativeMethod(art.JavaVM(), cls, "jnitest", "(Ljava/lang/String;)Z")
	if err != nil {
		t.Fatal(err)
	}

	obj := cls.NewInstance()

	LETTERS := []byte{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}

	const CO = 10

	sem := semaphore.NewWeighted(CO)
	var skip bool

	start := time.Now()
	for _, a := range LETTERS {
		if skip {
			break
		}
		for _, b := range LETTERS {
			if skip {
				break
			}
			for _, c := range LETTERS {
				if skip {
					break
				}
				sem.Acquire(context.TODO(), 1)
				str := gava.FakeString(string(a) + string(b) + string(c))
				go func() {
					r := method(obj, str)
					if r.(java.JBoolean) {
						skip = true
						t.Logf("Found: %s, off=%v", str, time.Since(start))
					}
					sem.Release(1)
				}()
			}
		}
	}
	sem.Acquire(context.TODO(), CO)
}
