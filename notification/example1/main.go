package main

import "github.com/gen2brain/beeep"


//响一声beep
func main(){
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}
}
