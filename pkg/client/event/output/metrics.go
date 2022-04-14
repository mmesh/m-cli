package output

import (
	"fmt"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/events"
	"mmesh.dev/m-api-go/grpc/resources/metrics"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/resources"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) ShowMetrics(em *metrics.EventMetrics) {
	output.SubTitleT2("Activity")

	t := table.New()

	t.AddRow(colors.Black("Successful Events"), colors.DarkGreen(fmt.Sprintf("%07.0f", em.SuccessCount)))
	t.AddRow(colors.Black("Failed Events"), colors.DarkRed(fmt.Sprintf("%07.0f", em.FailCount)))
	t.AddRow(colors.Black("First Activity"), colors.DarkWhite(time.Unix(em.FirstActivity, 0).String()))
	t.AddRow(colors.Black("Last Activity"), colors.DarkWhite(time.Unix(em.LastActivity, 0).String()))
	t.AddRow(colors.Black("Activity Index"), colors.DarkWhite(fmt.Sprintf("%.4f", em.ActivityIndex)))
	t.AddRow(colors.Black("Failure Probability"), api.FailureProbability(em))
	t.AddRow(colors.Black("Score"), colors.DarkWhite(fmt.Sprintf("%.2f", em.Score)))
	t.AddRow(colors.Black("Resource Rating"), colors.Black("[")+printRating(em.Rating)+colors.Black("]"))

	switch em.LastResult {
	case events.EventResult_SUCCESS:
		t.AddRow(colors.Black("Last Event Result"), output.StrEnabled("SUCCESS"))
	case events.EventResult_FAIL:
		t.AddRow(colors.Black("Last Event Result"), output.StrDisabled("FAIL"))
	}

	t.Render()
	fmt.Println()
}

func (api *API) FailureProbability(em *metrics.EventMetrics) string {
	if em.FailProbability < 20 {
		return colors.Green(fmt.Sprintf("%.2f%%", em.FailProbability))
	} else if em.FailProbability < 60 {
		return colors.Yellow(fmt.Sprintf("%.2f%%", em.FailProbability))
	} else {
		return colors.Red(fmt.Sprintf("%.2f%%", em.FailProbability))
	}
}

func printRating(rating string) string {
	switch rating {
	case resources.RatingA:
		return colors.DarkGreen(rating)
	case resources.RatingB:
		return colors.DarkYellow(rating)
	case resources.RatingC:
		return colors.Yellow(rating)
	case resources.RatingD:
		return colors.DarkRed(rating)
	case resources.RatingE:
		return colors.Red(rating)
	}

	return colors.Black("n/a")
}
