package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)
//var webcon *websocket.Conn

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

var Connect *websocket.Conn

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
    for {
    // read in a message
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
    // print out that message for clarity
        fmt.Println(string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }
    }
}


func sendM(don *websocket.Conn) {
var err error
    //msg := `Hi, the handshake it complete!`
msg := []byte("Let's ogogogo to talk something.")
    err = don.WriteMessage(websocket.TextMessage, msg)
    if err != nil {
        log.Println(err)
    }
}


func homePage(w http.ResponseWriter, r *http.Request) {
	
	//go func() {
	    // simulate a long task with time.Sleep(). 5 seconds
	///    time.Sleep(5 * time.Second)
	    sendM(Connect)
	    //err = Conn.WriteMessage(1, []byte("Privet"))
	    /*if err != nil {
		log.Println(err)
		fmt.Fprintf(w,"Ops!");
	    }*/
	    
	    // note that you are using the copied context "cCp", IMPORTANT
	    //log.Println("Done! in path " + cCp.Request.URL.Path)
	//}()
	
	fmt.Fprintf(w, "-----321_123o234me Page ;)++ Truper And Git)")
}


func wsEndpoint(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }

    // upgrade this connection to a WebSocket
    // connection
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
    }

    log.Println("Client Connected")
    err = ws.WriteMessage(1, []byte("Hi Client!"))
    if err != nil {
        log.Println(err)
    }

    Connect = ws
    // listen indefinitely for new messages coming
    // through on our WebSocket connection
    reader(ws)
}

func setupRoutes() {
    //http.HandleFunc("/", homePage)
    http.HandleFunc("/ws", wsEndpoint)
    http.HandleFunc("/", homePage)
}

func main() {
    fmt.Println("Hello World")
    setupRoutes()
    log.Fatal(http.ListenAndServe(":8181", nil))
}