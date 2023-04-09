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

var elasticityEp string = "/cloud/elasticity/"

type elasticitySetThresholdResp struct {
	Status bool `json:"status"`
}

type elasticityEnableResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

type elasticityDisableResp struct {
	Status bool `json:"status"`
}

var elasticityCmd = &cobra.Command{
	Use:   "elasticity",
	Short: "All commands related to elasticity",
}

var elasticitySetLowerThresholdCmd = &cobra.Command{
	Use:   "lower_threshold [pod_id] [value]",
	Short: "Set a lower elastic threshold for a given resource pod.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodPost, ManagerEp+elasticityEp+"lower/", nil)
		if err != nil {
			panic(err)
		}

		params := req.URL.Query()
		params.Add("pod_id", args[0])
		params.Add("lower_threshold", args[1])
		req.URL.RawQuery = params.Encode()

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Decode the response
		var response elasticitySetThresholdResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			panic(err)
		}

		// Print the response
		if response.Status {
			fmt.Print("Success!")
		} else {
			fmt.Print("Failed!")
		}
	},
}

var elasticitySetUpperThresholdCmd = &cobra.Command{
	Use:   "upper_threshold [pod_id] [value]",
	Short: "Set a upper elastic threshold for a given resource pod.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodPost, ManagerEp+elasticityEp+"upper/", nil)
		if err != nil {
			panic(err)
		}

		params := req.URL.Query()
		params.Add("pod_id", args[0])
		params.Add("upper_threshold", args[1])
		req.URL.RawQuery = params.Encode()

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Decode the response
		var response elasticitySetThresholdResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			panic(err)
		}

		// Print the response
		if response.Status {
			fmt.Print("Success!")
		} else {
			fmt.Print("Failed!")
		}
	},
}

var elasticityEnableCmd = &cobra.Command{
	Use:   "enable [pod_id] [min_node] [max_node]",
	Short: "Enable elasticity for a given pod, also need to specifiy the min and max amount of node in elastic mode",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodPost, ManagerEp+elasticityEp+"enable/", nil)
		if err != nil {
			panic(err)
		}

		params := req.URL.Query()
		params.Add("pod_id", args[0])
		params.Add("min_node", args[1])
		params.Add("max_node", args[2])
		req.URL.RawQuery = params.Encode()

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Decode the response
		var response elasticityEnableResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			panic(err)
		}

		// Print the response
		if response.Status {
			fmt.Print("Success!")
		} else {
			fmt.Print("Failed: ")
			fmt.Println(response.Msg)
		}
	},
}

var elasticityDisableCmd = &cobra.Command{
	Use:   "disable [pod_id]",
	Short: "Disable elasticity for a given pod",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodDelete, ManagerEp+elasticityEp+"disable/", nil)
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
		var response podRmResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			panic(err)
		}

		// Print the response
		if response.Status {
			fmt.Print("Success!")
		} else {
			fmt.Print("Failed!")
		}
	},
}

func init() {
	rootCmd.AddCommand(elasticityCmd)
	elasticityCmd.AddCommand(elasticitySetLowerThresholdCmd)
	elasticityCmd.AddCommand(elasticitySetUpperThresholdCmd)
	elasticityCmd.AddCommand(elasticityEnableCmd)
	elasticityCmd.AddCommand(elasticityDisableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// elasticityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// elasticityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
