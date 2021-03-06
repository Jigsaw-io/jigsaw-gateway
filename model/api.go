package model

type UserICOAPI struct {
	EmailHash       string `json:"emailHash"`
	EncryptedSecret string `json:"encryptedSecret"`
	PublicKey       string `json:"publicKey"`
	Alias           string `json:"alias"`
	TxnHash         string `json:"txnHash"`
}

type KnowledgeAPI struct {
	EmailHash       string `json:"emailHash"`
	PreviousTxn       string `json:"previousTxn"`
	EncryptedSecret string `json:"encryptedSecret"`
	PublicKey       string `json:"publicKey"`
	Alias           string `json:"alias"`
	TxnHash         string `json:"txnHash"`
	Type			string `json:"type"`
	KnowledgeHash	string `json:"knowledeHash"`
	XDR				string `json:"xdr"`
}

type ContributionAPI struct {
	EmailHash       string `json:"emailHash"`
	PreviousTxn       string `json:"previousTxn"`
	EncryptedSecret string `json:"encryptedSecret"`
	PublicKey       string `json:"publicKey"`
	Alias           string `json:"alias"`
	TxnHash         string `json:"txnHash"`
	Type			string `json:"type"`
	KnowledgeHash	string `json:"knowledeHash"`
	KnowledgeID		string `json:"knowledgeID"`
	XDR				string `json:"xdr"`
}

type VoteAPI struct {
	EmailHash       string `json:"emailHash"`
	PreviousTxn       string `json:"previousTxn"`
	EncryptedSecret string `json:"encryptedSecret"`
	PublicKey       string `json:"publicKey"`
	Alias           string `json:"alias"`
	TxnHash         string `json:"txnHash"`
	Type			string `json:"type"`
	ContributionID		string `json:"contributionID"`
	KnowledgeID		string `json:"knowledgeID"`
	KnowledgeHash	string `json:"knowledeHash"`
	XDR				string `json:"xdr"`
}

type LastTxnAPI struct{
	LastTxn       string `json:"lastTxn"`
}