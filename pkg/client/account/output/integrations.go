package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/services/thirdParty"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func showIntegrations(i *thirdParty.Integrations) {
	if i == nil {
		return
	}

	if i.Pagerduty == nil &&
		i.Github == nil &&
		i.Slack == nil &&
		i.DigitalOcean == nil &&
		i.GCP == nil &&
		i.Scaleway == nil {
		return
	}

	output.SubTitleT2("Integrations")

	t := table.New()

	if i.Clickup != nil {
		if i.Clickup.Enabled {
			t.AddRow(colors.Black("ClickUp"), output.StrEnabled("enabled"))
			// t.AddRow(colors.Black("ClickUp API Key"), colors.DarkWhite(i.Clickup.ApiKey), colors.Black(" "))
			// t.AddRow(colors.Black("ClickUp Settings List URL"), colors.DarkWhite(i.Clickup.SettingsListURL), colors.Black(" "))
		} else {
			t.AddRow(colors.Black("ClickUp"), output.StrDisabled("not-configured"))
		}
		// t.AddRow()
	}
	if i.Github != nil {
		if i.Github.Enabled {
			t.AddRow(colors.Black("GitHub"), output.StrEnabled("enabled"))
			// t.AddRow(colors.Black("GitHub Access Token"), colors.DarkWhite(i.Github.AccessToken), colors.Black(" "))
			// t.AddRow(colors.Black("GitHub Webhooks Secret"), colors.DarkWhite(i.Github.WebhooksSecret), colors.Black(" "))
		} else {
			t.AddRow(colors.Black("GitHub"), output.StrDisabled("not-configured"))
		}
		// t.AddRow()
	}
	if i.Pagerduty != nil {
		if i.Pagerduty.Enabled {
			t.AddRow(colors.Black("PagerDuty"), output.StrEnabled("enabled"))
			// t.AddRow(colors.Black("PagerDuty Integration Key"), colors.DarkWhite(i.Pagerduty.IntegrationKey), colors.Black(" "))
		} else {
			t.AddRow(colors.Black("PagerDuty"), output.StrDisabled("not-configured"))
		}
		// t.AddRow()
	}
	if i.Slack != nil {
		if i.Slack.Enabled {
			t.AddRow(colors.Black("Slack"), output.StrEnabled("enabled"))
			if len(i.Slack.AlertsWebhook) > 0 {
				t.AddRow(colors.DarkWhite(" • Alert Notifications"), output.StrEnabled("enabled"))
				// t.AddRow(colors.Black("Slack Alert Notifications Webhook"), colors.DarkWhite(i.Slack.AlertsWebhook), colors.Black(" "))
			} else {
				t.AddRow(colors.DarkWhite(" • Alert Notifications"), output.StrDisabled("not-configured"))
			}
			if len(i.Slack.OpsWebhook) > 0 {
				t.AddRow(colors.DarkWhite(" • General Notifications & Reports"), output.StrEnabled("enabled"))
				// t.AddRow(colors.Black("Slack General Notifications Webhook"), colors.DarkWhite(i.Slack.OpsWebhook), colors.Black(" "))
			} else {
				t.AddRow(colors.DarkWhite(" • General Notifications & Reports"), output.StrDisabled("not-configured"))
			}
		} else {
			t.AddRow(colors.Black("Slack"), output.StrDisabled("not-configured"))
		}
		// t.AddRow()
	}
	/*
		if i.Crisp != nil {
			if i.Crisp.Enabled {
				t.AddRow(colors.Black("Crisp"), output.StrEnabled("enabled"))
				// t.AddRow(colors.Black("Crisp WebsiteID"), colors.DarkWhite(i.Crisp.WebsiteID), colors.Black(" "))
			} else {
				t.AddRow(colors.Black("Crisp"), output.StrDisabled("not-configured"))
			}
			// t.AddRow()
		}
	*/
	/*
		if i.Slack != nil {
			if i.Slack.Enabled {
				t.AddRow("Slack:", output.StrEnabled("enabled"))
				t.AddRow("Slack Bot User OAuth Access Token:", colors.Black(i.Slack.BotUserOAuthAccessToken), colors.Black(" "))
				t.AddRow("Slack Verification Token:", colors.Black(i.Slack.VerificationToken), colors.Black(" "))
				t.AddRow("Slack Signing Secret:", colors.Black(i.Slack.SigningSecret), colors.Black(" "))
			} else {
				t.AddRow("Slack:", output.StrDisabled("not-configured"))
			}
		}
	*/

	if i.DigitalOcean != nil {
		if i.DigitalOcean.Enabled {
			t.AddRow(colors.Black("DigitalOcean"), output.StrEnabled("enabled"))
			// t.AddRow(colors.Black("DigitalOcean Token"), colors.DarkWhite(i.DigitalOcean.Token), colors.Black(" "))
			// t.AddRow(colors.Black("DigitalOcean Project"), colors.DarkWhite(i.DigitalOcean.ProjectName), colors.Black(" "))
		} else {
			t.AddRow(colors.Black("DigitalOcean"), output.StrDisabled("not-configured"))
		}
		// t.AddRow()
	}

	if i.GCP != nil {
		if i.GCP.Enabled {
			t.AddRow(colors.Black("Google Cloud"), output.StrEnabled("enabled"))
			// t.AddRow(colors.Black("Google Cloud ProjectID"), colors.DarkWhite(i.GCP.ProjectID), colors.Black(" "))
			// t.AddRow(colors.Black("Google Cloud JSONKey"), colors.DarkWhite("*********"))
			// t.AddRow("GCP JSON Key:", colors.Black(string(i.GCP.JSONKey)), colors.Black(" "))
		} else {
			t.AddRow(colors.Black("Google Cloud"), output.StrDisabled("not-configured"))
		}
		// t.AddRow()
	}

	if i.Scaleway != nil {
		if i.Scaleway.Enabled {
			t.AddRow(colors.Black("Scaleway Cloud"), output.StrEnabled("enabled"))
			// t.AddRow(colors.Black("Scaleway OrganizationID"), colors.DarkWhite(i.Scaleway.OrganizationID), colors.Black(" "))
			// t.AddRow(colors.Black("Scaleway ProjectID"), colors.DarkWhite(i.Scaleway.ProjectID), colors.Black(" "))
			// t.AddRow(colors.Black("Scaleway AccessKey"), colors.DarkWhite(i.Scaleway.AccessKey))
			// t.AddRow(colors.Black("Scaleway SecretKey"), colors.DarkWhite(i.Scaleway.SecretKey))
		} else {
			t.AddRow(colors.Black("Scaleway Cloud"), output.StrDisabled("not-configured"))
		}
		// t.AddRow()
	}

	t.Render()
	fmt.Println()
}
