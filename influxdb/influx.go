package influxdb

import (
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
	"time"
	"ingress-monitor/structures"
	"context"
)

func PingInflux() {

	// Create a new HTTPClient
	client := influxdb2.NewClient(influxURL, token)
	// Test the connection
	_, err := client.Ping(context.Background())
	if err != nil {
		fmt.Println("Error pinging InfluxDB:", err)
		return
	}

	fmt.Println("Connected to InfluxDB successfully!")
}

func StoreData(array []structures.ClientIngress) {
	// Create an HTTPClient with the desired configuration
	client := influxdb2.NewClient(influxURL, token)
	writeAPI := client.WriteAPIBlocking(org, bucket)
	for _, ingress := range array {
		// Create a new point batch
		p := influxdb2.NewPoint("stat", 
		map[string]string{"resource": "Ingress"},
		map[string]interface{}{"name": ingress.Name, "host": ingress.Host, "response": ingress.Response},
		time.Now())	

		err := writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			fmt.Println("Error writing to influx")
			fmt.Println(err)
		} else {
			fmt.Println("Data point stored in InfluxDB successfully! - Ingress name - " + ingress.Name)
		}
	}
}