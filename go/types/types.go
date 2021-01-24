package types

type Receipt struct {
	Timestamp string
	Price     string
	Volume    string
	Fee       string
}

// Order at KRW market with market price
type Order struct {
	Currency string
	Price    string
}

type Credential struct {
	Sensitive string
}

type ResponseAccount struct {
	body string
}

type ResponseOrder struct {
	body string
}
