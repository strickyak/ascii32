# ascii32
## A language named ` `.

Just started.  This much works:

```
echo ' 3 8 + ! ' | go run ascii32.go

2020/09/28 01:38:29 Step[0]: {1 1   3}
2020/09/28 01:38:29 Step[1]: {1 2 3 1}
2020/09/28 01:38:29 Step[2]: {1 3   3}
2020/09/28 01:38:29 Step[3]: {1 4 8 1}
2020/09/28 01:38:29 Step[4]: {1 5   3}
2020/09/28 01:38:29 Step[5]: {1 6 + 2}
2020/09/28 01:38:29 Step[6]: {1 7   3}
2020/09/28 01:38:29 Step[7]: {1 8 ! 2}
<1><1>2020/09/28 01:38:29 Step[8]: {1 9   3}
2020/09/28 01:38:29 Step[9]: {1 10  0}
2020/09/28 01:38:29 Final Stack: {[]}
```

where `!` prints the top of the stack (like `.` in Forth).

Notice the output in angle brackets: `<1><1>`
