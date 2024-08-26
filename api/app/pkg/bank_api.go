package pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

// Package pkg provides a bank API implementation.
func init() {
	viper.SetDefault("BANK_API_URL", "http://localhost:8081")
	viper.SetDefault("BANK_ID", 1)
}

// BankApi represents a bank API client.
type BankApi struct {
	bankID  int
	baseURL string
}

// NewBankApi creates a new instance of BankApi.
func NewBankApi() *BankApi {
	return &BankApi{
		bankID:  viper.GetInt("BANK_ID"),
		baseURL: viper.GetString("BANK_API_URL"),
	}
}

// GetBankID returns the bank ID as a string.
func (api *BankApi) GetBankID() string {
	return fmt.Sprintf("%d", api.bankID)
}

// Get sends a GET request to the specified endpoint and returns the response body as a string.
func (api *BankApi) Get(endpoint string) (string, error) {
	url := fmt.Sprintf("%s%s", api.baseURL, endpoint)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Post sends a POST request with the provided data to the specified endpoint and returns the response body as a string.
func (api *BankApi) Post(endpoint string, data []byte) (string, error) {
	url := fmt.Sprintf("%s%s", api.baseURL, endpoint)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Delete sends a DELETE request to the specified endpoint and returns the response body as a string.
func (api *BankApi) Delete(endpoint string) (string, error) {
	url := fmt.Sprintf("%s%s", api.baseURL, endpoint)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
