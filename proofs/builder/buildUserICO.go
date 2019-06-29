package builder

import (
	"encoding/json"
	"net/http"

	"github.com/zeemzo/jigsaw-gateway/model"

	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/zeemzo/jigsaw-gateway/constants"
	"github.com/zeemzo/jigsaw-gateway/dao"
	"github.com/zeemzo/jigsaw-gateway/proofs/executer/stellarExecuter"

	"github.com/zeemzo/jigsaw-gateway/api/apiModel"
)

type AbstractXDRBuilder struct {
	UserICOAPI model.UserICOAPI
}

/*BuildUserICO - WORKING MODEL
@author - Azeem Ashraf
@desc - use the parameter to get user public key to build payment transaction and submit it
@params - ResponseWriter,Request
*/
func (AP *AbstractXDRBuilder) BuildUserICO(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	object := dao.Connection{}

	///HARDCODED CREDENTIALS ISSUER KEYPAIR FOR JIGXU
	publicKey := constants.ISSUERPUB
	secretKey := constants.ISSUERSEC
	// var result model.SubmitXDRResponse

	//BUILD THE PAYMENT XDR
	tx, err := build.Transaction(
		build.SourceAccount{publicKey},
		build.TestNetwork,
		build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		build.Payment(
			build.Destination{AddressOrSeed: AP.UserICOAPI.PublicKey},
			build.CreditAmount{"JIGXU", publicKey, "100"},
		),
	)

	//SIGN THE GATEWAY BUILT XDR WITH GATEWAYS PRIVATE KEY
	GatewayTXE, err := tx.Sign(secretKey)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "Failed",
		}
		json.NewEncoder(w).Encode(result)
	}

	//CONVERT THE SIGNED XDR TO BASE64 to SUBMIT TO STELLAR
	txeB64, err := GatewayTXE.Base64()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "Failed",
		}
		json.NewEncoder(w).Encode(result)
	}

	//SUBMIT THE GATEWAY'S SIGNED XDR
	display1 := stellarExecuter.ConcreteSubmitXDR{XDR: txeB64}
	response1 := display1.SubmitXDR(false)

	if response1.Error.Code == 400 {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "Failed",
		}
		json.NewEncoder(w).Encode(result)
	} else {
	
		err2 := object.InsertTransaction(AP.TxnBody[i])
		if err2 != nil {

		} else {
		}

		if checkBoolArray(Done) {
			w.WriteHeader(http.StatusOK)
			result := apiModel.SubmitXDRSuccess{
				Status: "Success",
			}
			json.NewEncoder(w).Encode(result)
		}

		return
	}
}
