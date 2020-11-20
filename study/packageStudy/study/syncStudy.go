package study

import (
	"fmt"
	"sync"
	"time"
)

func way1(){
	for i := 0; i < 100 ; i++{
		go fmt.Println(i)
	}
	time.Sleep(time.Second)
}

func way2(){
	c := make(chan bool, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-c
	}
}

func way3(){
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func SyncMain() {
	fmt.Println("---------way1---------")
	way1()
	fmt.Println("---------way2---------")
	way2()
	fmt.Println("---------way3---------")
	way3()
}
