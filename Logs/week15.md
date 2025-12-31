# Golang learning log - Week 15
## Skills to learn:
- Git and Github fluency and learning the commands and navigation 
- Focusing on one thing at a time and doing it good enough (i have a tendency to jump around topics and not stick to one thing) (beeng more disciplined with my learning)
- cause -> effect -> fix logging for problems encountered
## Big plans:
- realizing some ideas for CLI todolist some eg:
 Persistence
    - save/load tasks as JSON or gob to a file; autosave on change; manual export/import
- Filtering & sorting
    - by status, priority, date, substring; sort by due date or creation time
- Metadata
    - tags, due dates, estimated time, recurrence, subtasks
- Notifications
    - due reminders (background goroutine + time.Sleep or ticker); optional Windows notifications via external tool
- Undo/redo
    - simple command history stack for destructive actions
- Batch operations
    - add multiple, remove range, bulk modify priorities


    ## Note on the past week's progress:
    been busy with life and work, i am transitioning to a new job currently didn't have much time to code or learn golang. i did focus on expanding my knowledge on golang modules and packages and how to structure my code better. so there was some learning that happened even tho i didn't code much. 
    ## Progress this week:
    1. implemnted a load from file function to load tasks from a json file on startup (very early prototype, needs better error handling and some overall tests along with refinement)
    2. learned about marshalling structs to json, along with marshal indent, i find it to be useful if someone wanted to make the file more readable and concise, especially with structs. along with writefile and readfile from the os package. i never knew something like this existed in golang standard library. it is beyond words beautiful
    3. split the id generator into its own function for better code structure inside another file. it took me about 2 hours to figure ou how to properly use the return value from the function without corrupting the main logic and the backbone of the program. it was much simpler than i thought it would be. 
    4. split the magikarp and welcome message into a separate function for better code structure and readability. focus on main being clean and concise.

    ## To do:
    1. implement save to file function to save tasks to a json file on every change
    2. adding persistence to the file, at the moment it gets overwritten on every save.
    3. converting the json data into the existing struct format for tasks, so that the rest of the code can work with it seamlessly eg: adding, removing, listing tasks. problem is gonna pop if it doubles all the data when loading from file. gotta find a way to make it so marshal overwrites at the end. And upon the launch of the program, it should load everything into the struct slice.
    ### Later on: 
    - making everything either functions or variables into their own packages for better code structure and readability. essentially tidy up the codebase. because right now i do not see someone reading all of that and making sense. I do but still.
    ## Challenges faced:
    1. understanding how to properly read from a file and write to a file in golang.
    2. understand the scope of variables within loops and so they do not reset, i got lost within my code and i did not see i was in a loop the whole time. later on i realized i was resetting the id generator every time i loaded from file, the return value was not being stored properly and thus got reset every time. the fix was simple, create the variable outside the loop, make that variable call the id generator function, and store it into itself and just let the struct slice append to it raw.
    ## End of week thoughts:
    overall decent progress i now have a clear goal to achieve for the next week which is to seperate everything into separate packages and their own roles, cli handling, logic handling, infrastructure handling etc. it is gonna be the very first step into making something scalable and maintainable.
    the priorities have shifted for now. after rewrite most of the codebase i will then do what i said the to do.