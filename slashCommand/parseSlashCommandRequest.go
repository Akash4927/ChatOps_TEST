package slashCommand

import "net/http"

type SlashCommand struct {
	ChannelID   string
	ChannelName string
	Command     string
	ResponseURL string
	TeamDomain  string
	TeamID      string
	Text        string
	Token       string
	TriggerID   string
	UserID      string
	UserName    string
}

func ParseSlashCommandRequest(r *http.Request) (SlashCommand, error) {
	var slashCommand SlashCommand
	if err := r.ParseForm(); err != nil {
		return slashCommand, err
	}

	slashCommand.ChannelID = r.PostFormValue("channel_id")
	slashCommand.ChannelName = r.PostFormValue("channel_name")
	slashCommand.Command = r.PostFormValue("command")
	slashCommand.ResponseURL = r.PostFormValue("response_url")
	slashCommand.TeamDomain = r.PostFormValue("team_domain")
	slashCommand.TeamID = r.PostFormValue("team_id")
	slashCommand.Text = r.PostFormValue("text")
	slashCommand.Token = r.PostFormValue("token")
	slashCommand.TriggerID = r.PostFormValue("trigger_id")
	slashCommand.UserID = r.PostFormValue("user_id")
	slashCommand.UserName = r.PostFormValue("user_name")

	return slashCommand, nil
}
