package iso20022

// Specifies the reason why the instruction or request has a pending processing status.
type PendingProcessingReason12 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason14Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason12) AddCode() *PendingProcessingReason14Choice {
	p.Code = new(PendingProcessingReason14Choice)
	return p.Code
}

func (p *PendingProcessingReason12) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
