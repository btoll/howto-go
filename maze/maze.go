package maze

import (
	"fmt"
	"math/rand"
	"time"
)

type Row []string
type Grid []Row

type Cell struct {
	R int
	C int
}

type Maze struct {
	Grid     Grid
	WallChar string
	Height   int
	Width    int
}

func GetRand(max int) int {
	min := 1
	return rand.Intn(max-min) + min
}

func New(height, width int) *Maze {
	return &Maze{
		Grid:     nil,
		WallChar: "+",
		Height:   (height * 2) + 1,
		Width:    (width * 2) + 1,
	}
}

func (m *Maze) Create() Grid {
	grid := make(Grid, m.Height)
	for i := range grid {
		grid[i] = make(Row, m.Width)
		for j := 0; j < m.Width; j += 1 {
			grid[i][j] = m.WallChar
		}
	}
	m.Grid = grid
	return grid
}

func (m *Maze) Draw() {
	for _, row := range m.Grid {
		fmt.Println(row)
	}
}

func (m *Maze) Generate() Grid {
	r, c := 1, 3
	track := []Cell{
		Cell{
			R: r,
			C: c,
		},
	}
	m.Grid[r][c] = " "
	rand.Seed(time.Now().Unix())

	for len(track) > 0 {
		current := track[len(track)-1]
		neighbors := m.GetNeighbors(current)
		if len(neighbors) == 0 {
			track = track[:len(track)-1]
		} else {
			neighbor := neighbors[0]
			m.Grid[neighbor.R][neighbor.C] = " "
			m.Grid[(neighbor.R+current.R)/2][(neighbor.C+current.C)/2] = " "
			track = append(track, neighbor)
		}
	}
	for _, entrance := range m.MakeEntrances() {
		m.Grid[entrance.R][entrance.C] = " "
	}
	return m.Grid
}

func (m *Maze) GetRandStep(max int) int {
	n := GetRand(max)
	for n%2 == 0 {
		n = GetRand(max)
	}
	return n
}

func (m *Maze) MakeEntrances() []Cell {
	//	side := GetRand()
	start := Cell{
		R: 0,
		C: m.GetRandStep(m.Width),
	}
	end := Cell{
		R: m.Height - 1,
		C: m.GetRandStep(m.Width),
	}
	return []Cell{start, end}
}

func (m *Maze) GetNeighbors(cell Cell) []Cell {
	if m.Grid == nil {
		panic("There is no grid!")
	}
	// n -> r, c - 1
	// s -> r, c + 1
	// e -> r + 1, c
	// w -> r - 1, c
	r := cell.R
	c := cell.C
	neighbors := []Cell{}
	if r > 1 && m.Grid[r-2][c] == m.WallChar {
		neighbors = append(neighbors, Cell{
			R: r - 2,
			C: c,
		})
	}
	if r < m.Width-2 && m.Grid[r+2][c] == m.WallChar {
		neighbors = append(neighbors, Cell{
			R: r + 2,
			C: c,
		})
	}
	if c > 1 && m.Grid[r][c-2] == m.WallChar {
		neighbors = append(neighbors, Cell{
			R: r,
			C: c - 2,
		})
	}
	if c < m.Height-2 && m.Grid[r][c+2] == m.WallChar {
		neighbors = append(neighbors, Cell{
			R: r,
			C: c + 2,
		})
	}
	rand.Shuffle(len(neighbors), func(i, j int) {
		neighbors[i], neighbors[j] = neighbors[j], neighbors[i]
	})
	return neighbors
}
