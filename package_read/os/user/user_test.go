package user_test

import (
	"fmt"
	"os/user"
	"testing"
)

func TestUser(t *testing.T) {
	u, _ := user.Current()
	gs, _ := u.GroupIds()
	fmt.Println(u.Uid, u.Gid, u.HomeDir, u.Username)
	fmt.Println(gs)
	g, _ := user.LookupGroup("admin")
	gp := g
	fmt.Println(gp.Gid, gp.Name)

	xiao := []string{
		"lehehe", "hello",
	}
	fmt.Println(xiao)
}
