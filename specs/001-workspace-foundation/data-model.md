# Data Model: Workspace Foundation

## Workspace Root

- **Description**: The authoritative root of the production Obsidian-compatible
  repository used by the assistant.
- **Fields**:
  - `path`: repository-relative or configured absolute path to the workspace
    root
  - `is_reference_example`: boolean flag indicating whether the root is the
    local `.workspace/` example
  - `obsidian_dir_present`: whether `.obsidian/` exists
  - `log_location`: path for repository-backed assistant logs
  - `config_location`: path for repository-backed assistant config
- **Validation Rules**:
  - `path` must resolve to a Git-managed Obsidian workspace root
  - `is_reference_example=true` must never imply production write authority
  - `log_location` and `config_location` must be inside repository-managed
    paths, not `runtime/`

## Sphere

- **Description**: A required or user-defined top-level life domain inside the
  workspace.
- **Fields**:
  - `name`: sphere directory name
  - `is_required_v1`: whether the sphere is one of `Работа`, `Дом`, `Авто`,
    `Здоровье`, `Семья`
  - `goals_note_path`: path to `Цели.md`
  - `tactical_goals_note_path`: path to `Тактические цели.md`
  - `operational_tasks_note_path`: path to `Операционные задачи.md`
  - `artifacts_dir_path`: path to `Артефакты/`
  - `custom_links`: optional links to monthly, annual, and 5-year planning
    anchors
  - `normalization_state`: `complete`, `missing_required_elements`, or
    `nonstandard_layout`
- **Validation Rules**:
  - Required `v1` spheres must expose all four mandatory artifacts
  - User-defined spheres may exist without being deleted or renamed
  - Missing required elements may be added automatically; existing content may
    not be overwritten automatically

## Planning Chain Anchor

- **Description**: Explicit contract anchor connecting operational work to
  monthly, annual, and 5-year horizons.
- **Fields**:
  - `source_note_path`: current note path where the link starts
  - `source_level`: `operational_task`, `monthly_goal`, `annual_goal`,
    `five_year_goal`
  - `target_note_path`: next note path in the planning chain
  - `link_status`: `documented`, `missing`, `proposed`
- **Validation Rules**:
  - The chain `Операционная задача -> Цель месяца -> Годовая цель -> Цель на 5
    лет` must have documented anchor points in the contract
  - Later stages may refine syntax, but the contract may not omit the levels
  - Missing chain anchors may be reported, but semantic relinking remains
    proposal-only unless explicitly confirmed

## Shared Artifact

- **Description**: A workspace-level note or directory that applies across
  spheres.
- **Fields**:
  - `name`: artifact name
  - `path`: repository path
  - `artifact_type`: `strategic_note`, `template`, `config_anchor`, or
    `log_anchor`
  - `is_required_v1`: whether the artifact is mandatory for `v1`
- **Required Instances**:
  - `Цели на 5 лет.md`
  - shared template location based on `!Общее/Шаблон сферы`
  - repository-managed assistant config location
  - repository-managed assistant log location
  - planning-chain anchor documentation for monthly, annual, and 5-year levels

## Sphere Template

- **Description**: Reference pattern for mandatory sphere structure and expected
  note semantics.
- **Fields**:
  - `source_path`: reference path in `.workspace/!Общее/Шаблон сферы`
  - `required_note_names`: `Цели.md`, `Тактические цели.md`,
    `Операционные задачи.md`
  - `required_directory_names`: `Артефакты/`
  - `semantic_guidance`: plain-language rules for long-term, monthly, and
    operational planning notes
- **Validation Rules**:
  - Template informs normalization but does not become a required production path
  - Semantic guidance may inform proposals but not justify automatic rewriting
    of existing user content

## Obsidian Client State

- **Description**: User-owned Obsidian settings and workspace layout files
  stored under `.obsidian/`.
- **Fields**:
  - `path`: `.obsidian/`
  - `files`: set of known settings files such as `app.json`,
    `appearance.json`, `core-plugins.json`, `graph.json`, `workspace.json`
  - `preservation_policy`: `preserve_unless_confirmed`
- **Validation Rules**:
  - Existing files must survive normalization unchanged unless the user confirms
    a replacement
  - Additional user settings files are allowed and remain authoritative
  - Required workspace artifacts must remain manually readable and editable
    through Obsidian after normalization

## Normalization Action

- **Description**: A planned or executed change used to bring a workspace into
  contract compliance.
- **Fields**:
  - `action_id`: stable identifier for audit linkage
  - `target_path`: repository path affected
  - `action_type`: `create`, `append_structure`, `propose`, `rename`, `move`,
    `archive`, `delete`, `overwrite`, `noop`
  - `autonomy_class`: `auto-allowed`, `proposal-only`, or
    `requires-confirmation`
  - `status`: `planned`, `executed`, `proposed`, `confirmed`, `rejected`,
    `failed`
  - `reason`: explanation tied to contract rules
  - `log_entry_id`: reference to the emitted audit record
- **State Transitions**:
  - `planned -> executed` for `auto-allowed` actions
  - `planned -> proposed -> confirmed|rejected` for gated actions
  - Any state -> `failed` when filesystem mutation or validation fails

## Audit Log Entry

- **Description**: Repository-backed record of a normalization observation or
  mutation.
- **Fields**:
  - `log_entry_id`: stable identifier
  - `timestamp`: execution time
  - `workspace_path`: resolved production workspace root
  - `action_ids`: related normalization action identifiers
  - `event_type`: `contract_violation`, `auto_fix`, `proposal`,
    `confirmed_change`, `error`
  - `summary`: user-readable description
  - `retention_days`: minimum 30
- **Validation Rules**:
  - Entries must remain readable from the repository for at least 30 days
  - Every executed or proposed mutation must map to at least one log entry

## Workspace Configuration

- **Description**: Repository-managed settings that tell the assistant where the
  production workspace lives and how to apply normalization policies.
- **Fields**:
  - `workspace_root_path`
  - `log_retention_days`
  - `default_sphere_priority_order`
  - `normalization_mode`: `report-only` or `apply-safe-structure`
  - `confirmation_required_for_obsidian_changes`: boolean
- **Validation Rules**:
  - `workspace_root_path` must be explicit and not inferred from `.workspace/`
  - `log_retention_days` must be at least 30 for `v1`
  - `normalization_mode` may never authorize destructive changes implicitly
