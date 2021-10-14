package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"github.com/panjf2000/ants"
)
var t chan int 
var wg sync.WaitGroup
var sample string = "./sample.out"
func readBlock(filePath string,n int){
	fileObj,err := os.Open(filePath)
	if err!= nil{
		fmt.Printf("open file failed,err:%v\n",err)
		return
	}
	defer fileObj.Close()
	//动态设置每次读取字节数
	fileInfo,_:=fileObj.Stat()
	m := fileInfo.Size()
	k := (int)(m/1000)
	t <-k
	buffer := make([]byte,1024*k)
	bytes,err := fileObj.ReadAt(buffer,(int64)(1024*k*(n-1)))
	if err==io.EOF{
		string_file := "./sample" + strconv.Itoa(n) + ".out"
		filewrite,err := os.OpenFile(string_file,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
		if err != nil {
			fmt.Printf("open file failed,err:%v\n",err)
			return
	   }
	   defer filewrite.Close()
	   filewrite.Write(buffer[:bytes])
	   os.Exit(1)
	}
	string_file := "./sample" + strconv.Itoa(n) + ".out"
	filewrite,err := os.OpenFile(string_file,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err != nil {
		 	fmt.Printf("open file failed,err:%v\n",err)
		 	return
		}
	defer filewrite.Close()
	filewrite.Write(buffer[:bytes])
	// wg.Done()
}
func main(){
	defer ants.Release()
	pool,_ :=ants.NewPool(1000)
	//依据内存大小调节并发数量
	c := <- t
	if c!=0{
		pool.Tune(100*c)
	}
	task := func(filePath string,n int){
		readBlock(filepath, n)
		wg.Done()
	}
	for n:=1;n<1000;n++{
		wg.Add(1)
		_ = pool.Submit(task)
	}
	wg.Wait()	
}