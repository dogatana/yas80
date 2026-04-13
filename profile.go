//go:build profile

package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func startCPUProfile(name string) func() {
	if name == "" {
		return func() {}
	}
	f, err := os.Create(name)
	if err != nil {
		fmt.Printf("cannot create %s\n", name)
		return func() {}
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("cannot start CpuProfile", err)
		return func() {}
	}

	return func() {
		pprof.StopCPUProfile()
		f.Close()
	}
}

func memProfile(name string) {
	if name == "" {
		return
	}
	f, err := os.Create(name)
	if err != nil {
		fmt.Printf("cannot create %s\n", name)
		return
	}
	defer f.Close()
	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Println("cannot write memory profile", err)
	}
}
