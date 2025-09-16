package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/leoarpin/spelling-bee-solver/internal/dict"
	"github.com/leoarpin/spelling-bee-solver/internal/solver"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// demander les 7 lettres
	fmt.Print("Quelles sont les 7 lettres présentes? ")
	lettersInput, _ := reader.ReadString('\n')
	lettersInput = strings.TrimSpace(lettersInput)

	// demander la lettre obligatoire
	fmt.Print("Quelle est la lettre obligatoire? ")
	centerInput, _ := reader.ReadString('\n')
	centerInput = strings.TrimSpace(centerInput)

	if len(centerInput) != 1 {
		log.Fatalf("Erreur: la lettre obligatoire doit être un seul caractère")
	}

	// charger le dictionnaire
	words, err := dict.Load("assets/words_cleaned.txt")
	if err != nil {
		log.Fatalf("Erreur chargement dictionnaire: %v", err)
	}

	// construire le puzzle
	p := solver.Puzzle{
		Center:  rune(centerInput[0]),
		Letters: solver.MakeLetterSet(lettersInput),
		MinLen:  4, // règle standard Spelling Bee
	}

	// résoudre
	results := solver.Solve(words, p)

	// afficher
	fmt.Println("\nLISTE DES MOTS JOUABLES:")
	for _, w := range results {
		fmt.Println(w)
	}
}
