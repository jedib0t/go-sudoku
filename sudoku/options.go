package sudoku

import (
	"math/rand"
	"time"
)

// Option helps customize the Grid.
type Option func(g *Grid)

var (
	defaultOpts = []Option{
		WithRNG(rand.New(rand.NewSource(time.Now().UnixNano()))),
		WithSubGrids([]SubGrid{
			{Locations: []Location{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}},
			{Locations: []Location{{0, 3}, {0, 4}, {0, 5}, {1, 3}, {1, 4}, {1, 5}, {2, 3}, {2, 4}, {2, 5}}},
			{Locations: []Location{{0, 6}, {0, 7}, {0, 8}, {1, 6}, {1, 7}, {1, 8}, {2, 6}, {2, 7}, {2, 8}}},
			{Locations: []Location{{3, 0}, {3, 1}, {3, 2}, {4, 0}, {4, 1}, {4, 2}, {5, 0}, {5, 1}, {5, 2}}},
			{Locations: []Location{{3, 3}, {3, 4}, {3, 5}, {4, 3}, {4, 4}, {4, 5}, {5, 3}, {5, 4}, {5, 5}}},
			{Locations: []Location{{3, 6}, {3, 7}, {3, 8}, {4, 6}, {4, 7}, {4, 8}, {5, 6}, {5, 7}, {5, 8}}},
			{Locations: []Location{{6, 0}, {6, 1}, {6, 2}, {7, 0}, {7, 1}, {7, 2}, {8, 0}, {8, 1}, {8, 2}}},
			{Locations: []Location{{6, 3}, {6, 4}, {6, 5}, {7, 3}, {7, 4}, {7, 5}, {8, 3}, {8, 4}, {8, 5}}},
			{Locations: []Location{{6, 6}, {6, 7}, {6, 8}, {7, 6}, {7, 7}, {7, 8}, {8, 6}, {8, 7}, {8, 8}}},
		}),
	}
)

// WithRNG customizes the Random Number Generator used by the Grid.
func WithRNG(rng *rand.Rand) Option {
	return func(g *Grid) {
		g.rng = rng
	}
}

// WithSubGrids customizes the way Sub-Grids are defined in the Grid.
func WithSubGrids(sgs []SubGrid) Option {
	return func(g *Grid) {
		g.subGrids = make([]SubGrid, len(sgs))
		for idx, sg := range sgs {
			g.subGrids[idx] = SubGrid{
				Locations: sg.Locations,
				grid:      g,
			}
		}
	}
}
