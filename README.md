From [`golint`](https://github.com/golang/lint/#purpose):
> We will not be adding pragmas or other knobs to suppress specific warnings, so do not expect or require code to be completely "lint-free".

`golint-free` is exactly like `golint` (it wraps golint) with the ability to suppress specific warning messages.

#### Install

`go get -u github.com/hnry/golint-free`

create `$HOME/.golint-free` to be a JSON with 2 fields:

- golint
  > String for file path to golint (must be absolute path, no ENV vars)

- ignore
  > An array of strings used for matching golint warnings that you want ignored

Example warning:

    testfile:25:6: don't use underscores in Go names; func Test_input should be TestInput
    testfile:40:6: exported type MockWriter should have comment or be unexported

If I wanted both to be suppressed I would edit `.golint-free` to look like this:

    {
      golint: '/home/hnry/go/bin/golint',
      ignore [
        "don't use underscores in Go names",
        "should have comment or be unexported"
      ]
    }

And it'll suppress warnings that match any lines you have.

#### Using golint-free as golint
Normally you should just use the command `golint-free` over `golint` but...
There are some other tools, plugins, addons, etc. that use golint but don't provide the option to change the path (for you to point to golint-free instead).

Personally I just renamed golint-free to golint and move golint to say 'golint-orig', then edit .golint-free to point to golint-orig.
