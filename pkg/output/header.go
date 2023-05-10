package output

import (
	"fmt"

	"mmesh.dev/m-lib/pkg/utils/colors"
)

func AppHeader(output bool) string {
	return logo(output)
}

func logo(output bool) string {
	website := fmt.Sprintf("%s %s", colors.Black("Main Website: "), colors.DarkWhite("https://mmesh.io"))
	docsite := fmt.Sprintf("%s %s", colors.Black("Documentation:"), colors.DarkWhite("https://docs.mmesh.io"))
	// discord := fmt.Sprintf("%s %s", colors.Black("Discord:      "), colors.DarkWhite("https://mmesh.io/discord"))

	if output {
		l1 := fmt.Sprintf(colors.DarkWhite("  ■   ▄  ▄▄ ▄▄ ▄▄ ▄▄ ▄▄▄▄ ▄▄▄▄ ▄  ▄ ") + colors.Green("│") + "\n")
		l2 := fmt.Sprintf(colors.DarkWhite("■  ██    █ ▄ █ █ ▄ █ █■   ▀  ▄ █▄▄█ ")+colors.Green("│")+" %s\n", website)
		l3 := fmt.Sprintf(colors.DarkWhite("  ▀   ■  ▀ ▀ ▀ ▀ ▀ ▀ ▀▀▀▀ ▀▀▀▀ ▀  ▀ ")+colors.Green("│")+" %s\n", docsite)
		return fmt.Sprintf("%s%s%s\n", l1, l2, l3)
	}

	fmt.Printf("  ■   ▄  ▄▄ ▄▄ ▄▄ ▄▄ ▄▄▄▄ ▄▄▄▄ ▄  ▄ " + colors.Green("│") + "\n")
	fmt.Printf("■  ██    █ ▄ █ █ ▄ █ █■   ▀  ▄ █▄▄█ "+colors.Green("│")+" %s\n", website)
	fmt.Printf("  ▀   ■  ▀ ▀ ▀ ▀ ▀ ▀ ▀▀▀▀ ▀▀▀▀ ▀  ▀ "+colors.Green("│")+" %s\n", docsite)
	fmt.Println()

	return ""
}

/*
func Logo0() {
	fmt.Printf("  ■   ▄  ▄   ▄ ▄   ▄ ▄▄▄▄ ▄▄▄▄ ▄    " + colors.Blue("│") + " \n")
	fmt.Printf("■  ██    █▀▄▀█ █▀▄▀█ █■■▀ ▀  ▄ █■■▄ " + colors.Blue("│") + " \n")
	fmt.Printf("  ▀   ■  ▀ ▀ ▀ ▀ ▀ ▀ ▀▀▀▀ ▀▀▀▀ ▀  ▀ " + colors.Blue("│") + " Discover more at " + colors.DarkWhite("https://mmesh.io") + "\n")
}

func Logo1() {
	fmt.Printf("  ■   ▄  ▄   ▄ ▄   ▄ ▄▄▄▄ ▄▄▄▄ ▄  ▄ │                  \n")
	fmt.Printf("■  ██    █▀▄▀█ █▀▄▀█ █■■  ▀■■▄ █■■█ │ Discover more at:\n")
	fmt.Printf("  ▀   ■  ▀ ▀ ▀ ▀ ▀ ▀ ▀▀▀▀ ▀▀▀▀ ▀  ▀ │ https://mmesh.io \n")
}

func Logo2() {
	fmt.Printf("  ■   ▄  ▄   ▄ ▄   ▄ ▄▄▄▄ ▄▄▄▄ ▄  ▄ ┌┐                 \n")
	fmt.Printf("■  ██    █▀▄▀█ █▀▄▀█ █■■  ▀■■▄ █■■█ ││ https://mmesh.io\n")
	fmt.Printf("  ▀   ■  ▀ ▀ ▀ ▀ ▀ ▀ ▀▀▀▀ ▀▀▀▀ ▀  ▀ └┘                 \n")
}

func Logo3() {
	fmt.Printf("  ■   ▄  ▄   ▄ ▄   ▄ ▄▄▄▄ ▄▄▄▄ ▄  ▄ ┌┐┌───┐                  \n")
	fmt.Printf("■  ██    █▀▄▀█ █▀▄▀█ █■■  ▀■■▄ █■■█ │││ ■ │ Discover more at:\n")
	fmt.Printf("  ▀   ■  ▀ ▀ ▀ ▀ ▀ ▀ ▀▀▀▀ ▀▀▀▀ ▀  ▀ └┘└───┘ https://mmesh.io \n")
}

func Logo4() {
	fmt.Printf("  ■   ▄  ▄   ▄ ▄   ▄ ▄▄▄▄ ▄▄▄▄ ▄  ▄   ┌┐┌───┐ \n")
	fmt.Printf("■  ██    █▀▄▀█ █▀▄▀█ █■■  ▀■■▄ █■■█   │││ ■ │ \n")
	fmt.Printf("  ▀   ■  ▀ ▀ ▀ ▀ ▀ ▀ ▀▀▀▀ ▀▀▀▀ ▀  ▀ ▀ └┘└───┘ \n")
}

func logo5() {
	fmt.Print("██  ██ ██  ██ ██████ ██████ ██  ██")
	fmt.Print("██████ ██████ ████   ██████ ██████")
	fmt.Print("██  ██ ██  ██ ██████ ██████ ██  ██")

	fmt.Print("┌┐  ┌┐┌┐  ┌┐┌▄▄┌▄▄ ┌┐  ┌┐")
	fmt.Print("││┌┐││││┌┐│││■ └■■┐││■■││")
	fmt.Print("└┘└┘└┘└┘└┘└┘└▀▀ ▀▀┘└┘  └┘")

	fmt.Print("╔╗┌┐  ┌┐┌┐  ┌┐╔═┌── ┌┐  ┌┐")
	fmt.Print("║║││┌┐││││┌┐││╠═└──┐│────│")
	fmt.Print("╚╝┘└┘└┘└┘└┘└┘└╚═ ──┘└┘  └┘")
}
*/
