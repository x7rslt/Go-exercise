package main

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

func main() {

	//Read a TOML document
	config, _ := toml.Load(`
	[postgress]
	user = "sqlname"
	password = "password"
	`)
	user := config.Get("postgress.user").(string)
	postgressConfig := config.Get("postgress").(*toml.Tree)
	password := postgressConfig.Get("password").(string)
	fmt.Println(user, password)

	//Use Unmarshal
	type Postgres struct {
		User     string
		Password string
	}

	type Config struct {
		Postgres Postgres
	}
	doc := []byte(`
	[Postgres]
	User = "Sqlname2"
	Password = "Password2"
	`)
	config2 := Config{}
	toml.Unmarshal(doc, &config2)
	fmt.Println(config2, config2.Postgres.User)

}
