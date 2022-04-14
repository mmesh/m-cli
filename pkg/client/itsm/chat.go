package itsm

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/rivo/tview"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/itsm"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-lib/pkg/mmp"
	"mmesh.dev/m-lib/pkg/mmp/cli"
	"mmesh.dev/m-lib/pkg/mmp/protos/chat"
	"mmesh.dev/m-lib/pkg/utils/msg"
	"mmesh.dev/m-lib/pkg/version"
)

func (api *API) Chat() {
	nxnc, grpcConn := grpc.GetNetworkAPIClient()
	defer grpcConn.Close()

	var wg sync.WaitGroup
	chatWaitc := make(chan struct{}, 2)

	wg.Add(1)
	go grpc.Control(nxnc, &wg, chatWaitc)

	// authKey, err := auth.GetAuthKey()
	// if err != nil {
	// 	msg.Alert("Invalid or inexistent authorization key. Login to refresh your token.")
	// 	os.Exit(1)
	// }

	// output.SectionHeader("mmesh service management")

	i := newChatSession()

	// mmp.NewChatInit(authKey, issue, mmp.TxControlQueue)
	chat.NewChatInit(i, mmp.TxControlQueue)

	fmt.Println()

	cli.Connected()

	output.CmdLog("Opening chat session via mmesh network..")

	go execClose(chatWaitc)

	chat.View = chat.NewView()

	hl1 := fmt.Sprintf("[gray::]%s [gray::]%s", version.CLI_NAME, version.GetVersion())
	hl2 := tview.TranslateANSI(output.AppHeader(i.AccountID, true))
	hdr := fmt.Sprintf("%s\n%s", hl1, hl2)

	chat.View.Header.SetText(hdr)

	Output().ChatIssueInfo(chat.View.IssueInfo, i)

	chat.View.Start()

	wg.Wait()
	chat.View.Stop()
}

func execClose(waitc chan struct{}) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-c
		switch sig {
		case os.Interrupt:
		case syscall.SIGINT:
		case syscall.SIGTERM:
		}

		msg.Info("Closing connection handlers...")
		waitc <- struct{}{}
	}()
}

func newChatSession() *itsm.Issue {
	opts := []string{
		"New ticket",
		"Existing ticket follow-up",
	}

	optID := input.GetSelect("Select option:", "", opts, survey.Required)

	var i *itsm.Issue

	if optID == opts[0] {
		a := account.GetAccount()

		// var issueClass itsm.IssueClass
		// p := provider.GetProvider(services.ProviderType_PROFESSIONAL_SERVICES)

		// issueClass = itsm.IssueClass_EXTERNAL_SERVICE
		issueClass := itsm.IssueClass_PLATFORM

		userEmail := viper.GetString("user.email")
		if len(userEmail) == 0 {
			userEmail = input.GetInput("User Email:", "", "", input.ValidEmail)
		}

		userNickname := strings.Split(userEmail, "@")[0]

		tm := time.Now().Unix()

		i = &itsm.Issue{
			AccountID: a.AccountID,
			// IssueID: "",
			// ServiceID:  p.ServiceID,
			// ProviderID: p.ProviderID,

			OwnerUserEmail:    userEmail,
			OwnerUserNickname: userNickname,
			Summary:           "Chat Session",
			Description:       "Customer service chat session",

			Class:        issueClass,
			IssueType:    itsm.IssueType_UNSPECIFIED,
			IssueSubtype: itsm.IssueSubtype_UNSPECIFIED_CHAT,

			Status:       itsm.IssueStatus_ISSUE_OPEN,
			CreationDate: tm,
			LastModified: tm,
		}
	} else {
		i = getIssue(false)
	}

	return i
}
