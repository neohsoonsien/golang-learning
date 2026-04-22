package mock

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestSetName(t *testing.T) {
	t.Log("TestSetName begins")

	student := Student{}

	err := student.SetName("John")

	assert.DeepEqual(t, err, nil)
	assert.DeepEqual(t, student, Student{Name: "John"})
}

func TestGetName(t *testing.T) {
	t.Log("TestGetName begins")

	student := Student{
		Name: "Tim",
		Id:   "A1234",
	}

	name, err := student.GetName()

	assert.DeepEqual(t, err, nil)
	assert.DeepEqual(t, *name, "Tim")
}

func TestSetId(t *testing.T) {
	t.Log("TestSetId begins")

	student := Student{}

	err := student.SetId("A1234")

	assert.DeepEqual(t, err, nil)
	assert.DeepEqual(t, student, Student{Id: "A1234"})
}

func TestGetId(t *testing.T) {
	t.Log("TestGetId begins")

	student := Student{
		Name: "Tim",
		Id:   "A1234",
	}

	id, err := student.GetId()

	assert.DeepEqual(t, err, nil)
	assert.DeepEqual(t, *id, "A1234")
}
