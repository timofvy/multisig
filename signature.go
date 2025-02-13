package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"flag"
	"log"
	"math/big"
	"net/http"
	"time"

	"main/abi/safe_abi"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/viper"
)

type MultisigTransactionRequest struct {
	To             common.Address `json:"to"`
	Value          *big.Int       `json:"value"`
	Data           []byte         `json:"data"`
	Operation      uint8          `json:"operation"`
	SafeTxGas      *big.Int       `json:"safeTxGas"`
	BaseGas        *big.Int       `json:"baseGas"`
	GasPrice       *big.Int       `json:"gasPrice"`
	GasToken       common.Address `json:"gasToken"`
	RefundReceiver common.Address `json:"refundReceiver"`
	Signatures     []byte         `json:"signatures"`
}

func LoadConfig() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.ReadInConfig() //nolint:errcheck
}

func main() {
	LoadConfig()

	txData := flag.String("tx-data", "", "Tx data")
	flag.Parse()

	err := Signature(*txData)
	if err != nil {
		log.Println(err)
		return
	}
}

func getProvider() (*ethclient.Client, error) {
	rpcClient, err := rpc.DialOptions(
		context.Background(),
		viper.GetString("rpcUrl"),
		rpc.WithHTTPClient(&http.Client{ //nolint:exhaustruct
			Timeout: 15 * time.Second,
		}),
	)
	if err != nil {
		return &ethclient.Client{}, err
	}

	connection := ethclient.NewClient(rpcClient)

	return connection, nil
}

func Signature(txData string) error {
	priv := viper.GetString("PRIV_KEY_SERVICE_BACKEND")
	multisigAddress := common.HexToAddress(viper.GetString("multisigAddress"))
	targetAddress := common.HexToAddress(viper.GetString("targetAddress"))

	provider, err := getProvider()
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(priv)
	if err != nil {
		return err
	}

	var data []byte

	if txData == "" {
		data = []byte{}
	} else {
		data, err = hex.DecodeString(txData[2:]) // remove the "0x" prefix
		if err != nil {
			return err
		}
	}

	contractCaller, err := safe_abi.NewSafeAbiCaller(
		multisigAddress,
		provider,
	)
	if err != nil {
		log.Println("1111", contractCaller)
		return err
	}

	log.Println("2222", err)

	callOpts := &bind.CallOpts{ //nolint:exhaustruct
		Pending: true,
	}

	// Получаем текущий nonce на мультисиге
	proxyNonce, err := contractCaller.Nonce(
		callOpts,
	)
	if err != nil {
		return err
	}

	safeTx := MultisigTransactionRequest{ //nolint:exhaustruct
		To:             targetAddress, // send to contract, example "Verifier"
		Value:          big.NewInt(0),
		Data:           data, // data with the parameters
		Operation:      0,
		SafeTxGas:      big.NewInt(0),
		BaseGas:        big.NewInt(0),
		GasPrice:       big.NewInt(0),
		GasToken:       common.Address{},
		RefundReceiver: common.Address{},
	}

	// Преобразуем данные
	rawSafeTx, err := contractCaller.EncodeTransactionData(
		callOpts,
		safeTx.To,
		safeTx.Value,
		safeTx.Data,
		safeTx.Operation,
		safeTx.SafeTxGas,
		safeTx.BaseGas,
		safeTx.GasPrice,
		safeTx.GasToken,
		safeTx.RefundReceiver,
		proxyNonce,
	)
	if err != nil {
		return err
	}

	// Преобразуем данные
	safeTxHashBytes := crypto.Keccak256(rawSafeTx)

	// Подписываем данные приватным ключом бэкенда
	safeTxSignature, err := signHash(safeTxHashBytes, privateKey)
	if err != nil {
		return err
	}

	log.Println("Signature: 0x" + hex.EncodeToString(safeTxSignature))

	return nil
}

func signHash(hash []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	sig, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, err
	}

	if sig[64] < 2 { //nolint:mnd
		sig[64] += 27 // Transform V from 0/1 to 27/28 according to the yellow paper
	}

	return sig, nil
}
