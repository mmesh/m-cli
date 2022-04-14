package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/client/event"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show(u *iam.User) {
	output.SectionHeader("IAM: User Details")
	output.TitleT1("User Information")

	ShowUser(u)
}

func ShowUser(u *iam.User) {
	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(u.AccountID))
	t.AddRow(colors.Black("User ID"), colors.DarkWhite(u.UserID))
	t.AddRow(colors.Black("Username"), colors.DarkWhite(u.Username))
	t.AddRow(colors.Black("Email"), colors.DarkWhite(u.Email))
	if len(u.NewEmail) > 0 {
		t.AddRow(colors.Black("New Email Requested"), colors.DarkWhite(u.NewEmail))
	}
	if u.EmailVerified {
		t.AddRow(colors.Black("Email Verified"), output.StrEnabled("yes"))
	} else {
		t.AddRow(colors.Black("Email Verified"), output.StrDisabled("no"))
	}

	var sshAuthKeyName string
	sshAuth := output.StrDisabled("disabled")
	if u.Credentials.SSH != nil {
		if u.Credentials.SSH.Key != nil && u.Credentials.SSH.Enabled {
			sshAuthKeyName = u.Credentials.SSH.Key.Name
			sshAuth = output.StrEnabled("enabled")
		}
	}
	t.AddRow(colors.Black("SSH Authentication"), sshAuth)

	// t.AddRow(colors.Black("Password"), colors.DarkWhite("********************"))

	if u.Credentials.TOTP.Enabled {
		t.AddRow(colors.Black("2-Factor Authentication"), output.StrEnabled("enabled"))
	} else {
		t.AddRow(colors.Black("2-Factor Authentication"), output.StrDisabled("disabled"))
	}

	if u.Status.Enabled {
		t.AddRow(colors.Black("User Status"), output.StrEnabled("enabled"))
	} else {
		t.AddRow(colors.Black("User Status"), output.StrDisabled("disabled"))
	}

	t.Render()
	fmt.Println()

	if u.SSHKeys != nil {
		output.SubTitleT2("SSH Keys")

		t := table.New()
		t.Header(colors.Black("NAME"), colors.Black("AUTH METHOD"), colors.Black("PUBLIC KEY"))

		for _, k := range u.SSHKeys {
			sshAuthKey := output.StrInactive("no")
			if sshAuthKeyName == k.Name {
				sshAuthKey = output.StrEnabled("yes")
			}
			t.AddRow(colors.DarkWhite(k.Name), sshAuthKey, output.Fit(k.PublicKey, 40))
		}

		t.Render()
		fmt.Println()
	}

	if u.RBAC != nil {
		if len(u.RBAC.SecurityGroups) > 0 {
			output.SubTitleT2("Security Groups")

			for _, sg := range u.RBAC.SecurityGroups {
				fmt.Printf(" ■ %s\n", colors.DarkGreen(sg))
			}
			fmt.Println()
		}

		if len(u.RBAC.ACLs) > 0 {
			output.SubTitleT2("ACLs")

			for _, acl := range u.RBAC.ACLs {
				fmt.Printf(" ■ %s\n", colors.DarkGreen(acl))
			}
			fmt.Println()
		}

		if len(u.RBAC.Roles) > 0 {
			output.SubTitleT2("Roles")

			for _, role := range u.RBAC.Roles {
				fmt.Printf(" ■ %s\n", colors.DarkGreen(role))
			}
			fmt.Println()
		}
	}

	if u.EventMetrics != nil {
		event.Output().ShowMetrics(u.EventMetrics)
	}
}
