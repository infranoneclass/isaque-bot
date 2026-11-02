package ports

type ConfigRepository interface {
	BindFromJSONFile(file []byte) error
	BindBotToken(token string)

	GetIfHoneypotOnlyKicks() bool
}
