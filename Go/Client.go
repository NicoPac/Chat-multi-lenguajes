// Cliente en lenguaje GO


package main

import "net"
import "fmt"
import "bufio"
import "os"

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
  
	// Mensaje a enviar
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Yo: ")
    text, _ := reader.ReadString('\n')

    // Envia al Server
    fmt.Fprintf(conn, usuario + text)

    // Espera respuesta
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print(message)
  }
}










// go run Client.go