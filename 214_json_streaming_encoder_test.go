package learngolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestJSONStreamingEncoder(t *testing.T) {
	w, _ := os.Create("sample_output.json")
	encoder := json.NewEncoder(w)

	product := Product{
		Id:       "P0002",
		Name:     "Pear",
		Price:    5000,
		ImageURL: "",
	}
	_ = encoder.Encode(product)

	fmt.Println(product)
}
