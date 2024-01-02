package channel

import (
	"testing"
)

func TestChannelOne(t *testing.T) {
	t.Log("TestChannelOne")

	// valid convert
	res := ChannelOne()
	expected := "message from ChannelTwo"
	if res != expected {
		t.Errorf("Expected value: %v, Actual: %v", res, expected)
	}
}
