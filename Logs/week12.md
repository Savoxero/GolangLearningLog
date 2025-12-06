# Golang learning log - Week 12
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
### Plans right now:
1. ~~Validate inputs early: check strconv.Atoi error immediately and only then run the search loop.
Check scanner.Scan() result and scanner.Err(); handle EOF gracefully~~
2. Improve user output: friendly messages and a formatted list command.
3. Add a simple persistence layer (JSON file save/load) so IDs survive restarts.
4. Centralize parsing/validation into one function and add unit tests for it and idgenerator.
5. Add bounds/length limits and duplicate detection to avoid bad inputs.
6. Expand help text and add examples for each command.
#Progress:
- [x] Validated inputs early and handled errors immediately. (EOF,I/O errors)
- Working on multiple line inputs for adding tasks.