/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

type res struct {
	Msg    string `json:"msg"`
	Status bool   `json:"status"`
}

func init() {
	rootCmd.AddCommand(initCmd)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:5550/cloud", nil)
	if err != nil {
		panic(err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf(string(body))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
