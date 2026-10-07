# Go-reloaded

Go-reloaded est un programme écrit en Go permettant de modifier et corriger automatiquement un texte à partir d'instructions présentes dans celui-ci.

## Fonctionnalités

Le programme permet de :

- convertir les nombres hexadécimaux avec `(hex)` ;
- convertir les nombres binaires avec `(bin)` ;
- convertir des mots en majuscules avec `(up)` ;
- convertir des mots en minuscules avec `(low)` ;
- mettre une majuscule initiale avec `(cap)` ;
- appliquer ces transformations à plusieurs mots avec `(up, n)`, `(low, n)` et `(cap, n)` ;
- corriger la ponctuation ;
- corriger les apostrophes ;
- transformer `a` en `an` lorsque le mot suivant commence par une voyelle ou un `h`.

## Utilisation

Le programme prend deux arguments :

1. le fichier contenant le texte à modifier ;
2. le fichier dans lequel enregistrer le résultat.

### Lancer le programme

```bash
go run . input.txt output.txt
```

Exemple :

```bash
go run . sample.txt result.txt
```

Le fichier `sample.txt` est lu, les transformations sont appliquées, puis le résultat est écrit dans `result.txt`.

## Structure du projet

```text
go-reloaded/
├── go.mod
├── main.go
├── hex.go
├── bin.go
├── case.go
├── punctuation.go
├── apostrophe.go
├── article.go
├── sample.txt
└── README.md
```

## Prérequis

- Go installé sur la machine.

## Tests

Pour tester le programme, créer un fichier texte contenant les différentes instructions, puis lancer :

```bash
go run . sample.txt result.txt
```

Le résultat peut ensuite être consulté dans le fichier de sortie.# go-reloaded-
