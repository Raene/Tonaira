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

 it returns data in the JSON form

```json
    {
    "data": {
        "address": string
},
    "success": true
}
```
  