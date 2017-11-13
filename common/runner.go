package common

import (
	"time"
	"os"
	"os/signal"
	"errors"
)

type Runner struct {
	tasks []func(int) //要执行的任务
	complete chan error //用于通知任务全部完成
	timeout <-chan time.Time //这些任务在多久内完成
	interrupt chan os.Signal //可以控制强制终止的信号
}

func New(tm time.Duration) *Runner  {
	return &Runner{
		complete:make(chan error),
		timeout:time.After(tm),
		interrupt:make(chan os.Signal,1),
	}
}

//新增一个任务
func (r *Runner) Add(fn... func(int)){
	r.tasks = append(r.tasks, fn...)
}

//执行超时的错误
var ErrorTimeOut = errors.New("执行超时啦")
//终端执行的错误
var ErrorInterrupt = errors.New("中断执行啦")

//判断是否终止
func (r *Runner) isInterrupt() bool  {
	select {
		case <-r.interrupt:
			signal.Stop(r.interrupt)
			return true
	default:
		return false;
	}
}

//执行任务
func (r *Runner) run() error  {
	for id, fn := range r.tasks {
		if r.isInterrupt() {
			return ErrorInterrupt
		}
		fn(id)
	}
	return nil
}

//开始执行
func (r *Runner) Start() error  {
	signal.Notify(r.interrupt, os.Interrupt)

	go func(){
		r.complete <- r.run()
	}()

	select {
		case err:=<-r.complete:
			return err
			case <- r.timeout:
			return ErrorTimeOut
	}

	return nil
}
