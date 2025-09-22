# pismo-transaction

Simple implementation of Pismo Transaction endpoints:
- `POST /accounts` - create accout
- `GET /accounts/{accountid}` - get account
- `POST /transactions`- create transaction

Business rule:
-Operation types:
1.CASH PURCHASE(negative amount)
2.INSTALLMENT PURCHASE(negative amount)
3.WITHDRAWAL (negative amount)
4.PAYMENT (positive amount)
-Transaction `EventDate` is set by server

##Requirements
-Go 1.20+ (module-aware)
-Docker (optional)

## Run locally

```bash
# build
go build -o pismo-test

#run
./pismo-test
#server listens on :8080