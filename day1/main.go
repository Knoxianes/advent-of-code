package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

func main() {
	cookie := &http.Cookie{
		Name:  "session",
		Value: "53616c7465645f5fb7f96dceb31e2ece648e84cc4ca1ecda9dffaadc0ee4e71eba58d1130a696741b67a35a4556179cd81d5d90f87751ec9c34c4d63e3582e87",
	}

	client := &http.Client{
		Jar:       &cookiejar.Jar{},
		Transport: &http.Transport{},
	}

	req, err := http.NewRequest("GET", "https://adventofcode.com/2015/day/1/input", nil)
	req.AddCookie(cookie)
	
	if err != nil {
		// handle error
	}


	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}

	body, err := io.ReadAll(resp.Body)
	sb := string(body)
	slicedString := strings.Split(sb,"")
	currentFloor := 0
	for i,value := range slicedString{
		if currentFloor < 0{
			fmt.Println(i)
			break
		}
		if value == "("{
			currentFloor++
		}else{
			currentFloor--
		}
	}

}
