package influxdb

import (
	"fmt"
	"github.com/influxdata/influxdb1-client/v2"
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