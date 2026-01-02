# Golang learning log - Week 16
## The big picture
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
## After rewriting the codebase(might be wrong about not doing it earlier but ehh well see):
 ## To do:
    1. implement save to file function to save tasks to a json file on every change
    2. adding persistence to the file, at the moment it gets overwritten on every save.
    3. converting the json data into the existing struct format for tasks, so that the rest of the code can work with it seamlessly eg: adding, removing, listing tasks. problem is gonna pop if it doubles all the data when loading from file. gotta find a way to make it so marshal overwrites at the end. And upon the launch of the program, it should load everything into the struct slice.
## Now:
- Rewriting the codebase and splitting everything into their own packages for better code structure and readability. after that is done, my main focus is going to be to refine this as much as possible and then move on to implementing the to do list features one by one. alot of work ahead of me but i am excited to tackle them all.
## Progress: 
- added KarpLogic.go file for my internal logic functions the core of the program
- reorganized main.go to just handle the CLI and call functions from KarpLogic.go i want main to only print messages and print errors that are returned from KarpLogic.go
- rewrote some of the code, i had to make some changes to how variables get their values, scopes etc. one issue i encountered was variable scoping within the add function, i had to create a var outside the if block to hold the task value so it can be printed into main flawlessly, i am likely to encounter more of these issues as i go along.
Problems faced:
- Karplogic.go
 - problem encountered: i was confused about double printing and double task adding but i realized that it was because i was calling the add function twice in main.go once for adding and once for printing. fixed that by removing the extra call. this took me 20 minutes, i thought there was a issue with 2 loops being called. Took me 20 minutes to figure that out ONE TINY function CALL smh.