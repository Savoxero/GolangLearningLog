# Golang learning log - Week 7
## Skills to learn:
- Git and Github fluency and learning the commands and navigation
- Planning and realizing project ideas
- Focusing on one thing at a time and doing it good enough (i have a tendency to jump around topics and not stick to one thing)
## Current plans:
### Calculator project:
- [] word counter using map[string] int (scrapped idea, don't think it's useful enough)
- [] a system to automatically exit if the input is repeated and the program does not execute as designed.
- removing nested loops and finding a way to replace them with functions or another clever way
### general learning:
- [] reading a shitton of documentation on git and programming generally
- noting each Cause -> EFFECT -> Fix 
## Habits ( gpt notes)
- Design before keysmashing. Write the program flow as pseudocode before touching Go. Your problem isn’t logic—it’s pre-planning.
## Using my strenghts (GPT notes)
- Channel your pattern intuition into writing design notes before code. When you see repetition, stop, sketch the logic, then rewrite.
---
- Turn your persistence into data: measure how many bugs you fix per hour, or how many features you validate with tests.
---
- Use your curiosity on internal mechanisms (how the Go runtime handles memory, how slices grow) rather than external tutorials. Understanding Go internals will later give you backend superpowers (performance tuning, concurrency safety, etc.).
---
- Exploit your self-awareness: build a routine that deliberately isolates one task (for example, 2-hour sprints with a single deliverable—history fix, slice test, etc.). Treat distractions like compilation errors.

## Cause -> Effect -> Fix log
### git problem
- Cause: deleted my README.md file by accident
- Effect: I was desperate to recover it and that made me learn git commands and and understand the directory structure better
- Fix: i had to go back to the strict directory it was deleted in, and just restore it, using git restore README.md inside the specific directory (check git status to see where you are, source control is also useful for understanding where the file was deleted at)
### calculator index condition problem
- Cause: my calculator was not working as intended, it was skipping some inputs and not counting them properly
- Effect: i tried searching for some fitting condition for the for loop, learning the limits of the conditions you can possibly make in go
- fix: for index = 0; index < len(history); index++ it uses the element count of the slice to determine how many times to loop, instead of trying to find a condition that matches the input value( i just needed something to track how many times the calculation has run, and i never knew u could use len() in the for loop condition)