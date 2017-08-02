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
	//"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rot13Cmd represents the rot13 command
var rot13Cmd = &cobra.Command{
	Use:   "rot13",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := strings.NewReader("Lbh penpxrq gur pbqr!")

		if len(args) == 1 {
			s = strings.NewReader(args[0])
		}

		r := rot13Reader{s}
		io.Copy(os.Stdout, &r)
	},
}

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	cypher_text := make([]byte, 8)
	cypher_read, cypher_err := rot.r.Read(cypher_text)
	if cypher_err != nil {
		return 0, cypher_err
	}
	written := 0

	for ; written < cypher_read; written++ {
		text := cypher_text[written]

		var start byte = 0

		if text >= 'A' && text <= 'Z' {
			start = 'A'
		} else if text >= 'a' && text <= 'z' {
			start = 'a'
		}

		if start > 0 {
			b[written] = ((text - start + 13) % 26) + start
		} else {
			b[written] = text
		}
	}
	return written, nil
}

func init() {
	RootCmd.AddCommand(rot13Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rot13Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rot13Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
