# Languages Game

### 002. Fizz Buzz
#### Level 1
Just telling whether a command-line argument is Fizz, Buzz, FizzBuzz or just a number

#### Level 2
Automatically produce fizzbuzzes at a fixed rate given by command-line args in milliseconds. If none is provided, use `1000` by default (1 second). A good chance to learn blocking threads (e.g. block for 1s)

### 003. Unit Conversions
#### Level 1
Receive three command-line arguments (anything else will be considered an error): `«input measure» to «output unit»`:

* The `«input measure»` carries both an amount and a unit (e.g. `21ºC`)
* The second argument is always `to`
* The third argument is only a unit, without an amount (e.g. `ºF`)

For this level, convert only between celsius and fahrenheit degrees, and return an error `Unrecognized unit: «unit»` for any other unit.

### 004. SemVer manager
Some ideas:
* Parse major, minor and patch
* Increase major, minor or patch
* Check whether one version is afer another
* Given a list of versions (maybe as a file?) parse a NPM version expresion (https://docs.npmjs.com/about-semantic-versioning#using-semantic-versioning-to-specify-update-types-your-package-can-accept) and choose an appropriate one.
