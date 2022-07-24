package iso20022

// Choice between a standard code and a proprietary code to specify the reason why the instruction/event has a pending status.
type PendingReason24Choice struct {

	// Standard code to specify the reason why the instruction/event has a pending status.
	Code *PendingReason13Code `xml:"Cd"`

	// Proprietary identification of the reason why the instruction/event has a pending status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PendingReason24Choice) SetCode(value string) {
	p.Code = (*PendingReason13Code)(&value)
}

func (p *PendingReason24Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
