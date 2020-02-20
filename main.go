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
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	title, body, project string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "scbo",
		Short: "you can easily create new scrapbox",
		Long:  "you can easily create new scrapbox",
		Run: func(c *cobra.Command, args []string) {
			fmt.Printf("Project: ")
			fmt.Scan(&project)
			fmt.Printf("Title: ")
			fmt.Scan(&title)
			fmt.Printf("Body: ")
			fmt.Scan(&body)

			url := "https://scrapbox.io/" + project + "/" + title + "?body=" + body
			fmt.Println(url)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
