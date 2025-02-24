package main

import (
	"fmt"
)

type Edge struct {
	source string
	target string
}

type Node struct {
	id      string
	typeStr string
}

func main() {
	edges := []Edge{
		{"input-node-1", "voyage-embed-node-2"},
		{"voyage-embed-node-2", "http-client-adapter-verticle-node-4"},
		{"http-client-adapter-verticle-node-4", "voyage-transform-node-3"},
		{"voyage-transform-node-3", "to-publish-verticle-node-10"},
		{"to-publish-verticle-node-10", "line-node-7"},
		{"to-publish-verticle-node-10", "facebook-node-8"},
		{"to-publish-verticle-node-10", "discord-node-9"},
		{"to-publish-verticle-node-10", "output-node-10"},
	}

	nodesMap := map[string]string{
		"input-node-1":                        "input",
		"voyage-embed-node-2":                 "org.maoz.prehandle.workers.neoai.aiclient.embedding.VoyageVerticle",
		"voyage-transform-node-3":             "org.maoz.prehandle.workers.neoai.aiclient.embedding.util.VoyageTransformVerticle",
		"http-client-adapter-verticle-node-4": "org.maoz.prehandle.workers.neoai.httpclient.HttpClientAdapterVerticle",
		"line-node-7":                         "org.maoz.prehandle.workers.neoai.notify.LineVerticle",
		"facebook-node-8":                     "org.maoz.prehandle.workers.neoai.notify.FacebookVerticle",
		"discord-node-9":                      "org.maoz.prehandle.workers.neoai.notify.DiscordVerticle",
		"to-publish-verticle-node-10":         "org.maoz.prehandle.workers.neoai.ebtransform.ToPublishVerticle",
		"output-node-10":                      "org.maoz.prehandle.workers.neoai.output.OutputVerticle",
	}

	nodes := []string{}
	for _, v := range nodesMap {
		nodes = append(nodes, v)
	}

	addressIn := []int{}
	addressOut := []int{}

	for _, edge := range edges {
		if edge.source == "input-node-1" {
			for i, node := range nodes {
				if node == nodesMap[edge.target] {
					addressIn = append(addressIn, i)
				}
			}
		}
		if edge.target != "line-node-7" && edge.target != "facebook-node-8" && edge.target != "discord-node-9" && edge.target != "output-node-10" {
			for i, node := range nodes {
				if node == nodesMap[edge.target] {
					addressOut = append(addressOut, i)
				}
			}
		}
	}

	fmt.Println("Nodes:", nodes)
	fmt.Println("addressIn:", addressIn)
	fmt.Println("addressOut:", addressOut)
}
