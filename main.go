package main

import (
	"fmt"
	"runtime"
	"sync"
	"net/http"
	"os"
	"github.com/Supparerk23/go-introduce/repositories"
	"github.com/labstack/echo/v4"
	"github.com/subosito/gotenv"
	)

var wg sync.WaitGroup

func main() {

	gotenv.Load()

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

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

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK,endMessage)
	})

	// Initial API
	api := e.Group("/api")

	propertyAPI := api.Group("/property")

	e.Logger.Fatal(e.Start(addr))
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