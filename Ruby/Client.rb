# Cliente en lenguaje RUBY

require 'socket'      # Obtenemos la biblioteca socket

hostname = 'localhost'
port = 2000

server = TCPSocket.open(hostname, port)

puts "Ingrese su nombre de usuario:"  #Coloca nombre de usuario
usuario = gets.chomp

loop {

  mensaje_recibido = server.gets   # Se lee cada línea de los datos recibidos del socket
  puts mensaje_recibido.chop      # Muestra los datos en la terminal
  
  print "Yo: "
  mensaje = gets.chomp            # Colocar el mensaje por consola
  server.puts usuario + ": " + mensaje +"     [" +Time.now.strftime('%H:%M')+ "]"    #Envia el mensaje con nombre de usuario y hora

}




# $ ruby clientRuby.rb 