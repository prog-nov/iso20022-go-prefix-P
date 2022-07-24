package iso20022

// Payment card performing the transaction.
type PaymentCard15 struct {

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData12 `xml:"PlainCardData,omitempty"`

	// Masked PAN to be printed the payment receipts or displayed to the cardholder. Masked digits may be absent or replaced by another character as '*'.
	MaskedPAN *Max30Text `xml:"MskdPAN,omitempty"`

	// Type of card product.
	CardProductType *CardProductType1Code `xml:"CardPdctTp,omitempty"`

	// Name of card product.
	CardProductName *Max35Text `xml:"CardPdctNm,omitempty"`
}

func (p *PaymentCard15) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard15) AddPlainCardData() *PlainCardData12 {
	p.PlainCardData = new(PlainCardData12)
	return p.PlainCardData
}

func (p *PaymentCard15) SetMaskedPAN(value string) {
	p.MaskedPAN = (*Max30Text)(&value)
}

func (p *PaymentCard15) SetCardProductType(value string) {
	p.CardProductType = (*CardProductType1Code)(&value)
}

func (p *PaymentCard15) SetCardProductName(value string) {
	p.CardProductName = (*Max35Text)(&value)
}
