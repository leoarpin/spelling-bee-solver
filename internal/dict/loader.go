package dict

import (
	"bufio"
	"os"
	"strings"
)

// Load lit un fichier dictionnaire (1 mot par ligne)
// et retourne une slice contenant tous les mots.
func Load(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var out []string
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		w := strings.TrimSpace(sc.Text())
		if w != "" {
			out = append(out, w)
		}
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return out, nil
}
