# Languages Game

### 001. Hello World
#### Level 1
Write a program that prints `"Hello World!"` to the screen (standard output).

#### Level 2
Receive a name from command-line arguments and print `Hello «name»!` to the screen. If no argument is provided, print `Hello World!`

#### Level 3
It takes two command-line arguments:

1. A name to greet.
2. A language code in ISO 639-1. One of: `en`, `es`, `fr`, `de` or `it`. Choose english if not present.

Greet the user using the name of the second argument (or `World` if not present) in the provided language. The salutations in each language are:

* `en`: `Hello, «name»!` or `Hello, World!`
* `es`: `¡Hola, «name»!` or `¡Hola, Mundo!`
* `fr`: `Bonjour, «name»!` or `Bonjour, le Monde!`
* `de`: `Hallo, «name»!` or `Hallo, Welt!`
* `it`: `Ciao, «name»!` or `Ciao, Mondo!`

##### Learning opportunities

* If your programming language has enumerations, they are a good fit to represent locale codes. Also, you can investigate how to turn a locale string provided as an argument into an enumeration value.

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
