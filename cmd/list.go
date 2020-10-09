/*
Copyright © 2020 Thomas Ruschival <thomas@ruschival.de>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	drclient "github.com/truschival/drcli/pkg/digitalrooster"
	"golang.org/x/net/context"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list a given resource [podcasts,alarms,radios]",
	Long: ``}


// radiosCmd represents the radios command
var list_radiosCmd = &cobra.Command{
	Use:   "radios",
	Short: "List configured internet radio steams",
	Long:  `Returns a list of all internet radio streams`,
	RunE:  readRadioList,
}

// podcastsCmd represents the podcasts command
var list_podcastsCmd = &cobra.Command{
	Use:   "podcasts",
	Short: "List all podcast rss feeds",
	Long:  `Returns a list of all configured podcast sources`,
	RunE:  readPodcastsList,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(list_radiosCmd)
	listCmd.AddCommand(list_podcastsCmd)
}

func SetupClient() *drclient.APIClient {
	configuration := drclient.NewConfiguration()
	configuration.Servers[0].URL= viper.Get("apiurl").(string)
	return drclient.NewAPIClient(configuration)
}

func readRadioList(cmd *cobra.Command, args []string) error {
	var len int32 = 1
	var offset int32 = 0
	
	api_client := SetupClient()

	resp, r, err := api_client.InternetRadioApi.IradioReadAll(
		context.Background()).Length(len).Offset(offset).Execute()
	if err != nil {
		jww.INFO.Printf("Error when calling `InternetRadioApi.IradioReadAll``: %v\n", err)
		jww.DEBUG.Printf("Full HTTP response: %v\n", r)
	}

	jww.INFO.Printf("Response from `RadiosApi.RadiosReadAll`: %v\n", resp)
	return err
}


// actual reading of podcast list
func readPodcastsList(cmd *cobra.Command, args []string) error {
	var len int32 =1
	var offset int32 =0
	api_client := SetupClient()	
	resp, r, err := api_client.PodcastsApi.PodcastsReadAll(
		context.Background()).Length(len).Offset(offset).Execute()
	if err != nil {
		jww.ERROR.Printf("Error when calling `PodcastsApi.PodcastsReadAll``: %v\n", err)
		jww.INFO.Printf("Full HTTP response: %v\n", r)
	}
	// response from `PodcastsReadAll`: []Podcast
	jww.INFO.Printf("Response from `PodcastsApi.PodcastsReadAll`: %v\n", resp)
	return err
}
