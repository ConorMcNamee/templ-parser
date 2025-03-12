# Calling Go functions dynamically

This is a tricky process and isn't exactly striaght forward. This will be a investigation into the packages required to pull this off and how exactly I'm going to go about it. 

From my initial research it looks like there are a number of libraries to look at that will achieve this. These libraries are:
    - Parser
    - Reflect
    - Plugin
Below you will see a deep dive into each of these pacakges and how they work and How it will affect this project

# Parser

The parser package implements a parser for go source files. Input can come in various forms and outputs a abstract syntax tree (AST) representing the go source code. Once we have all the files we want to be able to parse we can run `parser.ParseFile`

# Reflect

The reflect package implements runtime reflection allowing a program to manipulate objects with arbitrary types. This allows us to extract dynamic type information by calling TypeOf which returns a type