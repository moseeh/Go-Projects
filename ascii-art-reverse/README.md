# ASCII ART REVERSE

`Ascii-Art-Reverse` is a command-line program that converts ASCII art back into its corresponding string. It can work with three different banner styles: `standard`, `shadow`, and `thinkertoy`. Additionally, the program can convert a single string into ASCII art and print it to the terminal using a specified banner style.

## Table of Contents

* [Features](#features)
* [Installation](#installation)
* [Usage](#usage)
* [Banner Files](#banner-files)
* [Contributing](#contributing)
* [License](#license)
* [Contributors](#contributors)

## Features

* ASCII art reversal: coverts ASCII art back into its original string usinf predefined banner styles
* Multi-Banner Support: Works with three different banner styles: standard, shadow, and thinkertoy.
* String to ASCII Art Conversion: Convert any string to ASCII art and print it in the terminal with a specified banner style.
* Flexible Input Handling: Supports both file input (for reversing ASCII art) and direct string input (for generating ASCII art).
* Error Handling: Provides clear error messages for issues such as missing files or unsupported banner styles.

## Installation

### Prerequisites

* Go 1.18 or later

### Steps

* clone the repository into your local machine 

``` bash

git clone https://learn.zone01kisumu.ke/git/moonyango/ascii-art-reverse

```

* navigate into the project directory 

``` bash

cd ascii-art-reverse

```
## Usage 

* This program can be run in two main modes
  - Reverse Mode: Coverts an ASCII-art back to its corresponding string

``` bash

  go run . --reverse=file.txt

```
   - Art Mode: Coverts a single string into ASCII art and prints it to the terminal. One can choose to use a banner file or not. 

``` bash

  go run . [string] [banner]

```

## Banner files
Banner files with specific graphical template representations using ASCII will be provided
- standard.txt
- shadow.txt
- thinkertoy.txt

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any changes or improvements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributors

- [Moses Otieno](https://github.com/moseeh)
- [Kevin Wasonga](https://github.com/kevwasonga)
