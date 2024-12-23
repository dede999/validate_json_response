package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Product struct {
	Data        string  `json:"Data"`
	ProductName string  `json:"Produto Coletado"`
	Price       float32 `json:"Preço Coletado"`
	LinkUrl     string  `json:"Link URL"`
}

func (p Product) String() string {
	return fmt.Sprintf("Data: %s\nProduct Name: %s\nPrice: %f\nLink URL: %s\n-------\n", p.Data, p.ProductName, p.Price, p.LinkUrl)
}

func main() {
	thisMoment := time.Now()
	data, err := ioutil.ReadFile("data/data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	data_lines := strings.Split(string(data), "\n")

	var products []Product
	var product Product

	for _, line := range data_lines {
		if line == "" {
			continue
		}
		err := json.Unmarshal([]byte(line), &product)
		if err != nil {
			fmt.Println("Error unmarshalling:", err)
			return
		}
		products = append(products, product)
	}
	fmt.Println("Total products:", len(products))
	fmt.Println("Time taken:", time.Since(thisMoment))
}
