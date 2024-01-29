package learngolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Hanivan",
		MiddleName: "Rizky",
		LastName:   "Sobari",
		Hobbies:    []string{"Music", "Gaming", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonRequest := `{"FirstName":"Hanivan","MiddleName":"Rizky","LastName":"Sobari","Hobbies":["Music","Gaming","Coding"]}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestJSONArrayEndcodeComplex(t *testing.T) {
	customer := Customer{
		FirstName:  "Hanivan",
		MiddleName: "Rizky",
		LastName:   "Sobari",
		Addresses: []Address{
			{
				Street:     "Jl. Abc",
				Country:    "Indonesia",
				PostalCode: "98162317352",
			},
			{
				Street:     "Rumah ke 99",
				Country:    "Indonesia",
				PostalCode: "7123512389",
			},
		},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJSONArrayDecodeComplex(t *testing.T) {
	jsonString := `{"FirstName":"Hanivan","MiddleName":"Rizky","LastName":"Sobari","Hobbies":null,"Addresses":[{"Street":"Jl. Abc","Country":"Indonesia","PostalCode":"98162317352"},{"Street":"Rumah ke 99","Country":"Indonesia","PostalCode":"7123512389"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestOnlyJSONArrayDecodeComplex(t *testing.T) {
	jsonString := `[{"Street":"Jl. Abc","Country":"Indonesia","PostalCode":"98162317352"},{"Street":"Rumah ke 99","Country":"Indonesia","PostalCode":"7123512389"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}
