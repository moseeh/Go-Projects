# MY-LS-1

`my-ls` is a Go-based implementation of the `ls` command, which is widely used to display files and directories in a specified directory. Similar to the original `ls` command, when no directory is specified, it lists the files and directories of the present directory

## How it Works

### Basic Functionality:
By default, `my-ls` displays the content of the current directory similar to the standard `ls` command. If the directory is specified it lists the contents of that directory instead.

### Supported Flags:
- `-l`: Long format listing, showing detailed file information such as:
  - Permissions
  - Owner
  - Group
  - Size
  - Modification time
- `-R`: Recursively lists all subdirectories and their contents:
  - Shows each directory level.
  - Handles symlinks correctly.
- `-a`: Includes hidden files (files starting with `.`).
- `-r`: Reverses the order of the file listing.
- `-t`: Sorts files by modification time, displaying the most recently modified files first.


### Flag Combinations:

Similar to the original `ls`, multiple flags can be combined (e.g., `-ltr`), and the order of the flags doesn't matter. `my-ls` processes all specified flags and adjusts the output accordingly.

## Installation

1. Clone this repo to get started.

``` bash 

git clone https://learn.zone01kisumu.ke/git/nichotieno/my-ls-1

```
2. cd into the repos directory:

``` bash 

cd my-ls-1

```

## Usage 

``` bash

# List files in the current directory
go run ./cmd

# List files in a specific directory
go run ./cmd /path/to/directory

# List files in long format, including hidden files, sorted by time, in reverse order
go run ./cmd -altr


```

You can also combine directories paths and flags. since there are no limit to any number of arguments that can be passed.

## Additional Features

- File Coloring: Files and directories are color-coded to differentiate between various file types (directories, executable files, symlinks, etc.).
- Short Format Adaptation: The program automatically adjusts the display format based on terminal width to prevent truncated or misaligned outputs.

my-ls is a useful tool for understanding file system operations in Go, providing a solid foundation for building custom command-line utilities.
