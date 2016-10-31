package main

import (
	"encoding/json"	
	"log"
	"net/http"
)

func indexHandler (rw http.ResponseWriter, req *http.Request) {
	responseData := map[string]interface{} {
		"result" : "OK",
	}

	jsonEncoder := json.NewEncoder(rw)
	jsonEncoder.Encode(responseData)
}

func main() {
	http.HandleFunc("/", indexHandler);
	
	http.HandleFunc("/1", EulerOneHandler);
	http.HandleFunc("/2", EulerTwoHandler);
	http.HandleFunc("/3", EulerThreeHandler);
	http.HandleFunc("/4", EulerFourHandler);
	http.HandleFunc("/5", EulerFiveHandler);
	http.HandleFunc("/6", EulerSixHandler);
	http.HandleFunc("/7", EulerSevenHandler);
	http.HandleFunc("/8", EulerEightHandler);
	http.HandleFunc("/9", EulerNineHandler);
	http.HandleFunc("/10", EulerTenHandler);
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err);
	}
}