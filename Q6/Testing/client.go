package Testing

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/dchest/uniuri"
)

func Client() {
	rand.Seed(time.Now().UnixNano())
	counter := 1
	for {
		time.Sleep(5 * time.Second)
		postBody, _ := json.Marshal(map[string]int{
			"counter": counter,
		})
		requestBody := bytes.NewBuffer(postBody)
		//Leverage Go's HTTP Post function to make request
		req, _ := http.NewRequest("POST", "http://localhost:3000/", requestBody)
		randomWord := uniuri.New()
		req.Header.Set("X-RANDOM", randomWord)
		req.Header.Set("content-type", "application/json")

		client := http.Client{}
		resp, err := client.Do(req)
		//Handle Error
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()
		counter++
	}
}
