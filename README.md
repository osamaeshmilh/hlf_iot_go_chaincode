
# MedsData Chaincode

The `MedsData` chaincode represents a smart contract for managing medication data within a Hyperledger Fabric network. This chaincode includes functions to create, read, and query medication data, storing information such as batch numbers, warehouse numbers, IoT identifiers, temperature and humidity sensor IDs, timestamps, temperature readings, and humidity readings.

## Data Structure

The `MedsData` structure represents an individual record, with the following fields:

- `UniqueKey`: A string that uniquely identifies the record.
- `BatchNo`: Batch number associated with the medication.
- `WarehouseNo`: Warehouse number where the medication is stored.
- `IotId`: IoT device identifier associated with the record.
- `TemperatureSensorId`: Temperature sensor identifier.
- `HumiditySensorId`: Humidity sensor identifier.
- `Timestamp`: Timestamp of the record.
- `Temperature`: Temperature reading.
- `Humidity`: Humidity reading.

## Functions

### CreateMedsData

This function allows you to create a new `MedsData` record. It takes in the batch number, warehouse number, IoT ID, temperature sensor ID, humidity sensor ID, timestamp, temperature, and humidity. The unique key is generated based on the batch number and timestamp, and the data is then marshaled to JSON and stored in the ledger using this unique key.

```go
func (s *SmartContract) CreateMedsData(ctx contractapi.TransactionContextInterface, batchNo string, warehouseNo string, iotId string, temperatureSensorId string, humiditySensorId string, timestamp string, temperature float64, humidity float64) error
```

### ReadMedsData

This function reads a `MedsData` record from the ledger using the provided ID. If found, it unmarshals the JSON data into a `MedsData` structure and returns it.

```go
func (s *SmartContract) ReadMedsData(ctx contractapi.TransactionContextInterface, id string) (*MedsData, error)
```

### GetAllMedsData

This function queries all `MedsData` records in the ledger, unmarshals each record into a `MedsData` structure, and returns a slice of pointers to the retrieved records.

```go
func (s *SmartContract) GetAllMedsData(ctx contractapi.TransactionContextInterface) ([]*MedsData, error)
```

## Deployment

You can deploy this chaincode to your Hyperledger Fabric network by following the standard chaincode deployment procedures, making sure to meet all prerequisites required by your specific network configuration.
