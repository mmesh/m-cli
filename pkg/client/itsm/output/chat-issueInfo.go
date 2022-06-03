package output

import (
	"fmt"
	"io"

	"mmesh.dev/m-api-go/grpc/resources/itsm"
	"mmesh.dev/m-cli/pkg/output"
)

func (api *API) ChatIssueInfo(w io.Writer, i *itsm.Issue) {
	if len(i.IssueID) == 0 {
		i.IssueID = "-"
	}

	fmt.Fprintln(w)
	fmt.Fprintf(w, "[gray::]Account ID\n")
	fmt.Fprintf(w, " [silver::b]%s\n", i.AccountID)
	fmt.Fprintf(w, "[gray::]Ticket ID\n")
	fmt.Fprintf(w, " [silver::b]%s\n", i.IssueID)
	fmt.Fprintf(w, "[gray::]Service ID\n")
	fmt.Fprintf(w, " [silver::b]%s\n", i.ServiceID)
	fmt.Fprintf(w, "[gray::]Provider ID\n")
	fmt.Fprintf(w, " [silver::b]%s\n", i.ProviderID)
	fmt.Fprintf(w, "[gray::]Class\n")
	fmt.Fprintf(w, " [purple::][[fuchsia::b]%s[purple::b]]\n", issueClass(i.Class))
	fmt.Fprintf(w, "[gray::]Type\n")
	fmt.Fprintf(w, " [purple::][[fuchsia::b]%s[purple::b]]\n", issueType(i.IssueType))
	fmt.Fprintf(w, "[gray::]Subtype\n")
	fmt.Fprintf(w, " [purple::][[fuchsia::b]%s[purple::b]]\n", issueSubtype(i.IssueSubtype))
	fmt.Fprintf(w, "[gray::]Status\n")
	fmt.Fprintf(w, " [purple::][[fuchsia::b]%s[purple::b]]\n", issueStatus(i.Status))

	fmt.Fprintf(w, "[gray::]Creation Date\n")
	fmt.Fprintf(w, " [silver::b]%s\n", output.Datetime(i.CreationDate))
	fmt.Fprintf(w, "[gray::]Last Modified\n")
	fmt.Fprintf(w, " [silver::b]%s\n", output.Datetime(i.LastModified))

	fmt.Fprintln(w)

	fmt.Fprintf(w, "[gray::]Summary\n")
	fmt.Fprintf(w, "[silver::b]%s\n", i.Summary)
	fmt.Fprintln(w)
	fmt.Fprintf(w, "[gray::]Description\n")
	fmt.Fprintf(w, "[silver::b]%s\n", i.Description)

	fmt.Fprintln(w)
}
