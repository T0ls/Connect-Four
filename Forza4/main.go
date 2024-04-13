package main

import (
	. "fmt"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

var pl1 string
var pl2 string

type move struct {
	x int
	y int
}

func CheckOcc(i int, j int, result [4][2]int) bool {
	for z := 0; z < 4; z++ {
		if i == result[z][0] && j == result[z][1] {
			return true
		}
	}
	return false
}

func PrintCampoWin(campo [12][13]string, result [4][2]int) {
	c1 := color.New(color.FgGreen)
	c2 := color.New(color.FgRed)
	c3 := color.New(color.FgBlue)
	for i := 3; i < 9; i++ {
		switch i {
		case 3:
			Println("╔═══╦═══╦═══╦═══╦═══╦═══╦═══╗")
		default:
			Println("╠═══╬═══╬═══╬═══╬═══╬═══╬═══╣")
		}
		for j := 3; j < 10; j++ {
			Print("║ ")
			if CheckOcc(i, j, result) {
				c1.Print(campo[i][j])
			} else {
				if campo[i][j] == pl1 {
					c2.Print(campo[i][j])
				} else {
					c3.Print(campo[i][j])
				}
			}
			Print(" ")
		}
		Println("║")
		if i == 8 {
			Println("╚═══╩═══╩═══╩═══╩═══╩═══╩═══╝")
		}
	}
}

func PrintCampoPrev(campoPrev [12][13]string, x int, y int, player string) {
	c1 := color.New(color.FgYellow)
	c2 := color.New(color.FgRed)
	c3 := color.New(color.FgBlue)
	Println()
	if x == 3 {
		Print("  ")
	} else {
		Print(" ")
		for r := 1; r <= x-3; r++ {
			Print("    ")
		}
		Print(" ")
	}
	c1.Println("\u25BC")
	for i := 3; i < 9; i++ {
		switch i {
		case 3:
			Println("╔═══╦═══╦═══╦═══╦═══╦═══╦═══╗")
		default:
			Println("╠═══╬═══╬═══╬═══╬═══╬═══╬═══╣")
		}
		for j := 3; j < 10; j++ {
			if i == y && j == x {
				Print("║ ")
				c1.Print(player)
				Print(" ")
			} else {
				Print("║ ")
				if campoPrev[i][j] == pl1 {
					c2.Print(campoPrev[i][j])
				} else {
					c3.Print(campoPrev[i][j])
				}
				Print(" ")
			}
		}
		Println("║")
		if i == 8 {
			Println("╚═══╩═══╩═══╩═══╩═══╩═══╩═══╝")
		}
	}
}

func DisponibilitàMossa(campo [12][13]string, pos int) (int, bool) {
	max := false
	for i := 8; i >= 3; i-- {
		if campo[i][pos] == " " {
			return i, max
		}
	}
	return 0, true
}

func MossaPlayer(campo [12][13]string, player string) ([12][13]string, move) {
	campoPrev := campo
	var b bool
	var mossa move
	posH := 6 // va da 0 a 6
	posV, _ := DisponibilitàMossa(campo, posH)
	posPH := posH
	posPV := posV
	PrintCampoPrev(campoPrev, posH, posV, player)
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		// Eseguo il codice in base all'input dell'utente
		if key == keyboard.KeyArrowLeft || char == 'a' { // Left || a
			os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
			posH--
			if posH < 3 {
				posH = 9
			}
			posV, b = DisponibilitàMossa(campo, posH)
			if !b {
				campoPrev[posPV][posPH] = " "
				campoPrev[posV][posH] = player
				posPH, posPV = posH, posV
			} else {
				campoPrev[posPV][posPH] = " "
			}
			PrintCampoPrev(campoPrev, posH, posV, player)
		} else if key == keyboard.KeyArrowRight || char == 'd' { // Right || d
			os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
			posH++
			if posH > 9 {
				posH = 3
			}
			posV, b = DisponibilitàMossa(campo, posH)
			if !b {
				campoPrev[posPV][posPH] = " "
				campoPrev[posV][posH] = player
				posPH, posPV = posH, posV
			} else {
				campoPrev[posPV][posPH] = " "
			}
			PrintCampoPrev(campoPrev, posH, posV, player)
		} else if key == keyboard.KeyEnter || key == keyboard.KeySpace { // Enter || SpaceBar
			if campo[posV][posH] == " " {
				campoPrev[posV][posH] = player
				posPH, posPV = posH, posV
				mossa.y = posV
				mossa.x = posH
				break
			}
		}
		//pulisco la scelta
	}
	return campoPrev, mossa
}

func CheckLW(campo [12][13]string, player string, mossa move) (bool, [4][2]int) {
	//check ⟵
	if campo[mossa.y][mossa.x-1] == player && campo[mossa.y][mossa.x-2] == player && campo[mossa.y][mossa.x-3] == player {
		if campo[mossa.y][mossa.x+1-3] == player && campo[mossa.y][mossa.x+2-3] == player && campo[mossa.y][mossa.x+3-3] == player {
			return true, [4][2]int{
				{mossa.y, mossa.x},
				{mossa.y, mossa.x - 1},
				{mossa.y, mossa.x - 2},
				{mossa.y, mossa.x - 3},
			} //controlla inverso
		}
	}
	return false, [4][2]int{
		{0, 0},
	}
}

func CheckLD(campo [12][13]string, player string, mossa move) (bool, [4][2]int) {
	//check ↙
	if campo[mossa.y+1][mossa.x-1] == player && campo[mossa.y+2][mossa.x-2] == player && campo[mossa.y+3][mossa.x-3] == player {
		if campo[mossa.y+3][mossa.x-3] == player && campo[mossa.y+2][mossa.x-2] == player && campo[mossa.y+1][mossa.x-1] == player && campo[mossa.y][mossa.x] == player {
			return true, [4][2]int{
				{mossa.y, mossa.x},
				{mossa.y + 1, mossa.x - 1},
				{mossa.y + 2, mossa.x - 2},
				{mossa.y + 3, mossa.x - 3},
			} //controlla inverso
		}
	}
	return false, [4][2]int{
		{0, 0},
	}
}

func CheckRD(campo [12][13]string, player string, mossa move) (bool, [4][2]int) {
	//check ↘
	if campo[mossa.y+1][mossa.x+1] == player && campo[mossa.y+2][mossa.x+2] == player && campo[mossa.y+3][mossa.x+3] == player {
		if campo[mossa.y+3][mossa.x+3] == player && campo[mossa.y+2][mossa.x+2] == player && campo[mossa.y+1][mossa.x+1] == player && campo[mossa.y][mossa.x] == player {
			return true, [4][2]int{
				{mossa.y, mossa.x},
				{mossa.y + 1, mossa.x + 1},
				{mossa.y + 2, mossa.x + 2},
				{mossa.y + 3, mossa.x + 3},
			} //controlla inverso
		}
	}
	return false, [4][2]int{
		{0, 0},
	}
}

func Add(mossa move, x int, y int) move {
	mossa.x += x
	mossa.y += y
	return mossa
}

func CheckWin(campo [12][13]string, player string, mossa move) (bool, [4][2]int) {
	var b bool
	var result [4][2]int
	//check ↓
	if campo[mossa.y+1][mossa.x] == player && campo[mossa.y+2][mossa.x] == player && campo[mossa.y+3][mossa.x] == player {
		return true, [4][2]int{
			{mossa.y, mossa.x},
			{mossa.y + 1, mossa.x},
			{mossa.y + 2, mossa.x},
			{mossa.y + 3, mossa.x},
		}
	}

	//check ⟵
	b, result = CheckLW(campo, player, mossa)
	if b {
		return true, result
	} //check ⟵
	//check ⟵X⟶
	b, result = CheckLW(campo, player, Add(mossa, 1, 0))
	if b {
		return true, result
	} //check ⟵
	b, result = CheckLW(campo, player, Add(mossa, 2, 0))
	if b {
		return true, result
	} //check ⟵
	b, result = CheckLW(campo, player, Add(mossa, 3, 0))
	if b {
		return true, result
	} //check ⟵

	//check ↙
	b, result = CheckLD(campo, player, mossa)
	if b {
		return true, result
	}
	//check ↙X↗
	b, result = CheckLD(campo, player, Add(mossa, 1, -1))
	if b {
		return true, result
	} //check ↙
	b, result = CheckLD(campo, player, Add(mossa, 2, -2))
	if b {
		return true, result
	} //check ↙
	b, result = CheckLD(campo, player, Add(mossa, 3, -3))
	if b {
		return true, result
	} //check ↙

	//check ↘
	b, result = CheckRD(campo, player, mossa)
	if b {
		return true, result
	}
	//check ↘X↖
	b, result = CheckRD(campo, player, Add(mossa, -1, -1))
	if b {
		return true, result
	} //check ↘
	b, result = CheckRD(campo, player, Add(mossa, -2, -2))
	if b {
		return true, result
	} //check ↘
	b, result = CheckRD(campo, player, Add(mossa, -3, -3))
	if b {
		return true, result
	} //check ↘
	// Nessuna vittoria trovata
	return false, result
}

func ScambioPlayer(player string) string {
	if player == pl1 {
		return pl2
	}
	return pl1
}

func showMenu(options []string) int {
	c1 := color.New(color.FgRed)
	selectedOption := 0

	println(" Choose you input option: ")
	Println()
	// Loop per leggere l'input dell'utente
	for {
		for i, option := range options {
			if i == selectedOption {
				c1.Printf(" >> %d. %s << \n", i+1, option)
			} else {
				Printf("    %d. %s\n", i+1, option)
			}
		}
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		// Esegui il codice in base all'input dell'utente
		if key == keyboard.KeyArrowUp || char == 'w' {
			selectedOption--
			if selectedOption < 0 {
				selectedOption = len(options) - 1
			}
		} else if key == keyboard.KeyArrowDown || char == 's' {
			selectedOption++
			if selectedOption >= len(options) {
				selectedOption = 0
			}
		} else if key == keyboard.KeyEnter || key == keyboard.KeySpace {
			Println("You chose:", options[selectedOption])
			return selectedOption
		}
		println()
		os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
		println(" Choose you input option: ")
		Println()
	}
}

func ChooseCharacter() (string, string) {
	// p1
	Println("Choose your character player 1: ")
	options := []string{"X", "0", "♥", "†"}
	var p1, p2 string
	selectedOption1 := showMenu(options)
	switch selectedOption1 {
	case 0:
		p1 = "X"
	case 1:
		p1 = "0"
	case 2:
		p1 = "♥"
	case 3:
		p1 = "†"
	}
	// p2
	os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
	z := true
	for z {
		os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
		Println("Choose your character player 2: ")
		selectedOption2 := showMenu(options)
		switch selectedOption2 {
		case 0:
			p2 = "X"
		case 1:
			p2 = "0"
		case 2:
			p2 = "♥"
		case 3:
			p2 = "†"
		}
		if p2 != p1 {
			break
		}
	}
	return p1, p2
}

func Game1vs1() {
	os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
	var campo = [12][13]string{
		{"+", "+", "+", "+", "+", "+", "+", "+", "+", "+", "+", "+", "+"},
		{"+", "+", "+", "+", "+", "+", "+", "+", "+", "+", "+", "+", "+"},
		{"+", "+", "-", "-", "-", "-", "-", "-", "-", "-", "-", "+", "+"},
		{"+", "+", "|", " ", " ", " ", " ", " ", " ", " ", "|", "+", "+"},
		{"+", "+", "|", " ", " ", " ", " ", " ", " ", " ", "|", "+", "+"},
		{"+", "+", "|", " ", " ", " ", " ", " ", " ", " ", "|", "+", "+"},
		{"+", "+", "|", " ", " ", " ", " ", " ", " ", " ", "|", "+", "+"},
		{"+", "+", "|", " ", " ", " ", " ", " ", " ", " ", "|", "+", "+"},
		{"+", "+", "|", " ", " ", " ", " ", " ", " ", " ", "|", "+", "+"},
		{"+", "+", "-", "-", "-", "-", "-", "-", "-", "-", "-", "+", "+"},
		{"+", "+", "+", "+", "+", "+", "+", "+", "+", "+", "+", "+", "+"},
		{"+", "+", "+", "+", "+", "+", "+", "+", "+", "+", "+", "+", "+"},
	}
	pl1, pl2 = ChooseCharacter()
	player := pl1
	var result [4][2]int
	//Partita
	for {
		var mossa move
		campo, mossa = MossaPlayer(campo, player)
		var b bool
		b, result = CheckWin(campo, player, mossa)
		if b {
			os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
			break
		}
		player = ScambioPlayer(player)
		os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
	}
	Println("Il vincitore è:", player)
	PrintCampoWin(campo, result)
}

func main() {
	os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
	//Inizializzo la tastiera
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()
	options := []string{"1 vs 1", "1 vs Pc", "Quit"}
	selectedOption := showMenu(options)
	switch selectedOption {
	case 0:
		Game1vs1()
	case 1:
		Println("Work in Progress!")
		return
	case 2:
		return
	}
}

/*ToDo:
-Features:
	-choose your character
-Fix:
*/

/*Player list:
- X
- 0
- ♥
- †
*/

/*
  ↓	  ↓   ↓   ↓   ↓   ↓   ↓
╔═══╦═══╦═══╦═══╦═══╦═══╦═══╗
║   ║   ║   ║   ║   ║   ║   ║
╠═══╬═══╬═══╬═══╬═══╬═══╬═══╣
║   ║   ║   ║   ║   ║   ║   ║
╠═══╬═══╬═══╬═══╬═══╬═══╬═══╣
║   ║   ║   ║   ║   ║   ║   ║
╠═══╬═══╬═══╬═══╬═══╬═══╬═══╣
║   ║   ║   ║   ║   ║   ║   ║
╠═══╬═══╬═══╬═══╬═══╬═══╬═══╣
║   ║   ║   ║   ║   ║   ║   ║
╠═══╬═══╬═══╬═══╬═══╬═══╬═══╣
║   ║   ║   ║   ║   ║   ║   ║
╚═══╩═══╩═══╩═══╩═══╩═══╩═══╝
*/
