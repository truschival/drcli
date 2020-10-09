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
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	"os"
	fp "path/filepath"
)

var (
	cfgFile string
	home    string
)

const (
	defaultFileName string = ".drcli.yaml"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "drcli",
	Short: "CLI to configure and control a remote DigitalRooster",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		jww.ERROR.Println(err)
		os.Exit(1)
	}
}

func init() {
	jww.SetStdoutThreshold(jww.LevelTrace)
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", ".drcli.yaml",
		"config file (default is $HOME/.drcli.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	var err error
	home, err := homedir.Dir()
	if err != nil {
		jww.ERROR.Println(err)
		os.Exit(1)
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		fmt.Printf("using cfgFile: %s\n", cfgFile)
	} else {
		// Search config in home directory with name ".drcli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(defaultFileName)
	}

	initDefaultConfig()  // initialize variables, may be overwritten by config
	viper.AutomaticEnv() // read in environment variables that match
	
	if err := viper.ReadInConfig(); err != nil {
		if os.IsNotExist(err) {
			err := viper.WriteConfigAs(fp.Join(home, defaultFileName))
			if err != nil {
				jww.WARN.Printf(
					"Creating default config failed! %s\n ",
					err)
			}

		} else {
			jww.WARN.Printf("could not read config file %s %s\n",
				cfgFile, err)
		}
	}
}


func initDefaultConfig() {
	jww.TRACE.Println("> initDefaultConfig()")	
	viper.SetDefault("apiurl", "http://digitalrooster:6666/api/v1")
}
