package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SaveData(id int) {
	// id := 321006
	url := fmt.Sprintf("https://ncmym.edu.bd/honours/application/preview/%d", id)
	res, _ := http.Get(url)
	if res.StatusCode == 200 {
		filePath := fmt.Sprintf("./data/%d.html", id)
		body, _ := ioutil.ReadAll(res.Body)
		ioutil.WriteFile(filePath, body, 0644)
		fmt.Printf("\n%d - found", id)

	} else {
		fmt.Printf("\n%d - notfound", id)
	}
	res.Body.Close()

}

func main() {
	startId := 202328
	for id := startId; id < (startId * 10); id++ {
		SaveData(id)
	}
}
