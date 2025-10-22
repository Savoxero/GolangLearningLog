# Golang learning log - Week 6
## Skills to learn:
- Git and Github fluency and learning the commands and navigation
- Planning and realizing project ideas
- Focusing on one thing at a time and doing it good enough (i have a tendency to jump around topics and not stick to one thing)
## Current plans:
- im gonna refine the calculator and finish it. 
- [x] adding a way to track every past calculations(history with a slice) 
- [x] moving the calculations into its own package( not a feature )
- [x]continuous calculation
- [] word counter using map[string] int
- [] a system to automaticly exit if the input is repeated and the program does not execute as designed.
- after everything above is done Resume boot.dev trough becoming a member. still nowhere near being able to solve leetcode ;(
## Progress
- added comments to the calculator
- renamed the calculator to WoopCalc.go (Pokemon wooper reference) 
- made the operations and functions used in the calculator as a separate package. ( learned about explicitly mentioning the package like in this context u have to do calc.Add() and I'm pretty sure the functions have to start with a capital letter)
- Learned about len, slices, arrays, append
 - some notes about Arrays:
  - they are boxes. the number inside the [] defines how much places there is inside the array.
  - u can explicitly give a specific array a typed value, by stating which element should be the specific value.
  - 0 counts does not count as one so if its a[5] = 100 (0,0,0,0,0,100) kind of like after the specific number
  - [...] makes the compiler count how many arrays there is. and it's also used when u don't know how big u want the the array to be
  - : zeroes the specific places (1,2,2:,5) gonna print 1,2,0,0,5
- some notes about slices: 
- You can Temporarily grow them with `append`, slice them with `[:]`, and pass them around freely (they’re lightweight references).
- same as arrays 0 is 1, 1 is 2 they work in a referential manner
EG:
```go
 arr := [3]int{1, 2, 3}
sl := arr[1:]     // slice starting at index 1 → elements {2, 3}
sl[1] = 99    // 0, 1 (0 counts as the first element, and sl 1 corresponds to second element in the array/slice)    // transforms the arr [3] into 99 -> direct slice 
```
- added slice history to the calculator along with the command to see it, gotta add comments and do minor adjustments ( im still fuzzy about slices but they seem to be pretty easy to grasp, atleast when dealing with one)
- Got way more comfortable with git and the terminal, i use the terminal more often for most of the time and some keybinds, like replacing everything at once, searching (ctrl f), opening a new directory with the specified folder. etc.. much more to learn bout git

## Notes
- len is just counting how much elements the specific slice/array has. ( i misunderstood it as cap before)
- cap, the capacity of said array or slices, it expands by 2x generally, u can also do int[]{0,5} and this is gonna make it so the capacity is 5 but len is 0, u explicitly say that theres zero elements inside, if u dont it'll all be 5'se
## Progress on calculator ( i know there's alot to read but its mostly notes for myself)
- i've been trying to fix a specific loop issue that caused to print the error on the enter the operation part, turns out i just had to add a continue statement after printing the error, so it would skip the rest of the loop and go back to the start.
- minor tweaks to quality of life, like floating point numbers cut by 2 decimals so its readable and not a bunch of zeros that take so much of the terminal space
- i've come to realize that i did not need a if operation =! "+" && operation != "-" ... etc...
and all i need is a simple print with please enter a valid operation with continue right after it, so it would skip the rest of the loop and go back to the start.
- i've become more fluent with my computer terminal usage, i use it more often than the gui for git operations, i still have to learn alot but im getting there.
- understood debugging a little but much more to explore.
# End of week 6 TLDR:
- calculator is looking good, added history with slices, moved operations to their own package, added continuous calculations, error handling for invalid operations, minor quality of life tweaks like rounding floats to 2 decimals.
some stuff to do: 
- [] word counter using map[string] int
- [] a system to automaticly exit if the input is repeated and the program does not execute as designed.
- reading a shitton of documentation on git and programming generally
- removing nested loops and finding a way to replace them with functions or another clever way
Notes:
I've been busy with the job so i haven't had much time to code. its a shame because i know i can do so much more. kinda like the gym, even tho u exert urself to failure and beyond you still feel like u could do more. programming is the same way, u can always push urself further. even though its counterproductive sometimes.
