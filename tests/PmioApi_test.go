package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestMPIOClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
