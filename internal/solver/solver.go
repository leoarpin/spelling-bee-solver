package solver

import "sort"

// Puzzle représente un défi Spelling Bee.
type Puzzle struct {
	Center  rune          // la lettre obligatoire
	Letters map[rune]bool // ensemble des 7 lettres autorisées
	MinLen  int           // longueur minimale d’un mot (par défaut 4)
}

// MakeLetterSet construit un ensemble de lettres autorisées à partir d’une string.
func MakeLetterSet(s string) map[rune]bool {
	m := make(map[rune]bool, len(s))
	for _, r := range s {
		m[r] = true
	}
	return m
}

// validWord vérifie si un mot respecte les règles Spelling Bee.
func validWord(w string, p Puzzle) bool {
	if len(w) < p.MinLen {
		return false
	}
	hasCenter := false
	for _, r := range w {
		if !p.Letters[r] { // lettre non autorisée
			return false
		}
		if r == p.Center { // contient la lettre centrale
			hasCenter = true
		}
	}
	return hasCenter
}

// Solve retourne tous les mots valides du dictionnaire.
func Solve(words []string, p Puzzle) []string {
	var out []string
	for _, w := range words {
		if validWord(w, p) {
			out = append(out, w)
		}
	}
	sort.Strings(out) // tri alphabétique pour lisibilité
	return out
}
