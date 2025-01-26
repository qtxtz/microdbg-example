package main

import (
	crand "crypto/rand"
	"fmt"
	"io"
	"math/rand/v2"

	"github.com/wnxd/microdbg/debugger"
	"github.com/wnxd/microdbg/emulator"
)

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	data := make([]byte, length)
	for i := range data {
		data[i] = charset[rand.IntN(len(charset))]
	}
	return string(data)
}

func RandomBytes(length int) []byte {
	data := make([]byte, length)
	crand.Read(data)
	return data
}

func SyncPC(dbg debugger.Debugger) io.Closer {
	hook, _ := dbg.AddHook(emulator.HOOK_TYPE_CODE, func(ctx debugger.Context, addr, size uint64, data any) {
	}, nil, 1, 0)
	return hook
}

func PrintCode(dbg debugger.Debugger) io.Closer {
	hook, _ := dbg.AddHook(emulator.HOOK_TYPE_CODE, func(ctx debugger.Context, addr, size uint64, data any) {
		if mod, err := dbg.FindModuleByAddr(addr); err == nil {
			fmt.Printf("module: %s, offset: %08X, size: %d\n", mod.Name(), addr-mod.BaseAddr(), size)
		} else {
			fmt.Printf("address: %016X, size: %d\n", addr, size)
		}
	}, nil, 1, 0)
	return hook
}

func PrintBlock(dbg debugger.Debugger) io.Closer {
	hook, _ := dbg.AddHook(emulator.HOOK_TYPE_BLOCK, func(ctx debugger.Context, addr, size uint64, data any) {
		if mod, err := dbg.FindModuleByAddr(addr); err == nil {
			fmt.Printf("module: %s, offset: %08X, size: %d\n", mod.Name(), addr-mod.BaseAddr(), size)
		} else {
			fmt.Printf("address: %016X, size: %d\n", addr, size)
		}
	}, nil, 1, 0)
	return hook
}
