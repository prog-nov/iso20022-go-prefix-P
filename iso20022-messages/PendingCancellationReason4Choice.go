package iso20022

// Choice between a standard code or a proprietary code to specify the reason why a cancellation request sent for the related instruction is pending.
type PendingCancellationReason4Choice struct {

	// Standard code to specify the reason why a cancellation request sent for the related instruction is pending.
	Code *PendingCancellationReason4Code `xml:"Cd"`

	// Proprietary identification of the reason why a cancellation request sent for the related instruction is pending.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PendingCancellationReason4Choice) SetCode(value string) {
	p.Code = (*PendingCancellationReason4Code)(&value)
}

func (p *PendingCancellationReason4Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
