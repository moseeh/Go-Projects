# Statistical Functions in Go

This repository contains a program that contains functions for calculating average, variance, median and standard deviation of a set of numbers passed inside a text file.


## Functions

1. `CalcAverage(arr []float64) float64 ` : calculates the average of the given set of numbers
2. `CalcMedian(arr []float64) float64 ` : calculates the median of the given set of numbers
3. `CalcVariance(arr []float64) float64 ` : calculates the variance of the given set of numbers
4. `CalcStandardDev(n float64) float64 ` : calculates the standard deviation from the calculated variance

## Usage 

To use this program pass a .txt file as an argument as show below

``` bash

go run  . "example.txt"

```

## Requirements 
- The program requires the user to input at least one command line argument that needs only to be a .txt file.
- Failure to informs the user of the unavailability and the user can relaunch the proram again
- The program ignores other arguments passed and only works on the first argument 

## Installation 
- clone the repository and cd into it
``` bash 
git clone https://learn.zone01kisumu.ke/git/moonyango/math-skills
cd math-skills
go run . "input your file here"

```