# Basic go bazel project [![Test status](https://github.com/filmil/bazel-go-basic/workflows/Test/badge.svg)](https://github.com/filmil/bazel-go-basic/workflows/Test/badge.svg)

This is an empty go project that you can use for spinning off your own projects
that use `bazel` as a build system, and the go toolchain.  Of course you can add
other toolchains you need as your project grows.

- `bazel` pegged to 7.0.2
- `.bazelrc` added, with support for `user.bazelrc`
- Go toolchain set up with the latest go version at the moment (1.22.1).
- Added github workflow.
- LICENSE is present (Apache 2.0).
- README is present.
- `bazel build //...` works
- `bazel test //...` passes
- Workspace uses bzlmod (i.e. `MODULE.bazel` instead of `WORKSPACE`).

- Use `bazel run //:gazelle` to tidy up your build files.
- Use `bazel run //:buildifier` to format things.
