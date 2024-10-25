# Linear Stats

A Go program that performs linear regression analysis on numerical data from text files. The program calculates and outputs the linear regression line equation and Pearson correlation coefficient for the given dataset.

## Description

Linear Stats reads numerical data from a text file (one number per line), performs statistical calculations, and outputs:
- The linear regression line in the format: y = mx + b
- The Pearson correlation coefficient

The program assumes the x-values are the line numbers (0, 1, 2, 3, ...) and the y-values are the numbers provided in the file.

## Installation

```bash
git clone https://learn.zone01kisumu.ke/git/moonyango/linear-stats
cd linear-stats
```

## Usage

Run the program by providing a text file as a command-line argument:

```bash
go run ./cmd data.txt

```

### Input File Format
- Text file (.txt extension)
- One number per line
- Valid numeric values (integers or decimal numbers)

Example `data.txt`:
```
23.45
67.89
12.34
45.67
```

### Output Format
The program will output:
```
Linear Regression Line: y = [slope]x + [intercept]
Pearson Correlation Coefficient: [coefficient]
```

## Project Structure

```
linear-stats/
├── cmd/ 
|   └── main.go
├── utils/
│   ├── convtofloat.go  # Convert slice of strings to slice of floating point values
│   ├── isTXT.go        # confirm if file is a .txt file
|   ├── parseFile.go    # readfile and decode contents
|   ├── pearson.go      # calculate the pearson correlation
|   └── regression.go   # calculate the linear regression
├── go.mod
├── data.txt            # for test
└── README.md
```

## Functions

### Main Package
- Handles command-line arguments
- Validates input file
- Displays results

### Utils Package
- `IsTxt(filename string) bool`
  - Validates if the input file has .txt extension

- `ParseFile(filename string) []float64`
  - Reads and parses numbers from the input file
  - Returns a slice of float64 values

- `CalculateLinearRegression(data []float64) (slope, intercept float64)`
  - Calculates the linear regression line parameters
  - Returns slope (m) and y-intercept (b)

- `CalculatePearson(data []float64) float64`
  - Calculates the Pearson correlation coefficient
  - Returns a value between -1 and 1

## Error Handling

The program includes error handling for:
- Missing input file
- Invalid file extension
- File reading errors
- Invalid number formats
- Empty files


## Contributing

If you would like to contribute to this project, please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License.
