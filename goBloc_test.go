package goBlocUtilities

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"os"
	"reflect"
	"testing"
)

var cc *Client
var ee error

func TestMain(m *testing.M) {
	// Write code here to run before tests
	ee = godotenv.Load("vars.env")
	if ee != nil {
		log.Fatalf("authentication error: %v", ee)
	}

	cc, ee = Authenticate(os.Getenv("USERNAME"), os.Getenv("PASSWORD"), BaseUrl)
	if ee != nil {
		log.Fatalf("authentication error: %v", ee)
	}

	// Run tests
	exitVal := m.Run()

	// Write code here to run after tests

	// Exit with exit value from tests
	os.Exit(exitVal)
}

func TestAuthenticate(t *testing.T) {
	type args struct {
		username string
		password string
		baseUrl  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"no error",
			args{os.Getenv("USERNAME"),
				os.Getenv("PASSWORD"),
				BaseUrl},
			false},
		{"wrong password",
			args{os.Getenv("USERNAME"),
				"weronwdsd",
				BaseUrl},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Authenticate(tt.args.username, tt.args.password, tt.args.baseUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(gotC, tt.wantC) {
			//	t.Errorf("Authenticate() gotC = %v, want %v", gotC, tt.wantC)
			//}
		})
	}
}


func TestClient_GetElectricityOperators(t *testing.T) {
	operators := `[
        {
            "desc": "Port Harcourt Electricity Distribution Company",
            "id": "op_s3EXQWJjZA7gJP3AiwpJsM",
            "name": "PHEDC",
            "sector": "Electricity"
        },
        {
            "desc": "Jos Electricity Distribution Company",
            "id": "op_hFzVYtRLcheuGenZXB45Aq",
            "name": "JEDC",
            "sector": "Electricity"
        },
        {
            "desc": "Benin Electricity Distribution Company",
            "id": "op_eqGC74QFKYHTyHNtAJNFDF",
            "name": "BEDC",
            "sector": "Electricity"
        },
        {
            "desc": "Kano Electricity Distribution Company",
            "id": "op_s3Lb7zJJLh7VGF67wTut7U",
            "name": "KEDCO",
            "sector": "Electricity"
        },
        {
            "desc": "Eko Electricity Distribution Company",
            "id": "op_JDODSOE7fJjdhtX6v64TK",
            "name": "EKEDC",
            "sector": "Electricity"
        },
        {
            "desc": "Enugu Electricity Distribution Company",
            "id": "op_hNYNSsTbHp7YZAxn8tNHKR",
            "name": "EEDC",
            "sector": "Electricity"
        },
        {
            "desc": "Abuja Electricity Distribution Company",
            "id": "op_HndkGFuE7fJjh6UX6v64TJ",
            "name": "AEDC",
            "sector": "Electricity"
        },
        {
            "desc": "Ibadan Electricity Distribution Company",
            "id": "op_UmZ9d3NwtgG7SjATFiPFpk",
            "name": "IBEDC",
            "sector": "Electricity"
        },
        {
            "desc": "Kaduna Electricity Distribution Company",
            "id": "op_r49pjUdZaennuRw9Pbretd",
            "name": "KAEDCO",
            "sector": "Electricity"
        },
        {
            "desc": "Ikeja Electricity Distribution Company",
            "id": "op_te89Caom34mWbZbTPuuk6s",
            "name": "IKEDC",
            "sector": "Electricity"
        }
    ]`
	var eeee []Operator
	_ = json.Unmarshal([]byte(operators), &eeee)
	tests := []struct {
		name     string
		wantEops []Operator
		wantErr  bool
		cl       *Client
	}{
		// TODO: Add test cases.
		{"test 1",
			eeee,
			false,
			cc,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEops, err := tt.cl.GetElectricityOperators()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetElectricityOperators() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEops, tt.wantEops) {
				t.Errorf("GetElectricityOperators() gotEops = %v, want %v", gotEops, tt.wantEops)
			}
		})
	}
}

func TestClient_GetOneElectricityOperator(t *testing.T) {
	operator := `{
        "desc": "Port Harcourt Electricity Distribution Company",
        "id": "op_s3EXQWJjZA7gJP3AiwpJsM",
        "name": "PHEDC",
        "sector": "Electricity"
    }`
	var eeee *Operator
	_ = json.Unmarshal([]byte(operator), &eeee)

	tests := []struct {
		name    string
		id      string
		wantEo  *Operator
		wantErr bool
		cl      *Client
	}{
		{"phcn ph",
			"op_s3EXQWJjZA7gJP3AiwpJsM",
			eeee,
			false,
			cc},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEo, err := tt.cl.GetOneElectricityOperator(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOneElectricityOperator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEo, tt.wantEo) {
				t.Errorf("GetOneElectricityOperator() gotEo = %v, want %v", gotEo, tt.wantEo)
			}
		})
	}
}

func TestClient_GetElectricityProducts(t *testing.T) {

	products := `[
	{
	"category": "pctg_aLrTqmTXE8sdDavnKtxPJG",
	"desc": null,
	"fee_type": "FLEXIBLE",
	"id": "prd_HZ45Squnnhw9LiZFXNSGqL",
	"meta": null,
	"name": "PHEDC Bill Payment",
	"operator": "op_s3EXQWJjZA7gJP3AiwpJsM"
	},
	{
	"category": "pctg_aLrTqmTXE8sdDavnKtxPJG",
	"desc": null,
	"fee_type": "FLEXIBLE",
	"id": "prd_Uc28gg954YHftPPoZEVrCp",
	"meta": null,
	"name": "JEDC Bill Payment",
	"operator": "op_hFzVYtRLcheuGenZXB45Aq"
	},
	{
	"category": "pctg_aLrTqmTXE8sdDavnKtxPJG",
	"desc": null,
	"fee_type": "FLEXIBLE",
	"id": "prd_FiijU49GhTCWinMu2BujXi",
	"meta": null,
	"name": "BEDC Bill Payment",
	"operator": "op_eqGC74QFKYHTyHNtAJNFDF"
	},
	{
	"category": "pctg_aLrTqmTXE8sdDavnKtxPJG",
	"desc": null,
	"fee_type": "FLEXIBLE",
	"id": "prd_2ZhWnxEjn6Jxn8VRf484vX",
	"meta": null,
	"name": "KEDCO Bill Payment",
	"operator": "op_s3Lb7zJJLh7VGF67wTut7U"
	},
	{
	"category": "pctg_gT6gfD4Xc5bhbPYTnKPDz8",
	"desc": null,
	"fee_type": "FLEXIBLE",
	"id": "prd_XHSjeVHW9MNi3LvHkL2tQc",
	"meta": null,
	"name": "EKEDC Non Bill Payment",
	"operator": "op_JDODSOE7fJjdhtX6v64TK"
	},
	{
	"category": "pctg_aLrTqmTXE8sdDavnKtxPJG",
	"desc": null,
	"fee_type": "FLEXIBLE",
	"id": "prd_ayZHAUaCuLABCfvFLdf9Yc",
	"meta": null,
	"name": "EKEDC Bill Payment",
	"operator": "op_JDODSOE7fJjdhtX6v64TK"
	},
	{
	"category": "pctg_aLrTqmTXE8sdDavnKtxPJG",
	"desc": null,
	"fee_type": "FLEXIBLE",
	"id": "prd_Ja5yT5NTAYpWjSx5jZ7See",
	"meta": null,
	"name": "EEDC Bill Payment",
	"operator": "op_hNYNSsTbHp7YZAxn8tNHKR"
	},
	{
	"category": "pctg_aLrTqmTXE8sdDavnKtxPJG",
	"desc": null,
	"fee_type": "FLEXIBLE",
	"id": "prd_EqywNZpHahTFkrfFW2r8mv",
	"meta": null,
	"name": "AEDC Bill Payment",
	"operator": "op_HndkGFuE7fJjh6UX6v64TJ"
	},
	{
	"category": "pctg_aLrTqmTXE8sdDavnKtxPJG",
	"desc": null,
	"fee_type": "FLEXIBLE",
	"id": "prd_5dHvcgQK3WmXJGpcBuLzx8",
	"meta": null,
	"name": "KAEDCO Bill Payment",
	"operator": "op_r49pjUdZaennuRw9Pbretd"
	},
	{
	"category": "pctg_aLrTqmTXE8sdDavnKtxPJG",
	"desc": null,
	"fee_type": "FLEXIBLE",
	"id": "prd_NwZ5jYVCXbNxS2zz9UxZvo",
	"meta": null,
	"name": "IKEDC Bill Payment",
	"operator": "op_te89Caom34mWbZbTPuuk6s"
	}
	]`
	var eeee []Product
	_ = json.Unmarshal([]byte(products), &eeee)
	tests := []struct {
		name     string
		wantEops []Product
		wantErr  bool
		cl       *Client
	}{
		{"no error",
			eeee,
			false,
			cc},
		{"shitty client",
			nil,
			true,
			&Client{
				baseUrl:     BaseUrl,
				bearerToken: "eewe",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotEops, err := tt.cl.GetElectricityProducts()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetElectricityProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEops, tt.wantEops) {
				t.Errorf("GetElectricityProducts() gotEops = %v, want %v", gotEops, tt.wantEops)
			}
		})
	}
}

func TestClient_CustomerValidation(t *testing.T) {

	type args struct {
		operatorID string
		customerID string
		category   string
	}
	tests := []struct {
		name    string
		cl      *Client
		args    args
		wantErr bool
	}{
		{"invalid token",
			&Client{
				baseUrl:     BaseUrl,
				bearerToken: "kedu, odinma",
			},
			args{
				operatorID: "op_HndkGFuE7fJjh6UX6v64TJ",
				customerID: "04177215607",
				category:   "prepaid",
			},
			true},
		{"invalid params",
			cc,
			args{
				operatorID: "2",
				customerID: "3",
				category:   "prepaid",
			},
			true},

		{"valid params",
			cc,
			args{
				operatorID: "op_HndkGFuE7fJjh6UX6v64TJ",
				customerID: "04177215607",
				category:   "prepaid",
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := tt.cl.CustomerValidation(tt.args.operatorID, tt.args.customerID, tt.args.category)
			if err != nil {
				t.Logf("error msg: %v\n", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerValidation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_PaymentFulfilment(t *testing.T) {

	payer := NewPayer("Alade Chi Musa", "08123456789", "user@example.com")
	geol := NewGeolocation("6.425963", "3.440534")
	req := NewPaymentRequest("7a6345baaab6567d8aaf553fe9cabdd1c","",
		"",1000, 3, "prd_W4GaPHntujFVc9B686VUcd", 1,
		payer, "Ikeja, Lagos", geol)

	tests := []struct {
		name    string
		cl      *Client
		request *PaymentRequest
		wantErr bool
	}{
		{"test1", cc, req, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := tt.cl.ElectricityPaymentFulfilment(tt.request)
			if err != nil {
			t.Logf("error msg %v\n", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ElectricityPaymentFulfilment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestClient_ElectricityPaymentsReQuery(t *testing.T) {

	tests := []struct {
		name    string
		cl *Client
		ref string
		//wantPfr *PaymentFulfilmentResponse
		wantErr bool
	}{

		{"test 1",
			cc,
			"7a6345baaab6567d8aaf553fe9cabdd1c",
		false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := tt.cl.ElectricityPaymentsReQuery(tt.ref)
			if err != nil {
			t.Logf("error msg: %v\n", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ElectricityPaymentsReQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

//ELECTRICITY
func TestClient_GetTelecomsOperators(t *testing.T) {
	operators := `[
        {
            "desc": "Airtel Nigeria",
            "id": "op_Q4vPwCHzKj9E8CoCLs7TaC",
            "name": "Airtel",
            "sector": "Telecommunications"
        },
        {
            "desc": "Globacom (Glo Nigeria)",
            "id": "op_P9ids5CbcsWacHPb7c2mzp",
            "name": "Glo",
            "sector": "Telecommunications"
        },
        {
            "desc": "MTN Nigeria",
            "id": "op_93FFk4Khww9rYXgK2Xtr9E",
            "name": "MTN",
            "sector": "Telecommunications"
        },
        {
            "desc": "Visafone Nigeria",
            "id": "op_K5w2yzGJJmqrRx83qHbFXv",
            "name": "Visafone",
            "sector": "Telecommunications"
        },
        {
            "desc": "9Mobile Nigeria",
            "id": "op_jwes4uVEWVHtz8aSgyQ6Vy",
            "name": "9Mobile",
            "sector": "Telecommunications"
        }
    ]`
	var eeee []Operator
	_ = json.Unmarshal([]byte(operators), &eeee)
	tests := []struct {
		name     string
		wantEops []Operator
		wantErr  bool
		cl       *Client
	}{
		{"test 1",
			eeee,
			false,
			cc,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEops, err := tt.cl.GetTelecomsOperators()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTelecomsOperators() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEops, tt.wantEops) {
				t.Errorf("GetTelecomsOperators() gotEops = %v, want %v", gotEops, tt.wantEops)
			}
		})
	}
}

func TestClient_GetOneTelecomsOperatorsProducts(t *testing.T) {

	products := `[
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_jyKzZuF4VoXLBq5jvMpw8g",
            "meta": {
                "currency": "NGN",
                "data_expiry": "30 days",
                "data_value": "1.5GB",
                "fee": "1000.00"
            },
            "name": "MTN -- 1.5GB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_3UGD4NKt6XfgLz3YZhm9Ki",
            "meta": {
                "currency": "NGN",
                "data_expiry": "30 days",
                "data_value": "3.5GB",
                "fee": "2000.00"
            },
            "name": "MTN -- 3.5GB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_vywaaowDhZoPiqevERePj8",
            "meta": {
                "currency": "NGN",
                "data_expiry": "30 days",
                "data_value": "3GB",
                "fee": "1500.00"
            },
            "name": "MTN -- 3GB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_VWzE6L9k9KzCfvhawdz6ib",
            "meta": {
                "currency": "NGN",
                "data_expiry": "1 day",
                "data_value": "25MB",
                "fee": "50.00"
            },
            "name": "MTN -- 25MB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_XD6Zr4sUAs9vjv2QBGumbT",
            "desc": null,
            "fee_type": "RANGE",
            "id": "prd_GvERdWnr8B9gfQGqkQU58G",
            "meta": {
                "currency": "NGN",
                "maximum_fee": "300000.00",
                "minimum_fee": "1.00"
            },
            "name": "MTN -- Airtime",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_CVJ5i2VyB8tef4XRxbtkgm",
            "meta": {
                "currency": "NGN",
                "data_expiry": "1 day",
                "data_value": "1GB",
                "fee": "350.00"
            },
            "name": "MTN -- 1GB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_FsDUTmmqAwRXdvbcZtZzoe",
            "meta": {
                "currency": "NGN",
                "data_expiry": "90 days",
                "data_value": "150GB",
                "fee": "70000.00"
            },
            "name": "MTN -- 150GB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_HEpmiJwNwAgjsd9QDfPCF4",
            "meta": {
                "currency": "NGN",
                "data_expiry": "90 days",
                "data_value": "120GB",
                "fee": "50000.00"
            },
            "name": "MTN -- 120GB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_T8bTQdc6nuKNmJXGxWPghg",
            "meta": {
                "currency": "NGN",
                "data_expiry": "60 days",
                "data_value": "100GB",
                "fee": "30000.00"
            },
            "name": "MTN -- 100GB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_4r3DmcJkZwq54iAFyLtiC2",
            "meta": {
                "currency": "NGN",
                "data_expiry": "60 days",
                "data_value": "60GB",
                "fee": "20000.00"
            },
            "name": "MTN -- 60GB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_NoxCVBQgbpRJ5fgLLYcPF2",
            "meta": {
                "currency": "NGN",
                "data_expiry": "30 days",
                "data_value": "40GB",
                "fee": "15000.00"
            },
            "name": "MTN -- 40GB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_b7DDoK7kSRTo58GVaQfLcy",
            "meta": {
                "currency": "NGN",
                "data_expiry": "30 days",
                "data_value": "3GB",
                "fee": "1500.00"
            },
            "name": "MTN -- 3GB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_q84fz3jW8QAWb6kYwBbdHn",
            "meta": {
                "currency": "NGN",
                "data_expiry": "7 days",
                "data_value": "350MB",
                "fee": "300.00"
            },
            "name": "MTN -- 350MB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_nuys7tHSKyNEiE3n2YCpSc",
            "meta": {
                "currency": "NGN",
                "data_expiry": "30 days",
                "data_value": "2GB",
                "fee": "1200.00"
            },
            "name": "MTN -- 2GB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_jwxATdTYnFV8swhSGqrp3W",
            "meta": {
                "currency": "NGN",
                "data_expiry": "30 days",
                "data_value": "25GB",
                "fee": "10000.00"
            },
            "name": "MTN -- 25GB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_SUevLDJAPGwfyJZ9i9Ps4U",
            "meta": {
                "currency": "NGN",
                "data_expiry": "30 days",
                "data_value": "15GB",
                "fee": "6000.00"
            },
            "name": "MTN -- 15GB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_torSPJA47jpTooqoJwqwpb",
            "meta": {
                "currency": "NGN",
                "data_expiry": "30 days",
                "data_value": "11GB",
                "fee": "5000.00"
            },
            "name": "MTN -- 11GB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_Eqh6468LahSi5vHgS4UV62",
            "meta": {
                "currency": "NGN",
                "data_expiry": "7 days",
                "data_value": "750MB",
                "fee": "500.00"
            },
            "name": "MTN -- 750MB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_gveeMWyh7DHcYbrvkxgco3",
            "meta": {
                "currency": "NGN",
                "data_expiry": "2 days",
                "data_value": "200MB",
                "fee": "200.00"
            },
            "name": "MTN -- 200MB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_9He9oE4G28HHqoAXxLEHeD",
            "meta": {
                "currency": "NGN",
                "data_expiry": "1 day",
                "data_value": "75MB",
                "fee": "100.00"
            },
            "name": "MTN -- 75MB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        },
        {
            "category": "pctg_6dNCwhHijBGYQnZ4jXTRJs",
            "desc": null,
            "fee_type": "FIXED",
            "id": "prd_ZLvK4VwFtyicyLPgeYMxrg",
            "meta": {
                "currency": "NGN",
                "data_expiry": "30 days",
                "data_value": "6.5GB",
                "fee": "3500.00"
            },
            "name": "MTN -- 6.5GB",
            "operator": "op_93FFk4Khww9rYXgK2Xtr9E"
        }
    ]`
	var eeee []Product
	_ = json.Unmarshal([]byte(products), &eeee)
	tests := []struct {
		name     string
		operator string
		wantEops []Product
		wantErr  bool
		cl       *Client
	}{
		{"no error",
			"op_93FFk4Khww9rYXgK2Xtr9E",
			eeee,
			false,
			cc},
		{"shitty client",
			"op_93FFk4Khww9rYXgK2Xtr9E",
			nil,
			true,
			&Client{
				baseUrl:     BaseUrl,
				bearerToken: "eewe",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotEops, err := tt.cl.GetOneTelecomsOperatorsProducts(tt.operator)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOneTelecomsOperatorsProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEops, tt.wantEops) {
				t.Errorf("GetOneTelecomsOperatorsProducts() gotEops = %v, want %v", gotEops, tt.wantEops)
			}
		})
	}
}

func TestClient_TelecomsPaymentFulfilment(t *testing.T) {

	payer := NewPayer("Alade Chi Musa", "08123456789", "user@example.com")
	geol := NewGeolocation("6.425963", "3.440534")
	req := NewPaymentRequest("","whaterwds",
		"08103236508",1000, 3, "prd_ZhzAL5M7jXEdPSSunBpwX3", 1,
		payer, "Ikeja, Lagos", geol)

	tests := []struct {
		name    string
		cl      *Client
		request *PaymentRequest
		wantErr bool
	}{
		{"test1", cc, req, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := tt.cl.TelecomsPaymentFulfilment(tt.request)
			if err != nil {
				t.Logf("error msg %v\n", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("TelecomsPaymentFulfilment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestClient_TelecomsPaymentsReQuery(t *testing.T) {

	tests := []struct {
		name    string
		cl *Client
		ref string
		//wantPfr *PaymentFulfilmentResponse
		wantErr bool
	}{

		{"test 1",
			cc,
			"7a6345baaab6567d8aaf553fe9cabdd1c",
			false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := tt.cl.TelecomsPaymentsReQuery(tt.ref)
			if err != nil {
				t.Logf("error msg: %v\n", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("TelecomsPaymentsReQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestClient_TelecomsPaymentsReQueryByClientRef(t *testing.T) {

	tests := []struct {
		name    string
		cl *Client
		ref string
		//wantPfr *PaymentFulfilmentResponse
		wantErr bool
	}{

		{"test 1",
			cc,
			"7a6345baaab6567d8aaf553fe9cabdd1c",
			false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := tt.cl.TelecomsPaymentsReQueryByClientRef(tt.ref)
			if err != nil {
				t.Logf("error msg: %v\n", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("TelecomsPaymentsReQueryByClientRef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
