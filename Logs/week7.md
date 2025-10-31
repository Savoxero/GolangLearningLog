# Golang learning log - Week 7
## Skills to learn:
- Git and Github fluency and learning the commands and navigation
- Planning and realizing project ideas
- Focusing on one thing at a time and doing it good enough (i have a tendency to jump around topics and not stick to one thing)
- writing experimental code to ingrain new concepts, and test the limits of both my knowledge and the language itself.
## Current plans:
### Calculator project:
- [] word counter using map[string] int
- [] a system to automatically exit if the input is repeated and the program does not execute as designed.
- removing nested loops and finding a way to replace them with functions or another clever way
### general learning:
- [] reading a shitton of documentation on git and programming generally
- noting each Cause -> EFFECT -> Fix 
## Habits ( gpt notes)
- Design before keysmashing. Write the program flow as pseudocode before touching Go. Your problem isn’t logic—it’s pre-planning.

## Progress:
- learned about structs and how to use them (essentially a template for data types)
- bufio package for reading input:
 - bufionewscanner(os.stdin) essentially creates a scanner object that reads from standard input (the terminal) for user input.
 - scanner.scan analyzes and turns the text into tokens 
 - scanner.text() gets the tokens from the input and returns it as a string.
 - pointers and how to use them in some way, learned about the & and * operators but still not 100% sure how to use them effectively. it's a bit confusing and it's easy to get lost in the syntax.
 - more of strconv shenangians, converting strings to integers and vice versa. still i don't understand why the error handling is necessary for every conversion, seems a bit overkill for simple conversions. (what do i know tho)
 ## Notes:
 - im thinking about creating a markdown file to log all the progress i'm making on my current projects, say added this or that feature, learned this or that concept etc. but i already have this log. so i'm gonna just put big stuff here and small stuff in the project files themselves.


