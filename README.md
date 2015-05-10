# gofizzbuzz
This is a command line fizzbuzz implementation in go.

###### Installing
1. go install https://github.com/caustin/gofizzbuzz
 

###### Command line usage 
Invoke the program by executing gofizzbuzz:

    chris@Oxidation ~> gofizzbuzz

Command line parameters:

|parameter   |optional   |default   |example   |
|---|---|---|---|
|start   |yes   |1   |  gofizzbuzz --start 10 |
|last   |yes   |100   |  gofizzbuzz --last 10 |

    chris@Oxidation ~> gofizzbuzz --start 10 --last 30
    Buzz
    11
    Fizz
    13
    14
    FizzBuzz
    16
    17
    Fizz
    19
    Buzz
    Fizz
    22
    23
    Fizz
    Buzz
    26
    Fizz
    28
    29
    FizzBuzz
