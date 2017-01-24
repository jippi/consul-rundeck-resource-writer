package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

type nodes []*node

type node struct {
	Node     string   `yaml:"nodename"`
	Address  string   `yaml:"hostname"`
	OsArch   string   `yaml:"osArch"`
	OsFamily string   `yaml:"osFamily"`
	OsName   string   `yaml:"osName"`
	Username string   `yaml:"username"`
	Tags     []string `yaml:"tags"`
}

func main() {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = "resources.yaml"
	}
	log.Printf("Waiting for stdin...")

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Could not read stdin")
	}

	var n nodes
	err = json.Unmarshal(bytes, &n)
	if err != nil {
		log.Fatalf("Could not marshal input")
	}

	regex := regexp.MustCompile(`(\d+)$`)

	for _, nn := range n {
		nn.OsArch = "amd64"
		nn.OsFamily = "unix"
		nn.OsName = "Linux"
		nn.Username = "root"
		nn.Tags = append(nn.Tags, regex.ReplaceAllString(nn.Node, ""))
	}

	d, err := yaml.Marshal(&n)
	if err != nil {
		log.Fatalf("Could not marshal struct to YAML")
	}

	s := "---\n"
	s = s + string(d)

	f, err := os.Create(configFile)
	if err != nil {
		log.Fatalf("Could not create file")
	}

	f.Write([]byte(s))
}
