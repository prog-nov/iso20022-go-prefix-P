package iso20022

// Sensible data associated with the payment card performing the transaction.
type PlainCardData18 struct {

	// Primary Account Number (PAN) of the card.
	PAN *Min8Max28NumericText `xml:"PAN,omitempty"`

	// Identify a card or a payment token inside a set of cards with the same PAN.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt,omitempty"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`

	// ISO track 1 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The format is conform to ISO 7813, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track1 *Max76Text `xml:"Trck1,omitempty"`

	// ISO track 2 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The content is conform to ISO 7813, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track2 *Max37Text `xml:"Trck2,omitempty"`

	// ISO track 3 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The content is conform to ISO 4909, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track3 *Max104Text `xml:"Trck3,omitempty"`

	// Name of the cardholder stored on the card.
	CardholderName *Max45Text `xml:"CrdhldrNm,omitempty"`
}

func (p *PlainCardData18) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData18) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData18) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData18) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData18) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

func (p *PlainCardData18) SetTrack1(value string) {
	p.Track1 = (*Max76Text)(&value)
}

func (p *PlainCardData18) SetTrack2(value string) {
	p.Track2 = (*Max37Text)(&value)
}

func (p *PlainCardData18) SetTrack3(value string) {
	p.Track3 = (*Max104Text)(&value)
}

func (p *PlainCardData18) SetCardholderName(value string) {
	p.CardholderName = (*Max45Text)(&value)
}
