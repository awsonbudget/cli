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

var podEp string = "/cloud/pod"

type podLsResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Name  string `json:"name"`
		Id    int    `json:"id"`
		Nodes int    `json:"nodes"`
	} `json:"data"`
}

type podRegisterResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

type podRmResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

var podCmd = &cobra.Command{
	Use:   "pod",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var podLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodGet, ManagerEp+podEp, nil)
		if err != nil {
			panic(err)
		}

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Decode the response
		var response podLsResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			panic(err)
		}

		// Print the response
		if response.Status {
			fmt.Print("Success: ")
			fmt.Println(response.Msg)
			for _, pod := range response.Data {
				fmt.Printf("| ID: %d | Name: %s | Nodes: %d |\n",
					pod.Id, pod.Name, pod.Nodes)
			}
		} else {
			fmt.Print("Failed: ")
			fmt.Println(response.Msg)
		}
	},
}

var podRegisterCmd = &cobra.Command{
	Use:   "register [pod_name]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodPost, ManagerEp+podEp, nil)
		if err != nil {
			panic(err)
		}

		params := req.URL.Query()
		params.Add("pod_name", args[0])
		req.URL.RawQuery = params.Encode()

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Decode the response
		var response podRegisterResp
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

var podRmCmd = &cobra.Command{
	Use:   "rm [pod_name]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodDelete, ManagerEp+podEp, nil)
		if err != nil {
			panic(err)
		}

		params := req.URL.Query()
		params.Add("pod_name", args[0])
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
			fmt.Print("Success: ")
			fmt.Println(response.Msg)
		} else {
			fmt.Print("Failed: ")
			fmt.Println(response.Msg)
		}
	},
}

func init() {
	rootCmd.AddCommand(podCmd)
	podCmd.AddCommand(podLsCmd)
	podCmd.AddCommand(podRegisterCmd)
	podCmd.AddCommand(podRmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// podCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// podCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
