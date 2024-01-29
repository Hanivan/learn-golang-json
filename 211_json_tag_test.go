package learngolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageURL string `json:"image_url"`
}

func TestJSONTagEncode(t *testing.T) {
	product := Product{
		Id:       "P0001",
		Name:     "apel",
		Price:    2000,
		ImageURL: "http://localhost:8080/img/apel.png",
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"apel","price":2000,"image_url":"http://localhost:8080/img/apel.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	json.Unmarshal(jsonBytes, product)

	fmt.Println(product)
	fmt.Println(product.Name)
}
