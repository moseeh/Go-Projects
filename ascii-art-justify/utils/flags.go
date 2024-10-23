package utils

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Define the flags
var (
	ColorPtr  = flag.String("color", "", "coloring the output")
	AlignPtr  = flag.String("align", "", "set the alignment")
	OutputPtr = flag.String("output", "", "set the output file")
)

// ValidateFlagFormat checks if the flag is in the format --flagname=value
func ValidateFlagFormat() {
	valid := true

	// Check if flags are in the correct format
	flag.VisitAll(func(f *flag.Flag) {
		found := false
		for _, arg := range os.Args[1:] {
			if strings.HasPrefix(arg, "--"+f.Name+"=") {
				found = true
				break
			}
		}
		if !found && flagIsSet(f.Name) {
			valid = false
		}
	})

	if !valid {
		fmt.Printf("Usage: go run . [OPTION] [STRING] [BANNER]\n\n")
		fmt.Println("Example: go run . --align=right something standard")
		os.Exit(0)
	}
}

// flagIsSet checks if a specific flag was set
func flagIsSet(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// ParseFlags parses the flags and returns the non-flag arguments
func ParseFlags() []string {
	// Check for undefined flags
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			flagName := strings.TrimLeft(arg, "-")
			flagName = strings.SplitN(flagName, "=", 2)[0]
			if flag.Lookup(flagName) == nil {
				fmt.Printf("Flag not found: -%s\n", flagName)
				flag.Usage()
				os.Exit(1)
			}
		}
	}
	flag.Parse()
	ValidateFlagFormat()
	// return elements that aren't flags
	return flag.Args()
}
