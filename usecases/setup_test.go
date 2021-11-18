package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestInit(t *testing.T) {
	suite.Run(t, new(GroupHandlerSuite))
}