package goBlocUtilities

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/0sax/err2"
	"github.com/thoas/go-funk"
	"net/http"
)

const (
	BaseUrl                               = "https://sandbox.buildwithcore.com/v2.0"
	AuthEndpoint                          = "/authentication/token"
	ElectricityOperatorsEndpoint          = "/electricity/operators"
	OneElectricityOperatorEndpoint        = "/electricity/operators/"
	ElectricityProductsEndpoint           = "/electricity/products"
	ElectricityPaymentFulfillmentEndpoint = "/electricity/payments"

	TelecomsOperatorsEndpoint          = "/telecommunications/operators"
	OneTelecomsOperatorEndpoint        = "/telecommunications/operators/"
	TelecomsPaymentFulfillmentEndpoint = "/telecommunications/payments"

	//Categories
	PostPaid   = "postpaid"
	Prepaid    = "prepaid"
	SmartMeter = "smart-meter"
)

func categories() []string {
	return []string{PostPaid, Prepaid, SmartMeter}
}

type Client struct {
	baseUrl     string
	bearerToken string
}

func Authenticate(username, password, baseUrl string) (c *Client, err error) {

	c = &Client{
		baseUrl:     baseUrl,
		bearerToken: "",
	}

	b64 := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	headers := map[string]string{
		"Authorization": "Basic " + b64,
	}

	var resp Response

	err = makeRequest(http.MethodGet, baseUrl+AuthEndpoint, nil, headers, &resp)
	if err != nil {
		c = nil
		return
	}

	if resp.isSuccess() {
		lr, er := resp.loginResponse()
		if er != nil {
			err = er
			return
		}
		c.bearerToken = lr.Token
	} else {
		err = err2.NewClientErr(errors.New(resp.Error.Message), resp.Error.Message, resp.Error.Code)
	}

	return

}

//ELECTRICITY

func (c *Client) GetElectricityOperators() (eops []Operator, err error) {

	var resp Response
	err = c.standardRequest(http.MethodGet, ElectricityOperatorsEndpoint, nil, &resp)
	if err != nil {
		return
	}

	if resp.isSuccess() {
		eops, err = resp.operators()
		return
	} else {
		re := resp.Error
		err = err2.NewClientErr(errors.New(re.Message), re.Message, re.Code)
		return
	}
}

func (c *Client) GetOneElectricityOperator(id string) (eo *Operator, err error) {

	var resp Response
	err = c.standardRequest(http.MethodGet, OneElectricityOperatorEndpoint+id, nil, &resp)
	if err != nil {
		return
	}

	if resp.isSuccess() {
		eo, err = resp.operator()
		return
	} else {
		re := resp.Error
		err = err2.NewClientErr(errors.New(re.Message), re.Message, re.Code)
		return
	}
}

func (c *Client) GetElectricityProducts() (eops []Product, err error) {

	var resp Response
	err = c.standardRequest(http.MethodGet, ElectricityProductsEndpoint, nil, &resp)
	if err != nil {
		return
	}

	if resp.isSuccess() {
		eops, err = resp.products()
		return
	} else {
		re := resp.Error
		err = err2.NewClientErr(errors.New(re.Message), re.Message, re.Code)
		return
	}
}

func (c *Client) CustomerValidation(operatorID, customerID, category string) (
	cv *CustomerValidation, err error) {

	if !funk.ContainsString(categories(), category) {
		ss := fmt.Sprintf("unknown category: '%v'", category)
		err = err2.NewClientErr(errors.New(ss), ss, 400)
		return
	}

	var resp Response
	url := fmt.Sprintf("%v/%v/customers/%v?category=%v",
		OneElectricityOperatorEndpoint, operatorID, customerID, category)
	err = c.standardRequest(http.MethodGet, url, nil, &resp)
	if err != nil {
		return
	}

	if resp.isSuccess() {
		cv, err = resp.customerValidation()
		return
	} else {
		re := resp.Error
		err = err2.NewClientErr(errors.New(re.Message), re.Message, re.Code)
		return
	}
}

func (c *Client) ElectricityPaymentFulfilment(request *PaymentRequest) (
	pfr *PaymentFulfilmentResponse, err error) {

	var resp Response

	err = c.standardRequest(http.MethodPost, ElectricityPaymentFulfillmentEndpoint, request, &resp)
	if err != nil {
		return
	}

	if resp.isSuccess() {
		pfr, err = resp.paymentFulfillment()
		return
	} else {
		re := resp.Error
		err = err2.NewClientErr(errors.New(re.Message), re.Message, re.Code)
		return
	}
}

func (c *Client) ElectricityPaymentsReQuery(ref string) (
	pfr *PaymentFulfilmentResponse, err error) {

	var resp Response
	url := fmt.Sprintf("%v/%v", ElectricityPaymentFulfillmentEndpoint, ref)
	err = c.standardRequest(http.MethodGet, url, nil, &resp)
	if err != nil {
		return
	}

	if resp.isSuccess() {
		pfr, err = resp.paymentFulfillment()
		return
	} else {
		re := resp.Error
		err = err2.NewClientErr(errors.New(re.Message), re.Message, re.Code)
		return
	}
}

//TELECOMS

func (c *Client) GetTelecomsOperators() (tops []Operator, err error) {

	var resp Response
	err = c.standardRequest(http.MethodGet, TelecomsOperatorsEndpoint, nil, &resp)
	if err != nil {
		return
	}

	if resp.isSuccess() {
		tops, err = resp.operators()
		return
	} else {
		re := resp.Error
		err = err2.NewClientErr(errors.New(re.Message), re.Message, re.Code)
		return
	}
}

func (c *Client) GetOneTelecomsOperatorsProducts(id string) (topprods []Product, err error) {

	var resp Response
	err = c.standardRequest(http.MethodGet, OneTelecomsOperatorEndpoint+id+"/products", nil, &resp)
	if err != nil {
		return
	}

	if resp.isSuccess() {
		topprods, err = resp.products()
		return
	} else {
		re := resp.Error
		err = err2.NewClientErr(errors.New(re.Message), re.Message, re.Code)
		return
	}
}

func (c *Client) TelecomsPaymentFulfilment(request *PaymentRequest) (
	pfr *PaymentFulfilmentResponse, err error) {

	var resp Response

	err = c.standardRequest(http.MethodPost, TelecomsPaymentFulfillmentEndpoint, request, &resp)
	if err != nil {
		return
	}

	if resp.isSuccess() {
		pfr, err = resp.paymentFulfillment()
		return
	} else {
		re := resp.Error
		err = err2.NewClientErr(errors.New(re.Message), re.Message, re.Code)
		return
	}
}


func (c *Client) TelecomsPaymentsReQuery(ref string) (
	pfr *PaymentFulfilmentResponse, err error) {

	var resp Response
	url := fmt.Sprintf("%v/%v", TelecomsPaymentFulfillmentEndpoint, ref)
	err = c.standardRequest(http.MethodGet, url, nil, &resp)
	if err != nil {
		return
	}

	if resp.isSuccess() {
		pfr, err = resp.paymentFulfillment()
		return
	} else {
		re := resp.Error
		err = err2.NewClientErr(errors.New(re.Message), re.Message, re.Code)
		return
	}
}

func (c *Client) TelecomsPaymentsReQueryByClientRef(clientRef string) (
	pfr *PaymentFulfilmentResponse, err error) {

	var resp Response
	url := fmt.Sprintf("%v?client_reference=%v", TelecomsPaymentFulfillmentEndpoint, clientRef)
	err = c.standardRequest(http.MethodGet, url, nil, &resp)
	if err != nil {
		return
	}

	if resp.isSuccess() {
		pfr, err = resp.paymentFulfillment()
		return
	} else {
		re := resp.Error
		err = err2.NewClientErr(errors.New(re.Message), re.Message, re.Code)
		return
	}
}