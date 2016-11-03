package gorocket

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRocket_SendAndReceive(t *testing.T) {
	rocket := getDefaultClient(t)

	rooms, err := rocket.GetPublicRooms()
	assert.Nil(t, err)

	general := getRoom(rooms, "general")

	err = rocket.Send(general, "Test")
	assert.Nil(t, err)

	messages, err := rocket.GetMessages(general, nil)
	assert.Nil(t, err)

	message := findMessage(messages, testUserName, "Test")
	assert.NotNil(t, message)
}