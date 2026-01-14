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
- added multiple edge case handling and error handling to karpprint.go, remade the control flow so it does exactly what i expect it to do. EG: user has a task with the id of 1 and with the description of go take out the trash and id 2 is do the laundry, if the user types delete 2 & 3 & abc & "" it will delete task 2, ignore 3, ignore abc and ignore "" and print out that task 2 has been deleted. if the user types delete 3 it will throw an error that task with id of 3 doesn't exist. this is to avoid confusion and make the user experience better. single input is strict, multiple input is lenient.
## Problems faced:
- been trying to share the data to main in a efficient manner, i tried functions, but it was just empty output unecessary complexity. so i went with global variable for now, will refactor later. i do not like this decision. i need to somehow transport the slice data into main so the user knows what was deleted without necessary bloat.
- there's been alot of problems generally with edge cases and error handling, i am not gonna type them all out, but it made me stare at the screen for 10 minutes at a time. While going to the toilet i was wondering how the control flow should go on, in my spare time i also thought about it. generally pretty proud. i didn't use any external help for this one.
- mixing up logic and printing, need to separate them properly. karpprint.go is a mess right now. i will refactor it tommorow. split everything.
## Notes: 
- I realized that every design choice has a price to be paid for. I was trying to design stuff that really didn't go with the architecture of the program. for a little while. It's massive insight for the future, you can't have everything at once. a design choice comes with sacrifices.
- i rewrote some of the code in karpprint.go but i am mixing up logic, formatting and printing. i need to exactly separate each role. logic should be just logic, return an error if something goes wrong. otherwise return the data. as simple as that. formatting should take the data and format it into a string. printing should just print the input to the user in main.go from all the data gathered and mutations made in other packages.
- i feel great coding with music on. helps me focus and get into the zone faster. Problem solving is really enjoyable when you get into the flow state. 
## End of week thoughts:
- Pretty productive week overall. currently on a job search so i have lots of time to code and learn golang. 
- i do wonder how far i can take this project to. some of the idesa are: 
 - letting the user add a category to the struct, and then filter by category when listing tasks.
 - adding due dates and reminders via a background goroutine that checks for due dates and pop external notifications via a windows tool.
 -  detailed modification of tasks, like changing priority, adding tags, changing due dates etc.
 - undo/redo functionality via a command stack. 
 - and filtering and sorting tasks by different criteria.