# Servidor en lenguaje RUBY

require 'socket'                # Obtenemos la biblioteca socket

port= 2000
server = TCPServer.open(port)   # El socket escuchará el puerto deseado

  puts "Ingrese su nombre de usuario:"
  usuario = gets.chomp

loop {                          # Se ejecutará el server infinitamente

  client = server.accept        # Acepta la conexión del cliente
  

  while true

    mensaje_recibido = client.gets   # Se lee cada línea de los datos recibidos del socket
    puts mensaje_recibido.chop      # Muestra los datos en la terminal

    print "Yo:  "
    mensaje = gets.chomp          # Colocar el mensaje por consola
    client.puts usuario + ": " + mensaje +"     ["+ Time.now.strftime('%H:%M')+ "]"     #Envia el mensaje con nombre de usuario y hora
  
  end
}





# $ ruby servMundo.rb 



# $ ruby servMundo.rb 
