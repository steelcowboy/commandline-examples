# Grep
Grep is a utility to search files for strings of text. It is incredibly powerful
and is commonly used to search the output of commands for information.

## Searching a file
To search a single file for a string, run:
```
$ grep fmt goroutine.go
```
This will print out all lines that contain "fmt" somewhere in them.
If you want to see what line numbers these lines are, you can use grep's `-n`
flag:
```
$ grep -n fmt goroutine.go
```

## Searching the filesystem recursively 
To search files in a directory and all of its subdirectories, you can run:
```
$ grep -r fmt
```
This will print out the matching lines from any file that contains "fmt"
