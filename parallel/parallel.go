package parallel

import (
	"fmt"
	snake "snake/snake"
	"time"

	"github.com/eiannone/keyboard"
)

const FRAME_INTERVAL = time.Millisecond * 200

func ParallelTest() {

	var i keyboard.Key = 65514

	// go print(&i)
	go snake.RunSnake(&i)

	for {
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

func print(dir *keyboard.Key) {
	count := 0
	for {
		time.Sleep(FRAME_INTERVAL)
		fmt.Println(count, *dir)
		count++
	}
}
