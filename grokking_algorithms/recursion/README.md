### Recursion Recap

  - Recursion is when a function calls itself
  - Every recursive function has two cases: Base case and Recursive case
    - Example: 
      ```
      function foo(i) {
         if (i = 0){
          return true #<---base cases
        } else {
          foo( i - 1) # <--- recursive case
        }
      }  
      ```
  - A stack has two operations: push and pop
  - All function calls go onto the call stack
  - The call stack can get very large, which takes up a lot of memory

