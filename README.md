Run `./build.sh`

It will output

`<joker.core>:3513:15: Eval error: open base64.joke: file does not exist`

along with a stacktrace but after that you should see:

"Hello from the other namespace"

which proves the standalone binary has been created and works!

I've probably made a mistake with my changes to support embedding to cause the
'file does not exist' error.

The changes to joker are minimal:
https://github.com/timmw/joker/commit/d757aa0237537dc030cd30501d99db854c8c6f21
