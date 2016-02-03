From [`golint`](https://github.com/golang/lint/#purpose):
> We will not be adding pragmas or other knobs to suppress specific warnings, so do not expect or require code to be completely "lint-free".

`golint-free` is exactly like `golint` (it wraps golint) with the ability to suppress warning messages.

#### Install

`go get -u github.com/hnry/golint-free`

make sure you have your `$GOPATH` set and it is in your `$PATH`

edit `$HOME/.golint-free`, each line is a new entry to match against to suppress warnings from golint

Example warning:

    testfile:25:6: don't use underscores in Go names; func Test_input should be TestInput

    testfile:40:6: exported type MockWriter should have comment or be unexported

If I wanted both to be suppressed I would edit `.golint-free` to look like this:

    don't use underscores in Go names
    should have comment or be unexported

And it'll stop warnings that match any lines you have.
