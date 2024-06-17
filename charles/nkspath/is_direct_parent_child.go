package nkspath

import "path/filepath"

func IsDirectParentChild(parent, child string) bool {

	child = Normalize(child)
	childDir := filepath.Dir(child)

	return Normalize(parent) == Normalize(childDir)
}
