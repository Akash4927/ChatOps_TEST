package testRequest

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

// request type and API URL
var (
	Post  = "POST"
	Get   = "GET"
	Slash = "/slash"
)

// Slash Command Request key fields
var (
	ChannelID   = "channel_id"
	ChannelName = "channel_name"
	Command     = "command"
	ResponseURL = "response_url"
	TeamDomain  = "team_domain"
	TeamID      = "team_id"
	Text        = "text"
	Token       = "token"
	TriggerID   = "trigger_id"
	UserID      = "user_id"
	UserName    = "user_name"
)

// Slash command request value
var (
	ChannelIDValue   = "C123456"
	ChannelNameValue = "general"
	CommandValue     = "/test"
	EmptyText        = ""
	GetClusterText   = "Get Cluster"
	GetClustersText  = "Get Clusters"
	HelpText         = "HELP"
	ResponseURLValue = "http://localhost:80"
	TeamDomainValue  = "test"
	TeamIDValue      = "T123456"
	TokenValue       = "abcdefghij"
	TriggerIDValue   = "3160.123456789.89128918291"
	UserIDValue      = "U123456"
	UserNameValue    = "test"
)

// Request Objects
var (
	EmptyReq                      *http.Request
	ReqWithWrongVerificationToken *http.Request
	ReqWithHelpText               *http.Request
	ReqWithGetClusterText         *http.Request
	ReqWithGetClustersText        *http.Request
)

var err error

func init() {
	EmptyReq, err = http.NewRequest(Post, Slash, nil)
	if err != nil {
		log.Fatal(err)
	}

	emptyTextForm := url.Values{}
	emptyTextForm.Add(ChannelID, ChannelIDValue)
	emptyTextForm.Add(ChannelName, ChannelNameValue)
	emptyTextForm.Add(Command, CommandValue)
	emptyTextForm.Add(ResponseURL, ResponseURLValue)
	emptyTextForm.Add(TeamDomain, TeamDomainValue)
	emptyTextForm.Add(TeamID, TeamIDValue)
	emptyTextForm.Add(Text, EmptyText)
	emptyTextForm.Add(Token, TokenValue)
	emptyTextForm.Add(TriggerID, TriggerIDValue)
	emptyTextForm.Add(UserID, UserIDValue)
	emptyTextForm.Add(UserName, UserNameValue)

	ReqWithWrongVerificationToken, err = http.NewRequest(Post, Slash, strings.NewReader(emptyTextForm.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	ReqWithWrongVerificationToken.PostForm = emptyTextForm

	HelpTextForm := url.Values{}
	HelpTextForm.Add(ChannelID, ChannelIDValue)
	HelpTextForm.Add(ChannelName, ChannelNameValue)
	HelpTextForm.Add(Command, CommandValue)
	HelpTextForm.Add(ResponseURL, ResponseURLValue)
	HelpTextForm.Add(TeamDomain, TeamDomainValue)
	HelpTextForm.Add(TeamID, TeamIDValue)
	HelpTextForm.Add(Text, HelpText)
	HelpTextForm.Add(Token, TokenValue)
	HelpTextForm.Add(TriggerID, TriggerIDValue)
	HelpTextForm.Add(UserID, UserIDValue)
	HelpTextForm.Add(UserName, UserNameValue)

	ReqWithHelpText, err = http.NewRequest(Post, Slash, strings.NewReader(HelpTextForm.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	ReqWithHelpText.PostForm = HelpTextForm

	GetClusterForm := url.Values{}
	GetClusterForm.Add(ChannelID, ChannelIDValue)
	GetClusterForm.Add(ChannelName, ChannelNameValue)
	GetClusterForm.Add(Command, CommandValue)
	GetClusterForm.Add(ResponseURL, ResponseURLValue)
	GetClusterForm.Add(TeamDomain, TeamDomainValue)
	GetClusterForm.Add(TeamID, TeamIDValue)
	GetClusterForm.Add(Text, GetClusterText)
	GetClusterForm.Add(Token, TokenValue)
	GetClusterForm.Add(TriggerID, TriggerIDValue)
	GetClusterForm.Add(UserID, UserIDValue)
	GetClusterForm.Add(UserName, UserNameValue)

	ReqWithGetClusterText, err = http.NewRequest(Post, Slash, strings.NewReader(GetClusterForm.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	ReqWithGetClusterText.PostForm = GetClusterForm

	GetClustersForm := url.Values{}
	GetClustersForm.Add(ChannelID, ChannelIDValue)
	GetClustersForm.Add(ChannelName, ChannelNameValue)
	GetClustersForm.Add(Command, CommandValue)
	GetClustersForm.Add(ResponseURL, ResponseURLValue)
	GetClustersForm.Add(TeamDomain, TeamDomainValue)
	GetClustersForm.Add(TeamID, TeamIDValue)
	GetClustersForm.Add(Text, GetClustersText)
	GetClustersForm.Add(Token, TokenValue)
	GetClustersForm.Add(TriggerID, TriggerIDValue)
	GetClustersForm.Add(UserID, UserIDValue)
	GetClustersForm.Add(UserName, UserNameValue)

	ReqWithGetClustersText, err = http.NewRequest(Post, Slash, strings.NewReader(GetClustersForm.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	ReqWithGetClustersText.PostForm = GetClustersForm
}
