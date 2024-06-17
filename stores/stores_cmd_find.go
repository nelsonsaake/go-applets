package main

import (
	"projects/saelections/pkg/str"
	"projects/saelections/pkg/sysin"
	"projects/saelections/pkg/sysout"

	"github.com/spf13/cobra"
)

func handleFindCmd(q string) {
	cp, cf := sysout.Print, sysout.Fatal

	cerr := func(err error) {
		if err != nil {
			cf(err)
		}
	}

	pimgstostrs := func(pimgs []*Image) (imgs []string) {
		for _, pimg := range pimgs {
			if !str.Empty(pimg.Local) {
				imgs = append(imgs, pimg.Local)
			} else if !str.Empty(pimg.Url) {
				imgs = append(imgs, pimg.Url)
			}
		}
		return
	}

	printproduct := func(p Product) {
		cp(
			"---",
			p.Name,
			p.Price,
			p.Description,
			p.CategoryName,
			pimgstostrs(p.Images),
			p.Href,
		)
	}

	printprouducts := func(fps []Product) {
		for _, fp := range fps {
			printproduct(fp)
		}
	}

	printnoproductfound := func(fps []Product) {
		if len(fps) == 0 {
			cf("sorry, no item matched your query")
		}
	}

	handlefoundproducts := func(fps []Product, err error) {
		cp("Searching %s ...", q)
		cerr(err)
		printnoproductfound(fps)
		printprouducts(fps)
	}

	handlefoundproducts(RepoProductsFind(q))
}

func init() {
	var q string

	requireq := func() {
		if str.Empty(q) {
			q = sysin.Line("keyword")
		}
	}

	findCmd := &cobra.Command{}

	rootCmd.AddCommand(findCmd)

	// findCmd.Flags().StringVarP(&q, "keyword", "q", "", "search keyword")

	findCmd.Use = "find"

	findCmd.Short = "A search tool, for locating products in database"

	findCmd.Run = func(cmd *cobra.Command, args []string) {
		requireq()
		handleFindCmd(q)
	}
}
