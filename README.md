# MobilePulsa API Go Client

[![Coverage Status](https://coveralls.io/repos/github/pandudpn/mobile-pulsa-go/badge.png?branch=master)](https://coveralls.io/github/pandudpn/mobile-pulsa-go?branch=master)

This is Unofficial Project for Prepaid or Postpaid Product in Indonesia. This is created by [pandudpn](https://www.github.com/pandudpn).

- [Documentation](#documentation)
- [Installation](#installation)
- [Usage](#usage)

## Documentation

For the API Documentation, check [IAK Documentation](https://api.iak.id/docs/reference)

## Installation

Install with : 
```sh
go get -u github.com/pandudpn/mobile-pulsa-go
```

Then, import it using:

```go
import "github.com/pandudpn/mobile-pulsa-go"
```

## Usage

```go
package main

import (
    "log"

    "github.com/pandudpn/mobile-pulsa-go"
    "github.com/pandudpn/mobile-pulsa-go/topup"
)

func main() {
    apiKey := "api-key"
    userName := "username"
    appEnv := "development"

    opts := mobilepulsa.NewOption()
    opts.SetAPIKey(apiKey)
    opts.SetUsername(username)
    // set the environment
    if appEnv == "development" {
        opts.SetDevelopment()
    } else {
        opts.SetProduction()
    }

    // for an example, you want to buy pulse `Telkomsel` with nominal 1.000
    productCode := "htelkomsel1000"

    data := &topup.TopUpParam{
        RefID: "uuid-123",
        ProductCode: productCode,
        CustomerID: "target_phone_number",
    }

    topup, err := topup.CreatePayment(data, opts)
    if err != nil {
        // do something here
        log.Println("error create payment prepaid product", err)
        return
    }

    log.Println(topup.Message)
}
```

