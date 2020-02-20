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
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Projects []string
}

var configFile string = "$HOME/.scbo/config.yml"
var config Config

var (
	title, body string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "scbo",
		Short: "you can easily create new scrapbox",
		Long:  "you can easily create new scrapbox",
		Run: func(c *cobra.Command, args []string) {
			var project string
			var n int

			projects := config.Projects

			if len(projects) == 1 {
				project = projects[0]
				fmt.Println(project)

			} else {
				for i, p := range projects {
					fmt.Printf("%d  %v\n", i, p)
				}
				for {
					fmt.Printf("Project: ")
					fmt.Scan(&n)
					if 0 <= n && n < len(projects) {
						project = projects[n]
						fmt.Println(project)
						break
					}
					fmt.Println("Invalid project")
				}
			}

			fmt.Printf("Title: ")
			fmt.Scan(&title)
			fmt.Printf("Body: ")
			fmt.Scan(&body)

			url := "https://scrapbox.io/" + project + "/" + title + "?body=" + body
			fmt.Println(url)
			exec.Command("open", url).Start()
			pbcopy(url)
		},
	}

	cobra.OnInitialize(func() {
		viper.SetConfigName("config")
		viper.AddConfigPath("$HOME/.scbo")
		viper.SetConfigType("yml")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if err := viper.Unmarshal(&config); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func pbcopy(url string) {
	c1 := exec.Command("echo", url)
	c2 := exec.Command("pbcopy")

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	var out bytes.Buffer
	c2.Stdout = &out

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()
}
