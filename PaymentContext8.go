package iso20022

// Payment context in which the transaction is performed.
type PaymentContext8 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel3Code `xml:"TxChanl,omitempty"`

	// Indicates whether a message can be sent or not on an attendant display (attendant display present or not).
	AttendantMessageCapable *TrueFalseIndicator `xml:"AttndntMsgCpbl,omitempty"`

	// Language used to display messages to the attendant.
	// Reference: ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	AttendantLanguage *LanguageCode `xml:"AttndntLang,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`

	// Payment options the card acceptor can support.
	SupportedOption []*SupportedPaymentOption1Code `xml:"SpprtdOptn,omitempty"`
}

func (p *PaymentContext8) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext8) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext8) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext8) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext8) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel3Code)(&value)
}

func (p *PaymentContext8) SetAttendantMessageCapable(value string) {
	p.AttendantMessageCapable = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext8) SetAttendantLanguage(value string) {
	p.AttendantLanguage = (*LanguageCode)(&value)
}

func (p *PaymentContext8) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentContext8) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

func (p *PaymentContext8) AddSupportedOption(value string) {
	p.SupportedOption = append(p.SupportedOption, (*SupportedPaymentOption1Code)(&value))
}
