# gogitversion
Go package to assist in embedding git tag based versioning into an application

<img src="gogitit.png" width="100%" alt="gogitit logo created at https://pablo.buffer.com/" />

## Project Health Status

[![CircleCI](https://circleci.com/gh/prognosai/gogitversion.svg?style=shield)](https://circleci.com/gh/prognosai/gogitversion)

## Quickstart

```sh
> go get https://github.com/prognosai/gogitversion
```

```go
package main

import (
	"fmt"
	"github.com/prognosai/gogitversion"
)

func main() {
    ...
    fmt.Printf("version: %s\n", gogitversion.Get())
    ...
}
```

```sh
go build -ldflags "-X github.com/prognosai/gogitversion.version=0.0.1"
```

```sh
> ./myapp
version: 0.0.1
```

## Git Describe Versioning

When no embedded version is provided gogitversion falls back upon a versioning
technique using `git describe --tags --dirty`.

For more details on `git describe` check out: https://git-scm.com/docs/git-describe

This versioning technique will reference the most recent tag set using `git tag`.

For example:

```sh
# When the most recent tag matches the most recent commit with a clean working tree
> git describe --tags --dirty
v2018.07.03.115146

# When the most recent tag matches the most recent commit with a dirty working tree
> git describe --tags --dirty
v2018.07.03.115146-dirty

# When the most recent tag is 4 commits behind the most recent commit with a clean working tree
> git describe --tags --dirty
v2018.07.03.115146-4-gc3a421b

# When the most recent tag is 4 commits behind the most recent commit with a dirty working tree
> git describe --tags --dirty
v2018.07.03.115146-7-gc3a421b-dirty
```

This technique allows you to get very detailed information about the current code being executed within the application.

## Expected Usage

During local development fall back upon `git describe --tags --dirty` to provide
versioning information... by doing nothing except using `git tag`'s and go as you normally would.

```sh
> go build
```

When creating a release build embed the version as follows:

```sh
> go build -ldflags "-X github.com/prognosai/gogitversion.version=$(git describe --tags --dirty)"
```