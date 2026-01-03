# Vibe coding Golang interfaces one at a time

Afraid of using Cursor to manipulate your entire codebase? Want to introduce just a bit of vibes into your existing Go project without giving AI everything you have?

`govibeimpl` is here for you! It's a contract-first code generation tool that only vibes components, modules, and implementations to interfaces that you ask for, and allows its result to be immediately plugged into an existing project.

## How to use

Using this tool is very simple.

1. Install this tool via `go install github.com/yuedongze/govibeimpl/cmd/govibeimpl@latest`
2. Define the interface (e.g. `URLDownloader`) that you want `govibeimpl` to generate implementations for.
3. On top of the interfaces, add the following go generate directive (e.g. `//go:generate govibeimpl -name URLDownloader`)
4. Vibe generate all the code via `go generate ./...`, how simple is that!
5. Run `go mod tidy` to install any external packages introduced by the vibed code.
6. Profit?

## Examples

Check `examples` directory for working examples.
