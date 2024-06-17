package data

import "fmt"

type Data struct {
	GHC float64
	MB  float64
}

type DataByGHC []Data

func (opts DataByGHC) Len() int           { return len(opts) }
func (opts DataByGHC) Swap(i, j int)      { opts[i], opts[j] = opts[j], opts[i] }
func (opts DataByGHC) Less(i, j int) bool { return opts[i].GHC < opts[j].GHC }

func (data Data) String() string {

	return fmt.Sprintf("\t%6.2fMB - %6.2fGH,\n", data.MB, data.GHC)
}
