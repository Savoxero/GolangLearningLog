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


    ## Note on the past week 
    been busy with life and work, i am transitioning to a new job currently didn't have much time to code or learn golang. i did focus on expanding my knowledge on golang modules and packages and how to structure my code better. so there was some learning that happened even tho i didn't code much. 
    ## Progress this week:
    1. implemnted a load from file function to load tasks from a json file on startup (very early prototype, needs better error handling and some overall tests along with refinement)