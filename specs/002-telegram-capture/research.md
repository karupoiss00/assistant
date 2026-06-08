# Research: Telegram Capture

## Decision 1: Use a daily inbox file as the primary durable capture buffer

- **Decision**: Persist all inputs of the current day into one repository-backed
  daily inbox file, with audit logs stored separately.
- **Rationale**: This matches the clarification decision, keeps capture visible
  from Obsidian, and avoids prematurely inventing a richer classification
  scheme in Stage 2.
- **Alternatives considered**:
  - One file per input: rejected because it adds filesystem noise and was
    explicitly declined during clarification.
  - A single long-lived inbox file: rejected because daily grouping is easier to
    review and aligns better with later daily rituals.

## Decision 2: Confirm voice capture only after transcription and basic understanding

- **Decision**: Treat user-visible success for voice input as complete only
  after transcription and basic understanding are available.
- **Rationale**: The user explicitly wants the first bot reply to already show
  understanding of what was said, even if it adds delay.
- **Alternatives considered**:
  - Immediate transport-level confirmation before transcription: rejected
    because it gives a weaker UX signal.
  - Waiting for full downstream classification: rejected because Stage 2 is
    about reliable capture, not full semantic routing.

## Decision 3: Keep original voice files temporary, not durable

- **Decision**: Download and hold voice files only long enough to transcribe
  them, then delete them after success or failure handling.
- **Rationale**: This minimizes storage cost, reduces privacy exposure, and
  keeps durable state centered on transcript, metadata, and audit records.
- **Alternatives considered**:
  - Store all original audio durably: rejected because it was explicitly ruled
    out and is unnecessary for Stage 2 goals.
  - Never store audio even temporarily: rejected because transcription requires
    a transient processing artifact.

## Decision 4: Avoid semantic deduplication in Stage 2

- **Decision**: Save Telegram inputs as separate capture facts and defer complex
  semantic deduplication until real cases justify it.
- **Rationale**: Premature deduplication risks deleting or hiding real user
  intent, and the user explicitly prefers to solve that problem later.
- **Alternatives considered**:
  - Aggressive heuristic deduplication at intake: rejected because the rules
    would be speculative.
  - Ignoring repeat indicators entirely: rejected because transport-level repeats
    still need audit visibility.

## Decision 5: Use repository-backed logs as the canonical audit surface

- **Decision**: All capture events, retries, failures, and cleanup operations
  must be reflected in repository-backed logs correlated with Telegram message
  identifiers.
- **Rationale**: This follows the constitution and Stage 1 repository-first
  contract while keeping Telegram from becoming the only review interface.
- **Alternatives considered**:
  - Rely only on bot runtime logs: rejected because that would move durable
    history outside the repository.
  - Store only inbox content without event logs: rejected because retries and
    failures would be hard to audit.

## Decision 6: Fix the Stage 2 runtime as a Go webhook service

- **Decision**: Implement Stage 2 as a long-running Go webhook service rather
  than a polling worker or an always-on Codex runtime.
- **Rationale**: The feature spec already requires a webhook-based minimum
  deployable runtime, and fixing that choice now removes ambiguity around
  delivery semantics, retries, deployment shape, and validation strategy.
- **Alternatives considered**:
  - Polling worker: rejected because it conflicts with the accepted Stage 2
    webhook requirement and would alter retry and latency behavior.
  - Always-on Codex runtime: rejected because Stage 2 only needs a narrow
    capture service and should not depend on a broader orchestration layer.

## Decision 7: Use explicit workspace paths and append-only capture writes

- **Decision**: The runtime writes to explicit relative paths inside the
  production workspace, using append-only inbox and audit records as the
  durable capture surface.
- **Rationale**: This keeps the capture layer Obsidian-readable, preserves
  traceability, and removes ambiguity about which files change for each input.
- **Alternatives considered**:
  - Derive paths heuristically from local examples: rejected because `.workspace/`
    is only a reference example.
  - Store transcript or retry state only in runtime memory: rejected because it
    would violate the repository-first durability rule.
