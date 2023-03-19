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

var serverEp string = "/cloud/server/"

type serverLaunchResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		NodeId string `json:"node_id"`
		Port   int    `json:"port"`
	} `json:"data"`
}

type serverResumeResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		NodeId string `json:"node_id"`
		Port   int    `json:"port"`
	} `json:"data"`
}

type serverPauseResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		NodeId string `json:"node_id"`
		Port   int    `json:"port"`
	} `json:"data"`
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "All commands related to server",
}

var serverLaunchCmd = &cobra.Command{
	Use:   "launch [pod_id]",
	Short: "Launch all server nodes in a pod given the pod id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodPost, ManagerEp+serverEp+"launch/", nil)
		if err != nil {
			panic(err)
		}

		params := req.URL.Query()
		params.Add("pod_id", args[0])
		req.URL.RawQuery = params.Encode()

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Decode the response
		var response serverLaunchResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			panic(err)
		}

		// Print the response
		if response.Status {
			fmt.Print("Success: ")
			fmt.Println(response.Msg)
			for _, node := range response.Data {
				fmt.Printf("| NodeId: %s |\n| Port: %d\n",
					node.NodeId, node.Port)
			}
		} else {
			fmt.Print("Failed: ")
			fmt.Println(response.Msg)
		}
	},
}

var serverResumeCmd = &cobra.Command{
	Use:   "resume [pod_id]",
	Short: "Resume all server nodes in a pod given the pod id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodPost, ManagerEp+serverEp+"resume/", nil)
		if err != nil {
			panic(err)
		}

		params := req.URL.Query()
		params.Add("pod_id", args[0])
		req.URL.RawQuery = params.Encode()

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Decode the response
		var response serverResumeResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			panic(err)
		}

		// Print the response
		if response.Status {
			fmt.Print("Success: ")
			fmt.Println(response.Msg)
			for _, node := range response.Data {
				fmt.Printf("| NodeId: %s |\n| Port: %d\n",
					node.NodeId, node.Port)
			}
		} else {
			fmt.Print("Failed: ")
			fmt.Println(response.Msg)
		}
	},
}

var serverPauseCmd = &cobra.Command{
	Use:   "pause [pod_id]",
	Short: "Pause all server nodes in a pod given the pod id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodPost, ManagerEp+serverEp+"pause/", nil)
		if err != nil {
			panic(err)
		}

		params := req.URL.Query()
		params.Add("pod_id", args[0])
		req.URL.RawQuery = params.Encode()

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Decode the response
		var response serverPauseResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			panic(err)
		}

		// Print the response
		if response.Status {
			fmt.Print("Success: ")
			fmt.Println(response.Msg)
			for _, node := range response.Data {
				fmt.Printf("| NodeId: %s |\n| Port: %d\n",
					node.NodeId, node.Port)
			}
		} else {
			fmt.Print("Failed: ")
			fmt.Println(response.Msg)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(serverLaunchCmd)
	serverCmd.AddCommand(serverPauseCmd)
	serverCmd.AddCommand(serverResumeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
