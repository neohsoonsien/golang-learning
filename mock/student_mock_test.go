// this package is to be called from other packages
package mock

import (
	"errors"
	"testing"

	"gotest.tools/v3/assert"
)

func TestMockSetName_Success(t *testing.T) {
	t.Log("TestMockSetName_Success begins")

	mockStudent := MockStudent{
		SetNameFunc: func() error {
			return nil
		},
	}

	err := mockStudent.SetName("John")

	assert.DeepEqual(t, err, nil)
}

func TestMockSetName_Failed(t *testing.T) {
	t.Log("TestMockSetName_Failed begins")

	mockStudent := MockStudent{
		SetNameFunc: func() error {
			return errors.New("failed to 'SetName'")
		},
	}

	err := mockStudent.SetName("John")

	assert.Error(t, err, "failed to 'SetName'")
}

func TestMockGetName_Success(t *testing.T) {
	t.Log("TestMockGetName_Success begins")

	expectedName := "Tim"
	mockStudent := MockStudent{
		GetNameFunc: func() (*string, error) {
			return &expectedName, nil
		},
	}

	name, err := mockStudent.GetName()

	assert.DeepEqual(t, err, nil)
	assert.DeepEqual(t, *name, expectedName)
}

func TestMockGetName_Failed(t *testing.T) {
	t.Log("TestMockGetName_Failed begins")

	mockStudent := MockStudent{
		GetNameFunc: func() (*string, error) {
			return nil, errors.New("failed to 'GetName'")
		},
	}

	name, err := mockStudent.GetName()

	assert.Error(t, err, "failed to 'GetName'")
	assert.Assert(t, name == nil)
}