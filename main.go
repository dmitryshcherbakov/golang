package main

import (
    //"time"
    "strconv"
    "math/rand"
    "fmt"
    "log"
    "net/http"
    "io/ioutil"

    "github.com/gorilla/websocket"
)
//var webcon *websocket.Conn

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

type User struct {
    Connect *websocket.Conn
    client_key string
}
/*
type Connect_Users []User
*/

var userConnect []User

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


func sendM(don *websocket.Conn,message string) {
var err error
    //msg := `Hi, the handshake it complete!`
msg := []byte(message)
    err = don.WriteMessage(websocket.TextMessage, msg)
    if err != nil {
        log.Println(err)
    }
}

func goodGet() {

resp, err := http.Get("http://5.53.125.62/goget")
   if err != nil {
      log.Fatalln(err)
   }
//We Read the response body on the line below.
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
//Convert the body to type string
   sb := string(body)
   //sb = strconv.Itoa(rand.Intn(100))
   log.Printf(sb)
   //sendM(Connect,"Super!"+sb)
    for i, s := range userConnect {
	go func(i int,s User,sb string) { 
	str := strconv.Itoa(i)
	sendM(s.Connect,s.client_key+"=>"+str+" - Ok:"+sb) }(i,s,sb)
	//sendM(s.Connect,string(i)+"Ok"+sb);
    }
}

func goGet(w http.ResponseWriter, r *http.Request) {

    //time.Sleep(5 * time.Second)
    sb := strconv.Itoa(rand.Intn(100))
    fmt.Fprintf(w, sb)
}



func homePage(w http.ResponseWriter, r *http.Request) {
	
	//go func() {
	    // simulate a long task with time.Sleep(). 5 seconds
	///    time.Sleep(5 * time.Second)
	     go func() {
    		//time.Sleep(2 * time.Second)
		//sendM(Connect)
		goodGet()
	    }()
	    //err = Conn.WriteMessage(1, []byte("Privet"))
	    /*if err != nil {
		log.Println(err)
		fmt.Fprintf(w,"Ops!");
	    }*/
	    
	    // note that you are using the copied context "cCp", IMPORTANT
	    //log.Println("Done! in path " + cCp.Request.URL.Path)
	//}()
	
	fmt.Fprintf(w, "123-----321_123o234me Page ;)++ Truper And Git)")
}


func wsEndpoint(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }

    // upgrade this connection to a WebSocket
    // connection
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
    }

    params := r.URL.Query()
    client_key := string(params.Get("client_key"))

    log.Println("Client Connected")
    err = ws.WriteMessage(1, []byte("Hi Client!"))
    if err != nil {
        log.Println(err)
    }

    Connect = ws

    user_conn := User { Connect: ws, client_key: client_key }

    userConnect = append(userConnect,user_conn)

    // listen indefinitely for new messages coming
    // through on our WebSocket connection
    //go func() { reader(ws) }()
    reader(ws)
}

func setupRoutes() {
    http.HandleFunc("/goget", goGet)
    http.HandleFunc("/ws", wsEndpoint)
    http.HandleFunc("/", homePage)
}

func main() {
    //userConnect := []*websocket.Conn{}
    fmt.Println("Hello World")
    setupRoutes()
    log.Fatal(http.ListenAndServe(":8181", nil))
}