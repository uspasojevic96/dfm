/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	home, err := homedir.Dir()

	if err != nil {
		fmt.Println(err)
		return
	}

	path := home + string(os.PathSeparator) + ".dfm" + string(os.PathSeparator) + "repo"

	r, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:               args[0],
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	_ = r

	if err != nil {
		log.Fatal(err)
	}

}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
