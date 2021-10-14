package main

import (
	"fmt"
	"sync"
	"time"
	"github.com/panjf2000/ants"
)

func Task() {
	fmt.Println("Helloworld")
	time.Sleep(2*time.Second)

}
var wg sync.WaitGroup
func main(){
	defer ants.Release()
	pool,_ :=ants.NewPool(200)
	task := func(){
		Task()
		wg.Done()
	}
	for i:=0;i<10000;i++{
		wg.Add(1)
		_=pool.Submit(task)
	}
	wg.Wait()
}
