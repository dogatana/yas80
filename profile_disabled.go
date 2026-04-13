//go:build !profile

package main

func startCPUProfile(_ string) func() {
	return func() {}
}

func memProfile(_ string) {
	return
}
