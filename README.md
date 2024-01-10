# price-feeder-inspector
Script inspecting validator votes and missing currencies

## Vote Inspector Python Script 
This Python script fetches live submitted aggregated votes from validators and cross-checks them with oracle-accepted denominations. It displays the missing denominations of votes.

### How to run ? 
```python 
## For Canon 
$ python3 votes_inspector.py --network canon

## For mainnet 
$ python3 votes_inspector.py --network mainnet
## By Default It will runs for mainnet
$ python3 votes_inspector.py
```


## [Historical Oracle Votes Inspector](./inspector_go/README.md)

### Build 
```bash
$ cd inspector_go 
// build the binary
$ make build
```

### How to run ? 


```bash
$ ./build/price_inspector --config <CONFIG.yaml> --network canon
```