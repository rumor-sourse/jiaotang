package main

import (
	"fmt"
	"time"
	"sync"
)
var wg sync.WaitGroup
//定义任务类型Task
type Task struct{
	f func() error
}
//创建Task任务
func NewTask(arg_f func() error) *Task{
	t := Task{
		f: arg_f,
	}

	return &t
}

//执行Task的方法
func(t *Task) Execute(){
	t.f()
}

//定义一个Pool协程池的类型
type Pool struct{
	//对外的Task 入口 EntryChannel
	EntryChannel chan *Task
	//内部的Task 队列 JobsChannel
	JobsChannel chan *Task
	//协程池中最大的worker的数量
	worker_num int
}

//创建Pool函数
func NewPool(cap int) *Pool{
	//创建一个Pool
	p:= Pool{
		EntryChannel: make(chan *Task),
		JobsChannel: make(chan *Task),
		worker_num: cap,
	}
	return &p
}

//协程池创建一个Worker，并且让这个Worker去工作
func (p *Pool) worker(worker_ID int){
	//一个worker具体的工作
	//1、永久的从JobsChannel去取任务
	for task := range p.JobsChannel{
		//task就是当前worker从JobsChannel中拿到的任务
		//2、一旦取到这个任务，执行这个任务
		task.Execute()
		fmt.Println("worker ID",worker_ID,"执行完了一个任务")
	}
}

//让协程池，开始真正的工作，协程池一个启动方法
func (p *Pool)run(){
	//1、根据worker_num来创建worker去工作
	wg.Add(5)
	for i:=1;i<=p.worker_num;i++{
		go p.worker(i)
		wg.Done()
	}
	//2、从EntryChannel中去取任务，将取到的任务，发送给JobsChannel
	for task:= range p.EntryChannel{
		p.JobsChannel <- task
	}	
}
//测试协程池的工作
func main(){
	t := NewTask(func() error{
		fmt.Println(time.Now())
		return nil
	})
	p:= NewPool(5)
	task_num := 0
	go func(){
		for n:=0;n<10;n++{
		p.EntryChannel <-t
		task_num +=1
		fmt.Println("当前一共执行了",task_num,"个任务")
	}
	}()
	p.run()
	wg.Wait()
}