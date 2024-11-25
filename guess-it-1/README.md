# Guess It 1

## Description
"Guess It 1" is a program that predicts a range for the next number in a sequence based on previous inputs. It uses statistical calculations to make educated guesses about where the next number will fall.

## Objectives
The program:
1. Reads numbers from standard input
2. Calculates a predicted range for the next number
3. Outputs this range to standard output

## Instructions

### Input
The program receives numbers as standard input, one per line. For example:
```
189
113
121
114
145
110
```

### Output
For each number received (except the first two), the program outputs a range prediction for the next number in the following output:
```
lower_limit upper_limit
```
For example:
```
120 200
160 230
110 140
100 200
1 99
```

### Implementation Details
- The program uses at least the last two numbers to calculate the range for the next number.
- Statistical methods such as mean, variance, and standard deviation should be used in the calculations.
- The range is balanced between being small enough to be precise and large enough to often include the correct next number.

## Setup and Running
1. Ensure you have GO installed in your machine.
2. Clone the repository and cd into the project directory
``` bash
git clone https://learn.zone01kisumu.ke/git/moonyango/guess-it-1
cd guess-it-1
```
3. Run the program
``` bash
go run .
```
4. Input your numbers to see the predicted values

## Scripting
To make the program executable from the root folder of the provided tester, create a shell script named `script.sh` in the `student` folder with the following content:

```sh
#!/bin/sh
go run ./student/main.go
```

Make sure to make the script executable:

```bash
chmod +x student/script.sh
```

## Contributing
Contributions to this project are welcome. Please follow these steps to contribute:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/AmazingFeature`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
5. Push to the branch (`git push origin feature/AmazingFeature`)
6. Open a Pull Request

Please ensure your code adheres to the existing style and all tests pass before submitting a pull request.

## License
This project is licensed under the MIT License:

MIT License

Copyright (c) 2023 moonyango

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
