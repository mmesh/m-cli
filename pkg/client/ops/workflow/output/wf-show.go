package output

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/ops/workflow"
	"mmesh.dev/m-cli/pkg/client/event"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) Show(wf *workflow.Workflow) {
	output.SectionHeader("Ops: Workflow Details")
	output.TitleT1("Workflow Information")

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(wf.AccountID))
	t.AddRow(colors.Black("Project ID"), colors.DarkWhite(wf.ProjectID))
	t.AddRow(colors.Black("Workflow ID"), colors.DarkWhite(wf.WorkflowID))

	t.AddRow(colors.Black("Description"), colors.DarkWhite(wf.Description))

	if wf.Owner != nil {
		t.AddRow(colors.Black("Owner"), colors.DarkWhite(wf.Owner.Email))
	}

	if wf.Enabled {
		t.AddRow(colors.Black("Enabled"), output.StrEnabled("yes"))
	} else {
		t.AddRow(colors.Black("Enabled"), output.StrDisabled("no"))
	}

	if wf.Reviewed {
		t.AddRow(colors.Black("Reviewed"), output.StrEnabled("yes"))
	} else {
		t.AddRow(colors.Black("Reviewed"), output.StrDisabled("no"))
	}

	if wf.Approved {
		t.AddRow(colors.Black("Approved"), output.StrEnabled("yes"))
	} else {
		t.AddRow(colors.Black("Approved"), output.StrDisabled("no"))
	}

	t.Render()
	fmt.Println()

	if wf.Triggers != nil {
		output.SubTitleT2("Triggers")

		t = table.New()

		if wf.Triggers.Webhook.Enabled {
			t.AddRow(colors.Black("Workflow Webhook"), output.StrEnabled("enabled"))
			t.Render()

			// /api/v1/accounts/{accountToken}/webhooks/workflow/{wfToken}
			srv := viper.GetString("controller.authServer")
			if len(srv) > 0 {
				whURL := fmt.Sprintf("%s/api/v1/accounts/%s/webhooks/workflow/%s", srv, wf.AccountToken, wf.Token)
				fmt.Printf("\n%s %s\n", colors.Black("Workflow Webhook URL"), colors.DarkWhite(whURL))
			}
		} else {
			t.AddRow(colors.Black("Workflow Webhook"), output.StrDisabled("disabled"))
			t.Render()
		}

		fmt.Println()

		if wf.Triggers.GitEvents != nil {
			fmt.Printf(colors.Black("Git Events\n"))
			for _, gitEvt := range wf.Triggers.GitEvents.Events {
				fmt.Printf(" ■ %s\n", colors.DarkGreen(gitEvt))
			}
			fmt.Println()
		}

		if wf.Triggers.Schedule != nil {
			t = table.New()

			if len(wf.Triggers.Schedule.Crontab) > 0 {
				if wf.Triggers.Schedule.Enabled {
					t.AddRow(colors.Black("Workflow Crontab"), output.StrEnabled("enabled"), colors.DarkWhite(wf.Triggers.Schedule.Crontab))
				} else {
					t.AddRow(colors.Black("Workflow Crontab"), output.StrDisabled("disabled"), colors.DarkWhite(wf.Triggers.Schedule.Crontab))
				}
			}

			if wf.Triggers.Schedule.DateTime != nil {
				timestamp, err := utils.GetDateTime(wf.Triggers.Schedule.DateTime)
				if err != nil {
					msg.Errorf("Unable to get workflow schedule: %v", err)
					os.Exit(1)
				}
				tm := time.Unix(timestamp, 0)

				if wf.Triggers.Schedule.Enabled {
					t.AddRow(colors.Black("Workflow Schedule"), output.StrEnabled("enabled"), colors.DarkWhite(tm.String()))
				} else {
					t.AddRow(colors.Black("Workflow Schedule"), output.StrDisabled("disabled"), colors.DarkWhite(tm.String()))
				}
			}

			t.Render()
			fmt.Println()
		}
	}

	if wf.Jobs != nil {
		output.SubTitleT2("Jobs")

		t = table.New()

		for i, job := range wf.Jobs {
			jName := fmt.Sprintf("%s%s%s %s", colors.DarkMagenta("["), colors.Magenta(fmt.Sprintf("%d", i+1)), colors.DarkMagenta("]"), colors.Black("Name"))
			t.AddRow(jName, colors.DarkWhite(job.Name))
			t.AddRow(colors.Black("    Description"), colors.DarkWhite(job.Description))
			t.AddRow("", "")
			t.AddRow(colors.White("    Tasks"))
			t.AddRow(colors.Black("    ====="))
			for j, task := range job.Tasks {
				tName := fmt.Sprintf("    %s%s%s %s", colors.DarkMagenta("["), colors.Magenta(fmt.Sprintf("%d", j+1)), colors.DarkMagenta("]"), colors.Black("Name"))
				t.AddRow(tName, colors.DarkWhite(task.Name))
				t.AddRow(colors.Black("        Description"), colors.DarkWhite(task.Description))
				t.AddRow("", "")
				// t.AddRow(fmt.Sprintf("        %s", colors.InvertedBlack("Action Name")), colors.InvertedBlack("Command"), colors.InvertedBlack("Args"), colors.InvertedBlack("UID"), colors.InvertedBlack("GID"))
				t.AddRow(fmt.Sprintf("        %s", colors.Black("Action Name")), colors.Black("Command"), colors.Black("Args"), colors.Black("UID"), colors.Black("GID"))
				t.AddRow(fmt.Sprintf("        %s", colors.Black("-----------")), colors.Black("-------"), colors.Black("----"), colors.Black("---"), colors.Black("---"))
				for _, action := range task.Actions {
					actName := fmt.Sprintf("        %s", colors.DarkWhite(action.Name))
					actCmd := colors.DarkWhite(action.Command.Cmd)
					actArgs := colors.DarkWhite(strings.Join(action.Command.Args, " "))
					actUID := colors.DarkWhite(fmt.Sprintf("%d", action.Command.UID))
					actGID := colors.DarkWhite(fmt.Sprintf("%d", action.Command.GID))
					t.AddRow(actName, actCmd, actArgs, actUID, actGID)
				}
				t.AddRow("", "")
			}
		}

		t.Render()
		// fmt.Println()
	}

	if len(wf.Targets) > 0 {
		output.SubTitleT2("Targets")

		t = table.New()
		// t.Header(colors.Black("TENANT"), colors.Black("NETWORK"), colors.Black("SUBNET"), colors.Black("NODE"))
		// t.Header(colors.Black("Tenant"), colors.Black("Network"), colors.Black("Subnet"), colors.Black("Node"))

		// t.AddRow(colors.InvertedBlack("Tenant"), colors.InvertedBlack("Network"), colors.InvertedBlack("Subnet"), colors.InvertedBlack("Node"))
		t.AddRow(colors.Black("Tenant"), colors.Black("Network"), colors.Black("Subnet"), colors.Black("Node"))
		t.AddRow(colors.Black("------"), colors.Black("-------"), colors.Black("------"), colors.Black("----"))

		for _, n := range wf.Targets {
			t.AddRow(
				colors.DarkWhite(n.TenantID),
				colors.DarkWhite(n.NetID),
				colors.DarkWhite(n.VRFID),
				colors.DarkWhite(output.Fit(n.NodeID, 32)),
				// colors.DarkWhite(" "),
			)
		}

		t.Render()
		fmt.Println()
	}

	if wf.Notify != nil {
		output.SubTitleT2("Notifications")

		t = table.New()
		// t.Header(colors.Black("RECIPIENT"), colors.Black("CHANNEL"))
		// t.Header(colors.Black("Recipient"), colors.Black("Channel"))

		// t.AddRow(colors.InvertedBlack("Recipient"), colors.InvertedBlack("Channel"))
		t.AddRow(colors.Black("Recipient"), colors.Black("Channel"))
		t.AddRow(colors.Black("---------"), colors.Black("-------"))

		if wf.Notify.Email != nil {
			for _, r := range wf.Notify.Email.Recipients {
				t.AddRow(colors.DarkWhite(r.Email), output.StrNormal("email"))
			}
		}

		if wf.Notify.Slack != nil {
			for _, r := range wf.Notify.Slack.Recipients {
				t.AddRow(colors.DarkWhite(r.Email), output.StrNormal("slack"))
			}
		}

		t.Render()
		fmt.Println()
	}

	if len(wf.Reviewers) > 0 {
		output.SubTitleT2("Reviewers")

		t = table.New()
		// t.Header(colors.Black("RESPONSIBLE"), colors.Black("VALIDATION"), colors.Black("TIMESTAMP"))
		// t.Header(colors.Black("Responsible"), colors.Black("Validation"), colors.Black("Timestamp"))

		// t.AddRow(colors.InvertedBlack("Responsible"), colors.InvertedBlack("Validation"), colors.InvertedBlack("Timestamp"))
		t.AddRow(colors.Black("Responsible"), colors.Black("Validation"), colors.Black("Timestamp"))
		t.AddRow(colors.Black("-----------"), colors.Black("----------"), colors.Black("---------"))

		for _, reviewer := range wf.Reviewers {
			for _, rby := range wf.ReviewedBy {
				if rby.Responsible.Email == reviewer {
					tm := time.Unix(rby.Timestamp, 0)
					if rby.Validated {
						t.AddRow(colors.DarkWhite(reviewer), output.StrEnabled("REVIEWED"), colors.DarkWhite(tm.String()))
					} else {
						t.AddRow(colors.DarkWhite(reviewer), output.StrDisabled("REJECTED"), colors.DarkWhite(tm.String()))
					}
				}
			}
		}

		t.Render()
		fmt.Println()
	}

	if len(wf.Approvers) > 0 {
		output.SubTitleT2("Approvers")

		t = table.New()
		// t.Header(colors.Black("RESPONSIBLE"), colors.Black("VALIDATION"), colors.Black("TIMESTAMP"))
		// t.Header(colors.Black("Responsible"), colors.Black("Validation"), colors.Black("Timestamp"))

		// t.AddRow(colors.InvertedBlack("Responsible"), colors.InvertedBlack("Validation"), colors.InvertedBlack("Timestamp"))
		t.AddRow(colors.Black("Responsible"), colors.Black("Validation"), colors.Black("Timestamp"))
		t.AddRow(colors.Black("-----------"), colors.Black("----------"), colors.Black("---------"))

		for _, approver := range wf.Approvers {
			for _, aby := range wf.ApprovedBy {
				if aby.Responsible.Email == approver {
					tm := time.Unix(aby.Timestamp, 0)
					if aby.Validated {
						t.AddRow(colors.DarkWhite(approver), output.StrEnabled("APPROVED"), colors.DarkWhite(tm.String()))
					} else {
						t.AddRow(colors.DarkWhite(approver), output.StrDisabled("REJECTED"), colors.DarkWhite(tm.String()))
					}
				}
			}
		}

		t.Render()
		fmt.Println()
	}

	if wf.EventMetrics != nil {
		event.Output().ShowMetrics(wf.EventMetrics)
	}
}
