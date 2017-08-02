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

// sqrtCmd represents the sqrt command
var sqrtCmd = &cobra.Command{
	Use:   "sqrt",
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
		num, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Println("Error: Could not parse ", args[0], "as float.")
		}
		fmt.Println(Sqrt(num))
	},
}

func init() {
	RootCmd.AddCommand(sqrtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sqrtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sqrtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type ErrorNegativeSqrt float64

func (e ErrorNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot take the square root of a negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrorNegativeSqrt(x)
	}

	var n int = 10
	var z float64 = 1.0
	var zd float64 = 0.0

	for i := 0; i < n; i++ {
	  zd = (z*z - x) / (2*z)
	  z = z - zd
	}

	return z, nil
}
