# go-iap-galaxy

![](https://img.shields.io/badge/golang-1.19-blue.svg?style=flat)

go-iap-galaxy verifies the purchase receipt via GalaxyStore.


# Installation
```
go get github.com/coolishbee/go-iap-galaxy
```

# Quick Start

### In App Purchase via GalaxyStore

```go
import(
    "github.com/coolishbee/go-iap-galaxy"
)

func main() {
	client := galaxy.New()

	ctx := context.Background()
	resp, err := client.Verify(ctx, "purchaseId")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
```


# Support

This validator uses [Samsung IAP Server API](https://developer.samsung.com/iap/programming-guide/samsung-iap-server-api.html).