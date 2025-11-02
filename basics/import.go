package basics

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello Go standard library")
	resp, err := http.Get("https://google.com/")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}
