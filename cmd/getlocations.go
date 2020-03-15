/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// getlocationsCmd represents the getlocations command
var getlocationsCmd = &cobra.Command{
	Use:   "getlocations",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Getting all Rick and Morty Dimensions from REST API")
		getLocations()
	},
}

func init() {
	rootCmd.AddCommand(getlocationsCmd)
}

func getLocations() {
	apiurl := "https://rickandmortyapi.com/api/location"
	charClient := http.Client{}

	req, err := http.NewRequest(http.MethodGet, apiurl, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := charClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	rick := resp{}

	jsonErr := json.Unmarshal(body, &rick)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	for i := 0; i < len(rick.Results); i++ {
		fmt.Println(rick.Results[i].Dimension)
	}
}
