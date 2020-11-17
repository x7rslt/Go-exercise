package binary_test

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"testing"
)

type Gopher struct {
	Name string
	Age  string
}

func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
	err = binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
	if err != nil {
		return
	}
	size += 4
	n, err := w.Write([]byte(g.Name))
	size += int64(n)
	if err != nil {
		return
	}
	err = binary.Write(w, binary.BigEndian, int64(len(g.Age)))
	if err == nil {
		size += 4
	}
	return

}

func TestBinary(t *testing.T) {

	xiao := Gopher{Name: "xiaoshuao", Age: "12"}
	s, err := xiao.WriteTo(os.Stdout)
	fmt.Println(s, err)

}
