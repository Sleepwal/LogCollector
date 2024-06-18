package main

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	"time"
)

func main() {
	WritePoints()
	QueryDB()
}

func WritePoints() {
	bucket := "LogCollection"
	org := "Organization"
	token := "xCs_T4YxAzOWC2IxyJcTp_ExGynRXQgDFqjXy8opWPULIhCKJYL8SMu9HdvG78Kkta970pLflgivRZECsBb_Ew=="
	// Store the URL of your InfluxDB instance
	url := "http://localhost:8086"
	// Create new client with default option for server url authenticate by token
	client := influxdb2.NewClient(url, token)
	// User blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking(org, bucket)
	// Create point using full params constructor
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now())
	// Write point immediately
	writeAPI.WritePoint(context.Background(), p)
	// Ensures background processes finishes
	client.Close()
}

func QueryDB() {
	//bucket := "LogCollection"
	org := "Organization"
	token := "xCs_T4YxAzOWC2IxyJcTp_ExGynRXQgDFqjXy8opWPULIhCKJYL8SMu9HdvG78Kkta970pLflgivRZECsBb_Ew=="
	// Store the URL of your InfluxDB instance
	url := "http://localhost:8086"
	// Create cli
	cli := influxdb2.NewClient(url, token)
	// Get query cli
	queryAPI := cli.QueryAPI(org)
	// Get QueryTableResult
	result, err := queryAPI.Query(context.Background(), `from(bucket:"LogCollection")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`)
	if err == nil {
		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
			fmt.Printf("value: %v\n", result.Record().Value())
		}
		// Check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
	// Ensures background processes finishes
	cli.Close()
}
