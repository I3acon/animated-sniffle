package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Validator struct {
	URL              string `json:"URL"`
	ValidatorAddress string `json:"Validator_address"`
	Explorer         string `json:"Explorer"`
}

func main() {
	band := Validator{
		URL:              "https://api-enterprise.nodex.network/v1/rpc/band/commit?=height",
		ValidatorAddress: "3626ADA5339AEDA433DC76D449544CA09834DAE7",
		Explorer:         "https://www.cosmoscan.io/validator/bandvaloper13jgj2fu2wc0pff7wnsang3m9xt3k67u8zdd5n5",
	}

	// axelarTestnet := Validator{
	// 	URL:              "https://api-enterprise.nodex.network/v1/rpc/axelar-testnet/block?height",
	// 	ValidatorAddress: "7A33B63F34C9C8085A7B6720D9C6F3A4799013F7",
	// 	Explorer:         "https://testnet.axelarscan.io/validator/axelarvaloper1fzsmceetzcvtv3su6z53jjt8a7yc2qczpkh2t4",
	// }

	bandJSON, _ := json.Marshal(band)
	// axelarTestnetJSON, _ := json.Marshal(axelarTestnet)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	for {
		postData("https://node-bot-api-4rjcvtq2ea-as.a.run.app/uptime/commit", bandJSON, headers)
		time.Sleep(3 * time.Second)
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
