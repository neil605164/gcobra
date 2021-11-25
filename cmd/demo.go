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

// demoCmd represents the demo command
var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "demo 指令",
	Long:  `測試 cobra 新增指令功能，指令名稱為 demo`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取 float 标识符的值，默认为 false
		fstatus, _ := cmd.Flags().GetBool("float")
		if fstatus { // 如果为 true，则调用 floatAdd 函数
			floatAdd(args)
		} else {
			intAdd(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(demoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// demoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// demoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// 區域參數
	demoCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
}

// cmd/add.go
func intAdd(args []string) {
	var sum int
	// 循环 args 参数，循环的第一个值为 args 的索引，这里我们不需要，所以用 _ 忽略掉
	for _, ival := range args {
		// 将 string 转换成 int 类型
		temp, err := strconv.Atoi(ival)
		if err != nil {
			panic(err)
		}
		sum = sum + temp
	}
	fmt.Printf("Addition of numbers %s is %d\n", args, sum)
}

func floatAdd(args []string) {
	var sum float64
	for _, fval := range args {
		// 将字符串转换成 float64 类型
		temp, err := strconv.ParseFloat(fval, 64)
		if err != nil {
			panic(err)
		}
		sum = sum + temp
	}
	fmt.Printf("Sum of floating numbers %s is %f\n", args, sum)
}
