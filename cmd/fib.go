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

	"github.com/spf13/cobra"
)

// fibCmd represents the fib command
var fibCmd = &cobra.Command{
	Use:   "fib",
	Short: "Calculate fibbonacci numbers",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		f := fibonacci()
		for i := 0; i < 10; i++ {
			fmt.Println(f())
		}
	},
}

func init() {
	RootCmd.AddCommand(fibCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fibCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fibCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    var fn0 = 0
	var fn1 = 1
	var fn = 0
	var n = 0

	return func() int {
		if n == 0 {
			fn = fn0
		} else if n == 1 {
			fn = fn1
		} else if n > 1 {
			fn = fn1 + fn0
			fn0 = fn1
			fn1 = fn
		}
		n++

		return fn
	}
}
