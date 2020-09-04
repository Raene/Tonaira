# Tonaira

This is the backed-api for Tonaira. This api has GoMod enabled, so please enable yours, it uses **go 1.15**. Below available routes are listed together with what they expect to receive and what data they give in return.

/api/v1/

- <http://localhost:3000/api/v1/coin-stats>:
  - this route supports a GET verb @ <http://localhost:3000/api/v1/coin-stats/>, it returns json in the form

```go
{
     "data": {
        "CoinValue": coinValue,
        "ExchangeRate": xChangeRate
     }
    "message": "success"
    }
```

- <http://localhost:3000/api/v1/conflux>:
  - this route supports a POST verb @ <http://localhost:3000/api/v1/conflux>, it expects a transaction JSON object in the form

```json
{
  "accountNumber": string,
 "bank":          string,
 "sender":        string(not required),
 "senderEmail":   string(not required),
 "amount":        int,
 "network":       string,
 }
```

 it returns data in the JSON form, conflux hasnt gone public yet as at the time of building this, so it has no exchangeRate yet.

```json
    {
    "data": {
        "address": string,
        "exchangeRate": float32
},
    "success": true
}
```

` <http://localhost:3000/api/v1/blockchain>:
- POST: <http://localhost:3000/api/v1/blockchain/> it expects a transaction object same as the conflux api. and also returns data in the same format as the conflux api
  