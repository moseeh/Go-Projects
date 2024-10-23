# Ascii-art-justify

* This project is a command-line application that aligns ASCII art output according to specified options. 
* The alignment can be set to `center`, `left`, `right`, or `justify`, and the output will automatically adapt to the terminal size. 
* This ensures that the ASCII art remains properly aligned even when the terminal window is resized.


## Objectives

* Align ASCII art using different alignment options: `center`, `left`, `right`, `justify`.
* Adapt the ASCII art output to fit the terminal size dynamically.
* Ensure the program provides appropriate usage instructions when incorrect flags are used.

## Installation

To get started, clone the repository:
```bash
git clone https://learn.zone01kisumu.ke/git/moonyango/ascii-art-justify.git
```

```bash
cd ascii-art-justify
```

## Usage

To run the application, use the following command:

```bash
go run . --align=<type> <STRING> <BANNER>
```

### Alignment Options

* **center:** Aligns the ASCII art in the center of the terminal.
* **left:** Aligns the ASCII art to the left side of the terminal.
* **right:** Aligns the ASCII art to the right side of the terminal.
* **justify:** Stretches the ASCII art to fit the full width of the terminal, justifying the text.

### Examples

* Center Alignment:

```bash
go run . --align=center "hello" standard
```

* Left Alignment:

```bash
go run . --align=left "Hello There" standard
```

* Right Alignment:

```bash
go run . --align=right "hello" shadow
```

* Justify Alignment:

```bash
go run . --align=justify "how are you" shadow
```

### Incorrect Usage

* If an incorrect alignment option or format is used, the program will display the following message:

```bash
Usage: go run . [OPTION] [STRING] [BANNER]
Example: go run . --align=right something standard
```


### Contributions are welcome! Please follow these steps:

* Fork the repository.
* Create a new branch (git checkout -b feature-branch).
* Commit your changes (git commit -am 'Add new feature').
* Push to the branch (git push origin feature-branch).
* Create a new Pull Request.

### License

* This project is licensed under the MIT License. See the LICENSE file for details.