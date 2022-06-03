package output

import (
	"fmt"
	"strings"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/itsm"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

const mBot string = "m-bot"

func (api *API) Show(i *itsm.Issue) {
	output.SectionHeader("Support: Ticket Details")
	output.TitleT1("Ticket Information")

	issueID := strings.Split(i.IssueID, "_")[1]

	t := table.New()
	t.BulkData([][]string{
		{colors.Black("Account ID"), colors.DarkWhite(i.AccountID)},
		{colors.Black("Ticket ID"), colors.DarkWhite(issueID)},
		{colors.Black("Service ID"), colors.DarkWhite(i.ServiceID)},
		{colors.Black("Provider ID"), colors.DarkWhite(i.ProviderID)},
		{colors.Black("Owner"), colors.DarkWhite(i.OwnerUserEmail)},
		// {colors.Black("Class"), output.StrFixedField(issueClass(i.Class))},
		{colors.Black("Type"), output.StrFixedField(issueType(i.IssueType))},
		{colors.Black("Subtype"), output.StrFixedField(issueSubtype(i.IssueSubtype))},
		{colors.Black("Creation Date"), colors.DarkWhite(time.Unix(i.CreationDate, 0).String())},
		{colors.Black("Last Modified"), colors.DarkWhite(time.Unix(i.LastModified, 0).String())},
		{colors.Black("Status"), output.StrFixedField(issueStatus(i.Status))},
		// {colors.Black("Summary"), colors.DarkWhite(i.Summary)},
		// {colors.Black("Description"), colors.DarkWhite(i.Description)},
	})

	t.AddRow(colors.Black("Summary"), colors.DarkWhite(i.Summary))
	t.AddRow(colors.Black("Description"), colors.DarkWhite(i.Description))

	t.Render()

	fmt.Println()

	// if i.Comments == nil && i.ChatThreads == nil {
	// 	return
	// }

	issueComments(i)
	// issueChatThreads(i.ChatThreads)
}

func issueComments(i *itsm.Issue) {
	if i.Crisp == nil {
		return
	}

	if i.Crisp.Comments == nil {
		return
	}

	var n int
	switch i.IssueSubtype {
	case itsm.IssueSubtype_UNSPECIFIED_CHAT:
		n = len(i.Crisp.Comments)
	case itsm.IssueSubtype_UNSPECIFIED_CUSTOMER_SERVICE:
		n = len(i.Crisp.Comments)
	default:
		n = len(i.Crisp.Comments) - 2
	}
	if n == 0 {
		return
	}

	output.SubTitleT2("Activity")

	msg := fmt.Sprintf("There are %d comments related to this ticket. Show?", n)
	if !input.GetConfirm(msg, true) {
		fmt.Println()
		return
	}

	fmt.Println()

	// sep := strings.Repeat("─", 50)
	sep := strings.Repeat("-", 50)

	var m int

	for _, c := range i.Crisp.Comments {
		m++

		switch i.IssueSubtype {
		case itsm.IssueSubtype_UNSPECIFIED_CHAT:
		case itsm.IssueSubtype_UNSPECIFIED_CUSTOMER_SERVICE:
		default:
			if m < 3 {
				continue
			}
		}

		author := colors.DarkMagenta("-")
		if c.IsCustomer {
			if len(c.Nickname) > 0 {
				author = output.UserLocal(c.Nickname)
			}
		} else {
			if len(c.Nickname) > 0 {
				author = output.UserRemote(c.Nickname)
			}
		}

		var edited string
		if c.Edited {
			edited = output.StrNormal("edited")
		}

		tm := time.Unix(c.Timestamp, 0).String()
		hdr := fmt.Sprintf("%s %s %s", colors.Black(tm), author, edited)
		fmt.Println(hdr)
		fmt.Println(string(c.Text))

		switch i.IssueSubtype {
		case itsm.IssueSubtype_UNSPECIFIED_CHAT:
			if m < n {
				fmt.Println(colors.Black(sep))
			}
		case itsm.IssueSubtype_UNSPECIFIED_CUSTOMER_SERVICE:
			if m < n {
				fmt.Println(colors.Black(sep))
			}
		default:
			if m-2 < n {
				fmt.Println(colors.Black(sep))
			}
		}
	}

	fmt.Println()
}

/*
func issueComments(i *itsm.Issue) {
	if i.Comments == nil {
		return
	}

	var tmSort []string
	for tm := range i.Comments {
		tmSort = append(tmSort, tm)
	}
	sort.Strings(tmSort)

	t := table.New()
	// t.SetRowLine("─")

	sep := strings.Repeat("─", 50)
	n := len(i.Comments)
	var m int

	for _, tm := range tmSort {
		c := i.Comments[tm]
		m++
		author := output.UserLocal(fmt.Sprintf("@%s", i.AccountID))
		if c.Author != mBot {
			author = output.UserRemote(fmt.Sprintf("%s @%s", c.Author, i.ProviderID))
		}
		timestamp := time.Unix(c.Timestamp, 0).String()
		hdr := fmt.Sprintf("%s %s", colors.Black(timestamp), author)
		t.AddRow(hdr)
		t.AddRow(fmt.Sprintf("%s", c.Text))
		if m < n {
			t.AddRow(colors.Black(sep))
		}
	}

	t.Render()
	fmt.Println()
}

func issueChatThreads(chatThreads map[int64]*messaging.ChatThread) {
	if chatThreads == nil {
		return
	}

	n := len(chatThreads)
	if n == 0 {
		return
	}

	msg := fmt.Sprintf("There are %d chat conversations related to this issue. Show?", n)
	ok := input.GetConfirm(msg, true)
	if !ok {
		fmt.Println()
		return
	}

	showChatThread(chatThreads)
}

func showChatThread(chatThreads map[int64]*messaging.ChatThread) {
	var threadOptID string
	var threadsOpts []string
	threads := make(map[string]*messaging.ChatThread)

	for _, ct := range chatThreads {
		tm := time.Unix(ct.LastUpdated, 0).String()
		nComments := fmt.Sprintf("Messages: %d", len(ct.Comments))
		threadOptID = fmt.Sprintf("%s | %s", tm, nComments)
		threadsOpts = append(threadsOpts, threadOptID)
		threads[threadOptID] = ct
	}

	sort.Strings(threadsOpts)

	threadOptID = input.GetSelect("Chat:", "", threadsOpts, survey.Required)

	t := table.New()
	// t.SetRowLine("─")

	sep := strings.Repeat("─", 50)
	n := len(threads[threadOptID].Comments)
	var m int

	for timestamp, c := range threads[threadOptID].Comments {
		m++
		author := colors.DarkMagenta("-")
		if c.IsCustomer {
			if len(c.Nickname) > 0 {
				author = output.UserLocal(c.Nickname)
			}
		} else {
			if len(c.Nickname) > 0 {
				author = output.UserRemote(c.Nickname)
			}
		}
		tm := time.Unix(timestamp, 0).String()
		hdr := fmt.Sprintf("%s %s", colors.Black(tm), author)
		t.AddRow(hdr)
		t.AddRow(fmt.Sprintf("%s", c.Text))
		if m < n {
			t.AddRow(colors.Black(sep))
		}
	}

	fmt.Println()

	t.Render()
	fmt.Println()
}
*/

func issueStatus(s itsm.IssueStatus) string {
	switch s {
	case itsm.IssueStatus_ISSUE_OPEN:
		return "OPEN"
	case itsm.IssueStatus_ISSUE_CLOSED:
		return "CLOSED"
	}

	return "-"
}

func issueClass(c itsm.IssueClass) string {
	switch c {
	case itsm.IssueClass_PLATFORM:
		return "PLATFORM"
	case itsm.IssueClass_EXTERNAL_SERVICE:
		return "EXTERNAL SERVICE"
	}

	return "-"
}

func issueType(t itsm.IssueType) string {
	switch t {
	case itsm.IssueType_ASSISTANCE:
		return "ASSISTANCE"
	case itsm.IssueType_FEEDBACK:
		return "FEEDBACK"
	case itsm.IssueType_INCIDENT:
		return "INCIDENT"
	case itsm.IssueType_PROBLEM:
		return "PROBLEM"
	case itsm.IssueType_CHANGE:
		return "CHANGE"
	case itsm.IssueType_REQUEST:
		return "REQUEST"

	case itsm.IssueType_UNSPECIFIED:
		return "OTHER"
	}

	return "-"
}

func issueSubtype(st itsm.IssueSubtype) string {
	switch st {
	case itsm.IssueSubtype_ASSISTANCE_INFO:
		return "INFO"
	case itsm.IssueSubtype_ASSISTANCE_SUPPORT:
		return "SUPPORT"

	case itsm.IssueSubtype_FEEDBACK_GENERAL_COMMENTS:
		return "COMMENTS"
	case itsm.IssueSubtype_FEEDBACK_FEATURE_REQUEST:
		return "FEATURE REQUEST"

	case itsm.IssueSubtype_INCIDENT_BILLING:
		return "BILLING"
	case itsm.IssueSubtype_INCIDENT_PLATFORM:
		return "PLATFORM"

	case itsm.IssueSubtype_PROBLEM_PLATFORM:
		return "PLATFORM"

	case itsm.IssueSubtype_CHANGE_INTERNAL_REQUEST:
		return "INTERNAL REQUEST"

	case itsm.IssueSubtype_REQUEST_UNCATEGORIZED:
		return "UNCATEGORIZED"
	case itsm.IssueSubtype_REQUEST_ADVISORY_SERVICE:
		return "ADVISORY SERVICE"
	case itsm.IssueSubtype_REQUEST_SPECIAL_PROJECT:
		return "SPECIAL PROJECT"
	case itsm.IssueSubtype_REQUEST_SPECIAL_TASK:
		return "SPECIAL TASK"
	case itsm.IssueSubtype_REQUEST_SPECIFIC_FUNCTIONALITY:
		return "SPECIFIC FUNCTIONALITY"
	case itsm.IssueSubtype_REQUEST_EXTERNAL_SERVICE:
		return "EXTERNAL SERVICE"

	case itsm.IssueSubtype_UNSPECIFIED_CHAT:
		return "CHAT"
	case itsm.IssueSubtype_UNSPECIFIED_CUSTOMER_SERVICE:
		return "CUSTOMER SERVICE"
	case itsm.IssueSubtype_UNSPECIFIED_CHAT_EXTERNAL_SERVICE:
		return "CHAT EXTERNAL SERVICE"
	}

	return "-"
}
