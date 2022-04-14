package credentials

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/pquerna/otp/totp"
	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/status"
)

func SetCredentialsPassword(newUser bool) string {
	if !newUser {
		fmt.Println()

		if !input.GetConfirm("Change password?", false) {
			return ""
		}
	}

	for {
		pw1 := input.GetPassword("New Password:", "")
		pw2 := input.GetPassword("Confirm Password:", "")

		if pw1 == pw2 {
			return pw1
		}

		fmt.Printf("\nPasswords mismatch, please try again\n")
	}
}

func SetCredentialsTOTP(u *iam.User) *iam.TOTP {
	otp := u.Credentials.TOTP

	fmt.Println()

	userOpt := false

	var promptConfirm *survey.Confirm
	if !otp.Enabled {
		promptConfirm = &survey.Confirm{Message: "Do you want to enable 2-Factor Authentication?", Default: false}
	} else {
		promptConfirm = &survey.Confirm{Message: "Do you want to regenerate (or disable) the 2FA Token?", Default: false}
	}

	if err := survey.AskOne(promptConfirm, &userOpt, survey.WithValidator(survey.Required), survey.WithIcons(input.SurveySetIcons)); err != nil {
		status.Error(err, "Unable to get response")
	}

	if userOpt {
		fmt.Println()

		if otp.Enabled {
			if input.GetConfirm("Disable 2-Factor Authentication?", false) {
				otp.Enabled = false
				otp.Secret = ""
				return otp
			}
		}

		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      u.AccountID,
			AccountName: u.Email,
		})
		if err != nil {
			status.Error(err, "Unable to generate TOTP key")
		}

		// fmt.Printf("Issuer:       %s\n", key.Issuer())
		// fmt.Printf("Account Name: %s\n", key.AccountName())
		// fmt.Printf("Secret:       %s\n", key.Secret())

		// // Convert TOTP key into a PNG
		// var buf bytes.Buffer
		// img, err := key.Image(200, 200)
		// if err != nil {
		// 	status.Error(err, "Unable to generate QR image")
		// }
		// if err := png.Encode(&buf, img); err != nil {
		// 	status.Error(err, "Unable to generate PNG image")
		// }

		// // display the QR code to the user.
		// fmt.Println("Writing PNG to qr-code.png....")
		// if err := ioutil.WriteFile("qr-code.png", buf.Bytes(), 0644); err != nil {
		// 	status.Error(err, "Unable to write PNG file")
		// }

		qrObj := qrcodeTerminal.New()
		qrObj.Get(key.String()).Print()

		ok := false

		for !ok {
			fmt.Println()

			otpcode := input.GetInput("Confirm TOTP Code:", "", "", survey.Required)

			valid := totp.Validate(otpcode, key.Secret())
			if valid {
				otp.Secret = key.Secret()
				otp.Enabled = true

				ok = true
			} else {
				fmt.Printf("\nSorry, invalid code, please try again\n")
				otp.Secret = ""
				otp.Enabled = false
			}
		}
	}

	return otp
}

/*
func SetCredentialsSSH(u *iam.User) *iam.SSH {
	ssh := u.Credentials.SSH

	if ssh == nil {
		ssh = &iam.SSH{}
	}

	if len(ssh.PublicKey) > 0 {
		fmt.Println()

		if !input.GetConfirm("Change SSH public key?", false) {
			return ssh
		}
	}

	pubKey := input.GetInput("New SSH Public Key:", "", ssh.PublicKey, input.ValidSSHPublicKey)

	return &iam.SSH{
		PublicKey: pubKey,
	}
}
*/
