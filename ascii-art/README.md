# ASCII Art Generator

- This program, written in Go, converts a string into a graphic representation using ASCII characters. It handles numbers, letters, spaces, special characters, and newline characters and tab.
- The program cannot handle unprintable characters or non-enlish characters


## Usage

- Run the program from the command line and pass the string to convert as an argument:

```bash
$ go run . "Your string goes here"

``` 
- The program will then output the graphic representation of the string using ASCII characters.
- The program can also multiple strings passed as arguments from the command line. Joind the strings into one string and outputs the graphical representation as show:

``` bash 
$ go run . "string1" "string2" "string3" "string4"

```

## Dimensions 
- The output coverss 8 lines for every character-art with an allowance of a single line for spacing at the top of the art


## Requirements 
- The program requires the user to input at least one command line arguments and as many as the require for it to output the art. 
- Failure to informs the user of the unavailability and the user can relaunch the proram again


## Installation 
- clone the repository and cd into it
``` bash 
git clone https://github.com/moseeh/Go-Projects
cd ascii art 
go run . "input your string or strings"


```

## Banner files
Banner files with specific graphical template representations using ASCII will be provided
- standard.txt
- shadow.txt
- thinkertoy.txt