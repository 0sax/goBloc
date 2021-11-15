package goBlocUtilities

type PaymentRequest struct {
	ClientReference   string       `json:"client_reference,omitempty"`
	BeneficiaryMsisdn string       `json:"beneficiary_msisdn,omitempty"`
	PaymentReference  string       `json:"payment_reference,omitempty"`
	Amount            int          `json:"amount,omitempty"`
	ChannelId         int          `json:"channel_id,omitempty"`
	ProductId         string       `json:"product_id,omitempty"`
	ModeId            int          `json:"mode_id,omitempty"`
	Payer             *Payer       `json:"payer,omitempty"`
	Location          string       `json:"location,omitempty"`
	Geolocation       *Geolocation `json:"geolocation,omitempty"`
}

func NewPaymentRequest(
	paymentReference string,
	telecomsClientReference string,
	beneficiaryMsisdn string,
	amount int,
	channelId int,
	productId string,
	modeId int,
	payer *Payer,
	location string,
	geolocation *Geolocation) *PaymentRequest {

	return &PaymentRequest{
		PaymentReference: paymentReference,
		ClientReference:  telecomsClientReference,
		Amount:           amount,
		ChannelId:        channelId,
		ProductId:        productId,
		ModeId:           modeId,
		Payer:            payer,
		Location:         location,
		Geolocation:      geolocation,
		BeneficiaryMsisdn: beneficiaryMsisdn,
	}
}

type Payer struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func NewPayer(name string, phone string, email string) *Payer {
	return &Payer{Name: name, Phone: phone, Email: email}
}

type Geolocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func NewGeolocation(latitude string, longitude string) *Geolocation {
	return &Geolocation{Latitude: latitude, Longitude: longitude}
}
