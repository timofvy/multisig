package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/viper"
	"github.com/timofvy/multisig/abi/safe_abi"
	"github.com/timofvy/multisig/abi/safe_proxy_factory_abi"
)

const zeroAddress = "0x0000000000000000000000000000000000000000"

func LoadConfig() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.ReadInConfig() //nolint:errcheck
}

func getProvider() (*ethclient.Client, error) {
	rpcClient, err := rpc.DialOptions(
		context.Background(),
		viper.GetString("rpc_url"),
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

func main() {
	LoadConfig()

	owners := flag.String("owners", "", "Owners")
	threshold := flag.Int("threshold", 0, "Threshold")
	flag.Parse()

	signerAddrs := strings.Split(*owners, ",")

	var ownersAddrSlice []common.Address
	for _, addr := range signerAddrs {
		trimmedAddr := strings.TrimSpace(addr)
		ownerAddr := common.HexToAddress(trimmedAddr)

		ownersAddrSlice = append(ownersAddrSlice, ownerAddr)
	}

	err := sendDeployMultisig(ownersAddrSlice, *threshold)
	if err != nil {
		log.Println(err)
		return
	}
}

func sendDeployMultisig(
	owners []common.Address,
	threshold int,
) error {
	if threshold == 0 {
		return errors.New("threshold must be greater than 0")
	}

	if len(owners) == 0 {
		return errors.New("owners must be greater than 0")
	}

	priv := viper.GetString("private_key")
	safeProxyFactoryAddress := common.HexToAddress(viper.GetString("safe_proxy_factory"))
	safeAddress := common.HexToAddress(viper.GetString("safe"))

	contractAbi, err := abi.JSON(strings.NewReader(safe_abi.SafeAbiABI))
	if err != nil {
		return err
	}

	data, err := contractAbi.Pack("setup",
		owners,
		big.NewInt(int64(threshold)),
		common.HexToAddress(zeroAddress),
		[]byte{0},
		common.HexToAddress(zeroAddress),
		common.HexToAddress(zeroAddress),
		big.NewInt(0),
		common.HexToAddress(zeroAddress),
	)
	if err != nil {
		return err
	}

	provider, err := getProvider()
	if err != nil {
		return err
	}

	contractTransactor, err := safe_proxy_factory_abi.NewSafeProxyFactoryAbiTransactor(
		safeProxyFactoryAddress,
		provider,
	)
	if err != nil {
		return err
	}

	chainID, err := provider.ChainID(context.Background())
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(priv)
	if err != nil {
		return err
	}

	trOpts, err := bind.NewKeyedTransactorWithChainID(
		privateKey,
		chainID,
	)
	if err != nil {
		return err
	}

	transaction, err := contractTransactor.CreateProxyWithNonce(
		trOpts,
		safeAddress,
		data,
		big.NewInt(0),
	)
	if err != nil {
		return err
	}

	log.Println("Transaction sent: ", transaction.Hash().Hex())

	return nil
}
