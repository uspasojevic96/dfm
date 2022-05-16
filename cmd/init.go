/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/uspasojevic96/dfm/database"
	"github.com/uspasojevic96/dfm/pkg"
	"github.com/uspasojevic96/dfm/util"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes dotfile manager repository",
	Long:  `Initializes dotfile manager repository with specific dotfiles github repository.`,
	Run:   runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	path := util.GetDFMRepoPath()

	stat, _ := os.Stat(path)

	if stat != nil {
		log.Fatal("Repository already exists")
	}

	// url, err := url.Parse(args[0])
	//
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	configOptions, err := util.GetGitConfigOptions(args[0])

	if err != nil {
		log.Fatal(err)
	}

	_, err = git.PlainClone(path, false, &git.CloneOptions{
		URL: configOptions.URL,

		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth:              configOptions.Auth,
	})

	if err != nil {
		log.Fatal(err)
	}

	db := database.New()
	err = db.LoadDatabase()

	pkgs, err := pkg.LoadPackages(path)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		e := err
		err = db.SaveDatabase()

		if err != nil {
			log.Fatal(err)
		}

		log.Fatal(e)
	}

	for _, pkg := range pkgs {
		db.AddPackage(pkg)
		fmt.Print(pkg.Name)
	}

	err = db.SaveDatabase()

	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}
