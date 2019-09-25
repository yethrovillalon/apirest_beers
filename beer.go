package main

// StatusC ... < Some lines that describe your function>
type StatusC struct {
	Status      int    `json:"code"`
	Description string `json:"description"`
}

// Beer ... < Some lines that describe your function>
type Beer struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Brewery  string  `json:"brewery"`
	Country  string  `json:"country"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

// Beers ... < Some lines that describe your function>
type Beers []Beer

// BeerPrice ... < Some lines that describe your function>
type BeerPrice struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Brewery  string  `json:"brewery"`
	Country  string  `json:"country"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
	Quantity float64 `json:"quantity"`
	PriceBox float64 `json:"priceBox"`
}

// BeersBox ... < Some lines that describe your function>
type BeersBox []BeerPrice

// Retorno ... < Some lines that describe your function>
type Retorno struct {
	Status StatusC
	Beers  Beers
}

// RetornoBox ... < Some lines that describe your function>
type RetornoBox struct {
	Status   StatusC
	BeersBox BeersBox
}

// RetornoError ... < Some lines that describe your function>
type RetornoError struct {
	Status StatusC
}

// apiCashResponse ... < Some lines that describe your function>
type apiCashResponse struct {
	Success   bool               `json:"success"`
	Terms     string             `json:"terms"`
	Privacy   string             `json:"privacy"`
	Timestamp float64            `json:"timestamp"`
	Source    string             `json:"source"`
	Quotes    map[string]float64 `json:"quotes"`
}
