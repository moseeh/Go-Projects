# ascii-art-color

- Just like the `ascii-art` program, this program displays an ascii output of an input string but this time its colored. 
- The program colors the entire string or specific substring within the ascii-art format. 
- This tool allows you to highlight particular parts of a string in a specified color, making it useful for text decoration and visualization in the terminal.

## Features

- Colorize entire strings in ASCII-ART repredentation of words
- Colorize specific substrings in ASCII art representations of words.
- Supports customizable ASCII art for characters.
- Simple and efficient implementation.

## Usage

### Running the Program

To run the program, use the following command:

```sh
Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <letters to be colored> "something"

```

- The program can also work with one argument as shown below. In this case it provides an uncolored output

```sh
go run . "hello there"

```

- it can also color an entire string by passing one argument along with the flag `-color=` as shown below

```sh
go run . --color=<color> <enter your string here>
```
