# Palindrome Number

Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

`Time complexity : O(log{10}(n))`

`Space complexity : O(1)`

```Bash
$ go test -v -cover ./...
=== RUN   TestCase1
Expected: true, Output: true
--- PASS: TestCase1 (0.00s)
=== RUN   TestCase2
Expected: false, Output: false
--- PASS: TestCase2 (0.00s)
PASS
coverage: 72.7% of statements
ok      github.com/mxssl/LeetCodePalindromeNumberGolang       0.175s  coverage: 72.7% of statements
```
