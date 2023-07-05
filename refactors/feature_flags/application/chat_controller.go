package application

import "math/rand"

const NEW_CHAT_ROLLOUT_PERCENTAGE = 30

var RandomPercentageGenerator = func() int {
	return rand.Intn(100)
}

type ChatController struct{}

func (c *ChatController) PullMessages() string {
	if RandomPercentageGenerator() <= NEW_CHAT_ROLLOUT_PERCENTAGE {
		return "new chat"
	}
	return "old chat"
}
