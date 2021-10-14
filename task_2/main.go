package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// var wg sync.WaitGroup
var sample string = "./sample.out"
func readBlock(filePath string,n int){
	fileObj,err := os.Open(filePath)
	if err!= nil{
		fmt.Printf("open file failed,err:%v\n",err)
		return
	}
	defer fileObj.Close()
	//设置每次读取字节数
	buffer := make([]byte,1024*64)
	bytes,err := fileObj.ReadAt(buffer,(int64)(1024*64*(n-1)))
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
	for n:=1;n<1000;n++{
		// wg.Add(1)
		go readBlock(sample,n)
		time.Sleep(time.Millisecond)
	}
	// wg.Wait()	
}