// Servidor en lenguaje GO


package main

import "net"
import "fmt"
import "bufio"
import "os"
import "time"

func main() {

  var port string = "2000"
  ln, _ := net.Listen("tcp", ":"+port)

  // Coloca nombre de usuario
  fmt.Print("Ingrese su nombre de usuario: ")
  var usuario string
  fmt.Scanln(&usuario)
  usuario += ": "


  // Loop
  for { 

  // Acepta conexión
  conn, _ := ln.Accept()


    for{

      // Escribir mensaje a enviar
      reader := bufio.NewReader(os.Stdin)
      fmt.Print("Yo: ")
      text, _ := reader.ReadString('\n')
      hora := "             ["+(time.Now()).Format("15:04")+"]"

      text += hora
      // Enviar al Client
      fmt.Fprintf(conn, usuario + text)


      // Recibe mensaje
      message, _ := bufio.NewReader(conn).ReadString('\n')
      fmt.Print(string(message))

    }
  }
}












// go run Server.go