# goZigZag
ZigZag Encoding Library In Go

## Introduction
ZigZag Encoding is a way of encoding signed integers into unsigned ones. It works by doubling the absolute value of the number and subtracting 1 if the orginal number was negative. Thus the negative integers are encoded to odd numbers and the positive integers to even numbers.

## Using the package
```
package main

import (
	"fmt"
	"github.com/el10savio/zigzag"
)

func main() {
	number1 := -2
	number2 := 5
	fmt.Println(zigzag.Encode(number1)) // 3
	fmt.Println(zigzag.Encode(number2)) // 10
	fmt.Println(zigzag.Decode(3)) // -2
}
```

## Testing 
The package uses fuzz tests to verify the properties of the encode/decode functions.

## References
 - [Making all your integers positive with zigzag encoding](https://lemire.me/blog/2022/11/25/making-all-your-integers-positive-with-zigzag-encoding) [Daniel Lemire]
