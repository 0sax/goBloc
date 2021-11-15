package goBlocUtilities

import (
	"encoding/json"
	"time"
)

const (
	Success = "SUCCESS"
	Failure = "FAILURE"
)

type Error struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	UniqueCode string `json:"unique_code"`
}

type Response struct {
	Status string          `json:"status"`
	Error  Error           `json:"error"`
	Data   json.RawMessage `json:"data"`
}

func (r *Response) isSuccess() bool {
	return r.Status == Success
}

func (r *Response) operator() (eo *Operator, err error) {
	err = json.Unmarshal([]byte(r.Data), &eo)
	return
}

func (r *Response) operators() (eos []Operator, err error) {
	err = json.Unmarshal([]byte(r.Data), &eos)
	return
}

func (r *Response) products() (eps []Product, err error) {
	err = json.Unmarshal([]byte(r.Data), &eps)
	return
}

func (r *Response) loginResponse() (lr *LoginResponse, err error) {
	err = json.Unmarshal(r.Data, &lr)
	return
}

func (r *Response) customerValidation() (cv *CustomerValidation, err error) {
	err = json.Unmarshal(r.Data, &cv)
	return
}

func (r *Response) paymentFulfillment() (pfr *PaymentFulfilmentResponse, err error) {
	err = json.Unmarshal(r.Data, &pfr)
	return
}

type LoginResponse struct {
	ExpiresIn         int    `json:"expires_in"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	PasswordUpdatedAt string `json:"password_updated_at"`
	Token             string `json:"token"`
	UserCategory      string `json:"user_category"`
	UserCategoryId    string `json:"user_category_id"`
}

type Operator struct {
	Desc   string `json:"desc"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Sector string `json:"sector"`
}

type Product struct {
	Category string      `json:"category"`
	Desc     interface{} `json:"desc"`
	FeeType  string      `json:"fee_type"`
	Id       string      `json:"id"`
	Meta     interface{} `json:"meta"`
	Name     string      `json:"name"`
	Operator string      `json:"operator"`
}

type CustomerValidation struct {
		AccountNumber interface{} `json:"account_number"`
		Address       string      `json:"address"`
		Category      string      `json:"category"`
		Debt          interface{} `json:"debt"`
		Email         string      `json:"email"`
		Meta          struct {
			BusinessUnit        string    `json:"business_unit"`
			LastTransactionDate time.Time `json:"last_transaction_date"`
			Undertaking         string    `json:"undertaking"`
		} `json:"meta"`
		MeterNumber      interface{} `json:"meter_number"`
		MinimumPayable   string      `json:"minimum_payable"`
		Name             string      `json:"name"`
		PaymentReference string      `json:"payment_reference"`
		Phone            string      `json:"phone"`
		TariffClass      interface{} `json:"tariff_class"`
		TariffRate       interface{} `json:"tariff_rate"`
		UniqueId         string      `json:"unique_id"`
}

type PaymentFulfilmentResponse struct {
	AccountNumber string `json:"account_number"`
	Address       string `json:"address"`
	Amount        string `json:"amount"`
	BusinessUnit  string `json:"business_unit"`
	Created       string `json:"created"`
	CurrentDebt   string `json:"current_debt"`
	EReceipt      string `json:"e_receipt"`
	Meta          struct {
		CarriedOverAmount string `json:"carried_over_amount"`
		IncidentDeduction string `json:"incident_deduction"`
	} `json:"meta"`
	MeterNumber      string      `json:"meter_number"`
	Name             string      `json:"name"`
	OperatorReceipt  interface{} `json:"operator_receipt"`
	OperatorVat      string      `json:"operator_vat"`
	PaymentReference string      `json:"payment_reference"`
	PreviousDebt     string      `json:"previous_debt"`
	TariffClass      string      `json:"tariff_class"`
	TariffRate       interface{} `json:"tariff_rate"`
	Token            string      `json:"token"`
	TokenCost        string      `json:"token_cost"`
	Units            string      `json:"units"`
}