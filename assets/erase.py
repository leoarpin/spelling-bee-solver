# clean_dict.py
def clean_dictionary(input_file, output_file):
    with open(input_file, "r", encoding="utf-8") as f:
        words = f.readlines()

    # garder seulement les mots de 3 lettres ou plus
    cleaned = [w.strip() for w in words if len(w.strip()) > 3]

    with open(output_file, "w", encoding="utf-8") as f:
        for w in cleaned:
            f.write(w + "\n")

    print(f"{len(words) - len(cleaned)} mots supprimés, {len(cleaned)} mots gardés.")


if __name__ == "__main__":
    clean_dictionary("assets/words_alpha.txt", "words_cleaned.txt")
