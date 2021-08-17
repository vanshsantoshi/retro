package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"os"
	"strings"
	"net/http"
)

var entries = []string{}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Printf("err is not nil !")
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	    //serve the resource
	    fmt.Fprintf(w, "<table><tr><th>Time</th><th>Action</th><th>Data</th></tr>")
	    for i, _ := range entries {
	    	fmt.Fprintf(w, "%s", entries[len(entries) - i - 1])
	    }
	    fmt.Fprintf(w, "</table>")
	case "POST":
	    //add entry
	    body, err := ioutil.ReadAll(r.Body)
	    if err != nil {
	    	fmt.Fprintf(w, err.Error())
		}
		entry := strings.SplitN(string(body), "|", 3)
		new_entry := fmt.Sprintf("<tr><td>%s</td><td>%s</td> <td>%s</td></tr>", entry[0], entry[1], entry[2])
	    entries = append(entries, new_entry)
	    if len(entries) > 100 {
	    	entries = entries[1:]
	    }
	    fmt.Fprintf(w, "POST\n")
	default:
	    //do nothing
	}
}

