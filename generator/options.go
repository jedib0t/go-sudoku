package generator

import (
	"math/rand"

	"github.com/jedib0t/go-sudoku/sudoku"
)

// Option customized the behavior of one or more Generators.
type Option func(g Generator)

// WithBlockOrder customizes the order of priority of blocks to be filled.
//
// Applies to:
//  * BackTrackingGenerator
func WithBlockOrder(order []int) Option {
	return func(g Generator) {
		switch g.(type) {
		case *backTrackingGenerator:
			g.(*backTrackingGenerator).blockOrder = order
		}
	}
}

// WithDebug enables debug logging.
//
// Applies to:
//  * BackTrackingGenerator
//  * BruteForceGenerator
func WithDebug() Option {
	return func(g Generator) {
		switch g.(type) {
		case *backTrackingGenerator:
			g.(*backTrackingGenerator).baseGenerator.debug = true
		case *bruteForceGenerator:
			g.(*bruteForceGenerator).baseGenerator.debug = true
		}
	}
}

// WithRNG customizes the WithRNG used by the Generator.
//
// Applies to:
//  * BackTrackingGenerator
//  * BruteForceGenerator
func WithRNG(rng *rand.Rand) Option {
	return func(g Generator) {
		switch g.(type) {
		case *backTrackingGenerator:
			g.(*backTrackingGenerator).baseGenerator.rng = rng
		case *bruteForceGenerator:
			g.(*bruteForceGenerator).baseGenerator.rng = rng
		}
	}
}

// WithProgress enables the Generators to show each step of the solution. The
// 'waitInterval' value dictates how long to wait between each cycle.
//
// Applies to:
//  * BackTrackingGenerator
//  * BruteForceGenerator
func WithProgress(rpFn func(g, og *sudoku.Grid, attempts, cycles int)) Option {
	return func(g Generator) {
		switch g.(type) {
		case *backTrackingGenerator:
			g.(*backTrackingGenerator).baseGenerator.renderProgressFn = rpFn
		case *bruteForceGenerator:
			g.(*bruteForceGenerator).baseGenerator.renderProgressFn = rpFn
		}
	}
}

// WithSubGridOrder customizes the order of priority of SubGirds to be filled.
//
// Applies to:
//  * BruteForceGenerator
func WithSubGridOrder(order [][]int) Option {
	return func(g Generator) {
		switch g.(type) {
		case *bruteForceGenerator:
			g.(*bruteForceGenerator).subGridOrder = order
		}
	}
}

func defaultOptions() []Option {
	return []Option{
		WithBlockOrder(defaultRNG.Perm(81)),
		WithRNG(defaultRNG),
		WithSubGridOrder([][]int{
			{0, 0}, {0, 3}, {0, 6},
			{3, 0}, {3, 3}, {3, 6},
			{6, 0}, {6, 3}, {6, 6},
		}),
	}
}
