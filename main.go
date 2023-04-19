/*
Description:
This Go code is a minimalistic implementation of a regular expression engine that allows for pattern matching in text strings. The program takes in at least two arguments via the command line: a regular expression pattern and a text string to match against. If there are additional text strings, the program matches against each of them as well. The engine supports several regex metacharacters, including ^, $, *, and ..

Arguments:
regexp (string): The regular expression pattern to match against.
text (string): The text string to match against the regular expression pattern.

Returns:
string (matched): The function returns strings that match with the provided regex.

Comments:
The main() function checks that at least two arguments are provided to the program, and if not, it prints out a usage message and exits with a status code of 1.
The match() function takes in two string arguments and returns a boolean value indicating whether there is a match between the regular expression pattern and the text string.
The matchhere() function is a helper function that is called by match(). It also takes in two string arguments and returns a boolean value indicating whether there is a match between the regular expression pattern and the text string.
The matchstar() function is another helper function that is called by matchhere(). It takes in three string arguments and returns a boolean value indicating whether there is a match between the regular expression pattern and the text string.

*/

package main

import (
	"fmt"
	"os"
)

// The main function parses command line arguments and prints out any text that matches the given regular expression.
func main() {
	// If the number of arguments is less than three, print usage instructions and exit with status code 1.
	if len(os.Args) < 3 {
		fmt.Println("Usage: regex <regexp> <text1> [<text2> ...]")
		os.Exit(1)
	}

	// Assign the regular expression from the first argument to the `regexp` variable.
	regexp := os.Args[1]

	// Iterate over the remaining command line arguments and print any that match the regular expression.
	for _, text := range os.Args[2:] {
		if match(regexp, text) {
			fmt.Println(text)
		}
	}
}

// The match function determines whether the given regular expression matches the given text.
func match(regexp, text string) bool {
	// If the regular expression is empty, return true.
	if regexp == "" {
		return true
	}

	// If the regular expression starts with "^", match the regular expression at the beginning of the text.
	if regexp[0] == '^' {
		return matchhere(regexp[1:], text)
	}

	// Iterate over each character in the text and check if the regular expression matches starting from that character.
	for len(text) > 0 {
		if matchhere(regexp, text) {
			return true
		}
		text = text[1:]
	}

	// If the regular expression doesn't match any part of the text, return false.
	return false
}

// The matchhere function determines whether the given regular expression matches the given text starting at the beginning.
func matchhere(regexp, text string) bool {
	// If the regular expression is empty, return true.
	if regexp == "" {
		return true
	}

	// If the regular expression is "$" and has length 1, check if the text is empty.
	if regexp[0] == '$' && len(regexp) == 1 {
		return len(text) == 0
	}

	// If the regular expression has length greater than 1 and the second character is "*", match the rest of the regular expression repeatedly.
	if len(regexp) > 1 && regexp[1] == '*' {
		return matchstar(regexp[0], regexp[2:], text)
	}

	// If the text is non-empty and the first character of the regular expression is either "." or the same as the first character of the text, match the rest of the regular expression and the rest of the text.
	if len(text) > 0 && (regexp[0] == '.' || regexp[0] == text[0]) {
		return matchhere(regexp[1:], text[1:])
	}

	// If none of the above cases match, return false.
	return false
}

// The matchstar function matches a character followed by zero or more repetitions of the rest of the regular expression.
func matchstar(c byte, regexp, text string) bool {
	// Repeat until the text is empty or the first character of the text doesn't match `c`.
	for len(text) > 0 && (text[0] == c || c == '.') {
		// If the rest of the regular expression matches the remaining text, return true.
		if matchhere(regexp, text) {
			return true
		}
		text = text[1:]
	}
  
  return match(regexp, text)
  
}


