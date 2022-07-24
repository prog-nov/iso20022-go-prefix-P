package iso20022

// The status of an instruction, advice or request.
type PendingReason1 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason1) AddCode() *PendingReason1Choice {
	p.Code = new(PendingReason1Choice)
	return p.Code
}

func (p *PendingReason1) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
