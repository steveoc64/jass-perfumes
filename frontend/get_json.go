package main

import (
	"encoding/json"

	"github.com/gopherjs/jquery"
)

const apiServer = "http://ministrywasm.com"

func GetJSON(url string, data interface{}, f func()) {
	jquery.GetJSON(apiServer+url, func(d interface{}) {
		j, _ := json.Marshal(d)
		json.Unmarshal(j, data)
		go f()
	})
}
