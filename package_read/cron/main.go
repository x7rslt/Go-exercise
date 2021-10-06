package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

type TestTask struct {
	Name string
}

func (t *TestTask) Run() {
	fmt.Println("TestTask", t.Name)
}

type Test2Task struct {
	Name string
}

func (t *Test2Task) Run() {
	fmt.Println("Test2Task", t.Name)
}

func main() {
	// 新建一个定时任务对象
	// 根据cron表达式进行时间调度，cron可以精确到秒，大部分表达式格式也是从秒开始。
	//crontab := cron.New()  默认从分开始进行时间调度
	crontab := cron.New() //精确到秒
	//定义定时器调用的任务函数
	//定时任务
	spec := "*/5 * * * * ?" //cron表达式，每五秒一次

	//定义定时器调用的任务函数
	task := func() {
		fmt.Println("hello world", time.Now())
	}
	// 添加定时任务
	crontab.AddFunc(spec, task)

	// 添加多个定时器
	crontab.AddJob(spec, &TestTask{Name: "First Job"})
	crontab.AddJob(spec, &Test2Task{Name: "Second Job"})
	// 启动定时器
	crontab.Start()
	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer crontab.Stop()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
	select {
		case <-time.After(time.Second *10):
			return
	} //阻塞主线程停止
}