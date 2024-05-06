## **net-cat** 

### **Objectives**

This project shows the basics of the technology for implementing group chat using a TCP server. The project uses a client-server architecture, where the server creates the chat logic and interaction mechanisms between clients. The server simultaneously processes no more than 10 connections (clients). Each connection runs in parallel in the background using a goroutine mechanism, and communication between them is carried out using channels. To synchronize the work of goroutines when using shared resources, mutexes are used.

### **Instructions**

Procedure for the user:
<br>

1. Clone the project from the repository
2. Go to the project and enter the command: `go run .`
3. After successfully starting the server on the local machine with the default port 8989, open a new terminal window and enter the command: `nc localhost 8989`
5. Enter your nickname and press `Enter`
6. You are connected to the chat
