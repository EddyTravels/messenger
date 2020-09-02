package messenger

// Action is used to determine what kind of message a webhook event is.
type Action int

const (
	// UnknownAction means that the event was not able to be classified.
	UnknownAction Action = iota - 1
	// TextAction means that the event was a text message (May contain attachments).
	TextAction
	// DeliveryAction means that the event was advising of a successful delivery to a
	// previous recipient.
	DeliveryAction
	// ReadAction means that the event was a previous recipient reading their respective
	// messages.
	ReadAction
	// PostBackAction represents post call back
	PostBackAction
	// OptInAction represents opting in through the Send to Messenger button
	OptInAction
	// ReferralAction represents ?ref parameter in m.me URLs
	ReferralAction
	// AccountLinkingAction means that the event concerns changes in account linking
	// status.
	AccountLinkingAction
	// PassThreadControlAction means that thread control was passed to application
	PassThreadControlAction
	// TakeThreadControlAction means that thread control was taken away from application
	TakeThreadControlAction
	// RequestThreadControlAction means that secondary receiver is requesting thread control
	RequestThreadControlAction
	// StandbyAction means that event occurred but your application is not the current thread owner
	StandbyAction
)
