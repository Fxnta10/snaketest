package parallel

import (
	snake "snake/snake"

	"github.com/eiannone/keyboard"
)

func ParallelTest() {

	var i keyboard.Key = 65514
	var gameEnd bool = false

	// go print(&i)
	go snake.RunSnake(&i, &gameEnd)

	for {
		if gameEnd {
			return
		}
		i = getKey()

	}

}

func getKey() keyboard.Key {
	_, key, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}

	if key == keyboard.KeyArrowRight {
		return key
	}
	if key == keyboard.KeyArrowDown {
		return key
	}
	if key == keyboard.KeyArrowUp {
		return key
	}
	if key == keyboard.KeyArrowLeft {
		return key
	}
	return 0
}
