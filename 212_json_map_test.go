package learngolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONMapDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"apel","price":2000,"image_url":"http://localhost:8080/img/apel.png"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestJSONMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"name":  "Apel",
		"price": 2000,
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}
