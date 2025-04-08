package main

import (
	"fmt"

	"github.com/samirhembrom/blogaggregator/internal/config"
)

func main() {
	jsonData, err := config.Read()
	if err != nil {
		fmt.Printf("err")
	}
	err = config.SetUser(jsonData)
	if err != nil {
		fmt.Printf("err")
	}
	fmt.Printf("URL:%s ", jsonData.URL)
	fmt.Printf("user:%s ", jsonData.CurrentUser)
}
