package tests

import (
	"math/rand"
	"strconv"

	"github.com/ProcessMaker/pmio-sdk-go/pmio"
)

//Test case for AddGroup
func (suite *ClientTestSuite) TestAddGroup() {

	rnd := strconv.Itoa(rand.Int())
	groupAtt := pmio.GroupAttributes{}
	groupAtt.Code = "GroupCode" + rnd
	groupAtt.Title = "GroupTitle" + rnd
	groupAtt.Status = "ACTIVE"

	group := pmio.Group{}
	group.Attributes = groupAtt

	groupCreateItem := pmio.GroupCreateItem{Data: group}

	grp, _, err := suite.client.AddGroup(groupCreateItem)
	suite.Require().NoError(err)
	suite.NotEmpty(grp.Data.Id)
	suite.Equal(groupAtt.Code, grp.Data.Attributes.Code)
	suite.Equal(groupAtt.Title, grp.Data.Attributes.Title)
}
