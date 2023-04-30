package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// ================================================
// STRUCTS FOR HANDLING TABLE DATA
// ================================================

type Column struct {
	Name   string
	Tp     string
	Unique bool
}

type Table struct {
	Name          string
	PrimaryKeys   []Column
	ForeignKeys   []Column
	Columns       []Column
	RelatedTables []string
}

type Group struct {
	Name   string
	Label  string
	Tables []Table
}

// ================================================
// GENERIC SET
// ================================================

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Put(keys ...T) {
	for _, k := range keys {
		s[k] = struct{}{}
	}
}

func (s Set[T]) Has(key T) bool {
	_, exist := s[key]
	return exist
}

func (s Set[T]) Keys() []T {
	var keys []T
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

// ================================================
// COMMON FUNCTIONS
// ================================================

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
		err = filepath.Walk(input, func(path string, info fs.FileInfo, err error) error {
			if !info.IsDir() && filepath.Ext(path) == ".sql" {
				sqlFiles = append(sqlFiles, path)
			}
			return nil
		})

		if err != nil {
			continue
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
