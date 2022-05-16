package util

import (
	"errors"
	"os"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

type ConfigOptions struct {
	URL  string
	Auth *ssh.PublicKeys
}

func GetGitConfigOptions(url string) (*ConfigOptions, error) {
	if strings.HasPrefix(url, "git@") {
		path := getHomePath()

		path = path + string(os.PathSeparator) + ".ssh" + string(os.PathSeparator) + "id_rsa"

		if FileExists(path) {
			// sshKey, err := ioutil.ReadFile(path)
			// if err != nil {
			// 	return nil, err
			// }

			publicKey, err := ssh.NewPublicKeysFromFile("git", path, "")

			if err != nil {
				return nil, err
			}

			return &ConfigOptions{
				URL:  url,
				Auth: publicKey,
			}, nil
		}

		return nil, errors.New("SSH key does not exist")
	}

	return &ConfigOptions{
		URL:  url,
		Auth: nil,
	}, nil
}
