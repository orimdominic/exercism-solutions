package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

// Define the Garden type here.
type Garden struct {
	diagram  []string
	children []string
	chPlants map[string][]string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	dgmPlants := strings.Split(diagram, "\n")[1:]
	numOfRows := len(dgmPlants)
	if numOfRows != 2 {
		return nil, fmt.Errorf("wrong diagram format")
	}

	dgmPlantsCount := 0
	for _, v := range dgmPlants {
		dgmPlantsCount += len(v)
	}

	if dgmPlantsCount%4 != 0 {
		return nil, fmt.Errorf("odd number of cups")
	}

	ch := append(make([]string, 0), children...)
	sort.Strings(ch)

	chNameTable := make(map[string]bool)
	for _, name := range ch {
		_, ok := chNameTable[name]
		if ok {
			return nil, fmt.Errorf("duplicate name")
		}
		chNameTable[name] = true
	}

	for _, line := range dgmPlants {
		for _, r := range line {
			if r != 'G' && r != 'C' && r != 'R' && r != 'V' {
				return nil, fmt.Errorf("invalid cup codes")
			}
		}
	}

	charPlantMap := map[byte]string{
		'G': "grass",
		'C': "clover",
		'V': "violets",
		'R': "radishes",
	}

	chPlants := make(map[string][]string, 0)
	for i := 0; i < len(ch); i++ {
		chPlants[ch[i]] = []string{
			charPlantMap[dgmPlants[0][i*2]],
			charPlantMap[dgmPlants[0][i*2+1]],
			charPlantMap[dgmPlants[1][i*2]],
			charPlantMap[dgmPlants[1][i*2+1]],
		}
	}

	return &Garden{
		children: ch,
		diagram:  dgmPlants,
		chPlants: chPlants,
	}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := g.chPlants[child]

	return plants, ok
}
