package tests

import (
	"math/rand"
	"strconv"

	"github.com/ProcessMaker/pmio-sdk-go/pmio"
)

//Test case for FindProcesses
func (suite *ClientTestSuite) TestFindProcess() {
	//create process
	rnd := strconv.Itoa(rand.Int())
	processAtt := pmio.ProcessAttributes{}
	processAtt.Name = "ProcessName" + rnd
	processAtt.Status = "ACTIVE"
	processAtt.Type_ = "NORMAL"
	processAtt.DurationBy = "WORKING_DAYS"
	processAtt.DesignAccess = "PUBLIC"

	process := pmio.Process{}
	process.Attributes = processAtt

	processCreateItem := pmio.ProcessCreateItem{Data: process}

	procItem, _, err := suite.client.AddProcess(processCreateItem)
	suite.Require().NoError(err)
	suite.NotEmpty(procItem.Data.Id)

	//find all processes
	processCollection, _, err := suite.client.FindProcesses(1, 100)
	suite.Require().NoError(err)
	suite.NotZero(len(processCollection.Data))
}
