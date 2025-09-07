# bazel-tui

[![Test](https://github.com/filmil/bazel-tui/actions/workflows/test.yml/badge.svg)](https://github.com/filmil/bazel-tui/actions/workflows/test.yml)

A terminal user interface (TUI) for running commands and viewing their output in a windowed environment.

This application is built with Go and uses the following libraries:
- [tview](https.github.com/rivo/tview)
- [winman](https.github.com/epiclabs-io/winman)

## Usage

To run the application, you can use the following command:

```bash
bazel run //bin/tui1 -- -cmdline "your command here"
```
