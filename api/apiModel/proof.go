package apiModel

type SubmitXDRSuccess struct {
	Status string 
}


type TestTDP struct {
	XDR string
	// RawTDP string
}

type TestXDRSubmit struct {
	XDR        string
	Identifier string
	TdpId      string
	PublicKey  string
}

