# Quickstart: Workspace Foundation Validation

## Purpose

Use this guide to validate that the workspace contract is complete and that a
future implementation can normalize an existing Obsidian workspace safely.

## Prerequisites

- Feature branch `001-workspace-foundation`
- Review access to the reference workspace at `.workspace/`
- Access to the contract documents in [contracts/](./contracts/)
- Review access to [data-model.md](./data-model.md) and [plan.md](./plan.md)

## Scenario 1: Verify the documented `v1` filesystem contract

1. Open [contracts/workspace-filesystem-contract.md](./contracts/workspace-filesystem-contract.md).
2. Compare its required sphere list and shared artifacts with the actual
   reference tree under `.workspace/`.
3. Confirm that all five required spheres and `Цели на 5 лет.md` are covered.
4. Confirm the template sphere is documented as a reference pattern rather than
   a required production path.
5. Confirm the contract documents that required artifacts remain manually
   readable and editable through Obsidian.

**Expected outcome**:
- The contract covers every required `v1` sphere artifact and shared artifact.
- The production workspace path is explicitly configurable.
- Required artifacts remain operator-editable without depending on Telegram.

## Scenario 2: Validate safe normalization boundaries

1. Open [contracts/normalization-policy.md](./contracts/normalization-policy.md).
2. Review the `auto-allowed`, `proposal-only`, and
   `requires-confirmation` sections.
3. Walk through a hypothetical workspace where `Семья/Артефакты/` is missing
   but `Семья/Цели.md` already contains user-authored content.
4. Confirm the policy allows creating the missing directory but forbids
   rewriting the existing note automatically.
5. Walk through a hypothetical workspace with customized `.obsidian/workspace.json`.
6. Confirm the policy requires confirmation before any replacement or rewrite of
   `.obsidian/` files.
7. Walk through a hypothetical workspace that contains an extra custom sphere.
8. Confirm the policy preserves that sphere and does not delete or rename it
   automatically.

**Expected outcome**:
- Structural additions can happen safely without silent semantic changes.
- `.obsidian/` remains preserved unless the user explicitly approves changes.
- Additional user-defined spheres remain preserved by default.

## Scenario 3: Verify repository-backed source of truth and auditability

1. Open [contracts/workspace-config-contract.md](./contracts/workspace-config-contract.md).
2. Confirm the production workspace root is sourced from repository-managed
   configuration.
3. Open [contracts/normalization-policy.md](./contracts/normalization-policy.md)
   and inspect the normalization report contract.
4. Confirm that every executed or proposed action maps to a repository-backed
   audit record with at least 30 days of retention.
5. Cross-check the paths and retention statements against [config/README.md](/Users/rostislav.glizerin/projects/assistant/config/README.md) and [logs/README.md](/Users/rostislav.glizerin/projects/assistant/logs/README.md).

**Expected outcome**:
- No durable state depends on `runtime/`.
- The plan and contracts make repository logs and config inspectable without
  reading implementation code.

## Scenario 4: Validate planning-horizon linkage coverage

1. Open [contracts/workspace-filesystem-contract.md](./contracts/workspace-filesystem-contract.md).
2. Check that sphere notes and `Цели на 5 лет.md` are documented as the anchor
   points for the chain `Операционная задача -> Цель месяца -> Годовая цель ->
   Цель на 5 лет`.
3. Open [data-model.md](./data-model.md) and confirm that `Sphere`,
   `Shared Artifact`, `Planning Chain Anchor`, and `Normalization Action`
   support this linkage without forcing undocumented hidden state.

**Expected outcome**:
- The feature documents the required planning chain even though later stages
  will define the full linking mechanics.

## Scenario 5: Verify manual Obsidian editability

1. Open [contracts/workspace-filesystem-contract.md](./contracts/workspace-filesystem-contract.md).
2. Confirm the contract explicitly states that required artifacts must remain
   readable and editable through Obsidian.
3. Open [config/workspace.md](/Users/rostislav.glizerin/projects/assistant/config/workspace.md) and confirm the same rule is summarized for operators.
4. Confirm no contract requires Telegram as the only interface for durable
   edits.

**Expected outcome**:
- The workspace contract remains compatible with manual Obsidian editing.
