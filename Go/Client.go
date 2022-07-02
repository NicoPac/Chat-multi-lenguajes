// Cliente en lenguaje GO


package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

  // connect to server
  var port string = "2000"
  conn, _ := net.Dial("tcp", "127.0.0.1:"+port)



  fmt.Print("Ingrese su nombre de usuario: ")
  usuario := bufio.NewScanner(os.Stdin).Text()
  fmt.Println(usuario)



  for { 
  
	// what to send?
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Yo: ")
    text, _ := reader.ReadString('\n')

    // send to server
    fmt.Fprintf(conn, text)

    // wait for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server: "+message)
  }
}










// go run Client.go