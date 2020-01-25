package model


type SubmitXDRResponse struct {
	TDPID     string
	Identifier string
	PublicKey   string
	TXNID string
	TxnType string
	Error     Error
}
