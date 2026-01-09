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
- renamed and rewrote the remove function into delete. i added some proper error handling but went on a shortcut when it comes to the desired output, i will change that tommorow but for now it's functional. Not good.
- completely restructured karpdel.go, added alot of edge case handling and proper error handling. it took me about 2 days to get the ideas right but i think it's in a good place now. i made the design choice of ignoring empty input in a chain deletion and ignore the input that has text for ID's but when it's single deletion, it will throw an error. this is to avoid confusion and make the user experience better. i want the program to be as malleable as possible without breaking or being too strict.
- also refined the main.go file to accomodate the new changes.
## Problems faced:
- been trying to share the data to main in a efficient manner, i tried functions, but it was just empty output unecessary complexity. so i went with global variable for now, will refactor later. i do not like this decision. i need to somehow transport the slice data into main so the user knows what was deleted without necessary bloat.
- there's been alot of problems generally with edge cases and error handling, i am not gonna type them all out, but it made me stare at the screen for 10 minutes at a time. While going to the toilet i was wondering how the control flow should go on, in my spare time i also thought about it. generally pretty proud. i didn't use any external help for this one.
## Notes: 
- I realized that every design choice has a price to be paid for. I was trying to design stuff that really didn't go with the architecture of the program. for a little while. It's massive insight for the future, you can't have everything at once. a design choice comes with sacrfices.
## End of week thoughts: