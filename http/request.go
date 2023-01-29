package main

import (
	"fmt"
	"testing"
	"log"
	"net/http"
	"time"
	"os"
	"encoding/json"
	"encoding/csv"
	// "errors"
	// "bytes"
	"io"
	"strconv"
	"math/big"
	"context"
	"crypto/ecdsa"

	tusdt "cnx/cryptopay/commons/poc/tusdt/api"

	"github.com/joho/godotenv"
	"github.com/ethereum/go-ethereum"
	// "github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

// Default struct used by usdt example
type ETHWallet struct {
	PublicKey	string
	PrivateKey 	string
}

func (w *ETHWallet) getPrivateKey() (*ecdsa.PrivateKey, error) {
	pk, err := crypto.HexToECDSA(w.PrivateKey)
	if err != nil {
		return nil, err
	}
	return pk, nil
}

func (w *ETHWallet) Account() common.Address {
	return common.HexToAddress(w.PublicKey)
}

const (
	grpcNonce			= "122"			// grpcNonce has to be changed every time before te existing one expires
	firstCustomerId		= 200
	lastCustomerId		= 1199 
	transferAmountUsdt	= 4000000
)

// create a http/client once only, and reuse the TCP connection for all the functions
func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func ethClient(rpcUrl string) *ethclient.Client {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetClientToken(client *http.Client, grpcNonce string) (string, error) {

	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf(".env loading failure")
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
		log.Fatalf(".env loading failure")
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

	return wallet, nil
}

func TestCreateDepositWallets(t *testing.T) {

	client := httpClient()
	accessToken, _ := GetClientToken(client, grpcNonce)

	// walletId = 200-1199 are used for this test
	walletArray := make([]Wallet, 0)
	for customer := firstCustomerId; customer < lastCustomerId; customer++ {
		walletAddress, err := GenerateDepositWallet(client, accessToken, strconv.Itoa(customer))
		if err != nil {
			t.Fatal(err)
		}
		walletArray = append(walletArray, walletAddress)
	}
}

func GetDepositWallets(client *http.Client) ([]string) {

	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf(".env loading failure")
	}

	accessToken, _ := GetClientToken(client, grpcNonce)

	// walletId = [firstCustomerId, lastCustomerId) are used for this test
	walletArray := make([]string, 0)
	for customer := firstCustomerId; customer < lastCustomerId; customer++ {
		url := os.Getenv("BASE_URL") + "/api/wallets/" + strconv.Itoa(customer)

		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			fmt.Printf("Failed to create wallet address request for Customer Id: %v", customer)
		}
		request.Header.Set("Authorization", "bearer " + accessToken)
		request.Header.Set("Accept-Encoding", "gzip, deflate, br")
		request.Header.Set("Accept", "*/*")
		request.Header.Set("Connection", "keep-alive")
		request.Header.Set("Content-Type", "application/json")

		// make request to insert wallet for customerId
		response, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("Failed to get response from InsertWallet")
		}

		// Unmarshal the body into json
		var wallets Wallets
		err = json.Unmarshal(body, &wallets)
		walletArray = append(walletArray, wallets.Wallets.Wallet[0].WalletAddress)
	}

	return walletArray
}

func TransferUsdt(t *testing.T, client *ethclient.Client, usdtAddress common.Address, from *ETHWallet, to *ETHWallet, amount *big.Int) *types.Transaction {
	privateKey, _ := from.getPrivateKey()

	// suggest unit gas price
	unitGasPrice, err := client.SuggestGasPrice(context.Background()) // usually in gwei (1 eth = 10^9 gwei)
	if err != nil {
		t.Fatal(err)
	}

	// estimate gas
	estGas, err := client.EstimateGas(context.Background(),
		ethereum.CallMsg{
			From:     from.Account(),
			GasPrice: unitGasPrice,
			Value:    amount,
		})
	if err != nil {
		t.Errorf("Estimated Gas Needed %v", err)
	} else {
		t.Logf("estGas is %d", estGas)
	}

	nonce, err := client.PendingNonceAt(context.Background(), from.Account())
	if err != nil {
		t.Error(err)
	}

	// obtain the chainId from the NetworkID
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		t.Errorf("Error get chainid %v", err)
	}
	log.Printf("ChainId :%d", chainId)

	//Get signed transactor
	trx, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		t.Errorf("Error creating transactor %v", err)
	}

	trx.Nonce = big.NewInt(int64(nonce))
	trx.Value = big.NewInt(0) // we are not transfering the chain's native currency. Thus value is 0
	trx.GasPrice = unitGasPrice
	trx.GasLimit = estGas * 2 // too low will not get processed
	trx.Context = context.Background()

	// check usdt
	usdt, err := tusdt.NewTusdt(usdtAddress, client)
	if err != nil {
		t.Errorf("Error Accessing USDT Contract %v", err)
	}

	t.Logf("Transfer USDT from %s to %s trxValue %d, gasLimit %d, GasPrice %d, tokenValue %d", trx.From, to.Account(), trx.Value, trx.GasLimit, trx.GasPrice, amount)
	tx, err := usdt.Transfer(trx, to.Account(), amount)
	if err != nil {
		t.Errorf("Transfer USDT from %s to %s %v", trx.From, to.Account(), err)
	} else {
		t.Logf("Transfer transaction created %s %v\n From %s to %s", tx.Hash().Hex(), tx, from.Account().Hex(), to.Account().Hex())
	}

	return tx
}

//Current method is via ABI binding, eth book uses raw method
func TestTransferUsdt(t *testing.T) {

	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf(".env loading failure")
	}

	usdtToken := os.Getenv("USDT_TOKEN")
	infuraUrl := os.Getenv("INFURA_URL")
	sourceWallet := os.Getenv("SOURCE_WALLET")
	sourcePrivateKey := os.Getenv("SOURCE_PRIVATE_KEY")

	ethClient := ethClient(infuraUrl)
	httpClient := httpClient()

	var walletAddress []string
	walletAddress = GetDepositWallets(httpClient)

	transferAmountUsdt := big.NewInt(transferAmountUsdt) //amount in USDT, converted to token number

	senderWallet := &ETHWallet{
		PublicKey:  sourceWallet,
		PrivateKey: sourcePrivateKey,
	}

	// prepare the two-dimensioanal slice for the wallet transactions
	records := [][]string{}
	// open the file
	file, err := os.Create("transactions.csv")
    defer file.Close()
	if err != nil {
        log.Fatal("failed to open file", err)
    }

	for index, wallet := range walletAddress {

		receiverWallet := &ETHWallet{
			PublicKey:  wallet,
			PrivateKey: "",
		}

		tx := TransferUsdt(t, ethClient, common.HexToAddress(usdtToken), senderWallet, receiverWallet, transferAmountUsdt)

		t.Logf("tx sent: %s", tx.Hash().Hex())
		
		// append record for the transactions
		record := []string{}
		record = append(record, strconv.Itoa(index + firstCustomerId), wallet, tx.Hash().Hex())
		records = append(records, record)
	}

	// write the transaction records into the files
    writer := csv.NewWriter(file)
    err = writer.WriteAll(records)
    if err != nil {
        log.Fatal(err)
    }

}