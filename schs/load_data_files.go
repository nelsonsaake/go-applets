package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"math"
	"net/url"
	"path/filepath"
	"projects/saelections/pkg/str"
	"projects/saelections/pkg/ufs"
	"strings"

	"github.com/yosssi/gohtml"
)

const (
	dataDir        = "data"
	patternSuffix  = "_pattern"
	lineSeparator  = "\r\n"
	lineSeparator2 = "\n"
)

// ---

func isValidUrl(str string) bool {
	_, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}

	u, err := url.Parse(str)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

type FileEntry struct {
	Data  string
	Class string
}

func (fentry FileEntry) HTML() string {
	inner := fentry.Data
	if isValidUrl(fentry.Data) {
		inner = fmt.Sprintf("<a href='%[1]s'>%[1]s</a>", fentry.Data)
	}
	return fmt.Sprintf("<div class='%s'>%s</div>", fentry.Class, inner)
}

// ---

type Sch struct {
	Details []FileEntry
}

func (sch *Sch) HTML() string {
	buf := bytes.Buffer{}
	write := buf.WriteString

	write("<div class='sch'>")
	for _, entry := range sch.Details {
		write(entry.HTML())
	}
	write("</div>")

	return buf.String()
}

// ---

type File struct {
	datafile    string
	data        string
	dlines      []string
	patternfile string
	pattern     string
	plines      []string
	entries     []FileEntry
	Schs        []Sch
}

func (f *File) HTML() string {
	buf := bytes.Buffer{}
	write := buf.WriteString

	write("<div class='schs'>")
	for _, sch := range f.Schs {
		write(sch.HTML())
	}
	write("</div>")

	return gohtml.Format(buf.String())
}

// ---

func FileEntries(file File) []FileEntry {
	if len(file.dlines)%len(file.plines) != 0 {
		return nil
	}

	pl := func(i int) string {
		return file.plines[i%len(file.plines)]
	}

	entries := []FileEntry{}
	for i, dl := range file.dlines {
		entries = append(entries, FileEntry{Data: dl, Class: pl(i)})
	}

	return entries
}

func Files(ls map[string]string) []File {
	files := []File{}
	for p, pc := range ls {
		if strings.HasSuffix(p, patternSuffix) {
			d := strings.TrimSuffix(p, patternSuffix)
			if dc, ok := ls[d]; ok {
				file := File{datafile: d, data: dc, patternfile: p, pattern: pc}
				files = append(files, file)
			}
		}
	}
	return files
}

func min(i, j int) int {
	return int(math.Min(float64(i), float64(j)))
}

func Schs(file File) (schs []Sch) {
	var j int
	for i := 0; i < len(file.entries); i += len(file.plines) {
		j += min(len(file.plines), len(file.entries))
		schs = append(schs, Sch{file.entries[i:j]})
	}
	return
}

// ---

func WalkDataDirFunc(ls map[string]string) fs.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		ls[path], err = ufs.ReadFile(path)
		return err
	}
}

func Lines(s string) []string {
	raw := strings.Split(s, lineSeparator)
	if len(raw) == 1 && raw[0] == s {
		raw = strings.Split(s, lineSeparator2)
	}

	lines := []string{}
	for _, line := range raw {
		if !str.Empty(line) {
			lines = append(lines, strings.TrimSpace(line))
		}
	}

	return lines
}

func LoadData() []File {
	ls := map[string]string{}
	filepath.WalkDir(dataDir, WalkDataDirFunc(ls))

	files := Files(ls)
	for i := range files {
		file := &files[i]
		file.dlines = Lines(file.data)
		file.plines = Lines(file.pattern)
		file.entries = FileEntries(*file)
		file.Schs = Schs(*file)
	}

	return files
}

// ---
