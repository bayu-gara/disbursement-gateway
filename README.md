# Disbursement Gateway

Disbursement gateway has a feature to disburse e-wallet balance to physical bank account. This project was built using Golang 1.19

## Quick Start

### Run the project
Please make sure the port `8080` is not being used by other application
```
go mod tidy
make gorun
```

### Hit The API Using cURL
```
curl --location 'http://localhost:8080/disburse' \
--header 'Content-Type: application/json' \
--data '{
    "disbursement_data": [
        {
            "user_id": 1,
            "bank_code": "002",
            "bank_account_number": "019300020",
            "amount": 1000

        },
        {
            "user_id": 2,
            "bank_code": "014",
            "bank_account_number": "029300031",
            "amount": 1500

        },
        {
            "user_id": 3,
            "bank_code": "451",
            "bank_account_number": "218300020",
            "amount": 40000

        },
        {
            "user_id": 4,
            "bank_code": "451",
            "bank_account_number": "1112398476",
            "amount": 5000

        },
        {
            "user_id": 5,
            "bank_code": "014",
            "bank_account_number": "400018279",
            "amount": 8700

        },
        {
            "user_id": 6,
            "bank_code": "008",
            "bank_account_number": "66124633",
            "amount": 7000

        }
    ]
}'
```
### Hit The API Using Postman
**POST** `http:localhost:8080/disburse`
```
{
    "disbursement_data": [
        {
            "user_id": 1,
            "bank_code": "002",
            "bank_account_number": "019300020",
            "amount": 1000

        },
        {
            "user_id": 2,
            "bank_code": "014",
            "bank_account_number": "029300031",
            "amount": 1500

        },
        {
            "user_id": 3,
            "bank_code": "451",
            "bank_account_number": "218300020",
            "amount": 40000

        },
        {
            "user_id": 4,
            "bank_code": "451",
            "bank_account_number": "1112398476",
            "amount": 5000

        },
        {
            "user_id": 5,
            "bank_code": "014",
            "bank_account_number": "400018279",
            "amount": 8700

        },
        {
            "user_id": 6,
            "bank_code": "008",
            "bank_account_number": "66124633",
            "amount": 7000

        }
    ]
}
```

## Ideas
Starting from the main API, it receives a list of data that describes `from` and `to` where the money should be transferred. I actually think of another way where the client could send the csv file but for simplicity the list of data is in JSON format.

After the API receive the request, it will split the transaction to different worker. Since the disbursement transaction is processed multiple request at once, better we are not waiting each request until it is done. In this project i use goroutine for making asynchronous transaction but my first idea is using message queue/broker like [NSQ](https://nsq.io). Using NSQ will make sure the request/transaction being processed since it has retry mechanism. However, for simplicity of the environment setup i choose goroutine.

At last, after the transaction is done, we are able to see the status of the transaction (success/failed) through this API **GET** `http://localhost:8080/transaction?user_id=1`. It will list the transaction history of the specific user.

## Miscellaneous

### Transaction Status
| Name | Code |
| ----------- | ----------- |
| Success | 0 |
| Pending | 1 |
| Failed | 2 |
