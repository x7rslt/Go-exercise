package flag_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"github.com/pelletier/go-toml"

func Test(t *testing.T) {
	userhome, _ := os.UserHomeDir()
	fmt.Println(userhome)
	defaultconf := filepath.Join(userhome, ".ffufrc")
	fmt.Println(defaultconf)
	configData, err := ioutil.ReadFile(defaultconf)
	fmt.Println(configData,err)
}
