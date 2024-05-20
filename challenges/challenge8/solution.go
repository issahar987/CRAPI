package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	configurator "github.com/tomek-skrond/crapiconfigurator"
)

type ProductResponse struct {
	ProductID int     `json:"id"`
	Message   string  `json:"message"`
	Credit    float64 `json:"credit"`
}
type MoneyResponse struct {
	Credit float64 `json:"credit"`
}

func GetMoney(token string) float64 {
	req, err := http.NewRequest("GET", "https://crapi.bobaklabs.com:8443/workshop/api/shop/products", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req = configurator.ConfigureRequest(req, token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := configurator.CustomHttpClient().Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var accountBalance MoneyResponse
	if err := json.NewDecoder(resp.Body).Decode(&accountBalance); err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(accountBalance.Credit)

	return accountBalance.Credit
}
func ReturnProduct(product ProductResponse, token string) {
	productReturn := map[string]string{
		"status": "returned",
	}
	jsonReturn, err := json.Marshal(productReturn)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("https://crapi.bobaklabs.com:8443/workshop/api/shop/orders/%d", product.ProductID), bytes.NewBuffer(jsonReturn))
	if err != nil {
		log.Fatalln(err)
	}
	req = configurator.ConfigureRequest(req, token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := configurator.CustomHttpClient().Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	// fmt.Println(resp.StatusCode)

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading response body:", err)
	}

	fmt.Println(string(responseBody))

}

func BuyProduct(token string) (*ProductResponse, error) {
	productBuy := map[string]int{
		"product_id": 1, "quantity": 1,
	}
	jsonBuy, err := json.Marshal(productBuy)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	productReq, err := http.NewRequest("POST", "https://crapi.bobaklabs.com:8443/workshop/api/shop/orders", bytes.NewBuffer(jsonBuy))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	productReq = configurator.ConfigureRequest(productReq, token)
	productReq.Header.Add("Content-Type", "application/json")

	productResp, err := configurator.CustomHttpClient().Do(productReq)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer productResp.Body.Close()

	// fmt.Println(productResp.StatusCode)
	var product ProductResponse
	if err := json.NewDecoder(productResp.Body).Decode(&product); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Bought product:")
	fmt.Printf("%+v\n", product)
	return &product, nil
}
