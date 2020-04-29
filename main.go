package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const dateFormat = "20060102"
const titleFormat = "# DevJournal entry for %s on %s"
const titleDateFormat = "Monday, Jan _2 2006"

func getDate(dateStr string, offset int) (time.Time, error) {
	date := time.Now()
	if len(dateStr) > 0 {
		t, err := time.Parse(dateFormat, dateStr)
		if err != nil {
			return date, fmt.Errorf("date must be provided in %s format", dateFormat)
		}
		date = t
	}

	date = date.AddDate(0, 0, offset)
	return date, nil
}

func getFileName(dateStr string, offset int) (string, error) {
	date, err := getDate(dateStr, offset)
	if err != nil {
		return "", err
	}
	ts := date.Format(dateFormat)
	return fmt.Sprintf("%s-devjournal.md", ts), nil
}

func getTitle(author string, dateStr string, offset int) (string, error) {
	date, err := getDate(dateStr, offset)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(titleFormat, author, date.Format(titleDateFormat)), nil
}

func fileExists(filename string, fs fileSystem) bool {
	info, err := fs.Stat(filename)
	if fs.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func main() {
	authorAddr := flag.String("author", "Unknown Author", "The name of the author to which the journal entry belongs.")
	dateAddr := flag.String("date", "", "The date of the journal entry")
	dateOffsetAddr := flag.Int("offset", 0, "Open the journal, offset by the specified number of days.")
	flag.Parse()

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Unable to obtain user's home directory.", err)
	}

	dirpath := filepath.Join(homedir, "journal")
	fn, err := getFileName(*dateAddr, *dateOffsetAddr)
	if err != nil {
		log.Fatal(err)
	}
	fullPath := filepath.Join(dirpath, fn)

	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		os.MkdirAll(dirpath, 0700)
	}

	fs := defaultFileSystem{}
	if !fileExists(fullPath, fs) {
		f, err := os.Create(fullPath)
		if err != nil {
			log.Fatal("Unable to create journal file.", err)
		}

		line, err := getTitle(*authorAddr, *dateAddr, *dateOffsetAddr)
		if err != nil {
			log.Fatal(err)
		}
		f.WriteString(line)
	}

	fmt.Println("Opening journal file: ", fullPath)
	cmd := exec.Command("vim", "+", fullPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		log.Fatal("Command finished with error:", err)
	}
}
