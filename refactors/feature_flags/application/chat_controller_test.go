package application

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestChatNewFeature(t *testing.T) {
	RandomPercentageGenerator = func() int {
		return rand.Intn(NEW_CHAT_ROLLOUT_PERCENTAGE)
	}
	defer restoreRandomPercentageGenerator()

	chatController := ChatController{}
	assert.Equal(t, "new chat", chatController.PullMessages())
}

func TestChatOldFeature(t *testing.T) {
	RandomPercentageGenerator = func() int {
		return rand.Intn(100-NEW_CHAT_ROLLOUT_PERCENTAGE+1) + NEW_CHAT_ROLLOUT_PERCENTAGE
	}
	defer restoreRandomPercentageGenerator()

	chatController := ChatController{}
	assert.Equal(t, "old chat", chatController.PullMessages())
}

func restoreRandomPercentageGenerator() {
	RandomPercentageGenerator = func() int {
		return rand.Intn(100)
	}
}
