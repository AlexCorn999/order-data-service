package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/AlexCorn999/order-data-service/internal/domain"
)

func main() {
	var order domain.Order

	fData, err := os.ReadFile("./model.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(fData, &order)
	fmt.Printf("%+v", order)
}
