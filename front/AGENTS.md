# Frontend implementation guidelines

## Vue 3, Pinia, and shadcn-vue

- Write code developers can understand by following the UI flow.
- Keep UI state and UI-triggered actions in Pinia only when they must be shared across independent components or pages; otherwise keep the flow in its feature composable.
- Extract pure business rules, data transformations, and lengthy API workflows into clearly named functions using domain terms.
- Extract a component when it has a distinct visual responsibility or is reused; do not create one only to shorten a file.
- Avoid nested ternaries, complex inline types, and generic abstractions used only once.
- Structure each feature around its domain. Keep feature-specific `api/`, `model/`, `composables/`, `components/`, and `pages/` together under `features/<feature>/`; put code in `shared/` only when multiple features genuinely reuse it.
- Keep page components declarative: pages compose visual components and connect their events, but do not contain API calls, pagination setup, dialog state, or multi-step UI workflows.
- Put the state, loading/error handling, side effects, and named actions for one page or UI flow in a feature composable named `use<Domain><Flow>` (for example, `useCategoryList`). A composable may call feature API services; a Vue page or visual component may not.
- Keep visual components focused on one visual responsibility. They receive explicit props and emit named domain events; they do not call API services or coordinate unrelated dialogs, filters, and lifecycle actions.
- Put pure domain rules, validation, mapping, and data conversion in `model/` as named functions. These functions must not depend on Vue, Pinia, or HTTP.
- Prefer explicit control flow over clever syntax: use named `computed` values and `if`/`switch` statements instead of nested ternaries, long inline conditions, or action lookup abstractions that obscure behavior.
- For a single-statement `if`, omit braces and put its statement on the following line.
- Keep templates readable: do not compress sibling elements onto one line. Let the Biome formatter decide attribute wrapping, and move non-trivial template expressions into a named function or `computed` value.
- Do not extract solely to meet a line limit. Extract when a file has more than one responsibility; as a review signal, investigate pages over roughly 150 lines and visual components over roughly 200 lines.
- Preserve API contracts. Explain any necessary behavior change before implementing it.
- Follow the existing feature-based structure unless there is a clear reason to introduce a new pattern.
- Reuse the shared company-path helper for company-scoped API endpoints.
- Give request payloads and list parameters named types; keep their fields aligned with the API contract.
- Build request payloads in a named, explicitly typed variable before passing them to a service. Do not assemble object literals directly in service calls.
- Keep each service HTTP method easy to scan: use response generics, assign the response or data, then return it.
- Put individual backend requests in `services/api/`; put multi-request workflows in `services/use-cases/` as plain functions. Keep simple HTTP methods directly in service objects, and never introduce classes, interfaces, or wrappers solely for structure.
- Reuse API types in Vue and Pinia only when their shape and domain meaning match; otherwise convert explicitly at one boundary.
- In Vue templates, add a blank line between multiline siblings.

## Testing

- Work test-first (TDD): write a Vitest test that goes red on the missing behaviour or the bug, then write the code that turns it green. A bug fix starts with a test that reproduces it, plus the sibling cases of the same kind.
- Colocate tests as `<file>.test.ts` and build test data with the factories in `src/test-utils/factories.ts`.
- A pre-commit hook (`.githooks/pre-commit`, enabled by `pnpm install`) runs the whole suite, so a red test blocks the commit. Get the suite green before committing.
