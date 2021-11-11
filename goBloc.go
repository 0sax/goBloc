package goBlocUtilities

import (
	"encoding/base64"
	"errors"
	"github.com/0sax/err2"
	"net/http"
)

const (
	BaseUrl = "https://sandbox.buildwithcore.com/v2.0"
	AuthEndpoint                 = "/authentication/token"
	ElectricityOperatorsEndpoint = "/electricity/operators"
	OneElectricityOperatorEndpoint = "/electricity/operators/"

)

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

func (c *Client) GetElectricityOperators() (eops []ElectricityOperator, err error) {

	var resp Response
	err = c.standardRequest(http.MethodGet, ElectricityOperatorsEndpoint,nil, &resp)
	if err != nil {
	return
	}

	if resp.isSuccess() {
		eops, err = resp.electricityOperators()
		return
	} else {
		re := resp.Error
		err = err2.NewClientErr(errors.New(re.Message), re.Message, re.Code)
		return
	}
}

func (c *Client) GetOneElectricityOperator(id string) (eo *ElectricityOperator, err error) {

	var resp Response
	err = c.standardRequest(http.MethodGet, OneElectricityOperatorEndpoint+id,nil, &resp)
	if err != nil {
	return
	}

	if resp.isSuccess() {
		eo, err = resp.electricityOperator()
		return
	} else {
		re := resp.Error
		err = err2.NewClientErr(errors.New(re.Message), re.Message, re.Code)
		return
	}
}
