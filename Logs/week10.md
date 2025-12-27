# Golang learning log - Week 10
## Skills to learn:
- Git and Github fluency and learning the commands and navigation
- Planning and realizing project ideas
- Focusing on one thing at a time and doing it good enough (i have a tendency to jump around topics and not stick to one thing)
- cause -> effect -> fix logging for problems encountered
---
## Current plans: 
- Todolist CLI, progress has been annoyingly slow but steady
- resuming boot.dev. after im done with this todolist cli, im thinking about starting boot.dev again
- refining my basics, slices, structs, pointers, functions, control flow etc
- spending more time coding and experimenting rather than reading, and looking up stuff.(the ide has limits but i built my calculator essentially only with it)
---
## Habits ( gpt notes)
- Design before keysmashing. Write the program flow as pseudocode before touching Go.
- pre planning
## Using my strenghts (GPT notes)
- Channel your pattern intuition into writing design notes before code. When you see repetition, stop, sketch the logic, then rewrite.
- Exploit your self-awareness: build a routine that deliberately isolates one task 
--- 
##Progress:
- added more control flow to handle inputs from the user and invalid commands, my program now gives feedback to the user when they enter an invalid command or try to delete a task that doesn't exist. alot of crashes before this
- currently studying both for loops, slices and struct integration together. im trying to understand how to manipulate slices of structs, adding, deleting, updating struct fields inside a slice. long road ahead of me. to say the least
- added magikarp, welcome message
- rewrote ID system along with a wip id generator, bool flag to inform the user whether the task was found or not
- added more comments to my code to improve readability and understanding of the code flow

## Cause -> Effect -> Fix log:
- Cause: Program crashes when user tries to delete a task that doesn't exist.
  Effect: User frustration and inability to use the todo list effectively.
  Fix: Implemented input validation to check if the task index exists before attempting to delete it. Added error messages to inform the user of invalid inputs.
  - Cause: double printing of specific messages
  Effect: Cluttered output, confusing user experience.
  Fix: replaced continue with break where necessary to prevent redundant message printing.
  ---
  - Cause: the print message surrounding the deleted id has been inaccurate due to a misplacement of the code execution order.
  Effect: User confusion about which task was deleted.
  Fix: Adjusted the print statement to execute before the task deletion logic, even though it seems minor, it greatly improves clarity.
  Cause: ID system was referencing slice indexes directly never comparing the correct ID's unless the slice number matched the exactly generated ID, i used to generate the ID with len(tasks) + 1 so i was blind for some time about it but i read trough documents banged my head on the keyboard at 2 am and figured it out, it was a mess that needed a complete rewrite and it humbled me deeply, it took me 4 days to figure it out.
  Effect: Deleting tasks would sometimes remove the wrong task or cause index out of range errors, not to mention the printed tasks and their IDs were all over the place.
  Fix: Completely rewrote the ID system to ensure each task has a unique identifier that compares DIRECTLY with the user's generated ID. It was truly a journey.
  ## to do:
- implement saving and loading tasks from a file
- completely rewrite the ID system. it blocks on [0] and the printed tasks do not match the index in the slice len(tasks) +1 does not work efficiently
- there are alot of other things i want to add. GPT has given me some cool ideas like categorizing tasks, API integration, deadlines and reminders etc. but for now im just focusing on getting the basics down and making a functional todo list CLI app.


## Notes, and random bullshit:
```go for IDsegregator := len(tasks) - 1; IDsegregator >= 0; IDsegregator-- {
					tasks[IDsegregator].ID = tasks[IDfinder].ID
					for IDsegregator = len(tasks) + 1; IDsegregator >= tasks[IDfinder].ID; IDsegregator++ {
						tasks[IDsegregator].ID = tasks[IDfinder].ID
						break
					}
				}
			}
      ```
 absolutely disgusting piece of code. i barely understand it myself. i want to use it as a reference for what NOT to do when writing code. im trying to rewrite the ID system completely. its a mess right now. (tried to renumber the IDS after segregation but it just made everything worse) 
 absolute headache. 
# Summary of week 10:
This week has been a rollercoaster of emotions and learning experiences. I've made significant progress on my todolist CLI and it's been alot of misery looking at my screen with cluelessness and emptiness as to what to do next, but slowly but surely im getting there. The biggest challenge has been rewriting the ID system, which took me days of frustration and headbanging, but I finally cracked it. I've also improved the user experience by adding input validation and clearer feedback messages.
note: too lazy to fix the weird code block above, its just for reference of what NOT to do. Looking ahead, I plan to implement file saving/loading and continue refining the app. 
