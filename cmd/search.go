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
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

type AutoGenerated struct {
	Total    int         `json:"total"`
	Subtotal int         `json:"subtotal"`
	Page     int         `json:"page"`
	PerPage  int         `json:"per_page"`
	Search   interface{} `json:"search"`
	Sort     struct {
		By    interface{} `json:"by"`
		Order interface{} `json:"order"`
	} `json:"sort"`
	Results []struct {
		IP                       string      `json:"ip"`
		EnvironmentID            interface{} `json:"environment_id"`
		EnvironmentName          interface{} `json:"environment_name"`
		LastReport               interface{} `json:"last_report"`
		Mac                      string      `json:"mac"`
		RealmID                  interface{} `json:"realm_id"`
		RealmName                interface{} `json:"realm_name"`
		SpMac                    interface{} `json:"sp_mac"`
		SpIP                     interface{} `json:"sp_ip"`
		SpName                   interface{} `json:"sp_name"`
		DomainID                 int         `json:"domain_id"`
		DomainName               string      `json:"domain_name"`
		ArchitectureID           int         `json:"architecture_id"`
		ArchitectureName         string      `json:"architecture_name"`
		OperatingsystemID        int         `json:"operatingsystem_id"`
		OperatingsystemName      string      `json:"operatingsystem_name"`
		SubnetID                 int         `json:"subnet_id"`
		SubnetName               string      `json:"subnet_name"`
		SpSubnetID               interface{} `json:"sp_subnet_id"`
		PtableID                 int         `json:"ptable_id"`
		PtableName               string      `json:"ptable_name"`
		MediumID                 int         `json:"medium_id"`
		MediumName               string      `json:"medium_name"`
		Build                    bool        `json:"build"`
		Comment                  string      `json:"comment"`
		Disk                     string      `json:"disk"`
		InstalledAt              interface{} `json:"installed_at"`
		ModelID                  interface{} `json:"model_id"`
		ModelName                interface{} `json:"model_name"`
		HostgroupID              int         `json:"hostgroup_id"`
		HostgroupName            string      `json:"hostgroup_name"`
		OwnerID                  int         `json:"owner_id"`
		OwnerType                string      `json:"owner_type"`
		Enabled                  bool        `json:"enabled"`
		PuppetCaProxyID          interface{} `json:"puppet_ca_proxy_id"`
		Managed                  bool        `json:"managed"`
		UseImage                 interface{} `json:"use_image"`
		ImageFile                string      `json:"image_file"`
		UUID                     interface{} `json:"uuid"`
		ComputeResourceID        interface{} `json:"compute_resource_id"`
		ComputeResourceName      interface{} `json:"compute_resource_name"`
		ComputeProfileID         interface{} `json:"compute_profile_id"`
		ComputeProfileName       interface{} `json:"compute_profile_name"`
		Capabilities             []string    `json:"capabilities"`
		ProvisionMethod          string      `json:"provision_method"`
		PuppetProxyID            interface{} `json:"puppet_proxy_id"`
		Certname                 string      `json:"certname"`
		ImageID                  interface{} `json:"image_id"`
		ImageName                interface{} `json:"image_name"`
		CreatedAt                time.Time   `json:"created_at"`
		UpdatedAt                time.Time   `json:"updated_at"`
		LastCompile              interface{} `json:"last_compile"`
		GlobalStatus             int         `json:"global_status"`
		GlobalStatusLabel        string      `json:"global_status_label"`
		PuppetStatus             int         `json:"puppet_status"`
		BuildStatus              int         `json:"build_status,omitempty"`
		BuildStatusLabel         string      `json:"build_status_label,omitempty"`
		Name                     string      `json:"name"`
		ID                       int         `json:"id"`
		ConfigurationStatus      int         `json:"configuration_status,omitempty"`
		ConfigurationStatusLabel string      `json:"configuration_status_label,omitempty"`
	} `json:"results"`
}

const query = "https://foreman.localdomain/api/hosts"

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "searching the api",
	Long:  `a long description of searching the api.. - tbd. `,
	Run: func(cmd *cobra.Command, args []string) {
		var username = "admin"
		var passwd = "password"
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
		dat := AutoGenerated{}
		if err := json.Unmarshal(data, &dat); err != nil {
			panic(err)
		}
		fmt.Println(dat.Total)
		fmt.Println(dat.Results[0].IP)
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
