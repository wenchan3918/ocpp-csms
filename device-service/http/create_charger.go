/*
 * OCPP Device Service
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.0.0
 * Contact: gr.szalay@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chargerevolution/ocpp-device-service-api/types"
	"github.com/chargerevolution/ocpp-device-service/chargingstations"
)

func CreateNewCharger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	var newCharger types.Charger
	error := json.Unmarshal(body, &newCharger)
	if error != nil {
		fmt.Printf("could not unmarshal json: %s\n", error)
		return
	}

	fmt.Printf("New charger created: %v\n", newCharger)

	chargingstations.CreateCharger(&newCharger, newCharger.Id)

	w.Write([]byte("Success!"))
}
