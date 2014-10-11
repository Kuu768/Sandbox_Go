package main

import (
	"fmt"
)

func main() {
	hash := map[string]string{
		"name":    "hoge",
		"age":     "20",
		"address": "Osaka",
	}

	fmt.Println(hash)
	fmt.Println(hash["name"])
	fmt.Println(hash["age"])
	fmt.Println(hash["address"])

	hash["hoge"] = "Hoge"
	fmt.Println(hash)
}
