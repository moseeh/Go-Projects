package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Align(align, asciiArt, input string, contentLines []string) {
	width, err := getTerminalWidth()
	if err != nil {
		fmt.Println("Error getting terminal width:", err)
		return
	}
	if align == "justify" {
		justify(width, input, contentLines)
		return
	}
	asciiArtArr := strings.Split(asciiArt, "\n")
	for _, v := range asciiArtArr {
		spaces := width - len(v)
		if spaces < 0 {
			fmt.Println("Text longer than terminal size\n\nAdjust and try again or use shorter string")
			os.Exit(0)
		}
		switch align {
		case "left":
			fmt.Println(v)
		case "right":
			fmt.Println(strings.Repeat(" ", spaces) + v)
		case "center":
			leftSpaces := spaces / 2
			fmt.Println(strings.Repeat(" ", leftSpaces) + v)
		default:
			fmt.Println("Error: Unrecognized alignment option. Please use 'left', 'right', 'center', or 'justify'.")
			return
		}
	}
}

func getTerminalWidth() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	sArr := strings.Fields(string(output))
	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		return 0, err
	}
	return width, nil
}

func justify(width int, input string, contentLines []string) {
	input = ReplaceSpecial(input)
	sentArr := strings.Split(input, "\\n")
	for _, line := range sentArr {
		if line == "" {
			fmt.Println()
		} else {
			lineArr := strings.Fields(line)
			linelen := 0
			for _, v := range lineArr {
				n := 0
				for _, char := range v {
					n += len(contentLines[int(char-32)*9+4])
				}
				linelen += n
			}
			spaces := width - linelen
			if spaces < 0 {
				fmt.Println("Text longer than terminal size\n\nAdjust and try again or use shorter string")
				os.Exit(0)
			}
			gapCount := 0
			if len(lineArr) > 1 {
				gapCount = len(lineArr) - 1
			} else {
				gapCount = 1
			}
			spacePerGap := (spaces / gapCount)
			linesOfSlice := make([]string, 8)
			for j, word := range lineArr {
				for _, v := range word {
					for i := 1; i < 9; i++ {
						linesOfSlice[i-1] += contentLines[int(v-32)*9+i]
					}
				}
				if j != len(lineArr)-1 {
					for k, lineSlice := range linesOfSlice {
						linesOfSlice[k] = lineSlice + strings.Repeat(" ", spacePerGap)
					}
				}
			}
			fmt.Println(strings.Join(linesOfSlice, "\n"))
		}
	}

}
