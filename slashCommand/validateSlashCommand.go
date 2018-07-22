package slashCommand

import (
	"os"

	"github.com/mayadata-io/chat-go/environment"
)

// ValidateVerificationToken compare token received from request
// and token stored in environment variable, if it matches then
// it returns true else false
func (slashCommand *SlashCommand) ValidateVerificationToken() bool {
	if slashCommand.Token == os.Getenv(environment.VerificationToken) {
		return true
	}
	return false
}

// CheckSlashCommand compare command received from request
// and command stored in environment variable, if it matches then
// it returns true else false
func (slashCommand *SlashCommand) CheckSlashCommand() bool {
	if slashCommand.Command == os.Getenv(environment.SlashCommand) {
		return true
	}
	return false
}
