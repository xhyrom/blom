package resolver

import (
	"os"
	"path/filepath"
	"strings"
)

func (r *ModuleResolver) findModulePath(importPath string) (string, error) {
	if strings.HasPrefix(importPath, "std/") {
		stdPath := filepath.Join(r.StandardLibPath, importPath+".blom")
		if _, err := os.Stat(stdPath); err == nil {
			return stdPath, nil
		}
	}

	for _, searchPath := range r.SearchPaths {
		candidatePath := filepath.Join(searchPath, importPath+".blom")
		if _, err := os.Stat(candidatePath); err == nil {
			return candidatePath, nil
		}

		normalizedPath := filepath.Join(searchPath,
			strings.ReplaceAll(importPath, "/", string(os.PathSeparator))+".blom")
		if _, err := os.Stat(normalizedPath); err == nil {
			return normalizedPath, nil
		}
	}

	return "", os.ErrNotExist
}
