package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	addCmd := &cobra.Command{
		Use:   "scbo",
		Short: "Add scrapbox project",
		Long:  "Add scrapbox project",
		Run: func(c *cobra.Command, args []string) {
			makeConfigFile()
			var project string
			fmt.Scan(&project)

			content := []byte(project)
			ioutil.WriteFile("~/.scbo/config.yml", content, os.ModePerm)
		},
	}
}

func makeConfigFile() {
	if err := os.MkdirAll("~/.scbo/", 0777); err != nil {
		fmt.Println(err)
	}
	_, err := os.Stat("~/.scbo/config.yml")
	if err != nil {
		fmt.Println(err)
		content := []byte("projects:\n")
		ioutil.WriteFile("~/.scbo/config.yml", content, os.ModePerm)
	}
}
