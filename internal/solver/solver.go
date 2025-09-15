package solver

//find all words that can be formed with the letters given
import "sort"

type Puzzle struct {
	Center  rune
	Letters map[rune]bool
	MinLen  int
}

func MakeLetterSet(s string) map[rune]bool {
	m := make(map[rune]bool, len(s))
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r = r + 32 // to lower
		}
		m[r] = true
	}
	return m
}

func validWord(w string, p Puzzle) bool {
	if len(w) < p.MinLen {
		return false
	}
	hasCenter := false
	for _, r := range w {
		if !p.Letters[r] {
			return false
		}
		if r == p.Center {
			hasCenter = true
		}
	}
	return hasCenter
}

func Solve(words []string, p Puzzle) []string {
	out := make([]string, 0, 256)
	for _, w := range words {
		if validWord(w, p) {
			out = append(out, w)
		}
	}
	sort.Strings(out)
	return out
}
