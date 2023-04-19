 # RegGo
 
RegGo is a regex engine written in Go that allows you to match regular expressions against one or more input texts from the command line. It's a lightweight and easy-to-use tool that can help you search for specific patterns, validate user input, clean up data, extract data from text, and more.

# Installation 

To install RegGo, follow these simple steps:

1. Make sure you have Go installed on your system. If you don't, you can download it from the official website at https://golang.org/dl/.

2. Clone the RegGo repository from GitHub by running the following command in your terminal:

  `git clone https://github.com/hitbug-exe/RegGo.git`

3. Navigate to the RegGo directory by running the following command:

  `cd RegGo`

4. Build the executable file by running the following command:

  `go build`

This will create an executable file named "reggo" in the current directory.

5. (Optional) If you want to use RegGo from anywhere in your terminal, you can move the "reggo" executable file to a directory that's listed in your system's PATH environment variable.


# Usage

The program can be run by executing the following command in the terminal:

  `regex <regexp> <text1> [<text2> ...]`

where <regexp> is the regular expression to match and <text1>, <text2>, etc. are the input texts to match against.

For example:

  ` regex "^abc" "abc" "abcdef" "xyzabc" `

This command will match all input texts that start with "abc".


# References

https://www.cs.princeton.edu/courses/archive/spr09/cos333/beautiful.html

https://swtch.com/~rsc/regexp/regexp1.html

https://swtch.com/~rsc/regexp/dfa0.c.txt

# License

RegGo is released under the MIT License. Feel free to use, modify, and distribute it as you se