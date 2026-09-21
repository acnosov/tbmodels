# tbmodels

Shared Go library (module `github.com/acnosov/tbmodels`): msgp message types + NATS subjects for the mollybot workspace. Library only — no binary, no main.go.

## Files

- `models.go`, `surebet.go` — msgp structs; they ARE the NATS wire format shared by all services — changing fields/msg tags is a wire-compat change.
- `routes.go` — NATS subjects, naming `Producer.EventType`.
- `offers.go` — market-type constants + `OfferMap`.

## Codegen

- After editing structs in `models.go`/`surebet.go`, run `go generate ./...` (requires msgp binary).
- Never hand-edit `*_gen.go` / `*_gen_test.go` — msgp-generated.

## Verify (what CI runs)

- `go mod tidy -diff && go build ./... && go test -count=1 ./... && golangci-lint run`.
- `task check` = tidy-diff + lint + test-coverage + actionlint + typos (report-only — fix real typos or ignore via `.typos.toml`); `task run`/`build`/`generate` are broken template leftovers (no main.go; generate is watch-mode).

## Rules

- golangci-lint enables ALL linters (disable list in `.golangci.yaml`), gofumpt formatting — lint before committing; prek hooks also run go test + govulncheck + typos + gitleaks on commit.
- Conventional commits only: lowercase subject, ≤72 chars (`committed.toml`).
- Never tag manually — push to main, cocogitto auto-bumps and goreleaser releases the tag.
- Sibling services import this module via the parent `go.work` (no replace directives) — model changes ripple to every service; see `../AGENTS.md` for the workspace map.
- Ignore `_temp/` and `todo.md` (scratch).

## Skills

Before any Go coding, review, debugging, troubleshooting, or setup task,
load the `samber/cc-skills-golang@golang-how-to` skill first — it routes
to whichever other Go skills the task needs.
