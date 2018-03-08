// Copyright © 2017 Aljabr, Inc.
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
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	. "github.com/kocircuit/kocircuit/lang/go/eval"
	. "github.com/kocircuit/kocircuit/lang/go/weave"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Ko circuits compiled into this binary",
	Long: `Ko binaries provide a namespace of compiled circuits.
The list command lists all compiled circuits and their namespace paths.

Usage:
	ko list
`,
	Run: func(cmd *cobra.Command, args []string) {
		tables := [][]string{}
		tables = append(tables, EvalFaculty().StringTable("Eval")...)
		tables = append(tables, EvalIdiomRepo.StringTable("EvalIdiom")...)
		tables = append(tables, GoFaculty().StringTable("WeaveGo")...)
		tables = append(tables, GoIdiomRepo.StringTable("WeaveGoIdiom")...)
		printTable(os.Stdout, tables)
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

func printTable(w io.Writer, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Faculty", "Package", "Circuit", "Implemented by"})
	table.SetAutoMergeCells(false)
	table.SetRowLine(false)
	table.AppendBulk(data)
	table.Render()
}