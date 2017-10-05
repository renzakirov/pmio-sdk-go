package tests

import (
	"math/rand"
	"strconv"

	"github.com/ProcessMaker/pmio-sdk-go/pmio"
)

//Test case for DeleteProcess
func (suite *ClientTestSuite) TestDeleteProcess() {
	//create proccess
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

	processItem, _, err := suite.client.AddProcess(processCreateItem)
	suite.Require().NoError(err, "Process should be created")
	suite.Require().NotEmpty(processItem.Data.Id, "Process should be created")

	var res *pmio.ResultSuccess
	res, _, err = suite.client.DeleteProcess(processItem.Data.Id)
	suite.Require().NoError(err)
	suite.Equal("1202", res.Meta.Code, "TestDeleteProcess: result code expected")
}
