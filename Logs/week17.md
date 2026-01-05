# Golang learning log - Week 17
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
## Problems faced:
## Notes: 
## End of week thoughts: