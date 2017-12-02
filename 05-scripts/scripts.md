# Scripts
A script is simply a text file with commands for some sort of interpreter to,
well... interpret. You can have a Python interpreter, for example, that reads a
text file and sees if it's valid code -- if it is, it turns it into machine code
on the fly. Similarily, you can have a bash interpreter that looks for valid
bash.

# Bash
Bash is the shell you commonly use on UNIX systems, and aside from being able to
process commands also has a rich set of logical features that allow you to write
code that operates primarily on files.

## Hallo Welt
Here is the "Hallo Welt" script (because Hello World is too cliché):
```bash
#!/bin/bash

echo Hallo Welt
```
To run this script, run `chmod +x <script>` to make it executable, and then you
can execute it by typing `./<script>`.

As you can see, the "echo" command is exactly what you'd type into your
interactive bash interpreter, i.e. command shell. The first line is called a
shebang, and when a file is executed that has a shebang at the top it simply
means _run the command <program> <file>_

## Abusing shebang
Quick side note: because the shebang just says "take this filename and pass it
to the specified program, if you wrote a file that simply said:
```bash
#!/usr/bin/rm
```
Executing this file would simply cause it to delete itself!

# Quick reference

## Varibles
```bash
test=5 # Note you cannot have spaces around the equals sign
echo "$test" # $<var> gets the value of <var>
```
Note that when accessing varibles it's good practice to put double quotes around
them.

## Capturing program output into a variable
```bash
hour=$(date +%H) # Run the program date, ask for the current hour, save to var
```

## Read input into variables
```bash
echo -n "Howdy, please enter something awesome: " # Don't print a newline
read awesome # Read from the terminal and put response into varaible "awesome"
echo "You're right, $awesome is awesome!"
```

## Loops
For loop:
```bash
# Printing numbers 1-10
for i in {1..10}; do
    echo "$i"
done
```
While loop:
```bash
i="0" # Notice how we work with strings. Gross
while [ "$i" -lt "4" ]; do
    echo "Iteration $i"
    i=$[$i + 1]
done
```
_Note: at this point I hope you notice bash is not very good for numbers_

## If statements
```bash
hour=$(date +%H)

if [[ "$hour" -lt 10 ]]; then
    echo "It's early"
elif [[ "$hour" -lt 16 ]]; then
    echo "Good day!"
else
    echo "Getting late"
fi
```
Note that `[` is actually the command `test`. As such, you must always have a
space after it! In addition, `[[` is a command built into bash that is an
extension of tha test commmand; it is less portable (won't work on all systems)
but more powerful.
