// Servidor en lenguaje GO


package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

  var port string = "2000"
  ln, _ := net.Listen("tcp", ":"+port)

  // Coloca nombre de usuario
  fmt.Print("Ingrese su nombre de usuario: ")
  var usuario string
  fmt.Scanln(&usuario)
  usuario += ": "


  // Acepta conexión
  conn, _ := ln.Accept()


  // Loop
  for { 


    // Escribir mensaje a enviar
	  reader := bufio.NewReader(os.Stdin)
    fmt.Print("Yo: ")
    text, _ := reader.ReadString('\n')

    // Enviar al Client
    fmt.Fprintf(conn, usuario + text)
    


    // Recibe mensaje
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print(string(message))


  }
}












// go run Server.go