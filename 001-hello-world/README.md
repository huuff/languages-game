# Hello World
There are 2 variants of this exercise:

## Variants

### CLI
### Level 1
Write a program that prints `"Hello, World!"` to the screen (standard output).

#### Level 2
Receive a name from command-line arguments and print `Hello «name»!` to the screen. If no argument is provided, print `Hello, World!`

#### Level 3
It takes two command-line arguments:

1. A language code in ISO 639-1 (see the salutation table at the bottom).
2. A name to greet.

Greet the user using the name of the second argument (or the language equivalent of `World` if not present) in the provided language. Check the bottom salutation table to see the appropriate salutation for each language. Your solution should:

* If no arguments are given, print `Hello, World!`
* If one argument (the name) is given, print `Hello, «name»!`
* If two arguments are given (name and language):
    * If the second argument is a language of the table, return the correct salutation for that language.
    * If it isn't, exit with a non-zero exit code.

#### Level 4
TODO: Maybe receive parameters in any order with flags like `--name` or `--language`?

##### Learning opportunities
* If your programming language has enumerations, they are a good fit to represent locale codes. Also, you can investigate how to turn a locale string provided as an argument into an enumeration value.

### API

#### Level 1
Write a program that returns the string `Hello, World!` with a `200` code when receiving a `GET` request with `content-type: text/plain`

#### Level 2
Write a program that:

* When it receives a `GET` request on `/hello/«name»` returns `Hello, «name»!` 
* When it receives a `GET` request on `/hello` just returns `Hello, World!`

#### Level 3
Same as Level 2, but:

* The language of the salutation is provided in the request's `Accept-Language` header
* The same language is returned in a `Content-Language` header in the response
* If the language is not one of the known ones (see table), return a `400 (Bad Request)` error code.

## Salutations


| Language | Salutation with name | Salutation without name |
|----------|----------------------|-------------------------|
|`en`      |Hello, «name»!        |Hello, world!            |
|`es`      |¡Hola, «name»!        |¡Hola, Mundo!            |
|`fr`      |Bonjour, «name»!      |Bonjour, le Monde!       |
|`de`      |Hallo, «name»!        |Hallo, Welt!             |
|`it`      |Ciao, «name»!         |Ciao, Mondo!             |
