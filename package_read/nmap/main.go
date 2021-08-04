package main

import (
	"context"
	"fmt"
	"github.com/Ullaakut/nmap"
	"log"
	"time"
)

func main(){
	ctx,cancel:= context.WithTimeout(context.Background(),5*time.Minute)
	defer  cancel()

	scanner,err := nmap.NewScanner(
		nmap.WithTargets("127.0.0.1"),
		//nmap.WithPorts("90"),
		nmap.WithContext(ctx),
		)
	if err!= nil{
		log.Fatalf("unable to create nmap scanner:%v",err)
	}
	result,warnings,err := scanner.Run()

	if warnings != nil{
		log.Printf("warnings:\n %v",warnings)
	}
	for _,host := range result.Hosts{
		if len(host.Ports)==0||len(host.Addresses)==0{
			continue
		}
		fmt.Printf("Host %q:\n",host.Addresses[0])
		for _,port := range host.Ports{
			fmt.Printf("\tPort %d/%s %s %s\n",port.ID,port.Protocol,port.State,port.Service.Name)

		}
	}
	fmt.Printf("Nmap done:%d hosts up scanned in %3f seconds\n",len(result.Hosts),result.Stats.Finished.Elapsed)
}
