package learngolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestJSONStreamingDecoder(t *testing.T) {
	reader, _ := os.Open("product.json")
	decoder := json.NewDecoder(reader)

	product := &Product{}
	_ = decoder.Decode(product)

	fmt.Println(product)
}
