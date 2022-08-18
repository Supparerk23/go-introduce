package main

import (
	"fmt"
	"runtime"
	"sync"
	"go-introduce/go-module-test"
	)

var wg sync.WaitGroup

func main() {

	fmt.Println("OS\t\t",runtime.GOOS)
	fmt.Println("ARCH\t\t",runtime.GOARCH)
	fmt.Println("CPUs\t\t",runtime.NumCPU())
	fmt.Println("GOroutines\t\t",runtime.NumGoroutine())

	var allCPUs int = runtime.NumCPU() - 1

	wg.Add(allCPUs)

	for i := 0; i < allCPUs; i++ {
		go foo(i)
	}

	bar()

	fmt.Println("CPUs\t\t",runtime.NumCPU())
	fmt.Println("GOroutines\t\t",runtime.NumGoroutine())

	wg.Wait()

	var endMessage string = hello.Hello()

	fmt.Println("end : ",endMessage)
}

func foo(prefix int) {

	for i := 0; i < prefix; i++ {
		fmt.Println(prefix,"foo:",i)
	}

	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}