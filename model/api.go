package model

type UserICOAPI struct {
	EmailHash       string `json:"emailHash"`
	EncryptedSecret string `json:"encryptedSecret"`
	PublicKey       string `json:"publicKey"`
	Alias           string `json:"alias"`
	TxnHash         string `json:"txnHash"`
}
