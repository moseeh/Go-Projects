# ASCII Art Web Export File

ASCII Art Web Export File is a Go project that allows users to generate ASCII art from input text using various predefined banners through a stylized web interface, while being able to download a file with the art.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Algorithm](#usage)
- [Contributing](#contributing)

## Introduction

ASCII Art Web Export File provides a simple web interface to convert input text into ASCII art using different banner styles such as 'standard', 'shadow', and 'thinkertoy'. The project utilizes Go's HTTP package to handle requests and templates for rendering HTML responses. The web page is maintained and styled using CSS. It also utilizes HTTP Header to allow user to export a file. The entire application is containerized using Docker for easy deployment and consistency across different environments.

## Features

- **Generate ASCII Art**: Converts input text into ASCII art using selected banner styles.
- **Web Interface**: Accessible through a web browser, with forms for input and banner selection.
- **Error Handling**: Provides meaningful error messages for invalid input or file reading failures.
- **Dockerize**: Packaged as a Docker container for easy deployment and consistent runtime environment.
- **Export file**:  Provides a download button for the user to be able to download art in .txt format

## Project Structure
**The project structure is organized as follows:**

* `main.go`: Entry point of the application, sets up HTTP server and handles routing.
* `utils/`: Contains utility functions for file reading, ASCII art generation, and banner selection.
* `templates/`: HTML templates for rendering the web interface and displaying results.
* `static/`: CSS files for styling
* `banners/`: ASCII art files for different banner styles ('standard.txt', 'shadow.txt', 'thinkertoy.txt').
* `Dockerfile`: Instructions for building the Docker image.
* `go.mod`: Go module file for dependency management.


## Installation

1. **Clone the repository:**

   ```bash

   git clone https://learn.zone01kisumu.ke/git/moonyango/ascii-art-web-export-file
   cd ascii-art-web-export-file

   ```

## Usage 

1. **Run the program:**

   ```bash 
   go run .

   ```

2. **Acsess the web interface**

Open a web browser and go to http://localhost:8080

3. **Generate AsciiArt**
   * Enter your text in the input box
   * Select a banner style
   * Click on the 'Generate' button to see the ascii based on your input
   * Click on Download ART button to download a file with art based on your input

## ALGORITHM ASCII Art Web Generator

### FUNCTION ServeIndex(request, response):
    IF request method is not GET:
        Return 400 Bad Request
    IF request path is "/":
        Serve "templates/index.html"
    ELSE:
        Return 404 Not Found

### FUNCTION GenerateASCIIArt(request, response):
    IF request method is not POST:
        Return 405 Method Not Allowed
    
    input = Get "input" from form data
    banner = Get "banner" from form data
    
    IF input is empty:
        Return 400 Bad Request "Input is required"
    
    content = Read ASCII map file based on banner
    IF error reading file:
        Log error
        Return 500 Internal Server Error
    
    contentLines = Split content into lines
    
    art = Generate ASCII art from input and contentLines
    IF error generating art:
        Return 400 Bad Request
    
    Render "templates/result.html" with generated art
    IF error rendering template:
        Log error
        Return 500 Internal Server Error

### HELPER FUNCTIONS:
    ReadsFile(filename): Read and return file contents
    GetFile(banner): Return filename based on banner type
    SplitFile(content): Split content into lines
    DisplayText(input, contentLines): Generate ASCII art

### ERROR HANDLING:
    - Check for valid HTTP methods
    - Validate input presence
    - Handle file reading errors
    - Handle ASCII art generation errors
    - Handle template rendering errors

END ALGORITHM


## Contributing
**Contributions are welcome! If you'd like to contribute to this project, please follow these steps:**

1. Fork the repository and clone it locally.
2. Create a new branch for your feature or bug fix.
3. Make your changes and test them thoroughly.
4. Submit a pull request detailing your changes and improvements.

## Author
**This Project was Authored and maintaned by:**

 ![image](/images/moses.png)

 [Moses Onyango](https://learn.zone01kisumu.ke/git/moonyango) 


 ![ben](/images/bernad.png)

 [Bernad Okumu](https://learn.zone01kisumu.ke/git/bernaotieno)


