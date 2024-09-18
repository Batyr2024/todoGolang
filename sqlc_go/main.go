package main

import (
	"fmt"
	"sync"
)

var b func() func() (a bool, b bool, c interface{})

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	defer func() {
		if ok := recover(); ok != nil {
			fmt.Println(ok)
			main()
		}
	}()
	go in(&wg)
	wg.Wait()
	var a func() (a bool, b bool, c interface{})
	a = b()
	fmt.Println(a())
	fmt.Println("done")
}
func in(wg *sync.WaitGroup) {
	b = func() func() (a bool, b bool, c interface{}) {
		defer func(j int) { fmt.Println(j) }(3)
		return func() (a bool, b bool, c interface{}) {
			fmt.Println("pp")
			defer func(j int) { fmt.Println(j) }(0)
			a = true
			b = false
			c = fmt.Sprint("asd", "234", "00000")
			defer func(j int) { fmt.Println(j) }(1)
			defer func(j int) { fmt.Println(j) }(2)
			return
		}
	}
	fmt.Println("close in func")
	wg.Done()
}
