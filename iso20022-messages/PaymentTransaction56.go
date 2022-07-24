package iso20022

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransaction56 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed transaction.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalIdentification *Max35Text `xml:"RvslId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the reversed transaction.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	ReversedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RvsdInstdAmt,omitempty"`

	// Specifies if the creditor and/or debtor will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the reversal message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference20 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction56) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction56) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction56) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction56) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction56) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction56) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction56) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction56) AddOriginalTransactionReference() *OriginalTransactionReference20 {
	p.OriginalTransactionReference = new(OriginalTransactionReference20)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction56) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}
