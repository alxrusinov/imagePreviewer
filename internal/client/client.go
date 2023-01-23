package client

import (
	"fmt"
	"io"
	"net/http"
)

func DoClient() {
	client := http.Client{}

	resp, err := client.Get("http://ya.ru")

	if err != nil {
		fmt.Println("ERR", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

}
