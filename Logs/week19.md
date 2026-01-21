# Golang learning log - Week 19
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
- refactor the code structure to improve separation of concerns and maintainability.
- implement save/load to file functionality.
- test the new structure and ensure all functionalities work as expected.( my program relied on global variables too much, making it hard to test and maintain)
## Progress: 
- i am rewriting alot of the logic, i've been back and forth on how to structure it best. i was trying to build transactional semantics but i ABSOLUTELY despise it, i do not believe i will ever like this style of designing. maybe for very serious and complex applications but not for my simple todo list, i much prefer best effort semantics where if something fails, it fails and the user can retry or fix it.
- i have restructured the code into several packages:
  - main: handles CLI parsing and user interaction.
  - domain  : defines core data structures (Task, tasks) and business logic.
  - soon: persistence: manages saving/loading tasks to/from a file.

  - i need to document more stuff and learn how to use git add -p better, changes regarding a specific feature should be in a single commit. 
   - also adding comments to functions and methods to explain their purpose and usage.