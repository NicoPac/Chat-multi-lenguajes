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
  var usuario string
  fmt.Scanln(&usuario)
  usuario += ": "

  for { 
  
	// what to send?
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Yo: ")
    text, _ := reader.ReadString('\n')

    // send to server
    fmt.Fprintf(conn, usuario + text)

    // wait for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print(message)
  }
}










// go run Client.go