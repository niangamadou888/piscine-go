/* Projet sudoku

Team : Adama Niasse - Mouhamed Niang - Amadou Tidiane Diallo 

Réalisation d'un programme qui résoud un sudoku en Golang
Ce code est une implémentation en Go d'un programme pour résoudre des grilles de Sudoku. Il contient plusieurs fonctions qui sont utilisées pour vérifier la validité d'une grille de Sudoku et pour résoudre la grille en remplissant les cases vides.

Les fonctions principales sont:

    atoi(s string) int : Cette fonction convertit un chaîne de caractères en un entier.
    int_Table(input string) [9][9]int : Cette fonction convertit chaque argument en un tableau de 9 entiers.
    IsDuplicate(counter [10]int) bool : Cette fonction vérifie si un nombre est dupliqué sur une même ligne ou colonne.
    CaseVide(table *[9][9]int) bool : Cette fonction detecte les cases vides.
    isTableValid(table *[9][9]int) bool : Cette fonction vérifie si le tableau à résoudre est valide.
    Resolution(table *[9][9]int) bool : Cette fonction résout le sudoku en remplaçant chaque case vide par le nombre qu'il faut.

Il y a aussi des variables globales :

    count int : qui est utilisée pour compter le nombre de solutions possibles pour une grille donnée.
    parfait bool : qui est utilisée pour vérifier si la grille donnée est parfaite ou non.


*/


package main

import (
	"fmt"
	"os"
)

var count int
var parfait bool



//----------------------Cette fonction converti un string en int

func atoi(s string) int {
	nbr := 0
	ind := 0
	if len(s) == 0 {
		ind = ind + 0
	} else if s[0] == '-' || s[0] == '+' {
		ind = ind + 1
	}
	for ; ind < len(s); ind++ {

		if s[ind] < '0' || s[ind] > '9' {
			return 0
		}
		if s[0] == '-' {
			nbr = nbr*10 - int(s[ind]-'0')
		} else {
			nbr = nbr*10 + int(s[ind]-'0')
		}
	}

	return nbr
}


//----------------------Cette fonction converti chaque argument en un tableau de 9 entiers
func int_Table(input string) [9][9]int {
	table := [9][9]int{}
	i := 0
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			conv := atoi(string(input[i]))
			table[x][y] = conv
			i++
		}
	}
	return table
}


//----------------------Cette fonction vérifie si un nombre est dupliqué sur une meme ligne ou colonne
func IsDuplicate(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}


//-------------------------------------Cette fonction detecte les cages vides
func CaseVide(table *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == 0 {
				return true
			}
		}
	}
	return false
}


//----------------------------------------Cette fonction vérifie si le tableau à résoudre est valide
func isTableValid(table *[9][9]int) bool {

	for x := 0; x < 9; x++ {
		counter := [10]int{}
		for y := 0; y < 9; y++ {
			counter[table[x][y]]++
		}
		if IsDuplicate(counter) {
			return false
		}
	}

	for x := 0; x < 9; x++ {
		counter := [10]int{}
		for y := 0; y < 9; y++ {
			counter[table[y][x]]++
		}
		if IsDuplicate(counter) {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					counter[table[x][y]]++
				}
				if IsDuplicate(counter) {
					return false
				}
			}
		}
	}

	return true
}



//----------------------------------------Cette fonction résoud le sudoku en remplaçant chaque cage vide par son par le nombre qu'il faut
func Resolution(table *[9][9]int) bool {
	if count > 1 && !parfait {
		return false
	}
	if !CaseVide(table) {
		count++

		if parfait {
			return true
		}
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == 0 {
				parfait2 := false
				for element := 9; element >= 1; element-- {
					table[i][j] = element

					if isTableValid(table) {
						if Resolution(table) {
							if !parfait {
								parfait2 = true
							} else {
								return true
							}
						}
						table[i][j] = 0
					} else {
						table[i][j] = 0
					}
				}
				if parfait2 {
					return true
				}
				return false
			}
		}
	}
	return false
}

//---------------------Cette fonction prend en parametre un tableau de 9 linges 9 colonnes et l'affiche
func AfficherTable(table [9][9]int) {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if y != 8 {
				fmt.Printf("%d ", table[x][y])
			} else {
				fmt.Printf("%d", table[x][y])
			}
		}
		fmt.Println()
	}
}


func main() {
// -------------------------------------------------------Gestion des erreurs de saisie---------------------------------------
	if len(os.Args[1:]) != 9 {
		fmt.Println("Error")
		return
	}

	for _, arg := range os.Args[1:] {
		if len(arg) != 9 {
			fmt.Println("Error")
			return
		}

		for _, ch := range arg {
			if ch < '0' || ch > '9' {
				if ch != '.' {
					fmt.Println("Error")
					return
				}
			}
		}
	}

//--------------------------------------------------------------------
	arg := os.Args[1:]
	StringTable:= ""
	for _, v := range arg {
		for i := 0; i <= 8; i++ {
			if v[i] == '.' {
				StringTable= StringTable+ "0" // Pour remplacer les "." par des 0  Exemple: "167.8..2." => "167080020"
			} else {
				StringTable= StringTable+ string(v[i])
			}

		}
	}

	table := int_Table(StringTable) // Convertir notre string( "167.8..2.") en tableau de int [1 6 7 0 8 0 0 2 0]

	if Resolution(&table) {
		if count == 1 {
			parfait = true
			Resolution(&table)
			AfficherTable(table)
		} else {
			fmt.Println("Error")
		}
	} else {
		fmt.Println("Error")
	}

} 
