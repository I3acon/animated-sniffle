package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Body struct {
	URL                  string `json:"URL"`
	ValidatorAddress     string `json:"Validator_address"`
	Explorer             string `json:"Explorer"`
	MissedBlockThreshold int    `json:"MissedBlockThreshold"`
}

func main() {
	args := os.Args[1:]
	if len(args) != 5 {
		fmt.Println("Usage: main <URL> <Validator_address> <Explorer> <MissedBlockThreshold> <Frequency in ms>")
		return
	}
	intMissedBlockThreshold, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("Error converting MissedBlockThreshold to int:", err)
		return
	}
	body := Body{
		URL:                  args[0],
		ValidatorAddress:     args[1],
		Explorer:             args[2],
		MissedBlockThreshold: intMissedBlockThreshold,
	}

	bodyJSON, _ := json.Marshal(body)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	intFrequency, err := strconv.Atoi(args[4])
	if err != nil {
		fmt.Println("Error converting Frequency to int:", err)
		return
	}
	for {
		postData("http://localhost:9000/uptime/commit", bodyJSON, headers)
		time.Sleep(time.Duration(intFrequency) * time.Millisecond)
	}
}

func postData(url string, data []byte, headers map[string]string) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
}
