package main

import "fmt"

type Log struct {
	Id int `gorm:"primary_key" json:"id"`
	TaskId int `json:"task_id"`
	TaskName string `json:"task_name"`
	TaskType string `json:"task_type"`
	AllNum int `json:"all_num"`
	SuccessNum int `json:"success_num"`
	RunTime string `json:"run_time"`
	Status int `json:"status"`
	Error string `json:"error"`
	CreatedTime string `json:"created_time"`
}

func GetLog(pageNum int,pageSize int,maps interface{}) (log []Log){
	log = []Log{
		{1,2,"hello","world",100,1,"Now",200,"No problem","Now"},
	}
	fmt.Println(log)
	return log
}

func main(){
	pageNum := 1
	pageSize := 2
	maps := "hello"
	GetLog(pageNum,pageSize,maps)
}