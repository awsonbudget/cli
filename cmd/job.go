/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var jobEp = "/cloud/job"

type jobLsResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Name   string `json:"name"`
		Id     string `json:"id"`
		Status string `json:"status"`
		Node   string `json:"node"`
	} `json:"data"`
}

type jobLaunchResp struct {
	Status bool `json:"status"`
	Data   struct {
		Id string `json:"job_id"`
	} `json:"data"`
}

type jobAbortResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

type jobLogResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var jobLsCmd = &cobra.Command{
	Use:   "ls [node_id]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodGet, ManagerEp+jobEp, nil)
		if err != nil {
			panic(err)
		}

		if len(args) > 0 {
			params := req.URL.Query()
			params.Add("node_id", args[0])
			req.URL.RawQuery = params.Encode()
		}

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Decode the response
		var response jobLsResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			panic(err)
		}

		// Print the response
		if response.Status {
			fmt.Print("Success: ")
			fmt.Println(response.Msg)
			for _, job := range response.Data {
				fmt.Printf("| ID: %s | Name: %s | Status: %s | Node: %s |\n",
					job.Id, job.Name, job.Status, job.Node)
			}
		} else {
			fmt.Print("Failed: ")
			fmt.Println(response.Msg)
		}
	},
}

var jobLaunchCmd = &cobra.Command{
	Use:   "launch [job_name] [job_script]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Prepare the script
		file, err := os.Open(args[1])
		filename := filepath.Base(args[1])
		if err != nil {
			panic(err)
		}
		defer file.Close()

		buf := new(bytes.Buffer)
		io.Copy(buf, file)

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("job_script", filename)
		if err != nil {
			panic(err)
		}
		io.Copy(part, buf)

		err = writer.Close()
		if err != nil {
			panic(err)
		}

		// Build the request
		req, err := http.NewRequest(http.MethodPost, ManagerEp+jobEp, body)
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())

		params := req.URL.Query()
		params.Add("job_name", args[0])
		req.URL.RawQuery = params.Encode()

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Decode the response
		var response jobLaunchResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			panic(err)
		}

		// Print the response
		if response.Status {
			fmt.Print("Success: ")
			fmt.Println(response.Data.Id)
		} else {
			fmt.Print("Failed: ")
		}

	},
}

var jobAbortCmd = &cobra.Command{
	Use:   "abort [job_id]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodDelete, ManagerEp+jobEp, nil)
		if err != nil {
			panic(err)
		}

		params := req.URL.Query()
		params.Add("job_id", args[0])
		req.URL.RawQuery = params.Encode()

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Decode the response
		var response jobAbortResp
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

var jobLogCmd = &cobra.Command{
	Use:   "log [job_id]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Build the request
		req, err := http.NewRequest(http.MethodGet, ManagerEp+jobEp+"/log", nil)
		if err != nil {
			panic(err)
		}

		params := req.URL.Query()
		params.Add("job_id", args[0])
		req.URL.RawQuery = params.Encode()

		// Send the request
		res, err := Client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// Decode the response
		var response jobLogResp
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			panic(err)
		}

		// Print the response
		if response.Status {
			fmt.Print("Success: ")
			fmt.Println(response.Msg)
			fmt.Println(response.Data)
		} else {
			fmt.Print("Failed: ")
			fmt.Println(response.Msg)
		}
	},
}

func init() {
	rootCmd.AddCommand(jobCmd)
	jobCmd.AddCommand(jobLsCmd)
	jobCmd.AddCommand(jobLaunchCmd)
	jobCmd.AddCommand(jobAbortCmd)
	jobCmd.AddCommand(jobLogCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jobCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jobCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
