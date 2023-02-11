/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type nodeLsResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Name   string `json:"name"`
		Id     string    `json:"id"`
		Status string   `json:"status"`
		Pod    struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		} `json:"pod"`
	} `json:"data"`
}

type nodeRegisterResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

type nodeRmResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("node called")
	},
}

var nodeLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodGet, ManagerEp+"/cloud/node/", nil)
		if err != nil {
			fmt.Print("Failed: ")
			fmt.Println(err)
			return
		}

		params := req.URL.Query()
		if len(args) > 0 {
			params.Add("node_name", args[0])
		}
		req.URL.RawQuery = params.Encode()

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			fmt.Print("Failed: ")
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		// Decode the response
		var response nodeLsResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			fmt.Print("Failed: ")
			fmt.Println(err)
			return
		}

		// Print the response
		if response.Status {
			fmt.Print("Success: ")
			fmt.Println(response.Msg)
			for _, node := range response.Data {
                fmt.Printf("| ID: %s |\n| Name: %s | Status: %s | Pod: %s |\n", node.Id, node.Name, node.Status, node.Pod.Name)
			}
		} else {
			fmt.Print("Failed: ")
			fmt.Println(response.Msg)
		}
	},
}

var nodeRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
    Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodPost, ManagerEp+"/cloud/node/", nil)
		if err != nil {
			fmt.Print("Failed: ")
			fmt.Println(err)
			return
		}

		params := req.URL.Query()
		params.Add("node_name", args[0])
		if len(args) > 1 {
			params.Add("pod_name", args[1])
		}

		req.URL.RawQuery = params.Encode()

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			fmt.Print("Failed: ")
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		// Decode the response
		var response nodeRegisterResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			fmt.Print("Failed: ")
			fmt.Println(err)
			return
		}

		// Print the response
		if response.Status {
			fmt.Print("Success: ")
			fmt.Println(response.Msg)
		} else {
			fmt.Print("Failed: ")
			fmt.Println(response.Msg)
		}
	},
}



var nodeRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
    Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodDelete, ManagerEp+"/cloud/node/", nil)
		if err != nil {
			fmt.Print("Failed: ")
			fmt.Println(err)
			return
		}

		params := req.URL.Query()
		params.Add("node_name", args[0])
		req.URL.RawQuery = params.Encode()

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			fmt.Print("Failed: ")
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		// Decode the response
		var response podRmResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			fmt.Print("Failed: ")
			fmt.Println(err)
			return
		}

		// Print the response
		if response.Status {
			fmt.Print("Success: ")
			fmt.Println(response.Msg)
		} else {
			fmt.Print("Failed: ")
			fmt.Println(response.Msg)
		}
	},
}


func init() {
	rootCmd.AddCommand(nodeCmd)
	nodeCmd.AddCommand(nodeLsCmd)
	nodeCmd.AddCommand(nodeRegisterCmd)
	nodeCmd.AddCommand(nodeRmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
