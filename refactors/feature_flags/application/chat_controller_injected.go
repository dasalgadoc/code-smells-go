package application

type ChatControllerInjected struct {
	newChatRolloutPercentage int
}

func (c *ChatControllerInjected) PullMessages() string {
	if RandomPercentageGenerator() <= c.newChatRolloutPercentage {
		return "new chat"
	}
	return "old chat"
}

func NewChatControllerInjected(newChatRolloutPercentage int) *ChatControllerInjected {
	return &ChatControllerInjected{newChatRolloutPercentage: newChatRolloutPercentage}
}
