package mock

import (
	"testing"

	"gotest.tools/v3/assert"
)

type MockStudent struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func (s *MockStudent) GetName() (*string, error) {
	if s != nil {
		return &s.Name, nil
	}
	return nil, nil
}

func TestGetName(t *testing.T) {
	t.Log("TestGetName begins")

	mockStudent := MockStudent{
		Name: "Tim",
		Id:   "A1234",
	}

	name, err := mockStudent.GetName()

	assert.DeepEqual(t, err, nil)
	assert.DeepEqual(t, *name, "Tim")
}
