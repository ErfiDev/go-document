# Basic Questions

## what does "package main" mean

when we see the word package in go
we can think of a package as being like a project
or a workspace
a package is a collections of common source code files
so a package can have many related file inside of it each
file ending with a file extension of .go

the only requirement for every file inside of the package
is at the very first line of each file must declare the package that it blongs to

two types of packages

- Executable
- Reusable

  <br>
</br>

## what does "import" mean

the import statement is used to give our package
<br>

```
import "fmt"
```

<br>
so saying import fmt specificly means give my package main access to all of the code and all the functionality that is contained inside of other package called fmt

<br>
</br>

- fmt the name of the standard library package that is includes with the go programming language by default the fmt package is used to print out a lot of different information specificly to the terminal
  just give you a better sense of debugging and stuff like that

<br>
</br>

## what is that "func" thing

func is a short of the function in golang

```
    func main(){
        // code
    }
```

<br>

1. func is a function keyword
2. main is a function name
3. () is a function arguments
4. {} this is a function body
