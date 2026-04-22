package main

import (
	"fmt"
	"io"
	"net/http"
)

func GetPosts() {
	url := "https://jsonplaceholder.typicode.com/posts"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
func main() {
	GetPosts()
}
