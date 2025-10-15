# Golang learning log - Week 4
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
- len is just counting how long/how much space the specific slice/array has.