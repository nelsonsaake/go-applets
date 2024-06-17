package main

import "projects/saelections/pkg/sysout"

type Item struct {
	Name         string
	Photos       []string
	Price        string
	Units        string
	UnitPrice    string
	SubUnits     string
	SubUnitPrice string
}

type ReportData struct {
	Items [][]Item
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func GroupItemsPerPage(cs []Item, itemsPerPage int) (cspp [][]Item) {
	var j int
	for i := 0; i < len(cs); i += itemsPerPage {
		j += min(itemsPerPage, len(cs))
		cspp = append(cspp, cs[i:j])
	}
	return
}

func data() ReportData {
	items := []Item{}
	for id, item := range items {
		if len(item.Photos) > 1 {
			items[id].Photos = []string{item.Photos[0]}
		}
	}
	if err := loadJsonFile("data.json", &items); err != nil {
		sysout.Fatal(err)
	}
	ipp := GroupItemsPerPage(items, 5)
	sysout.Print(ipp)
	return ReportData{Items: ipp}
}
