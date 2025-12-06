# Golang learning log - Week 13
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
7. gotta add alot of MVP features before moving on to the big plans. and add some more intuitive QOL features and text.
--- 
## Progress:
- Added multiple input validations through "&" EG: ( add go buy a gallon of milk & text Andrew", this makes 2 tasks in one line) along with error handling, and generally tested edge cases and i put some control flow to the edge cases.
- completely restructured the entire remove logic along with adding multiple input through ("&") EG: ( remove 1 & 3 & 5 removes tasks with ID's 1,3,5) along with error handling and edge case testing. ( did it on my first try without any major issues) kinda proud of that. ^^
- added a multiple input to the check command, EG: ( check 2 & 4 & 6 ) else just type check and it'll show everything that is currently saved.
## Cause -> Effect -> Fix log:
1. Cause: while trying to implement the multiple input addition, i overgineered the input parsing and validation logic, making it convoluted and hard to follow.
   Effect: This led to complexity and made it vulnerable to multiple edge cases that were hard to track.
   Fix: made a simpler for loop to trace each slice of input, validate it, and then add it to the task list. This made the code more readable, maintainable and adaptable a single input works but so does multiple one's.
2. Cause: Forgot to apply `TrimSpace()` to split strings before processing in multi-delete loop.
   Effect: First delete worked, rest failed silently. Spent 2 hours debugging loop logic instead of data prep.
   Fix: Moved `TrimSpace()` to immediately after the split, before any other processing.