package models

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	//"github.com/pkg/errors"
)

func (this *Application) AddHouseItem(args []string) (string, error) {
	//	tempDataMap := make(map[string][]byte)
	//	tempDataMap["result"] = []byte("Transient data in AddHouseItem")

	var tempArgs [][]byte
	for i := 1; i < len(args); i++ {
		tempArgs = append(tempArgs, []byte(args[i]))
	}

	request := channel.Request{ChaincodeID: this.FabricSetup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]), []byte(args[4]), []byte(args[5]), []byte(args[6]), []byte(args[7])}}
	response, err := this.FabricSetup.client.Execute(request)
	if err != nil {
		// 资产转移失败
		return "", err
	}

	return string(response.TransactionID), nil
}
