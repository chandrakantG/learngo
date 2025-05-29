package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	// Defer is used to ensure that a function call is performed later in a program's execution, usually for cleanup.
	// It is often used to close files, release resources, or unlock mutexes.

	// // Example of using defer to close a file
	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close() // Ensures the file is closed when the function exits

	// // Perform operations with the file...

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	e1()
	e2()
	e3()

	//what are defer types in go?
	// In Go, there are no specific "defer types" as such, but the `defer` statement can be used with any function or method call.
	// The deferred function can be a regular function, a method, or an anonymous function (closure).
	// The key characteristics of deferred functions in Go are:
	// 1. **Function or Method Calls**: You can defer any function or method call, including those that take parameters and return values.
	// 2. **Anonymous Functions**: You can use anonymous functions (closures) with `defer`, allowing you to capture variables from the surrounding scope.
	// 3. **Parameters**: You can pass parameters to the deferred function, and it will capture the values of those parameters at the time of the defer statement.
	// 4. **Return Values**: Deferred functions can return values, but they do not affect the return values of the surrounding function.
	// 5. **Error Handling**: Deferred functions can be used to handle errors, but they should not be used to change the return value of the surrounding function.
	// 6. **Resource Management**: Deferred functions are commonly used for resource management, such as closing files, releasing locks, or cleaning up resources.
	// 7. **Execution Order**: Deferred functions are executed in the reverse order of their declaration, meaning the last deferred function declared is the first one executed.
	// 8. **Execution Timing**: Deferred functions are executed after the surrounding function returns, just before it returns to the caller.
	// 9. **Panic Recovery**: Deferred functions can be used to recover from panics, allowing you to handle errors gracefully.
	// 10. **Variable Capture**: Deferred functions capture the state of variables at the time of their declaration, not at the time of their execution.
	// 11. **Named Return Values**: Deferred functions can modify named return values, allowing you to change the return value of the surrounding function.
	// 12. **Pointer Parameters**: You can pass pointers to deferred functions, allowing them to modify the original variable.
	// 13. **Multiple Defer Statements**: You can have multiple `defer` statements in a function, and they will be executed in the reverse order of their declaration.
	// 14. **Defer in Loops**: Be cautious when using `defer` inside loops, as it can lead to unexpected behavior due to the deferred functions being executed after the loop completes.
	// 15. **Defer in Goroutines**: When using `defer` in goroutines, the deferred function will execute when the goroutine completes, not when the surrounding function returns.
	// 16. **Defer with Interfaces**: You can defer methods on interfaces, allowing you to call methods on interface types.
	// 17. **Defer with Structs**: You can defer methods on structs, allowing you to call methods on struct types.
	// 18. **Defer with Channels**: You can defer operations on channels, allowing you to perform actions on channels after the surrounding function returns.
	// 19. **Defer with Slices**: You can defer operations on slices, allowing you to perform actions on slices after the surrounding function returns.
	// 20. **Defer with Maps**: You can defer operations on maps, allowing you to perform actions on maps after the surrounding function returns.

	// Rule: When deferring a direct function call like defer fmt.Println(err), the arguments are evaluated immediately.
	// Rule: When deferring a function with parameters like defer func() { fmt.Println(err) }(), the arguments are evaluated at the time of execution.
	// Rule: Closures capture variables by reference, so they observe the latest value at the time of execution.
	// Rule: When deferring a closure with arguments, those arguments are evaluated immediately — even if the closure runs later.

	var a = acumulator()
	fmt.Printf("%d\n", a(1))
	fmt.Printf("%d\n", a(10))
	fmt.Printf("%d\n", a(100))

	fmt.Println("----------------------")
	var b = acumulator()

	fmt.Printf("%d\n", b(1))
	fmt.Printf("%d\n", b(10))
	fmt.Printf("%d\n", b(100))

	defer fmt.Println("defer main")
	var user = ""

	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success. err:", err)
			}
		}()

		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic("should set usr env")
			}

			fmt.Println("after panic")

		}()

	}()

	time.Sleep(5000)
	fmt.Println("end of main function")
	time.Sleep(5 * time.Second)

	for i := 0; i < 5; i++ {
		defer fmt.Println(i, 1)
	}

	for i := 0; i < 5; i++ {
		defer func() {
			defer fmt.Println(i, 2)
		}()
	}

	for i := 0; i < 5; i++ {
		defer func() {
			j := i
			fmt.Println(j, 3)
			fmt.Println(j, &j, i, &i, 3)
		}()
	}

	for i := 0; i < 5; i++ {
		j := i
		defer fmt.Println(j, 4)
	}

	for i := 0; i < 5; i++ {
		j := i
		defer func() {
			fmt.Println(j, 5)
		}()
	}

	for i := 0; i < 5; i++ {
		defer func(j int) {
			fmt.Println(j, 6)
		}(i)
	}
}

func acumulator() func(int) int {
	var x int
	return func(delta int) int {
		// check memory address of x
		fmt.Printf("(%p,%+v) ", &x, x)
		// check value of x
		fmt.Printf("(%+v,%+v) ", &x, x)
		// check memory address of delta
		fmt.Printf("(%p,%+v) ", &delta, delta)
		// check value of delta
		fmt.Printf("(%+v,%+v) ", &delta, delta)
		x += delta
		return x
	}
}

func e3() {
	var err error
	defer func(err error) {
		fmt.Println("Deferred function executed, err =", err)
	}(err)
	err = errors.New("defer3 error")
	fmt.Println("Function e3 executed, err =", err)
}

func e2() {
	var err error
	defer func() {
		fmt.Println("Deferred function executed, err =", err)
		// The deferred function will print the value of err when it is executed.
		// If err is nil, it will print "Deferred function executed, err = <nil>".
		// If you want to see the effect of an error, you can set err to a value before the deferred function runs.
	}()
	err = errors.New("defer2 error")
	fmt.Println("Function e2 executed, err =", err)
	// In this case, the deferred function will print "Deferred function executed, err = defer2 error".
	// The deferred function captures the value of err at the time it is executed, which is after the return statement.
	// Note: The deferred function runs after the surrounding function returns, so any changes made to err in the deferred function
	// will not affect the return value of e2.
	// output will be "Function e2 executed, err = defer2 error" followed by "Deferred function executed, err = defer2 error".
}

func e1() {
	var err error
	defer fmt.Println("Deferred function executed, err =", err)
	// The deferred function will print the value of err when it is executed.
	// However, since err is nil at this point, it will print "Deferred function executed, err = <nil>".
	// If you want to see the effect of an error, you can set err to a value before the deferred function runs.
	err = errors.New("defer1 error")
	fmt.Println("Function e1 executed, err =", err)
}

func f1() int {
	t := 5
	defer func() {
		t += 5 // This will modify t, but the change will not be visible outside this function why ?
		// because defer captures the variable t by value, not by reference.
		// The value of t at the time of the defer is what gets used when the deferred function runs.
		// Hence, the original t remains unchanged outside this deferred function.
		fmt.Println("Deferred function executed, t =", t)
	}()
	return t // The return value is set before the deferred function runs
	// In this example, the deferred function modifies t, but since t is passed by value to the deferred function,
	// the change does not affect the original t in f1. The return value of f1 is still 5.
	// If you want to modify the original variable, you can use a pointer or return the modified value.
	// If you want to see the effect of the deferred function, you can return a pointer or use a global variable.
	// Note: The deferred function runs after the surrounding function returns, so any changes made to t in the deferred function
	// will not affect the return value of f1.
	// output will be 5 because the deferred function modifies a copy of t, not the original t.
}

func f2() (r int) {
	defer func(r int) {
		r += 5 // This modifies the named return value r
		fmt.Println("Deferred function executed, r =", r)
	}(r)
	return 1
	// output will be 1 because the deferred function modifies a copy of r, not the original r.
}

func f3() (r int) {
	defer func(r *int) {
		*r += 5 // This modifies the value pointed to by r
		fmt.Println("Deferred function executed, r =", *r)

	}(&r)
	return 1 // Here, we pass a pointer to r, so the deferred function can modify the original value of r.
	// This way, the deferred function can change the value of r, and the change will be reflected in the return value.
	// The deferred function modifies the value of r by dereferencing the pointer, allowing it to change the original variable.
	// This is a common pattern when you want to ensure that the deferred function can modify the return value.
	// Note: The deferred function runs after the surrounding function returns, so any changes made to r in the deferred function
	// will affect the return value of f3.
	// output will be 6 because the deferred function adds 5 to the original value of r (which is 1).
}
