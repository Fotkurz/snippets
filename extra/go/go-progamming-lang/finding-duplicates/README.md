## Read duplicates

Read duplicate from Addison-Wesley Go Programming Lang book.

## How-To

Exec `go run main.go` passing in the files as args:

```bash
go run main.go input.txt
```

Outputs:

```
3       banana
2       grape
2       mango
4       apple
2       cherry
2       pear
```

Or using STDIN:

```bash
go run main.go
apple
orange
orange
strawberry
2       orange
```

PS: Press Ctrl+D on linux toend stdin