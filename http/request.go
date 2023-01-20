package main

import (
	"fmt"
	"testing"
	// "log"
	"net/http"
	"time"
	"os"
	"encoding/json"
	"errors"
	"bytes"
	"io"
	"strconv"

	"github.com/joho/godotenv"
)

type ClientToken struct {
	AccessToken		string    	`json:"access_token"`
	TokenType		string 		`json:"token_type"`
	Expiry			string		`json:"expires_in"`
	Error			string		`json:"error"`
}

type WalletType struct {
	CustomerId		string    	`json:"CustomerId"`
	WalletAddress	string 		`json:"WalletAddress"`
	WalletType		string		`json:"WalletType"`
	WalletId		string		`json:"WalletId"`
}

type Wallet struct {
	Wallet			WalletType			`json:"wallet"`
}

type WalletArray struct {
	Wallet			[]WalletType		`json:"wallet"`
}

type Wallets struct {
	Wallets			WalletArray			`json:"wallets"`
}

type ErrorType struct {
	Code		string    	`json:"code"`
	Message		string 		`json:"message"`
	Type		string		`json:"type"`
}

type Error struct {
	Error			ErrorType		`json:"error"`
}

// create a http/client once only, and reuse the TCP connection for all the functions
func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func GetClientToken(client *http.Client, grpcNonce string) (string, error) {
	
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(".env loading failure")
	}

	url := os.Getenv("BASE_URL") + "/api/auth/GetClientToken"

    request, err := http.NewRequest(http.MethodGet, url, nil)

    request.Header.Set("Authorization", os.Getenv("CREDENTIAL"))
    request.Header.Set("Accept", "*/*")
	request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Grpc-Metadata-Nonce", grpcNonce)

	responce, err := client.Do(request)
	if err != nil {
		return "", errors.New("Failed to get response from GetClientToken")
	}
	defer responce.Body.Close()
	body, err := io.ReadAll(responce.Body)

	var clientToken *ClientToken
	err = json.Unmarshal(body, &clientToken)
	if err != nil {
		return "", errors.New("Failed to parse ClientToken into json")
	} else if clientToken.Error == "Invalid nonce." {
		return "", errors.New("Invalid nonce")
	}

	return clientToken.AccessToken, nil
}

func GenerateDepositWallet(client *http.Client, accessToken string, customerId string) (Wallet, error) {

	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(".env loading failure")
	}

	url := os.Getenv("BASE_URL") + "/api/wallets"

	value := map[string]string{"CustomerId": customerId, "WalletType": "Deposit"}
	jsonData, err := json.Marshal(value)

    request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))

    request.Header.Set("Authorization", "bearer " + accessToken)
	request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-Type", "application/json")

	// make request to insert wallet for customerId
	response, err := client.Do(request)
	if err != nil {
		return Wallet{WalletType{"", "", "", ""}}, 
				errors.New("Failed to get response from InsertWallet")
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	// Unmarshal the body into json
	var wallet Wallet
	err = json.Unmarshal(body, &wallet)
	if err != nil {
		return Wallet{WalletType{"", "", "", ""}},
				errors.New("Wallet limit reached")
	}

	fmt.Println(wallet)

	return wallet, nil
}

func TestCreateDepositWallets(t *testing.T) {

	client := httpClient()
	// grpcNonce has to be changed every time before te existing one expires
	grpcNonce := "122"
	accessToken, _ := GetClientToken(client, grpcNonce)

	// walletId = 200-1199 are used for this test
	walletArray := make([]Wallet, 0)
	for customer := 1100; customer < 1200; customer++ {
		walletAddress, err := GenerateDepositWallet(client, accessToken, strconv.Itoa(customer))
		if err != nil {
			fmt.Println(err)
		}
		walletArray = append(walletArray, walletAddress)
	}
}

func TestGetDepositWallets(t *testing.T) {

	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(".env loading failure")
	}

	client := httpClient()

	// walletId = 200-1199 are used for this test
	walletArray := make([]string, 0)
	for customer := 300; customer < 310; customer++ {
		url := os.Getenv("BASE_URL") + "/api/wallets/" + strconv.Itoa(customer)

		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			fmt.Printf("Failed to create wallet address request for Customer Id: %v", customer)
		}
		request.Header.Set("Authorization", "bearer " + "6cfe0c65-62a9-4f90-88de-e3c332fa3bb2")
		request.Header.Set("Accept-Encoding", "gzip, deflate, br")
		request.Header.Set("Accept", "*/*")
		request.Header.Set("Connection", "keep-alive")
		request.Header.Set("Content-Type", "application/json")

		// make request to insert wallet for customerId
		response, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Failed to get response from InsertWallet")
		}

		// Unmarshal the body into json
		var wallets Wallets
		err = json.Unmarshal(body, &wallets)
		walletArray = append(walletArray, wallets.Wallets.Wallet[0].WalletAddress)
	}
	fmt.Println(len(walletArray))
}