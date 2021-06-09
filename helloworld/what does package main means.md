# What does package means ?
- packages == projects == Workspace
- packages are the place where the bunch of related go files live.
- consider the below example where  myapplication contains two packages client and server.

<pre>
myapplication/
    client/ 
    server/
</pre>

### Need for packages 
- To maintain our large codebase as packages makes it easy to navigate and access files.
- Gives the power of reusability,we can use are packages in multiple places.
- Portable
- It also paves way for advanced concepts we'll see it future.


 # Why is it named as "main" not any other name ?

- Generally Golang has 2 types of packages
    - Executable packages  ->  The one we use to execute of go code (binary file)
    - Resuable packages  ->  These are packages which contains helper function (Libraries, dependencies)

- So main packages are called by compiler which has  ```main()``` function which makes it has an entry point for the execution Go program
- main is special keyword
- Resuables packages will approriate names like ```strings,math,fmt,net``` etc...


# importing package

- we are allowed to import the packages(resuable type) which are created by developers and GoLang itself to access its functions and variables

#### Single package

```
import "fmt"
```

#### Multiple package

```
import "fmt"
import "math"
```

#### Short-Hand for Multiple package

```
import (
    "fmt"
    "math"
)
```

<b> Bonus :</b>
- ```fmt``` is a part of standard library
- used for I/O formating

