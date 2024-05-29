package utils

import "errors"

func FindColor(s string) ([3]int, error) {

	rgbColors := map[string][3]int{
		"red":         {255, 0, 0},
		"green":       {0, 255, 0},
		"blue":        {0, 0, 255},
		"yellow":      {255, 255, 0},
		"cyan":        {0, 255, 255},
		"purple":      {128, 0, 128}, // Magenta
		"white":       {255, 255, 255},
		"black":       {0, 0, 0},
		"orange":      {255, 165, 0},
		"pink":        {255, 192, 203},
		"brown":       {165, 42, 42},
		"gray":        {128, 128, 128},
		"navy":        {0, 0, 128},
		"lime":        {0, 255, 0},
		"maroon":      {128, 0, 0},
		"olive":       {128, 128, 0},
		"silver":      {192, 192, 192},
		"teal":        {0, 128, 128},
		"magenta":     {255, 0, 255},
		"lightgray":   {211, 211, 211},
		"darkgray":    {169, 169, 169},
		"lightblue":   {173, 216, 230},
		"lightgreen":  {144, 238, 144},
		"lightcyan":   {224, 255, 255},
		"lightyellow": {255, 255, 224},
		"lightpurple": {216, 191, 216},
		"thistle":     {216, 191, 216}, // Thistle
		"darkred":     {139, 0, 0},
		"darkgreen":   {0, 100, 0},
		"darkblue":    {0, 0, 139},
		"darkcyan":    {0, 139, 139},
		"darkmagenta": {139, 0, 139},
		"darkyellow":  {204, 204, 0}, // Olive
		"beige":       {245, 245, 220},
		"gold":        {255, 215, 0},
		"bronze":      {205, 127, 50},
		"violet":      {238, 130, 238},
		"indigo":      {75, 0, 130},
		"chartreuse":  {127, 255, 0},
		"aquamarine":  {127, 255, 212},
		"coral":       {255, 127, 80},
		"salmon":      {250, 128, 114},
		"turquoise":   {64, 224, 208},
		"khaki":       {240, 230, 140},
		"orchid":      {218, 112, 214},
		"plum":        {221, 160, 221},
		"tan":         {210, 180, 140},
		"peach":       {255, 218, 185},
		"mint":        {189, 252, 201},
		"lavender":    {230, 230, 250},
		"rose":        {255, 0, 127},
	}

	val, found := rgbColors[s]
	if !found {
		return [3]int{}, errors.New("color not available")
	}
	return val, nil
}
