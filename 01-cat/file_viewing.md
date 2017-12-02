# The cat command
cat will conCATinate files and dump them to its standard out.
The simplest example is the case of reading one file:
```
$ cat file1.txt
This is a cool file
```

# Multiple files
As the description implies, cat can also read multiple files and dump them to
standard out. This takes advantage of the fact that in UNIX, everything is a
file; since stdout is a file, by putting all of the file content there we've
effectively turned many files into one.
```
$ cat file1.txt
This is a cool file
$ cat file2.txt
This is another cool file
$ cat file3.txt
What if they were one?
$ cat file1.txt file2.txt file3.txt
This is a cool file
This is another cool file
What if they were one?
```
