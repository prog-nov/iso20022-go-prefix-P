package iso20022

// Payment card performing the transaction.
type PaymentCard12 struct {

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData10 `xml:"PlainCardData,omitempty"`

	// Bank identifier number of the issuer for routing purpose.
	IssuerBIN *Max15NumericText `xml:"IssrBIN,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3AlphaNumericText `xml:"CardCcyCd,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard12) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard12) AddPlainCardData() *PlainCardData10 {
	p.PlainCardData = new(PlainCardData10)
	return p.PlainCardData
}

func (p *PaymentCard12) SetIssuerBIN(value string) {
	p.IssuerBIN = (*Max15NumericText)(&value)
}

func (p *PaymentCard12) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard12) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3AlphaNumericText)(&value)
}

func (p *PaymentCard12) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}
