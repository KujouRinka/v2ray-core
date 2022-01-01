package base

// RootCommand is the root command of all commands
var RootCommand *Command

func init() {
	RootCommand = &Command{
		UsageLine: CommandEnv.Exec, // register v2ray binary name to RootCommand.
		Long:      "The root command",
	}
}

// RegisterCommand register a command to RootCommand
func RegisterCommand(cmd *Command) {
	RootCommand.Commands = append(RootCommand.Commands, cmd)
}
