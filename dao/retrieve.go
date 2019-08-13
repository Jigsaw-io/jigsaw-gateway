package dao

import (
	"github.com/zeemzo/jigsaw-gateway/api/apiModel"
	"github.com/zeemzo/jigsaw-gateway/model"

	"gopkg.in/mgo.v2/bson"

	// "fmt"

	"github.com/chebyrash/promise"
)

/*GetCOCbySender Retrieve All COC Object from COCCollection in DB by Sender PublicKey
@author - Azeem Ashraf
*/
func (cd *Connection) GetCOCbySender(sender string) *promise.Promise {
	result := []model.COCCollectionBody{}
	// p := promise.NewPromise()

	var p = promise.New(func(resolve func(interface{}), reject func(error)) {
		// Do something asynchronously.
		session, err := cd.connect()

		if err != nil {
			// fmt.Println(err)
			reject(err)

		}
		defer session.Close()

		c := session.DB("tracified-gateway").C("COC")
		err1 := c.Find(bson.M{"sender": sender}).All(&result)
		if err1 != nil || len(result) == 0 {
			// fmt.Println(err1)
			reject(err1)

		} else {
			resolve(result)

		}

	})

	return p

}

/*GetLastCOCbySubAccount Retrieve the Last COC Object from COCCollection in DB by SubAccount PublicKey
@author - Azeem Ashraf
*/
func (cd *Connection) GetLastCOCbySubAccount(subAccount string) *promise.Promise {
	result := model.COCCollectionBody{}
	result2 := apiModel.GetSubAccountStatusResponse{}
	// p := promise.NewPromise()

	var p = promise.New(func(resolve func(interface{}), reject func(error)) {
		// Do something asynchronously.
		session, err := cd.connect()

		if err != nil {
			// fmt.Println(err)
			reject(err)

		}
		defer session.Close()

		c := session.DB("tracified-gateway").C("COC")
		count, er := c.Find(bson.M{"subaccount": subAccount}).Count()
		if er != nil {
			// fmt.Println(er)
			reject(er)
		}

		err1 := c.Find(bson.M{"subaccount": subAccount}).Skip(count - 1).One(&result)
		if err1 != nil {
			// fmt.Println(err1)
			reject(err1)

		}
		result2.Receiver = result.Receiver
		result2.SequenceNo = result.SequenceNo
		result2.SubAccount = result.SubAccount
		if result.Status == "pending" {
			result2.Available = false
		} else {
			result2.Available = true
		}
		resolve(result2)

	})

	return p

}

/*GetCOCbyReceiver Retrieve All COC Object from COCCollection in DB by Receiver PublicKey
@author - Azeem Ashraf
*/
func (cd *Connection) GetCOCbyReceiver(receiver string) *promise.Promise {
	result := []model.COCCollectionBody{}
	// p := promise.NewPromise()

	var p = promise.New(func(resolve func(interface{}), reject func(error)) {
		// Do something asynchronously.
		session, err := cd.connect()

		if err != nil {
			// fmt.Println(err)
			reject(err)

		}
		defer session.Close()

		c := session.DB("tracified-gateway").C("COC")
		err1 := c.Find(bson.M{"receiver": receiver}).All(&result)
		if err1 != nil || len(result) == 0 {
			// fmt.Println(err1)
			reject(err1)

		} else {
			resolve(result)

		}

	})

	return p

}

/*GetCOCbyAcceptTxn Retrieve a COC Object from COCCollection in DB by Accept TXN
@author - Azeem Ashraf
*/
func (cd *Connection) GetCOCbyAcceptTxn(accepttxn string) *promise.Promise {
	result := model.COCCollectionBody{}
	// p := promise.NewPromise()

	var p = promise.New(func(resolve func(interface{}), reject func(error)) {
		// Do something asynchronously.
		session, err := cd.connect()

		if err != nil {
			// fmt.Println(err)
			reject(err)

		}
		defer session.Close()

		c := session.DB("tracified-gateway").C("COC")
		err1 := c.Find(bson.M{"accepttxn": accepttxn}).One(&result)
		if err1 != nil {
			// fmt.Println(err1)
			reject(err1)

		} else {
			resolve(result)

		}

	})

	return p

}

/*GetCOCbyRejectTxn Retrieve a COC Object from COCCollection in DB by Reject TXN
@author - Azeem Ashraf
*/
func (cd *Connection) GetCOCbyRejectTxn(rejecttxn string) *promise.Promise {
	result := model.COCCollectionBody{}
	// p := promise.NewPromise()

	var p = promise.New(func(resolve func(interface{}), reject func(error)) {
		// Do something asynchronously.
		session, err := cd.connect()

		if err != nil {
			// fmt.Println(err)
			reject(err)

		}
		defer session.Close()

		c := session.DB("tracified-gateway").C("COC")
		err1 := c.Find(bson.M{"rejecttxn": rejecttxn}).One(&result)
		if err1 != nil {
			// fmt.Println(err1)
			reject(err1)

		} else {
			resolve(result)

		}
	})

	return p

}

/*GetCOCbyStatus Retrieve All COC Object from COCCollection in DB by Status
@author - Azeem Ashraf
*/
func (cd *Connection) GetCOCbyStatus(status string) *promise.Promise {
	result := []model.COCCollectionBody{}
	// p := promise.NewPromise()

	var p = promise.New(func(resolve func(interface{}), reject func(error)) {
		// Do something asynchronously.
		session, err := cd.connect()

		if err != nil {
			// fmt.Println(err)
			reject(err)

		}
		defer session.Close()

		c := session.DB("tracified-gateway").C("COC")
		err1 := c.Find(bson.M{"status": status}).All(&result)
		if err1 != nil || len(result) == 0 {
			// fmt.Println(err1)
			reject(err1)

		} else {
			resolve(result)

		}

	})

	return p

}

/*GetLastCOCbyIdentifier Retrieve Last COC Object from COCCollection in DB by Identifier
@author - Azeem Ashraf
*/
func (cd *Connection) GetLastCOCbyIdentifier(identifier string) *promise.Promise {
	result := model.COCCollectionBody{}
	// result2 := apiModel.GetSubAccountStatusResponse{}
	// p := promise.NewPromise()

	var p = promise.New(func(resolve func(interface{}), reject func(error)) {
		// Do something asynchronously.
		session, err := cd.connect()

		if err != nil {
			// fmt.Println(err)
			reject(err)

		}
		defer session.Close()

		c := session.DB("tracified-gateway").C("COC")
		count, er := c.Find(bson.M{"identifier": identifier}).Count()
		if er != nil {
			// fmt.Println(er)
			reject(er)
		}

		err1 := c.Find(bson.M{"identifier": identifier}).Skip(count - 1).One(&result)
		if err1 != nil {
			// fmt.Println(err1)
			reject(err1)
		}

		resolve(result)

	})

	return p

}

/*GetLastKnowledgeByPublicKey Retrieve Last Transaction Object from TransactionCollection in DB by Identifier
@author - Azeem Ashraf
*/
func (cd *Connection) GetLastKnowledgeByPublicKey(publickey string) *promise.Promise {
	result := []model.KnowledgeAPI{}
	// p := promise.NewPromise()

	var p = promise.New(func(resolve func(interface{}), reject func(error)) {
		// Do something asynchronously.
		session, err := cd.connect()

		if err != nil {
			// fmt.Println(err)
			reject(err)

		}
		defer session.Close()

		c := session.DB("jigsaw-gateway").C("Knowledge")
		err1 := c.Find(bson.M{"publickey": publickey}).All(&result)
		if err1 != nil || len(result) == 0 {
			// fmt.Println(err1)
			reject(err1)

		} else {
			resolve(result[len(result)-1])

		}

	})

	return p

}


/*GetLastContributionByKnowledge Retrieve Last Transaction Object from TransactionCollection in DB by Identifier
@author - Azeem Ashraf
*/
func (cd *Connection) GetLastContributionByKnowledge(Knowledge string) *promise.Promise {
	result := []model.KnowledgeAPI{}
	// p := promise.NewPromise()

	var p = promise.New(func(resolve func(interface{}), reject func(error)) {
		// Do something asynchronously.
		session, err := cd.connect()

		if err != nil {
			// fmt.Println(err)
			reject(err)

		}
		defer session.Close()

		c := session.DB("jigsaw-gateway").C("Contributions")
		err1 := c.Find(bson.M{"knowledgeid": Knowledge}).All(&result)
		if err1 != nil || len(result) == 0 {
			// fmt.Println(err1)
			reject(err1)

		} else {
			resolve(result[len(result)-1])

		}

	})

	return p

}


/*GetLastVoteByContribution Retrieve Last Transaction Object from TransactionCollection in DB by Identifier
@author - Azeem Ashraf
*/
func (cd *Connection) GetLastVoteByContribution(Contribution string) *promise.Promise {
	result := []model.KnowledgeAPI{}
	// p := promise.NewPromise()

	var p = promise.New(func(resolve func(interface{}), reject func(error)) {
		// Do something asynchronously.
		session, err := cd.connect()

		if err != nil {
			// fmt.Println(err)
			reject(err)

		}
		defer session.Close()

		c := session.DB("jigsaw-gateway").C("Votes")
		err1 := c.Find(bson.M{"contributionid": Contribution}).All(&result)
		if err1 != nil || len(result) == 0 {
			// fmt.Println(err1)
			reject(err1)

		} else {
			resolve(result[len(result)-1])

		}

	})

	return p

}