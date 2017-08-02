// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
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
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// aCmd represents the a command
var aCmd = &cobra.Command{
	Use:   "a",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Error: A number is required")
			os.Exit(1)
		}
		num, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Error: Could not parse ", args[0], "as float.")
		}
		b := make([]byte, num)
		a := AAAReader{}
		a.Read(b)
		fmt.Println(string(b))
	},
}

type AAAReader struct{}

func (r AAAReader) Read(b []byte) (int, error) {
	to_write := len(b)
	written := 0
	for ; written < to_write; written++ {
		b[written] = 'A'
	}
	return written, nil
}

func init() {
	RootCmd.AddCommand(aCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
