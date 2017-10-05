package tests

import (
	"math/rand"
	"strconv"

	"github.com/ProcessMaker/pmio-sdk-go/pmio"
)

//Test case for AddUserToGroup
func (suite *ClientTestSuite) TestAddUserToGroup() {
	//create 3 users
	arrUsers := make([]string, 3, 3)
	for i := 0; i < 3; i++ {
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

		arrUsers[i] = userItem.Data.Id
	}

	userIds := pmio.UserIds{Users: arrUsers}
	groupAddUsersItem := pmio.GroupAddUsersItem{Data: userIds}

	//create Group
	rnd := strconv.Itoa(rand.Int())
	groupAtt := pmio.GroupAttributes{}
	groupAtt.Code = "GroupCode" + rnd
	groupAtt.Title = "GroupTitle" + rnd
	groupAtt.Status = "ACTIVE"

	group := pmio.Group{}
	group.Attributes = groupAtt

	groupCreateItem := pmio.GroupCreateItem{Data: group}

	grp, _, err := suite.client.AddGroup(groupCreateItem)
	suite.Require().NoError(err, "Group should be created")
	suite.Require().NotEmpty(grp.Data.Id, "Group should be created")

	//AddUsersToGroup
	var res *pmio.ResultSuccess
	res, _, err = suite.client.AddUsersToGroup(grp.Data.Id, groupAddUsersItem)
	suite.Require().NoError(err)
	suite.Equal("1021", res.Meta.Code, "User should be attached to the Group")
}
