# dux (du eXtended)

dux (du eXtended) is a tool to display disk usage. Though original du command has many options, dux has no options.
It may show what you wanted to see (largest directory, largest suffix, top largest files).

## install

```bash
go get github.com/ktrueda/dux
```

## usage

```bash
> dux .
Target Directory: /Users/ktrueda/oss/dux
🐘 File Size Group By suffix
                : ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 14,843,444 B
        .sample :  19,260 B
            .go :  6,347 B
         .pprof :  4,210 B
            .md :  1,774 B
           .sum :  823 B
           .mod :  376 B
.code-workspace :  117 B

🦒 File Size Group By directory
/Users/ktrueda/oss/dux/.git : ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 11,991,096 B
     /Users/ktrueda/oss/dux : ▇▇▇▇▇▇▇▇▇▇▇▇ 2,879,670 B
 /Users/ktrueda/oss/dux/lib :  5,585 B

🦛 Top Large size file
0 /Users/ktrueda/oss/dux/dux 2,871,608 B
1 /Users/ktrueda/oss/dux/.git/objects/7b/4d023402b41ca075f0d3e56c096ba3acc18b5a 1,906,420 B
2 /Users/ktrueda/oss/dux/.git/objects/c9/0070705d6bc0c23b47dcaaefe4aa1a9e8e626a 1,888,076 B
3 /Users/ktrueda/oss/dux/.git/objects/cd/da8ff162c7c25ee7d86b1828463af2a5904b65 1,621,550 B
4 /Users/ktrueda/oss/dux/.git/objects/60/cb31a5f0870820e9378012fcb2d57d59d3a356 1,621,252 B
5 /Users/ktrueda/oss/dux/.git/objects/e6/c58f32129a42f2cd45ab8ba43beb5be33fa36e 1,621,095 B
6 /Users/ktrueda/oss/dux/.git/objects/0f/623f7467415fe286c28110f3f18c19fd5d0bd3 1,620,599 B
7 /Users/ktrueda/oss/dux/.git/objects/d7/76ca0ad3eb2a3e42dd469bd2c40b30d73f3b16 1,616,861 B
8 /Users/ktrueda/oss/dux/.git/logs/HEAD 5,997 B
9 /Users/ktrueda/oss/dux/.git/objects/65/95eb89b394e228b7505c08ceb5b4e7c57df642 5,016 B
```