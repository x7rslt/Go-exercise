package read_test

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

//Read File Line By Line method 1
func TestReadLine1(t *testing.T){
	file,err := os.Open("dict.txt")
	if err != nil{
		fmt.Println("Open file Err:",err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err();err !=nil{
		fmt.Println("Line read err:",err)
	}
}


//Read File Line By Line
func TestReadLine(t *testing.T){
	file,err := os.Open("dict.txt")
	if err!= nil{
		fmt.Println("Fail to open file:",err)

	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string
	for fileScanner.Scan(){
		fileTextLines = append(fileTextLines,fileScanner.Text())

	}
	file.Close()
	for _,line := range fileTextLines{
		fmt.Println(line)
	}
}

//Read a file word by word
func TestReadWord(t *testing.T){
	// open file
	f, err := os.Open("dict.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file word by word using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		// do something with a word
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//Read a file in chunks
func TestReadLargeFile(t *testing.T){
	chunkSize := 10
	// open file
	f, err := os.Open("dict.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	buf := make([]byte, chunkSize)

	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if err == io.EOF {
			break
		}

		fmt.Println(string(buf[:n]))
	}
}

//Read an entire file
func TestReadEntire(t *testing.T){
	file,err := os.ReadFile("dict.txt")
	if err != nil{
		fmt.Println("Read Err:",err)
	}
	fmt.Println(string(file))
	fmt.Printf("%T",file)
}

//Read an entire file,输出不同格式
func TestReadAll(t *testing.T) {
	file, err := os.Open("dict.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data as hex:%x\n", data)
	fmt.Printf("Data as string:%s\n", data)
	fmt.Printf("Number of bytes read:%d", len(data))

}

//逐行读取.csv文件
func TestBufioread(t *testing.T) {
	reader := bufio.NewReader(os.Stdin) //不适合test，放在main函数；go run后，输入任意一行字符即返回打印。
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
}

func TestCsv(t *testing.T) {
	//准备读取文件
	fileName := "./top.csv"
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()

	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(row[1])
	}
}

func TestReadScan(t *testing.T) {
	wordfile, err := os.Open(`dict.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer wordfile.Close()

	done := make(chan bool)
	scanner := bufio.NewScanner(wordfile)
	go func() {
		for scanner.Scan() {
			pass := scanner.Text()
			fmt.Println(pass)
		}
		done <- true
	}()
	<-done
}

func TestScan(t *testing.T) {
	co := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		co[scanner.Text()]++
	}
	for line, n := range co {
		if n > 1 {
			fmt.Println(line, n)
		}
	}
}
