package main

import (
  "fmt"
	"net/http"
  "bytes"
  "io/ioutil"
)

func callAPI(url String,jsonStr []byte){



}
func main() {
    url := "http://127.0.0.1:8080/v1/movies"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"Name":"A","Desc":"test","Rating":10}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    //req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
