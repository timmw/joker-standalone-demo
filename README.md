This is a proof of concept for building a [joker](https://joker-lang.org)
program as a standalone binary. It's based on this
[gist](https://gist.github.com/pyr/bf63ac3a014517cafa371e96fdddfea0) by
[pyr](https://github.com/pyr).

---

Run `./build.sh`

After a 10-20 seconds or so you should see:

"Hello from the other namespace"

which proves the standalone binary has been created and works!

The changes to joker are minimal:
https://github.com/timmw/joker/compare/master...standalone-binary
