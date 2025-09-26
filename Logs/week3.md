#Update: Been Busy with real life stuff, school, work. I've had to focus almost entirely on my waitressing job and school work. And i couldn't find the time to stay consistent with my learning and my projects. It's time to lock in and get back to it. It feels really good to be back. But i also feel ashamed for not being consistent, its hard to forgive myself for that. But i must move on and keep going. if i want to be a backend developer. Sucess comes with hard work and discipline.
##Golang Learning Log - Week 3
On the track on learning Go. Trough boot.dev, i must say really good course, it allows u to experiment with stuff and go on ur own pace. the lessons spark my curiosity to try some projects for fun. The problem solving aspect and the assignments are BRILLIANT. I wouldn't design it any better.

### Stuff i learned this week:

- Printf statements such as:
  > \n = newline character
  > %d = decimal integer essentially int
  > %s = string
  > %f = float 64
- Syntax for declaring variables:
  Booleans, strings, int8 to 64 and the difference between the bytes they use up for memory efficeiency.
- variable bytes like strings use 1 byte per character. bool are fixed types declared by the system. int depends on the system architecture. 32 bit = 4 bytes, 64 bit = 8 bytes. so like int8 number: 122=1byte. unsigned integers being more efficient for memory 0-255=1byte. negative numbers aren't in the scope of uint types. Still fuzzy about calculating the memory usage of variables like inside a function or structs.
- constants and how theyre unchangeable values. u cannot overwrite them by any means, u can only declare them once. useful for values that are set in stone. u can also fuse constants together they don't work quite the same as variables.

#### concepts:

- truncating. the idea to replace a float with a integer value to either cut the decimal place or round it up, works vice versa. u can also do it with a string by converting it to a byte slice. example also below:
  func main() {
  str := "Hello, World!"

      // truncate to first 5 characters
      truncated := str[:5]

  example below:

```go
temperatureFloat := 88.26
temperatureInt := int64(temperatureFloat)
```

###Notes:

- I cant really note everything i've learned, ill share some projects and code snippets that i found interesting.
- I just learn on the go and note in obsidian anything relevant.
- I also use Claude to help me understand stuff and explain concepts to my liking in a manner that i prompted it to.
