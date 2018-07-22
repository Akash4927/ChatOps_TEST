package slashCommand

import (
	"strings"

	"github.com/mayadata-io/chat-go/response"
)

// Commands which chat server understands
var (
	Cluster  = "cluster"
	Clusters = "clusters"
	Get      = "get"
	Help     = "help"
)

// Response will be responsible for sending slash command response
func (slashCommand *SlashCommand) Response() []byte {
	if verified := slashCommand.ValidateVerificationToken(); !verified {
		return VerificationTokenMismatchResponse()
	}

	if matchCommand := slashCommand.CheckSlashCommand(); !matchCommand {
		return WrongCommandResponse()
	}

	return TextCommandResponse(slashCommand.Text)
}

// HelpCommandResponse will return help command response
func HelpCommandResponse() []byte {
	HelpCommandResponse := response.Response{
		ResponseType: response.InChannel,
		Text: ">>>*Get Clusters:* Lists all the Clusters\n" +
			"*Get Cluster <clustername>:* Show # cluster details\n" +
			"*Help:* Shows all available commands",
	}

	return response.JSONResponseOrSlashCommandErrorMessage(HelpCommandResponse)
}

// GetClusterCommandResponse will return detail of the requested cluster
func GetClusterCommandResponse() []byte {
	GetClusterCommandResponse := response.Response{
		ResponseType: response.InChannel,
		Text:         "*Get Cluster Command Response*",
	}

	return response.JSONResponseOrSlashCommandErrorMessage(GetClusterCommandResponse)
}

// GetClustersCommandResponse will return list of all the clusters available
func GetClustersCommandResponse() []byte {
	GetClustersCommandResponse := response.Response{
		ResponseType: response.InChannel,
		Text:         "*Get Clusters Command Response*",
	}

	return response.JSONResponseOrSlashCommandErrorMessage(GetClustersCommandResponse)
}

// TextCommandResponse function will deal with all the text received in the request
func TextCommandResponse(text string) []byte {
	if strings.ToLower(text) == Help {
		return HelpCommandResponse()
	}

	if strings.ToLower(text) == Get+" "+Clusters {
		return GetClustersCommandResponse()
	}

	if strings.ToLower(text) == Get+" "+Cluster {
		return GetClusterCommandResponse()
	}

	return response.SlashCommandErrorMessage
}

// VerificationTokenMismatchResponse will return verification token mismatch response
func VerificationTokenMismatchResponse() []byte {
	VerificationTokenMismatchResponse := response.Response{
		ResponseType: response.InChannel,
		Text:         "*Verification token mismatch*",
	}

	return response.JSONResponseOrSlashCommandErrorMessage(VerificationTokenMismatchResponse)
}

// WrongCommandResponse will return wrong command response
func WrongCommandResponse() []byte {
	WrongCommandResponse := response.Response{
		ResponseType: response.Ephemeral,
		Text:         "*Wrong Command*",
	}

	return response.JSONResponseOrSlashCommandErrorMessage(WrongCommandResponse)
}
