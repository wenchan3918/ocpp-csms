/*
 * OCPP Device Service
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.0.0
 * Contact: gr.szalay@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	sw "github.com/gregszalay/ocpp-csms/device-service/http"
	"github.com/gregszalay/ocpp-csms/device-service/pubsub"
)

func main() {
	log.Printf("Server started")

	var waitgroup sync.WaitGroup

	waitgroup.Add(1)
	go func() {
		fmt.Println("Inside my goroutine 1")
		router := sw.NewRouter()
		log.Fatal(http.ListenAndServe(":5000", router))
		waitgroup.Done()
	}()

	waitgroup.Add(1)
	go func() {
		fmt.Println("Inside my goroutine 2")
		pubsub.Subscribe()
		waitgroup.Done()
	}()

	waitgroup.Wait()

}
