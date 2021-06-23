package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	URLS := []string{
		"golang.org",
		"google.com",
		"amazon.com",
		"stackoverflow.com",
		"facebook.com",
	}

	c := make(chan string)
	for _,url := range URLS {
		go checkURL(url,c)
		fmt.Println(<- c)
	}
}



func checkURL(url string,c chan string)  {
	response := getUrlStatus(url)
		if response.Status != "200 OK"{
			// fmt.Println("The Site is Down")
			c <- url + "The Site might be down from chan"
		} else{
			// fmt.Println(url," is good...")
			c <- url + " is good from chan"
		}
}



func getUrlStatus(url string) *http.Response {
	response, err := http.Get("https://" + url)
	if err != nil{
		fmt.Println("Error : ",err)
		os.Exit(1)
	}
	return response
}
