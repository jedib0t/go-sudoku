# go-sudoku
<!----------------------------------------------------------------------------------------------------------------------
-- Please note that this file is auto-generated using the 'scripts/generate_readme.sh' script. Changes to this file will
-- be over-written in the next build. If you need to change this file, make the changes to 'scripts/README.md.template'
-- instead.
---------------------------------------------------------------------------------------------------------------------->

[![Build Status](https://github.com/jedib0t/go-sudoku/workflows/CI/badge.svg?branch=main)](https://github.com/jedib0t/go-sudoku/actions?query=workflow%3ACI+event%3Apush+branch%3Amain)

An implementation of Sudoku generators and solvers in GoLang.

## Usage

### Help
```
$ ./go-sudoku --help
go-sudoku: A GoLang based Sudoku generator and solver.

Usage: go-sudoku [flags] <action>

Actions:
--------
  * generate: Generate a Sudoku Grid and apply the specified difficulty on it
  * solve: Solve a Sudoku puzzle provided in a text (CSV) file

Examples:
---------
  * ./go-sudoku
  * ./go-sudoku -algorithm back-tracking -theme green -seed 42 generate
  * ./go-sudoku -format csv generate
  * ./go-sudoku -input puzzle.csv solve
  * ./go-sudoku -difficulty hard -format csv generate | ./go-sudoku solve

Optional Flags:
---------------
  -algorithm string
    	Algorithm (back-tracking/brute-force) (default "back-tracking")
  -debug
    	Enable Debug Logging?
  -difficulty string
    	Difficulty (none/easy/medium/hard/insane) (default "medium")
  -format string
    	Rendering Format (csv/table) (default "table")
  -help
    	Display this usage and help text
  -input string
    	File containing a Sudoku Puzzle in CSV format
  -no-color
    	Disable colors in rendering? [$NO_COLOR]
  -pattern string
    	Pattern to use instead of Difficulty (diamond/octagon/square/star/target/triangle)
  -progress
    	Show progress in real-time with an artificial delay?
  -seed int
    	RNG Seed (0 => random number based on time) [$SEED]
  -theme string
    	Table formatting theme (none/blue/cyan/green/magenta/red/yellow) (default "none")
  -type string
    	Sudoku Type (default/jigsaw/samurai) (default "default")
```

### Generator
*Note: all the commands below were run with SEED=42 to keep the results reproducible.*
```
$ ./go-sudoku generate
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃     8  5 │    6    │ 2     1  ┃
┃        7 │ 9  2    │          ┃
┃  6  3    │    5  1 │       4  ┃
┃ ─────────┼─────────┼───────── ┃
┃  2       │       8 │ 6  7     ┃
┃          │       2 │       9  ┃
┃     5    │    3  9 │          ┃
┃ ─────────┼─────────┼───────── ┃
┃  1  2  8 │       6 │ 9  5     ┃
┃  3  7  4 │ 2  9    │          ┃
┃  5       │ 1     7 │          ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃          Medium [42]          ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛

## custom difficulty
$ ./go-sudoku -difficulty insane generate
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃     8    │    6    │ 2        ┃
┃          │         │          ┃
┃          │         │          ┃
┃ ─────────┼─────────┼───────── ┃
┃  2       │       8 │    7     ┃
┃          │       2 │          ┃
┃          │       9 │          ┃
┃ ─────────┼─────────┼───────── ┃
┃        8 │         │    5     ┃
┃  3     4 │         │          ┃
┃          │ 1       │          ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃          Insane [42]          ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛

## custom pattern
$ ./go-sudoku -pattern diamond generate
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃          │    6    │          ┃
┃          │ 9  2  3 │          ┃
┃        2 │ 8     1 │ 7        ┃
┃ ─────────┼─────────┼───────── ┃
┃     4  9 │         │ 6  7     ┃
┃  8  6    │         │    1  9  ┃
┃     5  1 │         │ 4  2     ┃
┃ ─────────┼─────────┼───────── ┃
┃        8 │ 3     6 │ 9        ┃
┃          │ 2  9  5 │          ┃
┃          │    8    │          ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃          Diamond [42]         ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛

## samurai sudoku
$ ./go-sudoku -type samurai generate
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃     1    │    2  8 │ 5     6 │         │ 4  8    │    2  1 │          ┃
┃  5       │         │         │         │       9 │ 8  4  7 │    6  3  ┃
┃  6       │ 5       │    2    │         │    1    │    9    │       8  ┃
┃ ─────────┼─────────┼─────────┤         ├─────────┼─────────┼───────── ┃
┃     4    │    8  5 │ 3  6  9 │         │       8 │         │ 6        ┃
┃          │ 3       │ 2     4 │         │    4  6 │ 2       │    9     ┃
┃  2  5  3 │ 9       │    7  8 │         │ 9       │         │ 8        ┃
┃ ─────────┼─────────┼─────────┼─────────┼─────────┼─────────┼───────── ┃
┃     3  7 │ 4  6    │ 9     5 │ 7  6    │    3  1 │ 7  6    │ 4  5     ┃
┃     6  2 │         │    1    │ 9  2    │         │    3    │    1  2  ┃
┃  4     5 │         │       2 │         │         │ 5     2 │    8  6  ┃
┃ ─────────┴─────────┼─────────┼─────────┼─────────┼─────────┴───────── ┃
┃                    │                   │    7  3 │                    ┃
┃                    │         │ 4     2 │ 5       │                    ┃
┃                    │ 7       │    3  9 │    2  8 │                    ┃
┃ ─────────┬─────────┼─────────┼─────────┼─────────┼─────────┬───────── ┃
┃     7  6 │ 5  9    │       8 │ 3  4    │ 9  5    │ 1     4 │    3     ┃
┃  2     5 │       6 │    7    │ 2  9    │ 1       │    3  2 │       5  ┃
┃     4  1 │         │ 5  9    │ 1  8  7 │ 3     2 │ 5  7  6 │    8  1  ┃
┃ ─────────┼─────────┼─────────┼─────────┼─────────┼─────────┼───────── ┃
┃     1    │ 3  4    │       5 │         │ 7     9 │ 4  6    │    2  3  ┃
┃          │       7 │ 4  3  2 │         │    2    │    1    │          ┃
┃     3  2 │         │       7 │         │ 6  3  4 │ 2       │    1     ┃
┃ ─────────┼─────────┼─────────┤         ├─────────┼─────────┼───────── ┃
┃  6     7 │ 4  3  9 │         │         │         │ 7       │ 1        ┃
┃        4 │       1 │ 7  8    │         │ 4       │ 8     3 │ 6        ┃
┃     2    │    7    │    4    │         │ 2       │       1 │ 3        ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃                          Samurai Medium [42]                          ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### Solver
*Note: all the commands below were run with SEED=42 to keep the results reproducible.*
```
## generate a new puzzle and feed it to the solver to solve
$ ./go-sudoku -format csv generate | ./go-sudoku solve
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃             INPUT             ┃             OUTPUT            ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃     8  5 │    6    │ 2     1  ┃  9  8  5 │ 7  6  4 │ 2  3  1  ┃
┃        7 │ 9  2    │          ┃  4  1  7 │ 9  2  3 │ 8  6  5  ┃
┃  6  3    │    5  1 │       4  ┃  6  3  2 │ 8  5  1 │ 7  9  4  ┃
┃ ─────────┼─────────┼───────── ┃ ─────────┼─────────┼───────── ┃
┃  2       │       8 │ 6  7     ┃  2  4  9 │ 5  1  8 │ 6  7  3  ┃
┃          │       2 │       9  ┃  8  6  3 │ 4  7  2 │ 5  1  9  ┃
┃     5    │    3  9 │          ┃  7  5  1 │ 6  3  9 │ 4  2  8  ┃
┃ ─────────┼─────────┼───────── ┃ ─────────┼─────────┼───────── ┃
┃  1  2  8 │       6 │ 9  5     ┃  1  2  8 │ 3  4  6 │ 9  5  7  ┃
┃  3  7  4 │ 2  9    │          ┃  3  7  4 │ 2  9  5 │ 1  8  6  ┃
┃  5       │ 1     7 │          ┃  5  9  6 │ 1  8  7 │ 3  4  2  ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃       45 blocks to solve      ┃          3621 cycles          ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```
