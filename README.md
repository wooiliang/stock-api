# Stock API &middot; [![Build Status](https://travis-ci.org/wooiliang/stock-api.svg?branch=master)](https://travis-ci.org/wooiliang/stock-api)

This is the API to get the stock price from KLSE and SGX markets. You are required to deploy this API service into AWS Lambda (serverless, FaaS) before using it.

## Installation

1. Install the [serverless](https://serverless.com/) toolkit.
2. Configure your AWS profile in CLI.
3. Run `make`
4. Deploy to AWS Lambda by running `sls deploy`

## Examples

##### Request

```
curl -X GET \
  https://xxxxxxxxxx.execute-api.ap-southeast-1.amazonaws.com/dev/stocks/sgx/stock_c6l-sia \
  -H 'Cache-Control: no-cache' \
  -H 'Postman-Token: ecb056ba-bfe0-4e3b-93e1-b305f3ae5245'
```

##### Response

```
{
    "price": 10.69
}
```

#### For KLSE

1. Visit [https://klse.i3investor.com](https://klse.i3investor.com).
2. Browse the stock that you are looking for.
3. Copy the ID from the URL. Eg. http://klse.i3investor.com/servlets/stk/7054.jsp copy the `7054`
4. Form the GET request URL like this `https://xxxxxxxxxx.execute-api.ap-southeast-1.amazonaws.com/dev/stocks/klse/7054`
5. Get the price in the response.

#### For SGX

1. Visit [https://sginvestors.io](https://sginvestors.io).
2. Browse the stock that you are looking for.
3. Copy the ID from the URL. Eg. https://sginvestors.io/sgx/stock/z74-singtel/stock-info copy the `stock/z74-singtel`
4. Form the GET request URL like this `https://xxxxxxxxxx.execute-api.ap-southeast-1.amazonaws.com/dev/stocks/sgx/stock_z74-singtel` (Replace the `/` with `_`)
5. Get the price in the response.

## Contributing

### License

Stock API is [MIT licensed](./LICENSE).
