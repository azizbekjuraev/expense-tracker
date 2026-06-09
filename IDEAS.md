# Expense Tracker — Learn Go by Leveling Up This Project

A roadmap of features and refactors, ordered so each one teaches you a new Go concept. Pick any in a tier — they're independent. Have fun, break things, then fix them.

---

## Tier 1 — Clean up what you already have (Go fundamentals)

Small wins that teach you idiomatic Go without changing the project's shape.

- **Split `main.go` into packages.** Move types into `internal/expense/`, commands into `cmd/`, storage into `internal/storage/`. Learn: package layout, exported vs unexported identifiers, import paths.
- **Replace `log.Fatalf` everywhere with returned `error` values.** Bubble errors up to `main`. Learn: idiomatic error handling, `fmt.Errorf("...: %w", err)`, error wrapping.
- **Use `errors.Is` / `errors.As`.** Detect "file not found" instead of crashing on first run (`os.IsNotExist` → `errors.Is(err, os.ErrNotExist)`). Learn: sentinel errors, error trees.
- **Make `Amount` a `float64` or a dedicated `Money` type (cents as `int64`).** Real expenses have decimals. Learn: custom types, methods on types, why floats are bad for money.
- **Add a `Category` field** (food, transport, etc.) with validation. Learn: `iota` enums, `Stringer` interface, JSON marshaling of custom types.
- **Validate inputs.** Reject negative amounts, empty descriptions, future dates. Learn: early returns, guard clauses.
- **Atomic file writes.** Write to `expenses.json.tmp` then `os.Rename`. Prevents corruption if you Ctrl-C mid-write. Learn: `defer`, file primitives, race-condition thinking.
- **Run `golangci-lint` and fix everything.** Learn: idiomatic style, common pitfalls. Install: `brew install golangci-lint`.

## Tier 2 — Make it feel like a real tool (stdlib deep-dive)

- **Filter `list` by date range, category, min/max amount.** Flags: `--from`, `--to`, `--category`, `--min`. Learn: `time.Parse`, slice filtering, predicate functions.
- **Sort `list` by any column.** `--sort=amount --desc`. Learn: `sort.Slice`, comparator functions, closures.
- **Monthly / weekly summary.** `summary --month=2026-06`. Group by category too. Learn: `map[string]int`, time arithmetic, grouping patterns.
- **Budget command.** Set `budget set --monthly=50000`. Warn on `add` if you'd exceed it. Learn: persistent config, multiple JSON files, separation of concerns.
- **Export to CSV.** `export --format=csv > out.csv`. Learn: `encoding/csv`, `io.Writer` interface (so it works with files or stdout).
- **Import from CSV.** Mirror of above. Learn: streaming readers, line-by-line parsing.
- **Recurring expenses.** "Netflix, every 30 days." On `list`, auto-materialize missed occurrences. Learn: time math, state machines.
- **Undo last action.** Keep a tiny history stack in a sidecar file. Learn: stack data structure, command pattern.
- **Search.** `search "coffee"` with fuzzy match (try `github.com/sahilm/fuzzy`). Learn: third-party libs, ranking algorithms.

## Tier 3 — Concurrency (the fun Go superpower)

The CLI is too small to need concurrency, so invent reasons.

- **Watch mode.** `watch` command tails the file and re-renders the list when it changes. Learn: `fsnotify`, goroutines, `select`, `context.Context` for cancellation.
- **Concurrent currency conversion.** On `list --in=USD`, fan-out HTTP calls to an exchange-rate API for each unique currency. Learn: goroutines, `sync.WaitGroup`, `errgroup`, channels.
- **Parallel CSV import.** Import 100k rows by sharding across N workers. Learn: worker pools, bounded concurrency, channel patterns.
- **Background sync daemon.** `daemon start` runs in the background and periodically syncs to a remote endpoint. Learn: signal handling (`os/signal`), graceful shutdown, tickers.

## Tier 4 — Storage that isn't a JSON file

Each step teaches a different abstraction.

- **Define a `Storage` interface.** `Save`, `Load`, `Delete`, `Update`. Refactor JSON code behind it. Learn: interfaces, dependency injection, why interfaces belong on the consumer side in Go.
- **Add a SQLite backend.** Use `modernc.org/sqlite` (pure Go, no CGo). Same interface as JSON. Pick via `--storage=sqlite` flag. Learn: `database/sql`, prepared statements, migrations.
- **Add a Postgres backend.** Now you have three. Learn: connection pooling, `sql.DB` vs `sql.Conn`, env-based config.
- **Add caching layer.** A `cacheStorage` that wraps any storage. Learn: decorator pattern, in-memory caching, `sync.RWMutex`.

## Tier 5 — Testing (Go's killer feature)

- **Unit-test `loadExpenses` and `encodeExpenses`.** Use `t.TempDir()`. Learn: `testing` package, table-driven tests.
- **Test the CLI commands end-to-end.** Use `cobra`'s `SetArgs` and capture stdout. Learn: black-box CLI testing.
- **Add benchmarks.** `BenchmarkAdd1000`. Learn: `testing.B`, `go test -bench`.
- **Add fuzz tests.** Fuzz the CSV importer with `go test -fuzz`. Learn: Go 1.18+ fuzzing.
- **Mocks via interfaces.** Mock a `Clock` interface so `now := time.Now()` is testable. Learn: why interfaces make testing trivial.
- **Hit 80% coverage.** `go test -coverprofile=c.out && go tool cover -html=c.out`. Learn: coverage tooling.

## Tier 6 — Bring it online (Go for the web)

- **`serve` command.** Spin up an HTTP API exposing the same operations. `POST /expenses`, `GET /expenses`, etc. Learn: `net/http`, `http.ServeMux` (Go 1.22+ has path patterns), JSON request/response.
- **Add middleware.** Logging, request ID, panic recovery. Learn: middleware pattern as `func(http.Handler) http.Handler`.
- **Basic auth, then JWT.** Learn: HTTP auth, third-party JWT lib, why auth is hard.
- **WebSocket live updates.** Push new expenses to connected clients. Learn: `gorilla/websocket` or `nhooyr.io/websocket`, broadcast patterns.
- **TUI frontend.** Use `bubbletea` for an interactive terminal UI. Learn: Elm Architecture in Go, channels for event loops.
- **gRPC API.** Expose the same service over gRPC. Learn: protobuf, code generation, streaming RPCs.

## Tier 7 — Production polish

- **Structured logging with `log/slog`** (stdlib since Go 1.21). JSON logs in prod, text in dev. Learn: structured logging, log levels, context propagation.
- **Configuration with `viper` or just stdlib.** Support config file, env vars, flags — in that precedence. Learn: configuration layering.
- **Build for every platform.** `Makefile` with `GOOS=linux GOARCH=arm64 go build`. Learn: cross-compilation (Go's superpower).
- **Reproducible Docker image.** Multi-stage build, `FROM scratch` final image (~10MB). Learn: static linking, distroless images.
- **GitHub Actions CI.** Run tests, lint, build matrix across OSes. Learn: CI YAML, caching Go modules.
- **Release with `goreleaser`.** Auto-build binaries for every platform on git tag. Learn: release automation.
- **Embed assets with `embed`.** Embed a default config or HTML for the web UI directly into the binary. Learn: `//go:embed`.

## Tier 8 — Wild experiments

Things that aren't useful but are fun and teach weird corners of Go.

- **Plugin system with `plugin` package** OR **WASM plugins with `wazero`.** Let users write custom report generators. Learn: dynamic loading.
- **Generics-powered query DSL.** `Query[Expense]().Where(...).OrderBy(...)`. Learn: type parameters (Go 1.18+).
- **Custom JSON marshaler** that pretty-prints amounts as `$12.34`. Learn: `MarshalJSON`/`UnmarshalJSON`.
- **Make it a Telegram bot.** Forward "lunch 12" messages into expenses. Learn: long polling, bot APIs.
- **AI categorization.** Call the Claude API to auto-categorize from description. Learn: HTTP clients, streaming responses.
- **Encrypt the storage file** with a passphrase. Learn: `crypto/aes`, `crypto/rand`, key derivation.
- **Write your own minimal cobra-like CLI lib** and replace cobra. Learn: reflection, what your dependencies actually do.

---

## Suggested order if you want a path

1. Tier 1 entirely → idiomatic Go muscle memory.
2. Pick 3 from Tier 2 → stdlib confidence.
3. All of Tier 5 → testing is the unlock for everything after.
4. Tier 4 interfaces + SQLite → "now I think in Go."
5. Tier 6 HTTP server → most Go jobs want this.
6. Cherry-pick from Tier 3, 7, 8 for taste.

Each tier item should take 1–4 hours. If one drags past a day, drop it and try the next — momentum matters more than completeness when you're learning.
