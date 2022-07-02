// Servidor en lenguaje GO


package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

  var port string = "2000"
  ln, _ := net.Listen("tcp", ":"+port)


  
  fmt.Print("Ingrese su nombre de usuario: ")
  usuario := bufio.NewScanner(os.Stdin).Text()
  
  fmt.Println(usuario)



  // accept connection
  conn, _ := ln.Accept()


  // run loop forever (or until ctrl-c)
  for { 
    // get message, output
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message Received:", string(message))

	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Yo: ")
    text, _ := reader.ReadString('\n')

    // send to server
    fmt.Fprintf(conn,text)
  }
}












// go run Server.go