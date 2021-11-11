package goBlocUtilities

import (
	"encoding/json"
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

func (r *Response) electricityOperator() (eo *ElectricityOperator, err error) {
	err = json.Unmarshal([]byte(r.Data), &eo)
	return
}

func (r *Response) electricityOperators() (eos []ElectricityOperator, err error) {
	err = json.Unmarshal([]byte(r.Data), &eos)
	return
}

func (r *Response) loginResponse() (lr *LoginResponse, err error) {
	err = json.Unmarshal(r.Data, &lr)
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

type ElectricityOperator struct {
	Desc   string `json:"desc"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Sector string `json:"sector"`
}

type ElectricityProduct struct {
	Category string      `json:"category"`
	Desc     interface{} `json:"desc"`
	FeeType  string      `json:"fee_type"`
	Id       string      `json:"id"`
	Meta     interface{} `json:"meta"`
	Name     string      `json:"name"`
	Operator string      `json:"operator"`
}
