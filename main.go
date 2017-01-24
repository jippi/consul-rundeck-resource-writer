package main

import (
	"log"
	"os"
	"regexp"

	"time"

	api "github.com/hashicorp/consul/api"
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
	config := api.DefaultConfig()
	config.WaitTime = 5 * time.Minute

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal("Could not connect to Consul")
	}

	waitCh := make(chan string)
	updateCh := make(chan []*api.Node)

	go waiter(client, updateCh)
	go consumer(updateCh)

	<-waitCh
}

func configFile() string {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = "resources.yaml"
	}

	return configFile
}

func waiter(client *api.Client, ch chan []*api.Node) {
	q := &api.QueryOptions{WaitIndex: 0}

	for {
		nodes, meta, err := client.Catalog().Nodes(q)
		if err != nil {
			log.Printf("watch: unable to fetch nodes: %s", err)
			time.Sleep(10 * time.Second)
			continue
		}

		remoteWaitIndex := meta.LastIndex
		localWaitIndex := q.WaitIndex

		// only work if the WaitIndex have changed
		if remoteWaitIndex == localWaitIndex {
			log.Printf("Nodes index is unchanged (%d == %d)", localWaitIndex, remoteWaitIndex)
			continue
		}

		log.Printf("Nodes index is changed (%d <> %d)", localWaitIndex, remoteWaitIndex)

		ch <- nodes

		q = &api.QueryOptions{WaitIndex: remoteWaitIndex}
	}
}

func consumer(ch chan []*api.Node) {
	regex := regexp.MustCompile(`(\d+)$`)
	configFile := configFile()

	for {
		log.Println("Waiting for updates ...")
		consulNodes := <-ch
		var output nodes

		for _, consulNode := range consulNodes {
			var tnode node

			tnode.Node = consulNode.Node
			tnode.Address = consulNode.Address
			tnode.OsArch = "amd64"
			tnode.OsFamily = "unix"
			tnode.OsName = "Linux"
			tnode.Username = "root"
			tnode.Tags = append(tnode.Tags, regex.ReplaceAllString(consulNode.Node, ""))

			output = append(output, &tnode)
		}

		d, err := yaml.Marshal(output)
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

		log.Println("Wrote file")
	}
}
