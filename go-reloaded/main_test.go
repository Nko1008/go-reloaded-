package main

import "testing"

func TestProcessText(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"conversion hexadecimale", "1E (hex) fichiers ont ete ajoutes", "30 fichiers ont ete ajoutes"},
		{"conversion binaire", "Cela fait 10 (bin) ans", "Cela fait 2 ans"},
		{"deux conversions", "Ajoute 42 (hex) et 10 (bin) pour obtenir le resultat", "Ajoute 66 et 2 pour obtenir le resultat"},
		{"majuscules", "Pret, partez, go (up) !", "Pret, partez, GO!"},
		{"minuscules", "Arrete de CRIER (low)", "Arrete de crier"},
		{"premiere lettre en majuscule", "Le pont de Brooklyn bridge (cap)", "Le pont de Brooklyn Bridge"},
		{"majuscules sur deux mots", "Ce jeu est vraiment genial (up, 2)", "Ce jeu est VRAIMENT GENIAL"},
		{"majuscule sur trois mots", "bonjour tout le monde (cap,3)", "bonjour Tout Le Monde"},
		{"nombre superieur au nombre de mots", "tres bien (up, 10)", "TRES BIEN"},
		{"plusieurs transformations", "bonjour (cap) tout va bien (up)", "Bonjour tout va BIEN"},
		{"correction de la ponctuation", "Je suis passe par la ,et ensuite BOUM !!", "Je suis passe par la, et ensuite BOUM!!"},
		{"points de suspension", "Cette histoire est ... assez longue ,tu ne trouves pas ?", "Cette histoire est... assez longue, tu ne trouves pas?"},
		{"ponctuation en debut de texte", ", bonjour", ", bonjour"},
		{"nombres avec ponctuation", "Le prix est 3.14 euros et le numero est 1,000", "Le prix est 3.14 euros et le numero est 1,000"},
		{"apostrophes autour d'un mot", "Ce mot est : ' genial '", "Ce mot est: 'genial'"},
		{"apostrophes autour d'une phrase", "Elle a dit : ' je suis tres contente '", "Elle a dit: 'je suis tres contente'"},
		{"apostrophe dans un mot", "aujourd'hui tout va bien", "aujourd'hui tout va bien"},
		{"remplacement de a par an", "Il raconte a histoire incroyable", "Il raconte an histoire incroyable"},
		{"article avec majuscule", "C'etait A incroyable aventure", "C'etait An incroyable aventure"},
		{"article devant un h", "C'est a histoire", "C'est an histoire"},
		{"article devant une consonne", "C'est a livre", "C'est a livre"},
		{"traitement de plusieurs lignes", "1E (hex) fichiers\nbonjour (cap) tout le monde", "30 fichiers\nBonjour tout le monde"},
		{"conservation des lignes vides", "un\n\ndeux", "un\n\ndeux"},
		{"nombre hexadecimal invalide", "zz (hex) test", "zz (hex) test"},
		{"sujet hex", "1E (hex) files were added", "30 files were added"},
		{"sujet bin", "It has been 10 (bin) years", "It has been 2 years"},
		{"sujet up", "Ready, set, go (up) !", "Ready, set, GO!"},
		{"sujet low", "I should stop SHOUTING (low)", "I should stop shouting"},
		{"sujet cap", "Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"sujet up avec nombre", "This is so exciting (up, 2)", "This is SO EXCITING"},
		{"sujet ponctuation", "I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"sujet points de suspension", "I was thinking ... You were right", "I was thinking... You were right"},
		{"sujet apostrophes un mot", "I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"sujet apostrophes plusieurs mots", "As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
		{"sujet article", "There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"sujet hex et bin", "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},
		{"sujet article an", "There is no greater agony than bearing a untold story inside you.", "There is no greater agony than bearing an untold story inside you."},
		{"sujet ponctuation 2", "Punctuation tests are ... kinda boring ,what do you think ?", "Punctuation tests are... kinda boring, what do you think?"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processText(tt.input); got != tt.want {
				t.Errorf("processText(%q)\n obtenu : %q\n attendu : %q", tt.input, got, tt.want)
			}
		})
	}
}
