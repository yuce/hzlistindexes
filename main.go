package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/hazelcast/hazelcast-go-client"

	"github.com/yuce/hzlistindexes/internal"
)

func loadConfig(path string) (*hazelcast.Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config from: %s: %w", path, err)
	}
	var cfg hazelcast.Config
	if err = json.Unmarshal(b, &cfg); err != nil {
		return nil, fmt.Errorf("unmarshaling configuration: %w", err)
	}
	return &cfg, nil
}

func bye(msg string) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n", msg))
	os.Exit(1)

}

func main() {
	mapName := flag.String("m", "", "map name")
	cfgPath := flag.String("c", "", "configuration file (JSON)")
	flag.Parse()
	if *mapName == "" {
		bye("map name is required. Try -h")
	}
	//if len(os.Args) != 2 {
	//	bye(fmt.Sprintf("Usage: %s json-config-path", os.Args[0]))
	//	os.Exit(1)
	//}
	var cfg hazelcast.Config
	if *cfgPath != "" {
		c, err := loadConfig(*cfgPath)
		if err != nil {
			panic(err)
		}
		cfg = *c
	}
	client, err := hazelcast.StartNewClientWithConfig(context.Background(), cfg)
	if err != nil {
		panic(err)
	}
	ci := hazelcast.NewClientInternal(client)
	indexes, err := internal.GetIndexes(context.Background(), ci, *mapName)
	if err != nil {
		panic(err)
	}
	if len(indexes) == 0 {
		fmt.Printf("Found no indexes for :%s\n", *mapName)
	}
	for i, index := range indexes {
		fmt.Printf("%03d: name: %s, type: %d on key: %s, attrs: %v\n", i+1, index.Name, index.Type, index.BitmapIndexOptions.UniqueKey, index.Attributes)
	}
}
