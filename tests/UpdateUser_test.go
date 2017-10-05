package tests

import (
	"math/rand"
	"strconv"

	"github.com/ProcessMaker/pmio-sdk-go/pmio"
)

//Test case for UpdateUser
func (suite *ClientTestSuite) TestUpdateUser() {
	//create new user
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
	//save them parameters in addUser
	addUser, _, err := suite.client.AddUser(userCreateItem)
	suite.Require().NoError(err, "User should be created")
	suite.Require().NotEmpty(addUser.Data.Id, "User should be created")

	//change Username
	addUser.Data.Attributes.Username = "NewName" + rnd

	//update user with new Username
	updUser, _, err := suite.client.UpdateUser(addUser.Data.Id, pmio.UserUpdateItem{addUser.Data})

	suite.Require().NoError(err)
	suite.Equal(addUser.Data.Id, updUser.Data.Id)
	suite.Equal(addUser.Data.Attributes.Username, updUser.Data.Attributes.Username)
}
