package slashCommand

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mayadata-io/chat-go/testRequest"
)

func TestHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	var VerificationToken string
	var VerificationTokenKey = "VERIFICATION_TOKEN"
	var SlashCommandValue string
	var SlashCommandKey = "SLASH_COMMAND"
	tests := []struct {
		name   string
		args   args
		before func()
		after  func()
	}{
		{
			name: "When request is nil",
			args: args{
				w: httptest.NewRecorder(),
				r: testRequest.EmptyReq,
			},
			before: func() {},
			after:  func() {},
		},
		{
			name: "when verification token is wrong",
			args: args{
				w: httptest.NewRecorder(),
				r: testRequest.ReqWithWrongVerificationToken,
			},
			before: func() {},
			after:  func() {},
		},
		{
			name: "when verification token is correct and command is wrong",
			args: args{
				w: httptest.NewRecorder(),
				r: testRequest.ReqWithWrongVerificationToken,
			},
			before: func() {
				SlashCommandValue = os.Getenv(SlashCommandKey)
				err := os.Setenv(SlashCommandKey, "")
				if err != nil {
					t.Fatal(err)
				}
				VerificationToken = os.Getenv(VerificationTokenKey)
				err = os.Setenv(VerificationTokenKey, testRequest.TokenValue)
				if err != nil {
					t.Fatal(err)
				}
			},
			after: func() {
				err := os.Setenv(SlashCommandKey, SlashCommandValue)
				if err != nil {
					t.Fatal(err)
				}
				err = os.Setenv(VerificationTokenKey, VerificationToken)
				if err != nil {
					t.Fatal(err)
				}
			},
		},
		{
			name: "when both verification token and command are correct but text is empty",
			args: args{
				w: httptest.NewRecorder(),
				r: testRequest.ReqWithWrongVerificationToken,
			},
			before: func() {
				VerificationToken = os.Getenv(VerificationTokenKey)
				err := os.Setenv(VerificationTokenKey, testRequest.TokenValue)
				if err != nil {
					t.Fatal(err)
				}
			},
			after: func() {
				err := os.Setenv(VerificationTokenKey, VerificationToken)
				if err != nil {
					t.Fatal(err)
				}
			},
		},
		{
			name: "when both verification token and command are correct and text is help",
			args: args{
				w: httptest.NewRecorder(),
				r: testRequest.ReqWithHelpText,
			},
			before: func() {
				VerificationToken = os.Getenv(VerificationTokenKey)
				err := os.Setenv(VerificationTokenKey, testRequest.TokenValue)
				if err != nil {
					t.Fatal(err)
				}
			},
			after: func() {
				err := os.Setenv(VerificationTokenKey, VerificationToken)
				if err != nil {
					t.Fatal(err)
				}
			},
		},
		{
			name: "when both verification token and command are correct and text is Get Clusters",
			args: args{
				w: httptest.NewRecorder(),
				r: testRequest.ReqWithGetClustersText,
			},
			before: func() {
				VerificationToken = os.Getenv(VerificationTokenKey)
				err := os.Setenv(VerificationTokenKey, testRequest.TokenValue)
				if err != nil {
					t.Fatal(err)
				}
			},
			after: func() {
				err := os.Setenv(VerificationTokenKey, VerificationToken)
				if err != nil {
					t.Fatal(err)
				}
			},
		},
		{
			name: "when both verification token and command are correct and text is Get Cluster",
			args: args{
				w: httptest.NewRecorder(),
				r: testRequest.ReqWithGetClusterText,
			},
			before: func() {
				VerificationToken = os.Getenv(VerificationTokenKey)
				err := os.Setenv(VerificationTokenKey, testRequest.TokenValue)
				if err != nil {
					t.Fatal(err)
				}
			},
			after: func() {
				err := os.Setenv(VerificationTokenKey, VerificationToken)
				if err != nil {
					t.Fatal(err)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before()
			Handler(tt.args.w, tt.args.r)
			tt.after()
		})
	}
}
