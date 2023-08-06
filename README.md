# Medication IoT Data Chaincode

This chaincode is designed for Hyperledger Fabric and serves as a data registry for IoT sensors that track medication data. It receives data from various IoT sensors, including temperature and humidity readings, and stores them within the blockchain.

## Data Structure

The chaincode uses the following data structure to represent the medication data:

```go
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
