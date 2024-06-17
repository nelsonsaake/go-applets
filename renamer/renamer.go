package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"renamer/pkg/sysout"
	"renamer/pkg/ufs"
	"strings"
)

func isAGo() bool {
	if len(os.Args) > 1 {
		command := strings.ToLower(os.Args[1])
		acceptableCommands := []string{"dothething", "rename", "go", "do it", "doit"}
		for _, anAcceptableCommand := range acceptableCommands {
			if command == anAcceptableCommand {
				return true
			}
		}
	}
	return false
}

func main() {

	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {

		// must be directly related?
		if IsNotDirectlyRelated(dir, path) && mustBeDirectlyRelated {
			return nil
		}

		if d == nil {
			return nil
		}

		// skip dir
		if skipFiles && ufs.IsFile(path) {
			return nil
		}

		// skip file
		if skipDirs && ufs.IsDir(path) {
			return nil
		}

		// apply changes
		newName := d.Name()
		for from, to := range changes {
			newName = strings.ReplaceAll(newName, from, to)
		}

		// apply changes via funcs
		for _, f := range changerFuncs {
			newName = f(newName)
		}

		// if name is still the same
		oldName := d.Name()
		if newName == oldName {
			return nil
		}

		var (
			dir     = strings.TrimSuffix(path, d.Name()) // get dir
			newPath = filepath.Join(dir, newName)        // dir + name = newpath
		)

		if isAGo() {
			os.Rename(path, newPath) // rename path
		}
		sysout.Print("%s >>> %s\n", path, newPath) // print to oldPath and newPath to console

		return nil
	})

	if isAGo() {
		sysout.Print("[CHANGES APPLIED]")
	} else {
		sysout.Print("[PREVIEW OF CHANGES]")
	}
}
