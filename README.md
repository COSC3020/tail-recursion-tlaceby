[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/bHkMPWBv)

# Tail Recursion

In the lectures, we've seen code for a recursive implementation for computing
the Fibonacci numbers. Make this implementation tail-recursive. I have not
provided a template; depending on how you choose to implement the function, it
will have a different signature.

I have not provided any test code, but you can base yours on test code from
other exercises. Your tests must check the correctness of the result of running
the function and run automatically when you commit through a GitHub action.

Hint: It may help to have a look at the iterative dynamic programming
implementation. What changes between iterations of the loop?

## Runtime Analysis

Is the asymptotic complexity of tail-recursive Fibonacci different from the
non-tail-recursive version? Why, or why not, and what is the complexity
(worst-case $\Theta$)? Add your answer, including your reasoning, to this
markdown file.

**Solution**

The asymptotic complexity of **tail-recursive-fib** versus a non tail recursive version would remain the same. We will still make the same number of recursive calls. The only diference will be how the runtime may optimize those calls. However the overall complexity does not change as the number of calls and work done is identical.

**Complexity**

The overall worst-case $\Theta$ time complexity of the tail-recursive will be the same as that of an iterative implimentation. That is $\Theta(n)$.

The function `fib(n, a, b)` takes three parameters: `n`, `a`, and `b`.
The base cases are handled when n is 0 or 1, and the function returns a or b, respectively. These base cases have a constant time complexity of $\Theta(1)$.

For values of $n > 1$, the function recursively calls itself with $n - 1$, `b`, and $a + b$ as arguments.
The recursive call `fib(n-1, b, a+b)` decrements n by 1 and updates the values of a and b for the next recursive call.

Each call perform a constant amount of work, which includes the arithmetic operations and the function call itself.
The recursion continues until the base cases are reached, and the function returns the computed result.
The number of recursive calls made by the function is directly proportional to the value of n. For each value of n, the function makes n recursive calls before reaching the base cases.

This is similar to how a iterative version works. We will do a constant amount of work for each call and that call happens $n$ times. This in contrast to other recursive implimentations which repeat work. This method will not repeat any work other than the base case checks which are trivial.

Therefore the worst-case complexity of this implimentation is $\Theta(n)$.
