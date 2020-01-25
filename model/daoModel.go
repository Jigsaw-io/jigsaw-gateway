package model



type TransactionCollectionBody struct {
	Identifier      string
	TdpId           string
	ProfileID       string
	TxnHash         string
	PreviousTxnHash string
	FromIdentifier1 string
	FromIdentifier2 string
	ToIdentifier    string
	ItemCode        string
	ItemAmount      string
	PublicKey       string
	TxnType         string
	XDR             string
	Status          string
	MergeID         string
	Orphan          bool
}

