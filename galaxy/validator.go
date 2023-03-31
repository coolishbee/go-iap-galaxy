package galaxy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	ReceiptURL  string = "https://iap.samsungapps.com/iap/v6/receipt"
	ContentType string = "application/json"
)

type IAPResponse struct {
	// Unique identifier of the in-app item registered in Seller Portal
	ItemId string `json:"itemId"`

	// Unique identifier assigned to the in-app item payment transaction when it was successful
	PaymentId string `json:"paymentId"`

	// Unique identifier assigned to the purchase receipt
	OrderId string `json:"orderId"`

	// Package name of the app with a completed transaction
	PackageName string `json:"packageName"`

	// Title of the in-app item registered in Seller Portal
	ItemName string `json:"itemName"`

	// Brief explanation of the in-app item registered in Seller Portal
	ItemDesc string `json:"itemDesc"`

	// Date and time of the item purchase and payment transaction
	// ex: (YYYY-MM-DD HH:mm:ss GMT)
	PurchaseDate string `json:"purchaseDate"`

	// Total amount, including the in-app item price and all applicable taxes, billed to the user
	PaymentAmount string `json:"paymentAmount"`

	// Processing result of the request for the receipt
	// "success", "fail", "cancel"
	Status string `json:"status"`

	// Type of payment option used to purchase the item
	// "Credit Card", "Mobile Micro Purchase", "Prepaid Card", "PSMS", "Carrier Billing"
	PaymentMethod string `json:"paymentMethod"`

	// IAP operating mode in effect at the time of purchase
	// "TEST", "PRODUCTION"
	Mode string `json:"mode"`

	// For consumable in-app items only, whether or not the item
	// has been reported as consumed and is available for purchase again:
	// "Y". Consumed "N". Not Consumed
	ConsumeYN string `json:"consumeYN"`

	// Date and time the consumable item was reported as consumed
	// ex: (YYYY-MM-DD HH:mm:ss GMT)
	ConsumeDate string `json:"consumeDate"`

	// Model name of device that reported the item as consumed
	ConsumeDeviceModel string `json:"consumeDeviceModel"`

	// Transaction ID created by your app for security
	// Returned only if the pass-through parameter was set.
	PassThroughParam string `json:"passThroughParam"`

	// Currency code (3 characters) of the purchaser's local currency.
	// (for example, EUR, GBP, USD)
	CurrencyCode string `json:"currencyCode"`

	// Symbol of the purchaser's local currency
	// (for example, €, £, or $)
	CurrencyUnit string `json:"currencyUnit"`
}

type IAPError struct {
	Status       string `json:"status"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type Client struct {
	httpCli *http.Client
}

func New() *Client {
	client := &Client{
		httpCli: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	return client
}

func (c *Client) Verify(
	ctx context.Context,
	purchaseId string,
) (IAPResponse, error) {

	result := IAPResponse{}
	url := fmt.Sprintf("%s?purchaseID=%s", ReceiptURL, purchaseId)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return result, errors.New(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return result, errors.New(resp.Status)
	}

	respBody, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		fmt.Println(err)
	}

	if result.Status == "fail" {
		respError := IAPError{}
		err = json.Unmarshal(respBody, &respError)
		if err != nil {
			fmt.Println(err)
		}

		errMsg := fmt.Sprintf("status : %s errorCode : %d errorMessage : %s",
			respError.Status,
			respError.ErrorCode,
			respError.ErrorMessage)
		return result, errors.New(errMsg)
	}

	return result, nil
}
