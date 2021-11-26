/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"strconv"

	"github.com/spf13/cobra"
)

// evenCmd represents the even command
var evenCmd = &cobra.Command{
	Use:   "even",
	Short: "偶數相加",
	Long:  `針對用戶提供的清單，只對偶數進行相加`,
	Run: func(cmd *cobra.Command, args []string) {
		var evenSum int
		for _, ival := range args {
			temp, _ := strconv.Atoi(ival)
			if temp%2 == 0 {
				evenSum = evenSum + temp
			}
		}
		fmt.Printf("The even addition of %s is %d\n", args, evenSum)
	},
}

func init() {
	rootCmd.AddCommand(evenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// evenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// evenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
