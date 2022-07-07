// Cliente en lenguaje GO


package main

import "net"
import "fmt"
import "bufio"
import "os"
import "time"

func main() {

  // Conecta al servidor
  var port string = "2000"
  conn, _ := net.Dial("tcp", "127.0.0.1:"+port)


  fmt.Print("Ingrese su nombre de usuario: ")
  var usuario string
  fmt.Scanln(&usuario)
  usuario += ": "

  // Loop
  for { 

    // Recibe mensaje
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print(message)

	  // Escribir mensaje a enviar
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Yo: ")
    text, _ := reader.ReadString('\n')
    hora := "             ["+(time.Now()).Format("15:04")+"]"

    text += hora

    // Envia al Server
    fmt.Fprintf(conn, usuario + text)

  }
}










// go run Client.go