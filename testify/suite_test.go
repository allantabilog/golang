package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite                   // struct embedding: promotes all properties of suite.Suite to ExampleTestSuite
	VariableThatShouldStartAtFive int
}

// ensure that 	VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ExampleTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

// all methods that begin with "Test" are run as tests within a suite
func (suite *ExampleTestSuite) TestExample() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
	suite.Equal(5, suite.VariableThatShouldStartAtFive)
}

func (suite *ExampleTestSuite) TestExample2() {
	assert.Equal(suite.T(), true, !(false))
}

func (suite *ExampleTestSuite) TestExample3() {
	assert.Equal(suite.T(), 3*100/2, 150)
}

// in order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run()
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
