# Executing 

``` golang
$GOPATH/bin/hello 
go run hello.go //runs the executable as temporary
``` 
new main -> new folder -> new file
# Compiling 

``` golang
go install github.com/karinakozarova/project_folder
```
# Semicolons
Semicolons are optional.
The compiler is inserting the semicolon where they have to be
# Info
**Everything in go is a package:** <br>
package main contains a function main, which prints hello,world. <br> <br>
package main <br>
import "fmt" -> contains Printf <br>
All variables can be printed with %v <br>
All packages: <a href = "goland.org/pkg"> Packages</a> or
go doc fmt (shows fmt)

# Data types

## integer types
There is no automatic conversion between int types. <br>
int8, int16, int32, int64 <br>
uint8, uint16, uint32, uint64 <br>
<hr>

### machine dependent types

**uint, int, uintptr** <br>
Their sizes depend on the system. <br>

<hr>
rune data type - aliased to uint32(used to get to unicode encoded characters)
byte data type - aliased to uint8 (used to get to ascii encoded character)
<hr>

## floating point types
float32, float64 <br>
float64 - general use <br>

### two complex number types:
complex64, complex 128
<ul>

<li> specified by their bit size
<li> increase in bit width equates to an increase in precision, not range

</ul>

## strings
<ul>

<li> arrays of bytes/runes
<li> can be enclosed in " " or ` `
<li> a `string` can contain newlines
<li> a "string" can contain special charactesrs such as \n(new line) or \t(tab)
<li> can be be indexed
<li> the first index of a string is 0 (all arrays in Go are 0-based)

</ul>

## Booleans
3 boolean operators: &&,||,!  <br>
relational operators: ==, !=, >, <, >=, <= <br>
All relational operators return true||false <br>

## Declarations
Go is strongly typed so a variable always has only one specific type
variables are declared
``` go
var name type optional assignment
var hello strintg = "hello world"
```

If type was not given, the compiler will infer.
``` go
var age = 16 //int
var name = "user" //string
``` 

### := operator
Use type inference to assign a value to a new variable name. <br>
This syntax can't be used outside a function.

``` go
x := 16 //is the same as
var x = 16
``` 

If a non-typed var or := declaration is made:
<ul>
<li> anything that has no . or "" will be an int
<li> anything that appears as float will be typed as float64
<li> anything that appears as complex num will be complex128

</ul>

## Grouping multiple vars

``` golang
var(
    a=1
    a=6
) 

var a, b = 3, 5
a, b := 3,5
```

In global score,we need the var!

## Functions

### function declaration
``` golang
func func_name(parametres) return_type {
}
```
### returning simgle value
``` golang
func func_name(parametres) return_type {
      //Code goes here
      return result_as_return_Type
}
```
### returning multiple values

``` golang
func func_name(parametres) (return_type_1, return_type_2) {
      //Code goes here
      return result_as_return_Type1 result_as_return_Type2
}
```
### naming return values example
``` golang
func swap(b, a) (x int, y int) {
x, y = b,a
return //returns x and y
}
```

### variadic parameter list
``` golang
func adds(integers ...int) int {
//range
}
``` 
## Arrays

### Declaration and printing
``` golang
//declaration
variable_name := [array_count]type{numbers come here}
//printing
fmt.Printf("%v",variable_name)
``` 

### range and len of array

``` golang
for i:= range array_name
        //goes in for each element from the array
}
len(array)  // finds the length of the array
```
## Loops
There is only a for loop in Go

``` golang
for i := 0; i<count; i++ {
    //code here
}
for {
        //infinite loop
}
``` 

## If, else, else if

``` golang
if a==b {
            //some code here
} else {
            //something else here
}

if a==b {
            //some code here
} else if a>b {
            //something else here
}
else {
            //something else here
}
```
## Switch statement

``` golang
switch {
            case case_value>45564546:   /455... is a random value I choose
                    //something happens
            case other_Case_value = 4 :
                    //something other happens
            default:
                    //some default code for all other conditionals
}

switch variable_name {
        //same as withouth a variable
 }
 
 switch variable_name {
        case 1,2,3: 
        //do something here
        case 4,5,6:
        //do something else
        default:
        //default code for other conditions
 }
```
## Slices
slice - a reference to an underlying array
