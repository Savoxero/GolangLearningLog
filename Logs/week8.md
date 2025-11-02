# Golang learning log - Week 8
## Skills to learn:
- Git and Github fluency and learning the commands and navigation
- Planning and realizing project ideas
- Focusing on one thing at a time and doing it good enough (i have a tendency to jump around topics and not stick to one thing)
- cause -> effect -> fix logging for problems encountered
## Current plans: 
- Todolist CLI
- resuming boot.dev
- using git and github with more consistency
## Habits ( gpt notes)
- Design before keysmashing. Write the program flow as pseudocode before touching Go.
- pre planning
## Using my strenghts (GPT notes)
- Channel your pattern intuition into writing design notes before code. When you see repetition, stop, sketch the logic, then rewrite.
- Exploit your self-awareness: build a routine that deliberately isolates one task 
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
 - I understood github a little better, i am able to isolate a part of my project and upload it to github without necessary garbage around it, essentially isolating what matters.


=======
## Cause -> Effect -> Fix log
### git problem
- Cause: deleted my README.md file by accident
- Effect: I was desperate to recover it and that made me learn git commands and and understand the directory structure better
- Fix: i had to go back to the strict directory it was deleted in, and just restore it, using git restore README.md inside the specific directory (check git status to see where you are, source control is also useful for understanding where the file was deleted at)
### calculator index condition problem
- Cause: my calculator was not working as intended, it was skipping some inputs and not counting them properly
- Effect: i tried searching for some fitting condition for the for loop, learning the limits of the conditions you can possibly make in go
- fix: for index = 0; index < len(history); index++ it uses the element count of the slice to determine how many times to loop, instead of trying to find a condition that matches the input value( i just needed something to track how many times the calculation has run, and i never knew u could use len() in the for loop condition)
### input conversion problem
- Cause:tried to implement a code that strips down input from the user into seperate strings and find the command and the arguments. so EG: "add find my wallet on the street". i had trouble with seperating the add.
- Effect: i learned about strings.fields() which splits a string into a slice of strings based on spaces and then used strings.join() to join the arguments back together. EG: filter 
```go
:= strings.Fields(input)
		cmtokens := filter[0]
		argtokens := strings.Join(filter[1:], " ")
		input = cmtokens
        ``  // what this does is split the input into 2 seperate variables, the command and the arguments. it made me understand slices better as well. (slices need something to point towards to, so you can't just declare a slice and use it, it needs to be assigned to something first)
- Fix: used strings.fields() and strings.join() to split the input into command and arguments.