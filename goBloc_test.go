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
	var eeee []ElectricityOperator
	_ = json.Unmarshal([]byte(operators), &eeee)
	tests := []struct {
		name     string
		wantEops []ElectricityOperator
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"test 1",
			eeee,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEops, err := cc.GetElectricityOperators()
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
	var eeee *ElectricityOperator
	_ = json.Unmarshal([]byte(operator), &eeee)

	tests := []struct {
		name    string
		id    string
		wantEo  *ElectricityOperator
		wantErr bool
	}{
		{"phcn ph",
			"op_s3EXQWJjZA7gJP3AiwpJsM",
		eeee,
		false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEo, err := cc.GetOneElectricityOperator(tt.id)
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