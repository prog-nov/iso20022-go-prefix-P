package iso20022

// Product purchased to be paid.
type Product1 struct {

	// Product code of the item purchased.
	ProductCode *Max70Text `xml:"PdctCd"`

	// Unit of measure of the item purchased.
	UnitOfMeasure *UnitOfMeasure1Code `xml:"UnitOfMeasr,omitempty"`

	// Product quantity.
	ProductQuantity *DecimalNumber `xml:"PdctQty,omitempty"`

	// Price per unit of product.
	UnitPrice *ImpliedCurrencyAndAmount `xml:"UnitPric,omitempty"`

	// Monetary value of purchased product.
	ProductAmount *ImpliedCurrencyAndAmount `xml:"PdctAmt"`

	// Information on tax paid on the product.
	TaxType *Max35Text `xml:"TaxTp,omitempty"`

	// Additional information related to the product.
	AdditionalProductInformation *Max35Text `xml:"AddtlPdctInf,omitempty"`
}

func (p *Product1) SetProductCode(value string) {
	p.ProductCode = (*Max70Text)(&value)
}

func (p *Product1) SetUnitOfMeasure(value string) {
	p.UnitOfMeasure = (*UnitOfMeasure1Code)(&value)
}

func (p *Product1) SetProductQuantity(value string) {
	p.ProductQuantity = (*DecimalNumber)(&value)
}

func (p *Product1) SetUnitPrice(value, currency string) {
	p.UnitPrice = NewImpliedCurrencyAndAmount(value, currency)
}

func (p *Product1) SetProductAmount(value, currency string) {
	p.ProductAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (p *Product1) SetTaxType(value string) {
	p.TaxType = (*Max35Text)(&value)
}

func (p *Product1) SetAdditionalProductInformation(value string) {
	p.AdditionalProductInformation = (*Max35Text)(&value)
}
