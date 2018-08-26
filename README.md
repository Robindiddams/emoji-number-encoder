# the emoji number encoder!

who remembers phone numbers? I know I don't! now you can turn your phone number into just 4 easy-to-remember emojis! wow!

### what is this?
its a fun way to encode a phone number into 4 emojis!
try it out!
```bash
$ go run main.go 4085554607
😆🚲💃⤵️
$ go run main.go -d 😆🚲💃⤵️
4085554607
```

### how does it work?
in essance, its a base512 encoder, it breaks a number of max 36 bits down into four 9 bit chunks, each chunk is converted from binary to a number that references an emoji in a map

### can it do text?
No! It cant and it wont! If you want an emoji encoder that does text (and everything else 😅) checkout [keith-turner/ecoji](https://github.com/keith-turner/ecoji)