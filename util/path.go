package util

import (
	"os"

	"github.com/mitchellh/go-homedir"
)

func getHomePath() string {
	home, err := homedir.Dir()

	if err != nil {
		home, err = os.UserHomeDir()

		if err != nil {
			panic(err)
		}
	}

	return home
}

func GetDFMPath() string {
	return getHomePath() + string(os.PathSeparator) + ".dfm"
}

func GetDFMDatabasePath() string {
	return GetDFMPath() + string(os.PathSeparator) + "database.json"
}

func GetDFMInstallCachePath() string {
	return GetDFMPath() + string(os.PathSeparator) + "install"
}

func GetDFMRepoPath() string {
	return GetDFMPath() + string(os.PathSeparator) + "repo"
}
