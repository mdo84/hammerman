// Copyright © 2017 markus dollinger <markus@mdo.name>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Total   int `json:"total"`
	Results Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Env  string `json:"env"`
}

const query = "https://foreman.localdomain/api/hosts?search=core3.localdomain"

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "searching the api",
	Long:  `a long description of searching the api.. - tbd. `,
	Run: func(cmd *cobra.Command, args []string) {
		var username = "admin"
		var passwd = "ZyuzuDEHeTSkcAkp"
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		req, err := http.NewRequest("GET", query, nil)
		req.SetBasicAuth(username, passwd)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		output := Response{}
		err = json.Unmarshal(data, &output)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%s", output)
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
