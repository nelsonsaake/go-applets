package main

import (
	"path"
	"projects/saelections/pkg/sysout"
	"projects/saelections/pkg/ufs"

	"github.com/spf13/cobra"
)

type Item struct {
	Name         string   `csv:"name"`
	Photos       []string `csv:"photo"`
	Price        string   `csv:"price"`
	Units        string   `csv:"units"`
	UnitPrice    string   `csv:"unitprice"`
	SubUnits     string   `csv:"subunits"`
	SubUnitPrice string   `csv:"subunitprice"`
}

func handlePlaygroundCmd() {
	var (
		csvfile  = "repo/tosell.csv"
		outfile  = "report/data.json"
		items    = loadItemsFromCsvFile(csvfile)
		erritems = []string{}
	)

	productByName := func(name string) (p Product, err error) {
		p, err = RepoProductFind(name)
		if err != nil {
			erritems = append(erritems, name)
		}
		return
	}

	imagesByProductName := func(name string) (urls []string) {
		p, err := productByName(name)
		if err != nil {
			return
		}
		// if len(p.Images) == 0 {
		// 	ScrapeImageUrlsFor(p)
		// 	p, err = productByName(name)
		// 	if err != nil {
		// 		return
		// 	}
		// 	ScrapeImageFilesFor(p)
		// 	p, err = productByName(name)
		// 	if err != nil {
		// 		return
		// 	}
		// }
		for _, img := range p.Images {
			urls = append(urls, path.Join("..", img.Local))
		}
		return
	}

	for id := range items {
		items[id].Photos = imagesByProductName(items[id].Name)
	}

	sysout.Print("errors", len(erritems), erritems)

	if err := ufs.WriteFile(outfile, tojson(items)); err != nil {
		sysout.Fatal(err)
	}
}

func init() {
	playgroundCmd := &cobra.Command{}

	rootCmd.AddCommand(playgroundCmd)

	playgroundCmd.Use = "playground"

	// playgroundCmd.Short = "A search tool, for locating products in database"

	playgroundCmd.Run = func(cmd *cobra.Command, args []string) {
		handlePlaygroundCmd()
	}
}
