package dao

import (
	"github.com/Jigsaw-io/jigsaw-gateway/model"

	"gopkg.in/mgo.v2/bson"

	// "fmt"

	"github.com/chebyrash/promise"
)


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
	result := []model.ContributionAPI{}
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
	result := []model.VoteAPI{}
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