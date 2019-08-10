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

type AbstractICOBuilder struct {
	UserICOAPI model.UserICOAPI
}

/*BuildUserICOXLM - WORKING MODEL
@author - Azeem Ashraf
@desc - use the parameter to get user public key to build payment transaction for XLM and submit it
@params - ResponseWriter,Request
*/
func (AP *AbstractICOBuilder) BuildUserICOXLM(w http.ResponseWriter, r *http.Request) {

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
			build.NativeAmount{"100"},
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
		return
	}

	//CONVERT THE SIGNED XDR TO BASE64 to SUBMIT TO STELLAR
	txeB64, err := GatewayTXE.Base64()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "Failed",
		}
		json.NewEncoder(w).Encode(result)
		return
	}

	//SUBMIT THE GATEWAY'S SIGNED XDR
	display1 := stellarExecuter.ConcreteSubmitXDR{XDR: txeB64}
	response1 := display1.SubmitXDR(true)

	if response1.Error.Code == 400 {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "Failed: " + response1.Error.Message,
		}
		json.NewEncoder(w).Encode(result)
		return
	} else {

		AP.UserICOAPI.TxnHash = response1.TXNID
		err2 := object.InsertICOTransaction(AP.UserICOAPI)
		if err2 != nil {
			w.WriteHeader(http.StatusBadRequest)
			result := apiModel.SubmitXDRSuccess{
				Status: "Failed",
			}
			json.NewEncoder(w).Encode(result)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			result := apiModel.SubmitXDRSuccess{
				Status: "Success",
			}
			json.NewEncoder(w).Encode(result)
			return
		}
	}

}

/*BuildUserICOJIGXU - WORKING MODEL
@author - Azeem Ashraf
@desc - use the parameter to get user public key to build payment transaction for JIGXU and submit it
@params - ResponseWriter,Request
*/
func (AP *AbstractICOBuilder) BuildUserICOJIGXU(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	object := dao.Connection{}

	///HARDCODED CREDENTIALS ISSUER KEYPAIR FOR JIGXU
	publicKey := constants.JIGXUDISTRIBUTORPUB
	secretKey := constants.JIGXUDISTRIBUTORSEC
	// var result model.SubmitXDRResponse

	//BUILD THE PAYMENT XDR
	tx, err := build.Transaction(
		build.SourceAccount{publicKey},
		build.TestNetwork,
		build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		build.Payment(
			build.Destination{AddressOrSeed: AP.UserICOAPI.PublicKey},
			build.CreditAmount{"JIGXU", constants.JIGXUISSUERPUB, "30"},
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
		return
	}

	//CONVERT THE SIGNED XDR TO BASE64 to SUBMIT TO STELLAR
	txeB64, err := GatewayTXE.Base64()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "Failed",
		}
		json.NewEncoder(w).Encode(result)
		return
	}

	//SUBMIT THE GATEWAY'S SIGNED XDR
	display1 := stellarExecuter.ConcreteSubmitXDR{XDR: txeB64}
	response1 := display1.SubmitXDR(true)

	if response1.Error.Code == 400 {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "Failed: " + response1.Error.Message,
		}
		json.NewEncoder(w).Encode(result)
		return
	} else {

		AP.UserICOAPI.TxnHash = response1.TXNID
		err2 := object.InsertICOTransaction(AP.UserICOAPI)
		if err2 != nil {
			w.WriteHeader(http.StatusBadRequest)
			result := apiModel.SubmitXDRSuccess{
				Status: "Failed",
			}
			json.NewEncoder(w).Encode(result)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			result := apiModel.SubmitXDRSuccess{
				Status: "Success",
			}
			json.NewEncoder(w).Encode(result)
			return
		}
	}

}
