package tests

import (
	"io/ioutil"

	"github.com/ProcessMaker/pmio-sdk-go/pmio"
)

//Test case for ImportBpmnFile
func (suite *ClientTestSuite) TestImportBpmnFile() {
	//load BPMN file
	dat, err := ioutil.ReadFile("message_startevent.bpmn")
	suite.Require().NoError(err, "Imported BPMN-file not found")
	suite.Require().NotZero(len(dat), "Imported BPMN-file is empty")
	//create struct
	bpmnAtt := pmio.BpmnFileAttributes{}
	bpmnAtt.Bpmn = string(dat)

	bpmnFile := pmio.BpmnFile{}
	bpmnFile.Attributes = bpmnAtt

	bpmnImportItem := pmio.BpmnImportItem{Data: bpmnFile}

	var processCollection1 *pmio.ProcessCollection1
	//
	processCollection1, _, err = suite.client.ImportBpmnFile(bpmnImportItem)
	suite.Require().NoError(err)
	suite.NotEmpty(processCollection1.Meta)
	suite.Equal(2, len(processCollection1.Data))
}
