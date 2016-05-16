# parser-test
writing a parser in Go, used as a learning experience

### Explanation
I have had a basic understanding of how parsing a programming/markup language for quite a while, however I have never
actually written a parser. It seemed like a fun idea to mess around with it, and I have been programming with Go
recently, so I figured, why the hell not? I can learn more about Go, and write a parser!

### What it is
This will be a very simple [data serialization](https://en.wikipedia.org/wiki/Serialization) language (Primarily meant
for configuration), with a simple language specification.

### Language Spec

#### Variables
Variables can be Strings, Arrays, or Numbers (integers, or floats).

Example:
```
  integer = 1
  float = 1.337
  string = "Hello World"
  array = [
       "Element",
  ]
```

All variables are declared like so:
```
name = value
```

The spaces between them are not required, as they are ignored.

#### Arrays
Arrays should be comma separated. Arrays may contain multiple data types. The types are limited to Strings, and Numbers.
I would like to implement dictionaries, eventually. Single element arrays should end with a comma.

