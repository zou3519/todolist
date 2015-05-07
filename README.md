### todolist

by Richard Zou and Peter Ku

A top-down skiplist (todolist) is a type of dictionary that supports 
insert, search, and delete operations in O(log(n)) time.  
We implemented a todolist and benchmarked it against
other related data structures for our
Spring 2015
CS223 Probabalistic Algorithms class.

### Setup 

If you're new to Go, set up your GOPATH and Go workspace with
the instructions described in https://golang.org/doc/code.html.  Go
is rather picky about its workspace options but allows for some
very useful modularity as a result of this workspace organization.

When you are setup, get our code with the following command:

    go get github.com/zou3519/todolist

The following dependencies are necessary to build this project. 
- [stathat.com/c/treap](stathat.com/c/treap)
- [github.com/emirpasic/gods/trees/redblacktree](github.com/emirpasic/gods/trees/redblacktree)
To install a dependency, type
  
    go get stathat.com/c/treap

to install the dependency for treaps.

### Testing and running the code

Now that you have downloaded our code, type 

    go test

at the commandline to run all the tests. After the tests are successful, type

    go build

to compile the code into to a executable file, todolist.  