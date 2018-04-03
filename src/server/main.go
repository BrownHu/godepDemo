package main

import (
	"log"
	"github.com/tidwall/gjson"
)

func main() {
	jsonStr := `{"name":"YOOZOO" }`
	log.Println(gjson.Get(jsonStr, "name"))
}