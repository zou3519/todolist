### todolist

by Richard Zou and Peter Ku

A top-down skiplist (todolist) is a type of dictionary that supports 
insert, search, and delete operations in O(log(n)) time.  
We implemented a todolist and benchmarked it against
other related data structures for our
Spring 2015
CS223 Probabalistic Algorithms class.

### Setup 

You will need to set up Go and install the dependencies to run our
code. Luckily, this is a pretty straightforward process.
If you're new to Go, set up your GOPATH and Go workspace with
the instructions described in https://golang.org/doc/code.html.  Go
is rather picky about its workspace options but allows for some
very useful modularity as a result of this workspace organization.

When you have finished setting up Go, get our code with the following command:

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

### Running benchmarks

The main product of this repository is a benchmarking tool that
can benchmark a particular data structure that we have implemented
or have wrappers for.  Then benchmark will generate a 
csv file in ./Outputs/.  Run the benchmark with

    ./todolist -mode=<mode> -epsilon=<epsilon> -reps=<reps> -n=<n>

The parameters all have defaults and so are optional. Parameters
will only be used if the mode specified requires the parameters.

Parameters:
- epsilon: The epsilon constant used in todolist variants
- n: The maximum number of items to insert/search/delete
- reps: How many trials to perform for each value of n. The program
will average out the values automatically.
- mode: Which data structure to benchmark

Modes:
- epsilongraph: Generate a csv comparing the trade-off 
between search and insertion time as a function of epsilon.
- redblack: Benchmarks search, insert, and delete for a red-black tree
- treap: Benchmarks search, insert, and delete for a treap
- mapset: Benchmarks search, insert, and delete for a hash table
- skiplist: Benchmarks search, insert, and delete for a skiplist
- optimalbst: Benchmarks search for an optimal binary search tree
- linkedtodolist: Benchmarks search, insert, and delete for a 
standard todolist implemented with linked lists
- todolist: Benchmarks search, insert, and delete for a todolist implemented
with dynamic arrays
- todolist2: Benchmarks search, insert, and delete for a todolist implemented
with dynamic arrays with one adjustment
