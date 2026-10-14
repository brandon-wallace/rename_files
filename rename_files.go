package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// removeEmojis removes all emoji characters from the input string by filtering out
// runes that fall within known emoji Unicode ranges. Returns a string.
func removeEmojis(input string) string {
	result := []rune{}
	for _, r := range input {
		if !isEmoji(r) {
			result = append(result, r)
		}
	}
	return string(result)
}

// isEmoji checks if a rune is an emoji character by testing if it falls within
// any of the major Unicode emoji ranges including symbols, emoticons, and pictographs.
// Returns true or false.
func isEmoji(r rune) bool {
	return (r >= 0x1F300 && r <= 0x1F5FF) || // Miscellaneous Symbols and Pictographs
		(r >= 0x1F600 && r <= 0x1F64F) || // Emoticons
		(r >= 0x1F680 && r <= 0x1F6FF) || // Transport and Map Symbols
		(r >= 0x1F900 && r <= 0x1F9FF) || // Supplemental Symbols and Pictographs
		(r >= 0x2600 && r <= 0x26FF) || // Miscellaneous Symbols
		(r >= 0x2700 && r <= 0x27BF) || // Dingbats
		(r >= 0x1F800 && r <= 0x1F8FF) || // Supplemental Arrows-C
		(r >= 0x2B00 && r <= 0x2BFF) || // Miscellaneous Symbols and Arrows
		(r >= 0xFE00 && r <= 0xFE0F)
}

// normalizeFileName processes a file path to create a cleaned and normalized filename. Returns the new file path.
func normalizeFileName(oldPath string) string {
	directory := filepath.Dir(oldPath)
	base := filepath.Base(oldPath)
	ext := filepath.Ext(oldPath)
	filename := base[:len(base)-len(ext)]

	// Remove emojis
	newName := removeEmojis(filename)

	// Convert uppercase to lowercase
	newName = strings.ToLower(newName)

	// Remove square brackets with characters inbetween
	re := regexp.MustCompile(`\[[^\]]*\]`)
	newName = re.ReplaceAllString(newName, "")

	// Remove multiple underscore characters
	re = regexp.MustCompile(`_+`)
	newName = re.ReplaceAllString(newName, "_")

	// Replace spaces with underscore characters
	newName = strings.ReplaceAll(newName, " ", "_")

	// Replace _-_ pattern with underscore characters
	newName = strings.ReplaceAll(newName, "_-_", "_")

	// Replace _-_ pattern with underscore characters
	newName = strings.ReplaceAll(newName, "__", "_")

	// Remove comma characters
	newName = strings.ReplaceAll(newName, ",", "")

	// Remove octothorpe characters
	newName = strings.ReplaceAll(newName, "#", "")

	// Remove _-. characters from the beginning and ending of string
	newName = strings.Trim(newName, "_-.")

	return filepath.Join(directory, newName+strings.ToLower(ext))
}

func main() {
	dryRun := flag.Bool("dry-run", false, "boolean value to enable a dry-run")
	flag.Parse()

	// Add or remove extensions here.
	extensions := map[string]struct{}{
		".webm": {},
		".mp4":  {},
		".mkv":  {},
		".m4a":  {},
		".mp3":  {},
		".ogg":  {},
	}

	files := make([]string, 0)
	root := "."

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Print(err)
			return nil
		}

		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))
			if _, exists := extensions[ext]; exists {
				files = append(files, path)
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fileCount := 0

	for _, file := range files {
		newFile := normalizeFileName(file)

		// Skip the file if it has already been renamed.
		if file == newFile {
			continue
		}

		// Skip this file because it already exists.
		if _, err := os.Stat(newFile); err == nil {
			continue
		}

		if !*dryRun {
			err := os.Rename(file, newFile)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		fileCount++
		fmt.Println("New file name:", newFile)
	}

	fmt.Printf("\n %d files have been renamed.\n", fileCount)
}
