package http

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

type ErrorType struct {
	Code		string    	`json:"code"`
	Message		string 		`json:"message"`
	Type		string		`json:"type"`
}

type WalletArray struct {
	Wallet			WalletType		`json:"wallet"`
	Error			ErrorType		`json:"error"`
}

const (
	baseUrl = "baseUrl"
)

// create a http/client once only, and reuse the TCP connection for all the functions
func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func GetClientToken(client *http.Client, grpcNonce string) (string, error) {
	
	url := baseUrl + "/api/auth/GetClientToken"

    request, err := http.NewRequest(http.MethodGet, url, nil)

	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(".env loading failure")
	}

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

func InsertWallet(client *http.Client, accessToken string, customerId string) (WalletArray, error) {

	url := baseUrl + "/api/wallets"

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
		return WalletArray{WalletType{"", "", "", ""}, ErrorType{"", "Failed to get response from InsertWallet", ""}}, 
				errors.New("Failed to get response from InsertWallet")
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	// Unmarshal the body into json
	var wallets WalletArray
	err = json.Unmarshal(body, &wallets)
	if err != nil {
		return WalletArray{WalletType{"", "", "", ""}, ErrorType{"", "Failed to parse Wallet into json", ""}}, 
				errors.New("Failed to parse Wallet into json")
	} else if wallets.Error.Code == "15" {
		return WalletArray{WalletType{"", "", "", ""}, ErrorType{"", "Wallet limit reached", ""}},
				errors.New("Wallet limit reached")
	}
	fmt.Printf("%+v\n", wallets)

	return wallets, nil
}

func TestPrintToken(t *testing.T) {

	client := httpClient()
	// grpcNonce has to be changed every time before te existing one expires
	accessToken, _ := GetClientToken(client, "123")

	// walletId = 200-1200 are used for this test
	walletArray := make([]WalletArray, 0)	
	for customer := 200; customer < 1200; customer++ {
		walletAddress, err := InsertWallet(client, accessToken, strconv.Itoa(customer))
		if err != nil {
			fmt.Println(err)
		}
		walletArray = append(walletArray, walletAddress)
	}
}