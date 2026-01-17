# Golang learning log - Week 18
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
## Now:
- finish formatting within printing package, make add and remove also formatted properly.
- implement save/load to file functionality.
## Progress:
- printing formatting is done. it is a MVP for now.
- add and remove commands are formatted properly. With the new printing package, the output is much better.
- made the grammar a little better for some commands.
- clean separation of concerns between packages. printingvalidation and core logic are separated, print just prints what was validated and processed by other packages.
## Next:
- implement save/load to file functionality.