package credentials

import (
	"fmt"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/input"
)

func SetSSHKeys(u *iam.User, manageKeys bool) *iam.User {
	if !manageKeys {
		fmt.Println()

		if !input.GetConfirm("Manage your SSH keys?", false) {
			return u
		}
	}

	// output.Choice("Manage your SSH keys")

	var sshKeys []string
	var sshKeyName string

	if u.SSHKeys != nil {
		for _, k := range u.SSHKeys {
			sshKeys = append(sshKeys, k.Name)
		}
		sort.Strings(sshKeys)
	} else {
		u.SSHKeys = make(map[string]*iam.SSHKey)
	}

	sshKeys = append(sshKeys, input.NewResource)

	sshKeyName = input.GetSelect("SSH Key:", "", sshKeys, survey.Required)

	if sshKeyName == input.NewResource {
		sshKeyName = ""
	} else {
		if input.GetConfirm("Delete this SSH key?", false) {
			k := u.SSHKeys[sshKeyName]

			if u.Credentials.SSH != nil {
				if u.Credentials.SSH.Key != nil {
					if u.Credentials.SSH.Key.Name == k.Name &&
						u.Credentials.SSH.Key.PublicKey == k.PublicKey {
						u.Credentials.SSH = nil
					}
				}
			}

			delete(u.SSHKeys, sshKeyName)

			return u
		}
	}

	k := &iam.SSHKey{}

	k.Name = input.GetInput("Key Name:", "SSH key name", sshKeyName, input.ValidID)

	var publicKey string
	if key, ok := u.SSHKeys[sshKeyName]; ok {
		publicKey = key.PublicKey
	}
	k.PublicKey = input.GetInput("Public Key:", "", publicKey, input.ValidSSHPublicKey)

	u.SSHKeys[k.Name] = k

	if input.GetConfirm("Use this key for login?", false) {
		if u.Credentials.SSH == nil {
			u.Credentials.SSH = &iam.SSH{}
		}
		u.Credentials.SSH.Enabled = true
		u.Credentials.SSH.Key = k
	}

	return u
}

func GetSSHKey(userSSHKeys *iam.UserSSHKeys) *iam.SSHKey {
	if userSSHKeys.SSHKeys == nil {
		return nil
	}

	var sshKeys []string
	var sshKeyName string

	for _, k := range userSSHKeys.SSHKeys {
		sshKeys = append(sshKeys, k.Name)
	}
	sort.Strings(sshKeys)

	sshKeyName = input.GetSelect("SSH Key:", "", sshKeys, survey.Required)

	if k, ok := userSSHKeys.SSHKeys[sshKeyName]; ok {
		return k
	}

	return nil
}
