package iso20022

// Choice between an amount or an unspecified rate.
type RateAndAmountFormat42Choice struct {

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`
}

func (r *RateAndAmountFormat42Choice) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateAndAmountFormat42Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateValueType7Code)(&value)
}
