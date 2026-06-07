<!--
Sync Impact Report
Version change: template -> 1.0.0
Modified principles:
- Template Principle 1 -> I. Repository-First Truth
- Template Principle 2 -> II. Confirmed Autonomy
- Template Principle 3 -> III. Ritual-Driven Product Value
- Template Principle 4 -> IV. Traceable Operations
- Template Principle 5 -> V. Explicit Contracts
Added sections:
- Operational Constraints
- Delivery Workflow
Removed sections:
- None
Templates requiring updates:
- ✅ .specify/templates/plan-template.md
- ✅ .specify/templates/spec-template.md
- ✅ .specify/templates/tasks-template.md
- ✅ .specify/memory/constitution.md
- ⚠ pending .specify/templates/commands/*.md (directory absent; no updates applied)
Follow-up TODOs:
- None
-->
# Assistant Constitution

## Core Principles

### I. Repository-First Truth
All durable product state MUST live in the Obsidian-compatible repository:
sphere notes, goals, tasks, reflections, logs, schedules, and assistant
configuration. `runtime/` MAY store ephemeral execution artifacts, but it MUST
not become a source of truth. Every feature specification and implementation
plan MUST name which repository files are authoritative and how updates preserve
link integrity between day, month, year, and 5-year horizons. Rationale: the
product fails if Telegram feels useful but the repository stops being a reliable
memory and planning system.

### II. Confirmed Autonomy
Assistant actions MUST be classified before implementation as either
`auto-allowed`, `proposal-only`, or `requires-confirmation`. Any destructive or
strategic mutation, including deletions, archival, priority changes, goal
relinking, or creation of semantic links, MUST remain `proposal-only` or
`requires-confirmation` until the user explicitly approves it. Automatic actions
MUST be reversible from repository history and MUST be reported back to the user
in plain language. Rationale: the assistant is useful only while it reduces
overhead without silently steering the user's system.

### III. Ritual-Driven Product Value
Work MUST optimize for the recurring user rituals that define product value:
daily capture, morning check-in, evening review, weekly review, and monthly
planning. Every feature MUST state which ritual or planning horizon it improves,
what decision it helps the user make, and how it avoids bias toward work at the
expense of other life spheres. Features that do not strengthen a concrete
ritual, reduce loss of incoming tasks, or improve cross-sphere focus MUST be
treated as out of scope for near-term delivery. Rationale: the product succeeds
only if the user keeps returning to it every day and across longer review
cycles.

### IV. Traceable Operations
Every user request, assistant proposal, repository mutation, carry-over,
reprioritization, and system failure MUST be traceable through repository-backed
logs with at least 30 days of retention. Specifications and plans MUST define
what gets logged, which identifiers connect a Telegram interaction to file
changes, and how operators can inspect the decision path. State-changing
features MUST include verification for both the user-visible outcome and the
resulting log record. Rationale: a planning assistant without auditability
cannot be trusted or debugged.

### V. Explicit Contracts
Prompts, schemas, configuration, and orchestration rules MUST evolve through
explicit contracts rather than hidden behavior in code. New capabilities or
behavior changes MUST update the relevant spec artifacts, data schemas or file
formats, default configuration, and validation coverage before implementation is
considered complete. Ambiguity about input formats, note layout, scheduling, or
role responsibilities MUST be resolved in written artifacts rather than left to
runtime guesswork. Rationale: the product spans LLM behavior, repository
structure, and operator configuration, so implicit contracts decay quickly.

## Operational Constraints

The primary interface MUST remain Telegram, and durable outputs MUST remain
compatible with the repository structure described in `SPEC.md`. Manual editing
inside Obsidian MUST stay optional for daily operation. Integrations MAY enrich
context, but `v1` MUST not depend on YouTrack and MUST treat the calendar as the
main work-planning reference. Any feature that sends note content to an LLM MUST
minimize scope to the needed context and document the exposure boundary in the
specification or plan. Default schedules, sphere priorities, and logging
retention MUST remain configurable through repository-managed config files.

## Delivery Workflow

Feature specs MUST include: the target ritual or planning horizon, repository
artifacts affected, autonomy classification for every write path, measurable
success criteria, and acceptance scenarios for the primary flow and failure
paths. Implementation plans MUST fail the Constitution Check unless they define
authoritative files, logging behavior, config or schema changes, and validation
for user-visible outcomes plus repository mutations. Task lists MUST include the
work needed to update schemas, prompts, config, and logging when those contracts
change; tests are mandatory for state-changing or scheduling-critical flows.
Incremental delivery MUST prefer thin vertical slices that keep capture and
review workflows working end to end.

## Governance

This constitution overrides conflicting local conventions for specification,
planning, and implementation. Amendments MUST be recorded in
`.specify/memory/constitution.md`, include a semantic version decision, and
describe any required template or workflow propagation. Versioning policy:
MAJOR for incompatible principle or governance changes, MINOR for new principles
or materially expanded obligations, PATCH for clarifications that preserve the
existing meaning. Compliance review is mandatory in every spec, plan, task set,
and implementation review; constitution violations MUST be resolved by changing
the affected artifact or by explicitly amending the constitution, never by
silently ignoring the rule.

**Version**: 1.0.0 | **Ratified**: 2026-06-07 | **Last Amended**: 2026-06-07
