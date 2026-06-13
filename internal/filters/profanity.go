package filters

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	badWords map[string]struct{}
	once     sync.Once
	loadErr  error
)

func LoadBadWords(path string) error {

	once.Do(func() {

		file, err := os.Open(path)

		if err != nil {
			loadErr = fmt.Errorf(
				"open badwords file: %w",
				err,
			)
			return
		}

		defer file.Close()

		badWords = make(map[string]struct{})

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {

			word := strings.TrimSpace(
				strings.ToUpper(scanner.Text()),
			)

			if word == "" {
				continue
			}

			badWords[word] = struct{}{}
		}

		if err := scanner.Err(); err != nil {
			loadErr = fmt.Errorf(
				"scan badwords file: %w",
				err,
			)
			return
		}
	})

	return loadErr
}

func ContainsBadWords(text string) bool {

	normalized := strings.ToUpper(text)

	for word := range badWords {

		if strings.Contains(normalized, word) {
			return true
		}
	}

	return false
}
