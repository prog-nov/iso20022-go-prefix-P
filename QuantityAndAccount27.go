package iso20022

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount27 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown13 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount27) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount27) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount27) AddAccountOwner() *PartyIdentification36Choice {
	q.AccountOwner = new(PartyIdentification36Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount27) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount27) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount27) AddQuantityBreakdown() *QuantityBreakdown13 {
	newValue := new(QuantityBreakdown13)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount27) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}
