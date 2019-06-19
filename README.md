# Fee Service

`/api/services/fee/{cryptocurrency}`

`cryptocurrency` = {bitcoin, ethereum, litecoin, bitcoinCash, ethereumClassic, waves, stellar, token}

**Response: example UTXO Blockchain**
```
{
	"fromAddress": "35hK24tcLEWcgNA4JxpvbkNkoAcDGqQPsP",
	"amount": "1510.0012298742"
	"receiversCount": 1 \\ Optional. Default: 1
}
```

**Request example**
```
{
    "fee": 17176,
    "input": 1,
    "output": 2,
    "balance": 15100013195817,
    "maxAmount": 15100013142096,
    "maxAmountWithOptimalFee": 15100012253902,
    "isEnough": true,
    "isBadFee": false
}
```

**Response: example Ethereum based Blockchains**
```
{
	"fromAddress": "0x66666600e43c6d9e1a249d29d58639dedfcd9ade",
	"amount": "0.02323"
}
```

**Request example**
```
{
    "fee": 120960000000000,
    "gasPrice": 5760000000,
    "gas": 21000,
    "balance": 400992218860738302,
    "maxAmount": 400891418860738302,
    "maxAmountWithOptimalFee": 400871258860738302,
    "isEnough": true,
    "isBadFee": false
}
```