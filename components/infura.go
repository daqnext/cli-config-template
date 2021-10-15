package components

import (
	"context"
	"errors"
	"math/rand"

	fj "github.com/daqnext/fastjson"

	"github.com/ethereum/go-ethereum/ethclient"
)

type InfuraClient struct {
	ethClients []*ethclient.Client
}

/*
infura1
infura2
*/
func InitInfura(ConfigJson *fj.FastJson) (*InfuraClient, error) {

	var InfuraClients []*ethclient.Client

	////infura ini/////
	infura1, err := ConfigJson.GetString("infura1")
	if err != nil {
		return nil, errors.New("infura1 [string] in config.json not defined," + err.Error())
	}
	client1, err := ethclient.Dial(infura1)
	if err != nil {
		return nil, err
	}
	_, err1 := client1.HeaderByNumber(context.Background(), nil)
	if err1 != nil {
		return nil, err1
	}

	InfuraClients = append(InfuraClients, client1)

	infura2, err := ConfigJson.GetString("infura2")
	if err != nil {
		return nil, errors.New("infura2 [string] in config.json not defined," + err.Error())
	}
	client2, err := ethclient.Dial(infura2)
	if err != nil {
		return nil, err
	}
	_, err2 := client2.HeaderByNumber(context.Background(), nil)
	if err2 != nil {
		return nil, err2
	}

	InfuraClients = append(InfuraClients, client2)

	return &InfuraClient{InfuraClients}, nil
}

func (ic *InfuraClient) GetEthClient() *ethclient.Client {
	return ic.ethClients[rand.Intn(2)]
}
