# Tetris Optimizer

Tetris Optimizer is a Go program designed to read and validate tetromino shapes from a file, and arrange them into the smallest possible square grid. The program checks the validity of tetromino shapes and ensures that they can be placed on the grid without overlapping.

## Table of Contents
* [Introduction](#introduction)
* [Features](#features)
* [Installation](#installation)
* [Usage](#usage)
   * [Command Line](#command-line)
   * [Input File Format](#input-file-format)
* [Components](#components)
* [Examples](#examples)
* [Contributing](#contributing)
* [License](#license)
* [Author](#author)
* [Contributors](#contributors)

## Introduction

This project aims to solve the Tetris packing problem by reading tetromino shapes from an input file, validating them, and arranging them into the smallest possible square grid. It ensures that the shapes are valid and fit within the constraints of the grid.

## Features

* Read tetromino shapes from a text file.
* Validate the size and shape of tetrominoes.
* Trim unused lines from tetrominoes.
* Arrange tetrominoes into the smallest possible grid.
* Validate each line of the file to contain only valid characters (`#` and `.`).

## Installation

### Prerequisites
* Go 1.18 or later

### Steps
1. Clone the repository and cd into :

   ``` bash
   git clone https://github.com/moseeh/tetris-optimizer.git
   cd tetris-optimizer
   ```
2. Create a .txt file with a list of tetrominoes and run the program

### Usage

* Run the program with the input file containing tetromino shapes as show :

#### Command Line

``` bash 
go run . sample.txt 
```
where sample.txt is the .txt file containing tetrominoes

#### Input File Format

The input file should contain tetromino shapes represented by `#` and `.` characters. Each tetromino should be a 4x4 grid, with shapes separated by a blank line. For example:

```
#...
###.
....
....

....
..#.
..#.
..##
```

## Components

## Main Program
**File:** `main.go`

This is the entry point of the application. It reads the input file, validates tetromino shapes, and arranges them on the grid.

## Utilities
**Directory:** `utils`

Contains various utility functions for reading files, validating shapes, trimming tetrominoes, and placing them on the grid.

Functions:
* **ReadParseTetromino:** Reads the tetromino shapes from a file and validates each line.
* **IsTetrominoValidShape:** Checks if the tetromino shapes are valid based on connectivity and block count.
* **ValidateTetrominoSize:** Ensures each tetromino is a 4x4 grid with exactly four `#` characters.
* **TrimTetromino:** Removes unused lines from each tetromino.
* **PlaceTetromino:** Arranges the tetrominoes on the smallest possible grid.


## Examples

### Example Input File
`sample.txt`:

```
##..
##..
....
....

.#..
.#..
##..
....
```

`output:`

```
AAB
AAB
.BB
```

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any changes or improvements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.



## Author

- [Moses Otieno](https://github.com/moseeh)

## Contributors

- [Moses Otieno](https://github.com/moseeh)