/*
Copyright (C) 2016 Red Hat, Inc.

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

package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configUnsetCmd = &cobra.Command{
	Use:   "unset PROPERTY_NAME",
	Short: "unsets an individual value in a minishift config file",
	Long:  "unsets PROPERTY_NAME from the minishift config file.  Can be overwritten by flags or environmental variables",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintf(os.Stdout, "usage: minishift config unset PROPERTY_NAME")
			os.Exit(1)
		}
		err := unset(args[0])
		if err != nil {
			fmt.Fprintln(os.Stdout, err)
		}
	},
}

func init() {
	ConfigCmd.AddCommand(configUnsetCmd)
}

func unset(name string) error {
	m, err := ReadConfig()
	if err != nil {
		return err
	}
	delete(m, name)
	return WriteConfig(m)
}
