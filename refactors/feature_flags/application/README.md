# 🏁 Feature flags

Feature flags are a way of enabling or disabling functionality in an application. They are useful for a number of reasons, including:

- Allowing developers to merge unfinished or untested code into the main codebase without affecting the application's behavior.
- Allow to perform A/B testing.
- Allow to perform canary releases.

There is a lot of ways to implement feature flags, here the most common:

## ⚡️ Application parameters

The simplest way to implement a feature flag is to redirect a portion of the traffic to a different features. A naive implementation would be to use a random number generator to decide which feature to use.

```go
func (c *ChatController) PullMessages() string {
    if rand.Intn(100) <= NEW_CHAT_ROLLOUT_PERCENTAGE {
        return "new chat"
    }
	
    return "old chat"
}
```

In this code, each petition has an x% chance to use the new chat determined by `NEW_CHAT_ROLLOUT_PERCENTAGE`.

This approach has a problem with deterministic tests, because we can't exactly predict the result (without using mocks).
See `chat_controller_test.go` to see how test can be written in feature flags.

## 🧪 Injected application parameter 

```go
type ChatControllerInjected struct {
    newChatRolloutPercentage int
}

func (c *ChatControllerInjected) PullMessages() string {
    if RandomPercentageGenerator() <= c.newChatRolloutPercentage {
        return "new chat"
    }
	
    return "old chat"
}
```

## 📚 Bring from repository or environment variable

