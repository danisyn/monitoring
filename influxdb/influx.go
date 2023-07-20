package influxdb

import (
	"fmt"
	"github.com/influxdata/influxdb1-client/v2"
	"log"
	"time"
	"ingress-monitor/structures"
)

func PingInflux() {

	// Create a new HTTPClient
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     influxURL,
		Username: username,
		Password: password,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB HTTP client:", err)
		return
	}
	defer cli.Close()

	// Test the connection
	_, _, err = cli.Ping(0)
	if err != nil {
		fmt.Println("Error pinging InfluxDB:", err)
		return
	}

	fmt.Println("Connected to InfluxDB successfully!")
}

func StoreData(array []structures.ClientIngress) {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     influxURL,
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal("Error creating InfluxDB HTTP client:", err)
	}
	defer cli.Close()

	for _, ingress := range array {
		// Create a new point batch
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database:  database,
			Precision: "s", // Set the precision of timestamps (s: seconds, ms: milliseconds, etc.)
		})
		if err != nil {
			log.Fatal("Error creating batch points:", err)
		}

		tags := map[string]string{"IngressName": ingress.Name, "IngressHost": ingress.Host} // Tags associated with the data point
		fields := map[string]interface{}{"IngressResponse": ingress.Response} // Field values of the data point
		pt, err := client.NewPoint("ingress_liveness", tags, fields, time.Now())
		if err != nil {
			log.Fatal("Error creating data point:", err)
		}

		bp.AddPoint(pt)

		// Write the batch to InfluxDB
		if err := cli.Write(bp); err != nil {
			log.Fatal("Error writing batch to InfluxDB:", err)
		}

		fmt.Println("Data point stored in InfluxDB successfully! - Ingress name - " + ingress.Name)
	}

}