package webg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func WebGatherer() {
	requestServer()
}
func requestServer() {
	resp, err := http.Get("http://capregsoft.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// fmt.Printf("\nWebserver said: `%s`", string(body))
	file, _ := os.Create("webOutput.html")
	fmt.Fprint(file, string(body))
	file.Close()
}
