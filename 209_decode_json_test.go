package learngolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonRequest := `{"FirstName":"Hanivan","MiddleName":"Rizky","LastName":"Sobari"}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}
