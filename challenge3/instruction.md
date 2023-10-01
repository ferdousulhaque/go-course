## Challenge: Working with Flow Control
This challenge has us reading a local file to count 

### Acceptance Criteria

1. Use standard library `os` to read the provided data file and panic if there's an error. BONUS: use `os.Args` to get the filename to read
2. Handle panics with a deferred recovery
3. Use a map to track the count of words, letters, numbers and symbols
4. Use `for-range` loops
5. Use conditionals