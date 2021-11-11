package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

//当前目录下，go install 可生成greet2的执行文件
func main() {
	app := &cli.App{
		Name: "greet2",
		Usage: "fight the loneliness!",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
