package iso20022

// Choice of format for the processing status.
type ProcessingStatus3Choice struct {

	// A cancellation request from yourself for this instruction is pending waiting for further processing (only as an response to an SecuritiesTransactionStatusQuery). The pending status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	PendingCancellation *PendingStatus4Choice `xml:"PdgCxl"`

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus3Choice `xml:"AckdAccptd"`

	// Processing of the instruction/request is pending.
	PendingProcessing *PendingProcessingStatus1Choice `xml:"PdgPrcg"`

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionStatus3Choice `xml:"Rjctd"`

	// Instruction/Request is accepted but in repair.
	Repair *RepairStatus3Choice `xml:"Rpr"`

	// Instruction has been cancelled (only as an response to an SecuritiesTransactionStatusQuery). The status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	Cancelled *CancellationStatus4Choice `xml:"Canc"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`

	// Cancellation request from your counterparty for this transaction is pending waiting for your cancellation request or your consent.
	CancellationRequested *NoSpecifiedReason1 `xml:"CxlReqd"`

	// Modification request from your counterparty for this transaction is pending waiting for your cancellation request or your consent.
	ModificationRequested *NoSpecifiedReason1 `xml:"ModReqd"`
}

func (p *ProcessingStatus3Choice) AddPendingCancellation() *PendingStatus4Choice {
	p.PendingCancellation = new(PendingStatus4Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus3Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus3Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus3Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus3Choice) AddPendingProcessing() *PendingProcessingStatus1Choice {
	p.PendingProcessing = new(PendingProcessingStatus1Choice)
	return p.PendingProcessing
}

func (p *ProcessingStatus3Choice) AddRejected() *RejectionStatus3Choice {
	p.Rejected = new(RejectionStatus3Choice)
	return p.Rejected
}

func (p *ProcessingStatus3Choice) AddRepair() *RepairStatus3Choice {
	p.Repair = new(RepairStatus3Choice)
	return p.Repair
}

func (p *ProcessingStatus3Choice) AddCancelled() *CancellationStatus4Choice {
	p.Cancelled = new(CancellationStatus4Choice)
	return p.Cancelled
}

func (p *ProcessingStatus3Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	p.Proprietary = new(ProprietaryStatusAndReason1)
	return p.Proprietary
}

func (p *ProcessingStatus3Choice) AddCancellationRequested() *NoSpecifiedReason1 {
	p.CancellationRequested = new(NoSpecifiedReason1)
	return p.CancellationRequested
}

func (p *ProcessingStatus3Choice) AddModificationRequested() *NoSpecifiedReason1 {
	p.ModificationRequested = new(NoSpecifiedReason1)
	return p.ModificationRequested
}
