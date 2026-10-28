package pingcommand

type PingCommand struct{}

func NewPingCommand() *PingCommand {
	return &PingCommand{}
}

func (c *PingCommand) Execute() string {
	return "Pong!"
}
