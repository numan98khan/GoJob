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
	"fmt"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up tej on a remote machine",
	Run: func(cmd *cobra.Command, args []string) {
		setup(cmd)
	},
}

// This is the function that will carry out all the implementations
func setup(cmd *cobra.Command) error {
    queue, _ := cmd.Flags().GetString("queue")

	fmt.Println("setup called")

    fmt.Println("Argument :  " + queue)

    return nil
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	setupCmd.Flags().StringP("destination", "", "", "Machine to SSH into; [user@]host[:port]")
	rootCmd.MarkFlagRequired("destination")
	setupCmd.Flags().StringP("queue", "", "", "Directory for tej's files")
	setupCmd.Flags().StringP("runtime", "r", "", "runtime to deploy on the server if the queue doesn't\nexist. If unspecified, will auto-detect what is\nappropriate, and fallback on 'default'.")
	setupCmd.Flags().StringP("make-link", "", "", "")
	setupCmd.Flags().StringP("make-default-link", "", "", "")
	setupCmd.Flags().StringP("force", "", "", "")
	setupCmd.Flags().StringP("only-links", "", "", "")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
