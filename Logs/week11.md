# Golang learning log - Week 11
## Skills to learn:
- Git and Github fluency and learning the commands and navigation 
- Focusing on one thing at a time and doing it good enough (i have a tendency to jump around topics and not stick to one thing) (beeng more disciplined with my learning)
- cause -> effect -> fix logging for problems encountered
## Current plans:
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
 ## Progress:
 - added a much simpler function for a ID generator for tasks, it was much too overcomplicated with conditions that were not needed
 - added a way to check the saved tasks inside the slice, linked with the id, an option to view all tasks or explicitly calling out a task by its id.
 ## Notes:
 - Analysis and thinking about a basic structure leads to much better results, after writing boilerplate or some prototype u can see what needs to be changed or improved, writing with some future complexity in mind is not the way to go, starting simple atleast for me is a better approach.
 - the biggest weakness of my todo app is the fact that the tasks are not saved anywhere, it only works when the program is running, so persistence is the next big step, and that would require me to rewrite and revamp alot of stuff.
 - functions for things like adding, removing, viewing tasks are getting more complex as i add more features are gonna be necessary, i am gonna put them in a seperate file for better organization and code readability.