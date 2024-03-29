/*
Copyright © 2023 Joey Yu <xiaowei.yu@mail.mcgill.ca>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var nodeEp string = "/cloud/node/"

type nodeLsResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Name   string `json:"node_name"`
		Id     string `json:"node_id"`
		Type   string `json:"node_type"`
		Status string `json:"node_status"`
		Pod    struct {
			Name string `json:"pod_name"`
			Id   string `json:"pod_id"`
		} `json:"pod_data"`
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

type nodeLogResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "All commands related to node",
}

var nodeLsCmd = &cobra.Command{
	Use:   "ls [pod_id]",
	Short: "List all nodes in a specific pod. If no pod is given, all nodes will be listed",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodGet, ManagerEp+nodeEp, nil)
		if err != nil {
			panic(err)
		}

		if len(args) > 0 {
			params := req.URL.Query()
			params.Add("pod_id", args[0])
			req.URL.RawQuery = params.Encode()
		}

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Decode the response
		var response nodeLsResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			panic(err)
		}

		// Print the response
		if response.Status {
			fmt.Print("Success: ")
			fmt.Println(response.Msg)
			for _, node := range response.Data {
				fmt.Printf("| ID: %s |\n| Name: %s | Type: %s | Status: %s | Pod: %s |\n",
					node.Id, node.Name, node.Type, node.Status, node.Pod.Name)
			}
		} else {
			fmt.Print("Failed: ")
			fmt.Println(response.Msg)
		}
	},
}

var nodeRegisterCmd = &cobra.Command{
	Use:   "register [node_type] [node_name] [pod_id]",
	Short: "Register a node with a given type, name and a target pod id. The type can either be 'job' or 'server'",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodPost, ManagerEp+nodeEp, nil)
		if err != nil {
			panic(err)
		}

		params := req.URL.Query()
		params.Add("node_type", args[0])
		params.Add("node_name", args[1])
		params.Add("pod_id", args[2])
		req.URL.RawQuery = params.Encode()

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Decode the response
		var response nodeRegisterResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			panic(err)
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
	Use:   "rm [node_id]",
	Short: "Remove a specific node given its name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodDelete, ManagerEp+nodeEp, nil)
		if err != nil {
			panic(err)
		}

		params := req.URL.Query()
		params.Add("node_id", args[0])
		req.URL.RawQuery = params.Encode()

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Decode the response
		var response nodeRmResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			panic(err)
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

var nodeLogCmd = &cobra.Command{
	Use:   "log [node_id]",
	Short: "Output the log of a specific node",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodGet, ManagerEp+nodeEp+"log/", nil)
		if err != nil {
			panic(err)
		}

		params := req.URL.Query()
		params.Add("node_id", args[0])
		req.URL.RawQuery = params.Encode()

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Decode the response
		var response nodeLogResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			panic(err)
		}

		// Print the response
		if response.Status {
			fmt.Print("Success: ")
			fmt.Println(response.Data)
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
	nodeCmd.AddCommand(nodeLogCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
