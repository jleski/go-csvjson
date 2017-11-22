package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"encoding/json"
)

func routeCsv2Json(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // respond only to HTTP POSt
		jsonString, err := json.Marshal(r.URL.Query())
		if err != nil {
			fmt.Fprintln(w, "error: unable to parse json")
		} else {
			fmt.Fprintf(w, "%s\n", jsonString)
		}
	}
}

func routeJson2Csv(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // respond only to HTTP POST
		i := 0
		keys := make([]string, len(r.URL.Query()))
		values := make([]string, len(r.URL.Query()))
		for k, v := range r.URL.Query() {
			values[i] = strings.Join(v,"")
			keys[i] = k
			i++
		}
		fmt.Fprintf(w, "%s\n", strings.Join(keys, ",")) // output comma delimited header
		fmt.Fprintf(w, "%s\n", strings.Join(values,",")) // output comma delimited values
	}
}

func main() {
	http.HandleFunc("/json2csv", routeJson2Csv) // route
	http.HandleFunc("/csv2json", routeCsv2Json) // route
	err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil) // setup https listening
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}