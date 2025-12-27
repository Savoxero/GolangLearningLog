# Golang learning log - Week 9
## Skills to learn:
- Git and Github fluency and learning the commands and navigation
- Planning and realizing project ideas
- Focusing on one thing at a time and doing it good enough (i have a tendency to jump around topics and not stick to one thing)
- cause -> effect -> fix logging for problems encountered
## Current plans: 
- Todolist CLI, progress has been annoyingly slow but steady
- resuming boot.dev. after im done with this todolist cli, im thinking about starting boot.dev again
- using git and github with more consistency
## Habits ( gpt notes)
- Design before keysmashing. Write the program flow as pseudocode before touching Go.
- pre planning
## Using my strenghts (GPT notes)
- Channel your pattern intuition into writing design notes before code. When you see repetition, stop, sketch the logic, then rewrite.
- Exploit your self-awareness: build a routine that deliberately isolates one task 
## Progress:
- parsing user input into commands and arguments using strings.fields() and strings.join()
- basic todo list functionality: add, delete tasks: 
 - add: appends a new task to the todo list slice
 - delete: removes a task from the todo list slice based on its index using slicing and range( removes the task based on user input index, adjusted for 0-based indexing)(so 1 matches index 0 in the slice)
 - added some control flow to handle invalid inputs from the user, empty string inputs, removing non existing tasks
 - been understanding slices better, how they work, how to manipulate them along with for loops and range, it's a topic that took me a long time to actually understand. how outer loops wait for inner loops to finish, controlling flow with break and continue statements etc.. not gonna yap u know the gist of it
 ## Notes: 
 - I'll be honest, the amount of stuff i had to google and look up to get this todo list working is a bit overwhelming. but i guess that's how learning works, you hit a wall, you find a way around it. i feel like a absolute fraud that doesn't know anything. i try to understand some concepts, i read the theory, documentations. it all makes sense. but then i come around to my code and it's like all of that vanishes.
 - the amount of stuff to do is immense, i still haven't mastered slices, structs, pointers, and for loops in some instances. breaking and continuing makes sense to me but designing the flow is all about trial and error for me rn.