package tests

import (
	"math/rand"
	"strconv"

	"github.com/ProcessMaker/pmio-sdk-go/pmio"
)

//Test case for FindUserById
func (suite *ClientTestSuite) TestFindUserById() {
	//create user
	rnd := strconv.Itoa(rand.Int())
	userAtt := pmio.UserAttributes{}
	userAtt.Username = "user" + rnd
	userAtt.Password = "pass" + rnd
	userAtt.Firstname = "Name" + rnd
	userAtt.Lastname = "Last" + rnd
	userAtt.Email = userAtt.Firstname + "@" + userAtt.Lastname + ".com"

	user := pmio.User{}
	user.Attributes = userAtt
	user.Type_ = "user"

	userCreateItem := pmio.UserCreateItem{Data: user}

	userItem, _, err := suite.client.AddUser(userCreateItem)
	suite.Require().NoError(err, "User should be created")
	suite.Require().NotEmpty(userItem.Data.Id, "User should be created")

	//store user Id
	userID := userItem.Data.Id

	//find user whith stored Id
	userItem, _, err = suite.client.FindUserById(userID)

	suite.Require().NoError(err)
	suite.Equal(userID, userItem.Data.Id)
	suite.Equal(userAtt.Username, userItem.Data.Attributes.Username)
	suite.Equal(userAtt.Firstname, userItem.Data.Attributes.Firstname)
	suite.Equal(userAtt.Lastname, userItem.Data.Attributes.Lastname)
}
