# dux (du eXtended)

dux (du eXtended) is a tool to display disk usage. Though original du command has many options, dux has no options.
It may show what you want to see.

## install

```bash
go get github.com/ktrueda/dux
```

## usage

```bash
> dux .
Target Directory: /Users/ktrueda/oss/dux
File Size Group By suffix
                : ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 7,768,143 B
        .sample :  19,260 B
            .go :  5,909 B
            .md :  1,407 B
           .sum :  602 B
           .mod :  292 B
.code-workspace :  117 B

File Size Group By directory
/Users/ktrueda/oss/dux/.git : ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 4,908,813 B
     /Users/ktrueda/oss/dux : ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 2,881,340 B
 /Users/ktrueda/oss/dux/lib :  5,577 B

Top Large size file
0 /Users/ktrueda/oss/dux/dux 2,853,112 B
1 /Users/ktrueda/oss/dux/.git/objects/cd/da8ff162c7c25ee7d86b1828463af2a5904b65 1,621,550 B
2 /Users/ktrueda/oss/dux/.git/objects/60/cb31a5f0870820e9378012fcb2d57d59d3a356 1,621,252 B
3 /Users/ktrueda/oss/dux/.git/objects/0f/623f7467415fe286c28110f3f18c19fd5d0bd3 1,620,599 B
4 /Users/ktrueda/oss/dux/log 25,478 B
5 /Users/ktrueda/oss/dux/.git/hooks/pre-rebase.sample 4,898 B
6 /Users/ktrueda/oss/dux/lib/util/util.go 3,716 B
7 /Users/ktrueda/oss/dux/.git/hooks/update.sample 3,610 B
8 /Users/ktrueda/oss/dux/.git/hooks/fsmonitor-watchman.sample 3,327 B
9 /Users/ktrueda/oss/dux/.git/logs/HEAD 1,697 B
```