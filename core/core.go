package core

import (
"errors"
"os"
"os/signal"
"time"
)

var ErrorTimeOut = errors.New("Timeout")
var ErrorInterrupt = errors.New("Interrupt")

type Runner struct {
	interrupt chan os.Signal    //发送的信号，用来终止程序
	complete chan error         //用来通知任务全部完成
	timeout <- chan time.Time   //程序的超时时间
	tasks []func(string)   //要执行的任务
}

func New(tm time.Duration) *Runner {
	return &Runner{
		complete: make(chan error),
		timeout: time.After(tm),
		interrupt: make(chan os.Signal,1),
	}
}

func (r *Runner) Add(task ...func(string)){
	r.tasks = append(r.tasks,task...)
}

func (r *Runner) run(url string) error{
	for _,task := range r.tasks{
		if r.isInterrupt() {
			return ErrorInterrupt
		}
		task(url)
	}
	return nil
}

func (r *Runner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

func (r *Runner) Start(urls []string) error{
	signal.Notify(r.interrupt,os.Interrupt)
	go func() {
		r.complete <- r.run(urls[0])
	}()

	select {
	case err := <- r.complete :
		return err
	case <-r.timeout:
		return ErrorTimeOut
	}
}
