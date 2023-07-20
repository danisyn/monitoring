package main

import (
	"fmt"
	"net/http"
	"ingress-monitor/structures"
)

var array []structures.ClientIngress

func pingger(hosts []structures.ClientIngress) []structures.ClientIngress {

	for i, v := range hosts {
		resp, err := http.Get("https://" + v.Host + "/")
		if err != nil {
			fmt.Println("Error:", err)
		}
		defer resp.Body.Close()
	
		// Extract the HTTP response code
		responseCode := resp.StatusCode
	
		// Print the HTTP response code
		//fmt.Println("HTTP Response Code:", responseCode)

		array[i].Response = responseCode

	}

	return array

}