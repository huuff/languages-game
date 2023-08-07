# Languages Game

### 001. Hello World
#### Level 1
Write a program that prints `"Hello World!"` to the screen (standard output).

#### Level 2
Receive a name from command-line arguments and print `Hello «name»!` to the screen. If no argument is provided, print `Hello World!`

#### Level 3
Print a different greeting depending on the time of the day, that is:

( TODO: Maybe remove this... it's hard to fake time to test it )

* Before 12:00 : `Good morning, «name»!`
* After 12:00 and before 18:00 : `Good afternoon, «name»!`
* After 18:00 : `Good evening, «name»!`

### 002. Fizz Buzz
#### Level 1
Just telling whether a command-line argument is Fizz, Buzz, FizzBuzz or just a number

#### Level 2
Automatically produce fizzbuzzes at a fixed rate given by command-line args. A good chance to learn blocking threads (e.g. block for 1s)


### 003. Currency Conversion Service
(TODO: Maybe using an API?)

### 004. SemVer manager
Some ideas:
* Parse major, minor and patch
* Increase major, minor or patch
* Check whether one version is afer another
* Given a list of expressions (maybe as a file?) parse a NPM version expresion (https://docs.npmjs.com/about-semantic-versioning#using-semantic-versioning-to-specify-update-types-your-package-can-accept) and choose an appropriate one.
