package main

import "time"
import "github.com/cheggaaa/pb/v3"

func main(){
	count := 100000
	bar := pb.StartNew(count)

	for i:= 0;i<count;i++{
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.Finish()
}
