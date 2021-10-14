package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main(){
	fileObj,err := os.Open("./sample.out")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n",err)
		return
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for n:=1;;n++{
	for i:=1;i<=1000;i++{
	line,err := reader.ReadString('\n')
	if err!=nil{
		fmt.Printf("read file by bufio failed,err:%v\n",err)
		return 
	}
	if err == io.EOF{
		return
	}
	string_file := "./sample" + strconv.Itoa(n) + ".out"
	filewrite,err := os.OpenFile(string_file,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n",err)
		return
	}
	defer filewrite.Close()
	filewrite.Write([]byte(line))
	}
	}	
}