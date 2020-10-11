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
	"fmt"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	"github.com/truschival/drcli/pkg/digitalrooster"
	"golang.org/x/net/context"
	"strconv"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list a given resource [podcasts,alarms,radios]",
	Long:  ``}

// radiosCmd represents the radios command
var list_radiosCmd = &cobra.Command{
	Use:   "radios [<offset> [length]]",
	Short: "List configured internet radio steams",
	Long:  `Returns a list of all internet radio streams`,
	RunE:  readRadioList,
}

// podcastsCmd represents the podcasts command
var list_podcastsCmd = &cobra.Command{
	Use:   "podcasts [<offset> [length]]",
	Short: "List all podcast rss feeds",
	Long:  `Returns a list of all configured podcast sources`,
	RunE:  readPodcastsList,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(list_radiosCmd)
	listCmd.AddCommand(list_podcastsCmd)
}

func SetupClient() *digitalrooster.APIClient {
	configuration := digitalrooster.NewConfiguration()
	configuration.Servers[0].URL = viper.Get("apiurl").(string)
	jww.DEBUG.Printf("Using Server: %s\n", configuration.Servers[0].URL)
	return digitalrooster.NewAPIClient(configuration)
}

func readRadioList(cmd *cobra.Command, args []string) error {
	api_client := SetupClient()
	req := api_client.InternetRadioApi.IradioReadAll(
		context.Background())

	if len(args) >= 1 {
		if offset, err := strconv.Atoi(args[0]); err != nil {
			jww.ERROR.Printf("offset argument must be numeric %s\n", err)
		} else {
			jww.DEBUG.Printf("Offset: %d\n", int32(offset))
			req = req.Offset(int32(offset))
		}
	}
	if len(args) == 2 {
		if length, err := strconv.Atoi(args[1]); err != nil {
			jww.ERROR.Printf("length argument must be numeric %s\n", err)
		} else {
			jww.DEBUG.Printf("length: %d\n", int32(length))
			req = req.Length(int32(length))
		}
	}
	resp, r, err := req.Execute()

	if err != nil {
		jww.INFO.Printf("Error when calling `InternetRadioApi.IradioReadAll``: %v\n", err)
		jww.DEBUG.Printf("Full HTTP response: %v\n", r)
		return err
	}

	for _, s := range resp {
		fmt.Printf("%37.38s | %20.20s | %s\n", s.GetId(), s.GetName(), s.GetUrl())
	}
	return err
}

// actual reading of podcast list
func readPodcastsList(cmd *cobra.Command, args []string) error {
	api_client := SetupClient()
	req := api_client.PodcastsApi.PodcastsReadAll(
		context.Background())
	if len(args) >= 1 {
		if offset, err := strconv.Atoi(args[0]); err != nil {
			jww.ERROR.Printf("offset argument must be numeric %s\n", err)
		} else {
			jww.DEBUG.Printf("Offset: %d\n", int32(offset))
			req = req.Offset(int32(offset))
		}
	}
	if len(args) == 2 {
		if length, err := strconv.Atoi(args[1]); err != nil {
			jww.ERROR.Printf("length argument must be numeric %s\n", err)
		} else {
			jww.DEBUG.Printf("length: %d\n", int32(length))
			req = req.Length(int32(length))
		}
	}
	resp, r, err := req.Execute()

	if err != nil {
		jww.ERROR.Printf("Error when calling `PodcastsApi.PodcastsReadAll``: %v\n", err)
		jww.INFO.Printf("Full HTTP response: %v\n", r)
		return err
	}

	for _, s := range resp {
		fmt.Printf("%37.38s | %20.20s | %s\n", s.GetId(), s.GetTitle(), s.GetUrl())
	}
	return err

}
