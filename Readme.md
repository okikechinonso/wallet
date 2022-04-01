
# Wallet
This demonstrate how wallet a works. The following technologies were used during the implementation of this project:
MySQL,
Redis,
Gin.

Run the command to download all dependencies.
```
go mod tidy
```

Before you run the code be sure to you add a ``.env`` file  in the root directory of project that contains the following:
```
DB_PORT=<your db port>
DB_USER=<your db username>
DB_PASS=<your db password>
PORT=<your server port>
```

Run the application using:
```
modd 
```
or
```bigquery
go run ./...
```

####After running the code, the database is populated automatically with the default wallet address being 1

### CreditWallet(Method:POST)

```
localhost:<PORT>/api/v1/{:wallet_id}/credit
```
Payload
```
{
"amount": <enter amount(intergers or decimals)> 
}
```

### DebitWallet(Method:POST)
Endpoint
```
localhost:<PORT>/api/v1/{:wallet_address}/debit
```
payload
```
{
"amount": <enter amount(intergers or decimals)> 
}
```
### GetBalance(Method:GET)
Endpoint
```
localhost:<PORT>/api/v1/{:wallet_address}/balance
```


## Tests
Testing is done using the GoMock framework. The ``gomock`` package and the ``mockgen``code generation tool are used for this purpose.
If you installed the dependencies using the command given above, then the packages would have been installed. Otherwise, installation can be done using the following commands:
```
go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen
mockgen -source=internal/ports/wallet.go -destination=mock/mock_db.go -package=mock
```
run all the test files using:
```bigquery
go test -v ./...
```