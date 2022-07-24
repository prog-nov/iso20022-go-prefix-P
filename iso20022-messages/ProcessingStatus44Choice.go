package iso20022

// Choice of format for the processing status.
type ProcessingStatus44Choice struct {

	// Request has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus7Choice `xml:"AckdAccptd"`

	// Modification Request has been rejected for further processing.
	Rejected *RejectionOrRepairStatus14Choice `xml:"Rjctd"`

	// Modification request was completed.
	Completed *ProprietaryReason1 `xml:"Cmpltd"`

	// Modification request will not be executed.
	Denied *DeniedStatus10Choice `xml:"Dnd"`

	// Modification is pending. It is not known at this time whether modification can be affected.
	Pending *PendingStatus13Choice `xml:"Pdg"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (p *ProcessingStatus44Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus7Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus7Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus44Choice) AddRejected() *RejectionOrRepairStatus14Choice {
	p.Rejected = new(RejectionOrRepairStatus14Choice)
	return p.Rejected
}

func (p *ProcessingStatus44Choice) AddCompleted() *ProprietaryReason1 {
	p.Completed = new(ProprietaryReason1)
	return p.Completed
}

func (p *ProcessingStatus44Choice) AddDenied() *DeniedStatus10Choice {
	p.Denied = new(DeniedStatus10Choice)
	return p.Denied
}

func (p *ProcessingStatus44Choice) AddPending() *PendingStatus13Choice {
	p.Pending = new(PendingStatus13Choice)
	return p.Pending
}

func (p *ProcessingStatus44Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	p.Proprietary = new(ProprietaryStatusAndReason1)
	return p.Proprietary
}
