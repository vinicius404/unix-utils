package ls

import (
	"fmt"
	"io"
	"io/fs"
	"slices"
	"strings"
)

type LSConfig struct {
	Path               string
	ShowHiddenFiles    bool
	DirectoryIndicator bool
	ReverseOrder       bool
}

func defaultLSConfig() *LSConfig {
	return &LSConfig{
		Path:               ".",
		ShowHiddenFiles:    false,
		DirectoryIndicator: false,
		ReverseOrder:       false,
	}
}

// ListDirFiles lists the files in the specified directory and writes the output to the provided writer.
func ListDirFiles(fileSystem fs.FS, w io.Writer, config *LSConfig) error {
	if config == nil {
		config = defaultLSConfig()
	}

	dir, err := fs.ReadDir(fileSystem, config.Path)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", config.Path, err)
	}

	filteredEntries := filterEntries(dir, config)
	sortedEntries := sortEntries(filteredEntries, config)
	fmt.Fprint(w, buildOutputString(sortedEntries, config))

	return nil
}

func filterEntries(dir []fs.DirEntry, config *LSConfig) []fs.DirEntry {
	filteredEntries := []fs.DirEntry{}

	for _, entry := range dir {
		if !config.ShowHiddenFiles && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		filteredEntries = append(filteredEntries, entry)
	}

	return filteredEntries
}

func sortEntries(dir []fs.DirEntry, config *LSConfig) []fs.DirEntry {
	reverseAlphabetical := func(a, b fs.DirEntry) int {
		switch {
		case a.Name() < b.Name():
			return 1
		case a.Name() == b.Name():
			return 0
		case a.Name() > b.Name():
			return -1
		}

		return 0
	}

	if config.ReverseOrder {
		slices.SortFunc(dir, reverseAlphabetical)
	}

	return dir
}

func buildOutputString(dir []fs.DirEntry, config *LSConfig) string {
	var output strings.Builder

	for i, f := range dir {
		output.WriteString(f.Name())

		if config.DirectoryIndicator && f.IsDir() {
			output.WriteString("/")
		}

		if i < len(dir)-1 {
			output.WriteString("  ")
		} else {
			output.WriteString("\n")
		}
	}

	return output.String()
}
