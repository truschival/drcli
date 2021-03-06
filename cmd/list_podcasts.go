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
)

// podcastsCmd represents the podcasts command
var list_podcastsCmd = &cobra.Command{
	Use:   "podcasts",
	Short: "List all podcast rss feeds",
	Long: ``,
	RunE: func(cmd *cobra.Command, args []string) (error) {
		fmt.Println("list podcasts called")
		return nil
	},
}

func init() {
	listCmd.AddCommand(list_podcastsCmd)
}
