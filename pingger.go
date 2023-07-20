package main

import (
	"fmt"
	"net/http"
)

func pingger(hosts []ClientIngress) {

	for i, v := range hosts {
		resp, err := http.Get("https://" + v.Host + "/")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()
	
		// Extract the HTTP response code
		responseCode := resp.StatusCode
	
		// Print the HTTP response code
		//fmt.Println("HTTP Response Code:", responseCode)

		array[i].Response = responseCode

	}

	fmt.Println(array)

}