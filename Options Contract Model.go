package models

type OptionsType string
const (
    Call OptionsType = "call"
    Put  OptionsType = "put"
)

type LongShort string
const (
    Long  LongShort = "long"
    Short LongShort = "short"
)

type OptionsContract struct {
    Type          OptionsType `json:"type"`
    StrikePrice   float64     `json:"strike_price"`
    Bid           float64     `json:"bid"`
    Ask           float64     `json:"ask"`
    ExpirationDate string     `json:"expiration_date"`
    LongShort     LongShort   `json:"long_short"`
}
