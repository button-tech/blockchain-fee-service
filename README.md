# Blockchain Fee Service
- [About](#about)
    - [Supported currencies](#supported-currencies) 
- [Installation](#installation)
- [Run](#run)
- [Run Docker container](#run-docker-container)
- [Docker build locally](#docker-build-locally)
- [Usage](#usage)
  - [Bitcoin](#bitcoin)
  - [BitcoinCash](#bitcoincash)
  - [Litecoin](#litecoin)
  - [Ethereum](#ethereum)
  - [EthereumClassic](#ethereumClassic)
  - [Token](#token)
  - [Waves](#waves)
  - [Stellar](#stellar)
## About
Rest api for:
1. Calculating actual fee
2. Getting maximal sending amount
3. Getting maximal sending amount with optimal fee 
4. Getting balance

**Supported currencies**
- Bitcoin
- BitcoinCash
- Litecoin
- Ethereum
- EthereumClassic
- Waves
- Stellar
- Ethereum ERC20 tokens

## Getting Started
### Installation

```
go get github.com/button-tech/blockchain-fee-service
cd  ~/go/src/github.com/button-tech/blockchain-fee-service
```
OR
```
git clone https://github.com/button-tech/blockchain-fee-service.git
cd blockchain-fee-service
GO111MODULE=on go mod tidy
```
### Run
```
go run main.go
```
## Run Docker container
```
docker run -d -p 8080:8080 buttonwallet/blockchain-fee-service
```
## Docker build locally
```
chmod +x build.sh
acc=nickname ./build.sh

```
OR
```
sudo docker build -t blockchain-fee-service
```
## Usage `/fee/{cryptocurrency}`
#### Bitcoin `POST /fee/bitcoin`
**Request body**
```
{
	"fromAddress": "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2",
	"amount": "0.34007244"
}
```
**Response**
```
{
    "fee": 226728,
    "balance": 65911084,
    "maxAmountWithOptimalFee": 34007244,
    "maxAmount": 65413013,
    "isEnough": true,
    "isBadFee": false,
    "input": 16,
    "output": 1
}
```

#### BitcoinCash`POST /fee/bitcoinCash`
**Request body**
```
{
	"fromAddress": "bitcoincash:qqhsxzztfhntg8srupvxu2wtuxh3rykdrchfuc6ekj",
	"amount": "0.002583021"
}
```
**Response**
```
{
    "fee": 678,
    "balance": 2584101,
    "maxAmountWithOptimalFee": 2577738,
    "maxAmount": 2583021,
    "isEnough": true,
    "isBadFee": false,
    "input": 1,
    "output": 2
}
```

#### Litecoin `POST /fee/litecoin`
**Request body**
```
{
	"fromAddress": "LZnXf5KUQTaPFZe2Bb3YW3Li1kWdL4s6gX",
	"amount": "0.002583021"
}
```
**Response**
```
{
    "fee": 2034,
    "balance": 4825099,
    "maxAmountWithOptimalFee": 4779731,
    "maxAmount": 4808203,
    "isEnough": true,
    "isBadFee": false,
    "input": 1,
    "output": 2
}
```

#### Ethereum `POST /fee/ethereum`
**Request body**
```
{
	"fromAddress": "0x1E1b3f95c992901f9ba898140E46965DF67a1F2a",
	"amount": "0.002583021"
}
```
**Response**
```
{
    "fee": 378000000000000,
    "balance": 563155068262885000,
    "maxAmountWithOptimalFee": 562777068262885000,
    "maxAmount": 562840068262885000,
    "isEnough": true,
    "isBadFee": false,
    "gasPrice": 18000000000,
    "gas": 21000
}
```

#### EthereumClassic `POST /fee/ethereumClassic`
**Request body**
```
{
	"fromAddress": "0x9EAb4b0fC468A7f5D46228bf5a76cB52370d068D",
	"amount": "0.002583021"
}
```
**Response**
```
{
    "fee": 504000000000000,
    "balance": 15967035509449030793,
    "maxAmountWithOptimalFee": 15966531509449030793,
    "maxAmount": 15966615509449030793,
    "isEnough": true,
    "isBadFee": false,
    "gasPrice": 24000000000,
    "gas": 21000
}
```

#### Token `POST /fee/token`
**Request body**
```
{
	"fromAddress": "0x9ae49c0d7f8f9ef4b864e004fe86ac8294e20950",
	"tokenAddress": "0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359",
	"amount": "0.12928069852065452"
}
```
**Response**
```
{
    "fee": 378000000000000,
    "balance": 289655836120696853,
    "maxAmountWithOptimalFee": 17510668875340780548,
    "maxAmount": 0,
    "isEnough": true,
    "isBadFee": false,
    "gasPrice": 18000000000,
    "gas": 36914,
    "tokenBalance": 17510668875340780548
}
```

#### Waves `POST /fee/waves`
**Request body**
```
{
	"fromAddress": "3P9cCF2czAVDyfAxCEVc1Y3WHRLA57xZaua",
	"amount": "0.11773"
}
```
**Response**
```
{
    "fee": 300000,
    "balance": 11773300000,
    "maxAmountWithOptimalFee": 11773000000,
    "maxAmount": 11773000000,
    "isEnough": true,
    "isBadFee": false
}
```

#### Stellar `POST /fee/stellar`
**Request body**
```
{
	"fromAddress": "GC7ONLMFZIWLGKWC5THMOCLQER7E63UHOTFJ3RFS6JYLMHIU32CFWGPV",
	"amount": "23.23"
}
```
**Response**
```
{
    "fee": 100,
    "balance": 13499000000,
    "maxAmountWithOptimalFee": 13398999900,
    "maxAmount": 13398999900,
    "isEnough": true,
    "isBadFee": false
}
```


