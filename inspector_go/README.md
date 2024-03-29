# Oracle Price Votes Inspector


## Configuration 

> Note: if there is no validators list on configuration, it will inspect votes for all validators 

1. Add validator's info for checking missing oracle votes
2. Add a window parameter to specify the number of blocks to check from the latest height

### Ex: 
```yaml
---
validators:
- validator: umeevaloper1wusrw8nzgs6d8d47xqej5xqttvsu2geftfgu5f
  moniker: Stakewolle.com | Auto-compound
- validator: umeevaloper16kukyg8l6qse6n2rq24qr6qd3ylgnnzmut8v3u
  moniker: Coinpayu
window: 20
networks:
  canon:
    rpc: https://canon-4.rpc.network.umee.cc:443
    grpc: tcp://canon-4.network.umee.cc:11000
  mainnet:
    rpc: http://umee.rpc.m.stavr.tech:10457
    grpc: umee-grpc.polkachu.com:13690
  local:
    rpc: http://localhost:26657
    grpc: http://localhost:9090


``````

## How to run ? 

### Without build
```
// For the mainnet
$ go run main.go --network mainnet --config config.json

// For the canon
$ go run main.go --network canon --config config.json
```

## With build the binary 
```
// build the binary
$ go build -o ./build/ . 
```
Run the binary 
```
// For the mainnet
$ ./build/price_inspector --network mainnet --config config.json

// For the canon
$ ./build/price_inspector --network canon --config config.json
```
