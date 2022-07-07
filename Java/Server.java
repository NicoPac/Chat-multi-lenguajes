import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.Scanner;
import java.text.Format;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.Calendar;



public class Server {
    public static void main(String[] args){
        final ServerSocket serverSocket ; // Socket utilizado para enviar y recibir información como Server
        final Socket clientSocket ;
        final BufferedReader in;    // Se utiliza para leer datos del socket
        final PrintWriter out;      // Se utiliza para escribir datos del socket
        final Scanner sc=new Scanner(System.in); // Lee lo escrito por teclado

        try {
            int port = 2000;
            serverSocket = new ServerSocket(port);
            clientSocket = serverSocket.accept();
            out = new PrintWriter(clientSocket.getOutputStream());
            in = new BufferedReader (new InputStreamReader(clientSocket.getInputStream()));

            // Solicita nombre de usuario
            System.out.println("Ingrese su nombre de usuario: ");
            String usuario= (new Scanner(System.in)).nextLine();


            // Enviar mensaje
            Thread sender= new Thread(new Runnable() {
                String msg; // Variable que contiene lo ingresado por teclado por el usuario
                @Override   
                public void run() {
                    while(true){

                        Format hora = new SimpleDateFormat("HH.mm");
                        String horaStr = hora.format(new Date());

                        msg = sc.nextLine(); //Lee lo escrito por teclado
                        out.println(usuario + ": " +msg + "     ["+ horaStr +"]");    // Prepara para enviar el nombre de usuario, mensaje y la hora
                        out.flush();   // Envía el mensaje
                    }
                }
            });
            sender.start();


            // Recibir mensaje
            Thread receive= new Thread(new Runnable() {
                String msg ;
                @Override
                public void run() {
                    try {
                        msg = in.readLine();
                        
                        while(msg!=null){
                            System.out.println(msg);
                            msg = in.readLine();
                        }

                        System.out.println("Cliente desconectado");

                        out.close();
                        clientSocket.close();
                        serverSocket.close();
                    } catch (IOException e) {
                        e.printStackTrace();
                    }
                }
            });
            receive.start();

            
        } catch (IOException e) {
            e.printStackTrace();
        }


    }
}


/*
javac Server.java

java Server

*/
