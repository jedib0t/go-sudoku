# go-sudoku
<!----------------------------------------------------------------------------------------------------------------------
-- Please note that this file is auto-generated using the 'scripts/generate_readme.sh' script. Changes to this file will
-- be over-written in the next build. If you need to change this file, make the changes to 'scripts/README.md.template'
-- instead.
---------------------------------------------------------------------------------------------------------------------->

[![Build Status](https://github.com/jedib0t/go-sudoku/workflows/CI/badge.svg?branch=main)](https://github.com/jedib0t/go-sudoku/actions?query=workflow%3ACI+event%3Apush+branch%3Amain)

An implementation of Sudoku game, generators and solvers for the command line in
GoLang.

## Usage

### Game
```
$ ./go-sudoku -help
It works!
```

### Generator
*Note: all the commands below were run with the environment variable `SEED=42`
to keep the results reproducible.*
```
$ ./go-sudoku-generator generate
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
$ ./go-sudoku-generator -difficulty insane generate
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
$ ./go-sudoku-generator -pattern diamond generate
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
$ ./go-sudoku-generator -type samurai generate
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
*Note: all the commands below were run with the environment variable `SEED=42`
to keep the results reproducible.*
```
## generate a new puzzle and feed it to the solver to solve
$ ./go-sudoku -format csv generate | ./go-sudoku solve
It works!
```
