package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type MedsData struct {
	BatchNo             string    `json:"batchNo"`
	WarehouseNo         string    `json:"warehouseNo"`
	IotId               string    `json:"iotId"`
	TemperatureSensorId string    `json:"temperatureSensorId"`
	HumiditySensorId    string    `json:"humiditySensorId"`
	Timestamp           string `json:"timestamp"`
	Temperature         float64   `json:"temperature"`
	Humidity            float64   `json:"humidity"`
}

func (s *SmartContract) CreateMedsData(ctx contractapi.TransactionContextInterface, batchNo string, warehouseNo string, iotId string, temperatureSensorId string, humiditySensorId string, timestamp string, temperature float64, humidity float64) error {
	data := MedsData{
		BatchNo:             batchNo,
		WarehouseNo:         warehouseNo,
		IotId:               iotId,
		TemperatureSensorId: temperatureSensorId,
		HumiditySensorId:    humiditySensorId,
		Timestamp:           timestamp,
		Temperature:         temperature,
		Humidity:            humidity,
	}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(batchNo, dataJSON)
}

func (s *SmartContract) ReadMedsData(ctx contractapi.TransactionContextInterface, id string) (*MedsData, error) {
	dataJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if dataJSON == nil {
		return nil, fmt.Errorf("the data for id %s does not exist", id)
	}

	var data MedsData
	err = json.Unmarshal(dataJSON, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *SmartContract) GetAllMedsData(ctx contractapi.TransactionContextInterface) ([]*MedsData, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var medsDataList []*MedsData
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var medsData MedsData
		err = json.Unmarshal(queryResponse.Value, &medsData)
		if err != nil {
			return nil, err
		}
		medsDataList = append(medsDataList, &medsData)
	}

	return medsDataList, nil
}
