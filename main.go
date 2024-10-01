package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	HEIGHT    int = 50
	WIDTH     int = 20
	FRAMETIME int = 200
)

func generate_grid() [][]bool {
	new_grid := make([][]bool, WIDTH)
	for i := range new_grid {
		new_grid[i] = make([]bool, HEIGHT)
	}
	return new_grid
}

func print_grid(g [][]bool) {
	var grid string
	for i := 0; i < WIDTH; i++ {
		for j := 0; j < HEIGHT; j++ {
			if g[i][j] == true {
				grid += "*"
			} else {
				grid += " "
			}
		}
		grid += "\n"
	}
	fmt.Print(grid)
}

func count_cells_and_update(g [][]bool) [][]bool {
	pos := [8][2]int {
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0},           {1,  0},
		{-1, 1},  {0,  1}, {1,  1},
	}

	new_grid := make([][]bool, WIDTH)
	for i := range new_grid {
		new_grid[i] = make([]bool, HEIGHT)
	}
	
	for i := 0; i < WIDTH; i++ {
		for j := 0; j < HEIGHT; j++ {
			amount := 0
			for k := 0; k < 8; k++ {
				ni := i + pos[k][0]
				nj := j + pos[k][1]

				if ni >= 0 && ni < WIDTH && nj >= 0 && nj < HEIGHT && g[ni][nj] {
					amount++
				}
			}

			if !g[i][j] && amount == 3 { // Toda célula morta com exatamente três vizinhos vivos torna-se viva (nascimento).
				new_grid[i][j] = true
			} else if g[i][j] && amount < 2 { // Toda célula viva com menos de dois vizinhos vivos morre por isolamento.
				new_grid[i][j] = false
			} else if g[i][j] && amount > 3 { // Toda célula viva com mais de três vizinhos vivos morre por superpopulação.
				new_grid[i][j] = false
			} else if g[i][j] && (amount == 2 || amount == 3) { // Toda célula viva com dois ou três vizinhos vivos permanece viva.
				new_grid[i][j] = true
			}

			amount = 0
		}
	}

	return new_grid
}

func generate_random_grid(amount int) [][]bool {
	new_grid := make([][]bool, WIDTH)
	for i := range new_grid {
		new_grid[i] = make([]bool, HEIGHT)
	}

	rand.Seed(time.Now().Unix())
	for i := 0; i < amount; i++ {
		numh := rand.Intn(HEIGHT)
		numw := rand.Intn(WIDTH)
		new_grid[numw][numh] = true
	}

	return new_grid
}

func run_game_in_terminal(g[][]bool) {
	grid := g
	for {
		fmt.Print("\033[H\033[2J") // Limpa o terminal
		print_grid(grid)
		grid = count_cells_and_update(grid)
		time.Sleep(time.Duration(FRAMETIME) * time.Millisecond)
	}
}

func main() {
	grid := generate_grid()
	grid = generate_random_grid(300)
	run_game_in_terminal(grid)
}
