package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/joshuarubin/go-sway"
)

type Input struct {
	Vendor int64    `json:"vendor"`
	Names  []string `json:"names"`
}

func main() {
	ctx := context.Background()
	client, err := sway.New(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	tree, err := client.GetInputs(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	uniqVendor := make(map[int64]map[string]struct{})
	for _, item := range tree {
		if item.Vendor == 0 {
			continue
		}
		if _, ok := uniqVendor[item.Vendor]; !ok {
			uniqVendor[item.Vendor] = make(map[string]struct{})
			uniqVendor[item.Vendor][item.Name] = struct{}{}
		} else {
			uniqVendor[item.Vendor][item.Name] = struct{}{}
		}
	}

	var items []Input

	for vendor, names := range uniqVendor {
		item := Input{
			Vendor: vendor,
		}
		for name := range names {
			item.Names = append(item.Names, name)
		}
		items = append(items, item)
	}

	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Vendor < items[j].Vendor
	})

	data, _ := json.MarshalIndent(items, "", "	")
	fmt.Println(string(data))
}
