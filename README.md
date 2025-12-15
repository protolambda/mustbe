# mustbe

It must be... yet another Go test assertion library.

What makes this one unique:
- Assertion interface:
  - Plug in custom assertions
  - Customize where assertion errors go
  - Check with `Context`, for control over long-running assertions, or extra assertion attributes
- Type-safety: generics to avoid the `any` type
- No duplication or confusion between `Error/assert` and `FailNow/require`.
  Assertions are objects, `Must` is immediate, and results are handled by the caller.
- Readable:
  - `t.Must(be.Positive(v))`: "must be X" shorthand
  - `t.Must(assertion.Equal{Expected: a, Got: b})`: optional named args
- Small improvements over other assertion libraries:
  - `InDelta` without precision-loss or unnecessary float conversion.
  - `Eventually` that handles a panicking inner function.
  - `Equal` (shallow) and `DeepEqual` separated.
  - No unnecessary out-there assertions like YAML-encoded-comparison.
    You can bring your own `Assertion` implementations when really needed.
- No external dependencies

## Usage

See [example test](./must_example_test.go).

## License

MIT License, see [LICENSE](./LICENSE) file.
