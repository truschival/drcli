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
	"golang.org/x/net/context"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a single entry of a resource type [podcasts,alarms,radios]",
	Long:  ``}

// radiosCmd represents the radios command
var get_radioCmd = &cobra.Command{
	Use:   "radios <UUID>",
	Short: "get radio <UUID> ",
	Long:  `Returns a single radio station identified by <UUID>`,
	RunE:  getRadio,
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(get_radioCmd)
}

func getRadio(cmd *cobra.Command, args []string) error {
	jww.TRACE.Printf("get.getRadio(%v)\n", args)
	id := args[0]
	api_client := SetupClient()

	resp, r, err := api_client.RadiosApi.IradioReadOne(
		context.Background(), id).Execute()
	if err != nil {
		jww.INFO.Printf("Error when calling `InternetRadioApi.IradioReadAll``: %v\n", err)
		jww.DEBUG.Printf("Full HTTP response: %v\n", r)
		return err
	}
	fmt.Printf("%37.38s | %20.20s | %s\n", resp.GetId(), resp.GetName(), resp.GetUrl())

	return err
}
