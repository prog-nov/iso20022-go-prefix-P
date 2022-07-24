package iso20022

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation1 struct {

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmountOrUnknownChoice `xml:"Val"`

	// Type of value in which the price is expressed.
	ValueType *PriceValueType2Code `xml:"ValTp,omitempty"`

	// Type and information about a price.
	Type *TypeOfPrice5Code `xml:"Tp"`

	// Place from which the price was obtained.
	SourceOfPrice *PriceSourceFormatChoice `xml:"SrcOfPric,omitempty"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTimeChoice `xml:"QtnDt,omitempty"`

	// Indicates whether the price is expressed as a yield. The absence of Yielded means it is not applicable.
	Yielded *YesNoIndicator `xml:"Yldd,omitempty"`
}

func (p *PriceInformation1) AddValue() *PriceRateOrAmountOrUnknownChoice {
	p.Value = new(PriceRateOrAmountOrUnknownChoice)
	return p.Value
}

func (p *PriceInformation1) SetValueType(value string) {
	p.ValueType = (*PriceValueType2Code)(&value)
}

func (p *PriceInformation1) SetType(value string) {
	p.Type = (*TypeOfPrice5Code)(&value)
}

func (p *PriceInformation1) AddSourceOfPrice() *PriceSourceFormatChoice {
	p.SourceOfPrice = new(PriceSourceFormatChoice)
	return p.SourceOfPrice
}

func (p *PriceInformation1) AddQuotationDate() *DateAndDateTimeChoice {
	p.QuotationDate = new(DateAndDateTimeChoice)
	return p.QuotationDate
}

func (p *PriceInformation1) SetYielded(value string) {
	p.Yielded = (*YesNoIndicator)(&value)
}
