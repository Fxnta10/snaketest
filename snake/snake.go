package snake

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

const MAP_SIZE = 15
const _FRAME_INTERVAL = time.Millisecond * 1000

func RunSnake(dir *keyboard.Key) {

	var gamemap [MAP_SIZE][MAP_SIZE]string
	for i := 0; i < len(gamemap); i++ {
		for j := 0; j < len(gamemap[0]); j++ {
			gamemap[i][j] = " "
		}
	}

	var snakepos [][2]int
	snakepos = append(snakepos, [2]int{MAP_SIZE / 2, MAP_SIZE / 2}) //headcolumn
	snakepos = append(snakepos, [2]int{MAP_SIZE / 2, (MAP_SIZE / 2) - 1})
	snakemap := sliceToMap(snakepos)

	foodrow, foodcolumn := getFood(gamemap)
	var gamend bool
	tempkey := "65514"

	for !gamend {
		time.Sleep(_FRAME_INTERVAL)
		tempkey = fmt.Sprint(*dir)
		gamemap[foodrow][foodcolumn] = "▣"
		printMap(gamemap, snakemap)
		cords := changePos(snakepos[0][0], snakepos[0][1], tempkey)
		if cords[0] == foodrow && cords[1] == foodcolumn {
			gamemap[foodrow][foodcolumn] = " "
			foodrow, foodcolumn = getFood(gamemap)
		} else { //else if cords[0] != foodrow && cords[1] != foodcolumn { //if food not eaten
			snakepos = pop(snakepos)
		}
		snakepos = pushBack(snakepos, cords)
		snakemap = sliceToMap(snakepos) //map
		gamend = checkEnd(snakepos)

	}
}

func pop(snakepos [][2]int) [][2]int {
	return snakepos[:len(snakepos)-1]
}

func pushBack(snakepos [][2]int, value [2]int) [][2]int {

	snakepos = append(snakepos, [2]int{0, 0})

	copy(snakepos[1:], snakepos)

	snakepos[0] = value
	return snakepos
}

func getKey() (string, string) {
	char, key, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	// fmt.Println("...")
	return fmt.Sprint(char), fmt.Sprint(key)

}

func changePos(headrow int, headcolumn int, key string) [2]int {
	if key == "65517" { // up
		headrow--
	}
	if key == "65516" { //down
		headrow++
	}
	if key == "65515" { //left
		headcolumn--
	}
	if key == "65514" { //right
		headcolumn++
	}
	var cords [2]int
	cords[0] = headrow
	cords[1] = headcolumn
	return cords
}

func printMap(gamemap [MAP_SIZE][MAP_SIZE]string, snakemap map[[2]int]bool) {

	clearScreen()
	var cords [2]int
	// row := 1
	// column := 0
	fmt.Println(strings.Repeat("-", MAP_SIZE))

	for i := 0; i < MAP_SIZE; i++ {
		for j := 0; j < MAP_SIZE; j++ {
			// rowlenght := len(snakepos)
			// columnlenght := len(snakepos[0])
			// if i > rowlenght || column > columnlenght { //to prevent index out of range errors
			// 	fmt.Print(gamemap[i][j])
			// 	continue
			// }
			cords[0] = i
			cords[1] = j
			if checkSnake(snakemap, cords) == "head" {
				fmt.Print("◈") //the head of the snake

			}
			if checkSnake(snakemap, cords) == "body" {
				fmt.Print("▫")
			}
			if checkSnake(snakemap, cords) == "no snake" {
				fmt.Print(gamemap[i][j])
			}
			// fmt.Println(i, " ", j)
			// if i == snakepos[0][0] && j == snakepos[0][1] {
			// 	fmt.Print("◈") //the head of the snake
			// 	continue
			// } else if i == snakepos[row][column] && j == snakepos[row][column+1] {
			// 	fmt.Print("▫")
			// 	row++
			// 	column = 0
			// 	break
			// } else {

			// 	fmt.Print(gamemap[i][j])
			// }
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("-", MAP_SIZE))

}

func getFood(gamemap [MAP_SIZE][MAP_SIZE]string) (int, int) {
	foodrow := rand.Intn(MAP_SIZE)
	foodcolumn := rand.Intn(MAP_SIZE)
	for gamemap[foodrow][foodcolumn] != " " {
		foodrow = rand.Intn(MAP_SIZE)
		foodcolumn = rand.Intn(MAP_SIZE)
	}
	return foodrow, foodcolumn
}
func checkEnd(snakepos [][2]int) bool {
	headrow := snakepos[0][0]
	headcolumn := snakepos[0][1]
	for i := 1; i < len(snakepos); i++ { // when head and the body conincide
		for j := 0; j < 1; j++ {
			if headrow == snakepos[i][j] && headcolumn == snakepos[i][j+1] {
				fmt.Println("GAME OVER !!!")
				return true
			}
		}
	}
	if headrow < 0 || headcolumn < 0 || headrow >= MAP_SIZE || headcolumn >= MAP_SIZE { //the head is out of the map
		fmt.Println("GAME OVER !!!")
		return true
	}
	return false
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func checkSnake(snakemap map[[2]int]bool, cords [2]int) string {
	val, ok := snakemap[cords]
	if !ok {
		return "no snake"
	} else if val {
		return "head"
	} else {
		return "body"
	}
}

func sliceToMap(snakepos [][2]int) map[[2]int]bool {
	snakemap := make(map[[2]int]bool)

	for i, val := range snakepos {
		if i == 0 {
			snakemap[val] = true
		} else {
			snakemap[val] = false
		}
	}

	return snakemap
}
