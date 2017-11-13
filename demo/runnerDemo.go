package main

import (
	"github.com/woshicixide/golangstudy/common"
	"log"
	"time"
	"os"
)

func main() {
	log.Println("现在开始执行任务")
	tm := 3*time.Second
	r := common.New(tm)

	r.Add(createTask(), createTask(), createTask())

	if err:=r.Start();nil != err{
		switch err {
		case common.ErrorTimeOut:
			log.Println(err)
			os.Exit(1)
		case common.ErrorInterrupt:
			log.Println(err)
			os.Exit(2)
		}
	}

	log.Println("任务结束")
}

func createTask() func(int){
	return func (id int){
		log.Printf("我正在执行任务%d",id)
		time.Sleep(time.Duration(id)*time.Second)
	}
}
