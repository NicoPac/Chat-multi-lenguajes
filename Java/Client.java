import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;
import java.util.Scanner;
import java.text.Format;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.Calendar;

public class Client {
    public static void main(String[] args){
        final Socket clientSocket; // socket utilizado para enviar y revicir información del server
        final BufferedReader in;   // Se utiliza para leer datos del socket
        final PrintWriter out;     // Se utiliza para escribir datos del socket
        final Scanner sc = new Scanner(System.in); // Lee lo escrito por teclado

        try {
            String hostname = "localhost";
            int port = 2000;
            clientSocket = new Socket(hostname,port);
            out = new PrintWriter(clientSocket.getOutputStream());
            in = new BufferedReader(new InputStreamReader(clientSocket.getInputStream()));


            System.out.println("Ingrese su nombre de usuario: ");
            String usuario= (new Scanner(System.in)).nextLine();

            Thread sender = new Thread(new Runnable() {
                String msg;
                @Override
                public void run() {
                    while(true){

                        Format hora = new SimpleDateFormat("HH.mm");
                        String horaStr = hora.format(new Date());

                        msg = sc.nextLine();
                        out.println(usuario + ": " +msg + "     ["+ horaStr +"]");
                        out.flush();
                    }
                }
            });
            sender.start();
            Thread receiver = new Thread(new Runnable() {
                String msg;
                @Override
                public void run() {
                    try {
                        msg = in.readLine();
                        while(msg!=null){
                            System.out.println(msg);
                            msg = in.readLine();
                        }
                        System.out.println("Servidor fuera de servicio");
                        out.close();
                        clientSocket.close();
                    } catch (IOException e) {
                        e.printStackTrace();
                    }
                }
            });
            receiver .start();
    }catch (IOException e){
        e.printStackTrace();
        }
    }
}



/*
javac Client.java

java Client
*/