package tests

import (
	"math/rand"
	"strconv"

	"github.com/ProcessMaker/pmio-sdk-go/pmio"
)

//Test case for AddTask
func (suite *ClientTestSuite) TestAddTask() {

	//create new process (AddProcess)
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

	//create new task
	taskAtt := pmio.TaskAttributes{}
	taskAtt.Name = "TaskName" + rnd
	taskAtt.Description = "Description Task Name " + rnd
	taskAtt.ProcessId = processItem.Data.Id
	taskAtt.Type_ = "SERVICE-TASK"
	taskAtt.AssignType = "CYCLIC"
	taskAtt.TransferFly = true

	task := pmio.Task{}
	task.Attributes = taskAtt
	task.Type_ = "task"

	taskCreateItem := pmio.TaskCreateItem{Data: task}

	taskItem, _, err := suite.client.AddTask(processItem.Data.Id, taskCreateItem)

	suite.Require().NoError(err)
	suite.NotEmpty(taskItem.Data.Id)
	suite.Equal(taskItem.Data.Attributes.Name, taskAtt.Name)
	suite.Equal(taskItem.Data.Attributes.ProcessId, processItem.Data.Id)
}
