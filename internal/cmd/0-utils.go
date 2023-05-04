package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func getSqlFiles(inputPaths ...string) ([]string, error) {
	var sqlFiles []string
	for _, input := range inputPaths {
		// Get input stat
		stat, err := os.Stat(input)
		if err != nil {
			continue
		}

		// If input is not dir, halt
		if !stat.IsDir() {
			if filepath.Ext(stat.Name()) == ".sql" {
				sqlFiles = append(sqlFiles, input)
			}
			continue
		}

		// Look for SQL files inside dir
		dirItems, err := os.ReadDir(input)
		if err != nil {
			continue
		}

		for _, item := range dirItems {
			path := filepath.Join(input, item.Name())
			if !item.IsDir() && filepath.Ext(path) == ".sql" {
				sqlFiles = append(sqlFiles, path)
			}
		}
	}

	if len(sqlFiles) == 0 {
		return nil, fmt.Errorf("no sql files found")
	}

	return sqlFiles, nil
}

func writeOutput(data []byte, dst string, isRaw bool) error {
	// If there are no destination path, just write in stdout
	if dst == "" {
		_, err := os.Stdout.Write(data)
		return err
	}

	// If destination has no extension, put it
	if filepath.Ext(dst) == "" {
		if isRaw {
			dst += ".d2"
		} else {
			dst += ".svg"
		}
	}

	// Write to destination file
	return os.WriteFile(dst, data, os.ModePerm)
}
