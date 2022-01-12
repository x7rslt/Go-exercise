package slice_test

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

// 比较两个slice，求相同值,返回newslice
func TestSameItems(t *testing.T){
	var names = []string{"F5", "F7", "C6", "H5", "G5"}
	var board = []string{"B4", "D4", "G5", "F5", "F7", "C6"}

	results := []string{} // slice to store the result

	for i := 0; i < len(names); i++ {
		for k := 0; k < len(board); k++ {
			if names[i] != board[k] {
				continue   //继续循环
			}
			// append a value in result only if
			// it exists both in names and board
			results = append(results, names[i])
		}
	}
	fmt.Printf("%v %d \n", results, len(results))
}

// 比较两个slice，求相同值,返回去除不同的原slice
func TestSameItem(t *testing.T){
	names := []string{"F5", "F7", "C6", "H5", "G5"}
	board := []string{"B4", "D4", "G5", "F5", "F7", "C6"}
	for i:=0;i<len(names);{
		exist := false
		for _,item := range board{
			if names[i] == item{
				exist = true
				break
			}
		}
		if !exist{
			names = append(names[:i],names[i+1:]...)
		}else{
			i++   // 第一次疏忽，忘了递增，陷入死循环，惭愧
		}
	}
	fmt.Println("Slice same item :",names)
}
// 比较两个slice，求不同值,返回去新slice
func TestSliceDifferentItem(t *testing.T){
	names := []string{"F5", "F7", "C6", "H5", "G5"}
	board := []string{"B4", "D4", "G5", "F5", "F7", "C6"}
	newnames := []string{}
	for i:=0;i<len(names);{
		exist := false
		for _,item :=range board{
			if names[i] == item{
				exist = true
				break
			}
		}
		if !exist{
			newnames = append(newnames, names[i])
		}
		i++
	}
	fmt.Println("Slice different item:",newnames)
}

// 比较两个slice，组合成新的slice
func TestSliceUnion(t *testing.T){
	names := []string{"F5", "F7", "C6", "H5", "G5"}
	board := []string{"B4", "D4", "G5", "F5", "F7", "C6"}
	for _, b := range board {
		exist := false
		for _, n := range names {
			if n == b {
				exist = true
				break
			}
		}
		if !exist {
			names = append(names, b)
		}
	}
	fmt.Println(names)
}

func TestSlice(t *testing.T) {
	a := make([]int, 32)
	b := a[1:16]
	a = append(a, 1)
	a[2] = 42
	fmt.Println(a, b)

	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100

	bar := foo[1:4]
	bar[1] = 99
	fmt.Println(foo, bar)
}

func TestSliceCap(t *testing.T) {
	path := []byte("AAAA/BBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex:]

	fmt.Println("dir1:", string(dir1))
	fmt.Println("dir2:", string(dir2))

	dir1 = append(dir1, "suffix"...)
	fmt.Println("dir1:", string(dir1))
	fmt.Println("dir2:", string(dir2))

	//修改内存覆盖,使用Full Slice Expression
	path2 := []byte("AAAA/BBBBBBBB")
	sepIndex2 := bytes.IndexByte(path2, '/')
	Newdir1 := path2[:sepIndex2:sepIndex2]
	Newdir2 := path2[sepIndex2:]

	fmt.Println("Newdir1:", string(Newdir1))
	fmt.Println("Newdir2:", string(Newdir2))

	Newdir1 = append(Newdir1, "suffix"...)

	fmt.Println("Newdir1:", string(Newdir1))
	fmt.Println("Newdir2:", string(Newdir2))
}

func TestReflect(t *testing.T) {
	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2))
	//prints: m1 == m2: true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2))
	//prints: s1 == s2: true
}

type Person struct {
	Name   string
	Sexual string
	Age    int
}

func PrintPerson(p *Person) {
	fmt.Println(p.Name, p.Sexual, p.Age)
}

func (p *Person) Print() {
	fmt.Println(p.Name, p.Age, p.Sexual)
}
//在 Go 语言中，使用“成员函数”的方式叫“Receiver”，这种方式是一种封装，因为 PrintPerson()本来就是和 Person强耦合的，所以，理应放在一起。
func TestStruct(t *testing.T) {
	var p = Person{
		Name:   "Xiao",
		Sexual: "Male",
		Age:    12,
	}

	PrintPerson(&p)
	p.Print()
}
