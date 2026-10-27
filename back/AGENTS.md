# Backend Guidelines

## Stack and Structure

- Use Go, Gin, and sqlc.
- Organize features into business modules under internal/modules.
- Follow the root AGENTS.md and any applicable nested AGENTS.md files.
- Inspect existing code before changing its structure.
- Apply these guidelines to the requested scope. Avoid unrelated refactoring.

Preserve the following module structure:

```text
internal/modules/<module>/
├── app/
│   ├── register_user.go
│   └── confirm_email_verification.go
├── domain/
│   ├── user.go
│   ├── repository.go
│   └── errors.go
├── infra/
│   └── postgres/
│       ├── queries/
│       │   └── users.sql
│       ├── sqlc/
│       ├── mapper.go
│       └── repository.go
└── transport/
    └── http/
        ├── handler.go
        ├── routes.go
        ├── requests.go
        └── responses.go
```

File names above are examples for the users module. Name files according
to the responsibilities of each module.

- app/: application use cases, inputs, outputs, and application ports.
- domain/: entities, business rules, domain errors, and repository contracts.
- infra/postgres/: repository implementations, mappings, SQL queries,
  and sqlc-generated code.
- transport/http/: Gin handlers, routes, HTTP requests, and HTTP responses.

Use these directory names consistently. Do not rename app to application
or transport/http to http.

Keep use cases directly in app/. Do not introduce an additional usecases
directory unless the task explicitly requires reorganizing this layer.

Use the existing shared infrastructure location. If none exists, use
internal/platform for database connections, email delivery, the shared
error type, and centralized HTTP error handling.

Create directories only when they contain actual implementation.
Do not generate empty architectural scaffolding.

## Dependency Boundaries

- Keep domain independent of app, infra, and transport.
- Keep app independent of infra, transport, Gin, and database drivers.
- Never access the database directly from HTTP handlers.
- Keep sqlc-generated types inside the persistence boundary.
- Define explicit API request and response types.
- Keep HTTP binding tags and transport validation out of application types.
- Inject dependencies through constructors.
- Define small interfaces according to consumer needs.
- Keep domain repository contracts in domain/repository.go.
- Define application-specific capability interfaces, such as email delivery,
  in app when consumed there.
- Implement persistence contracts in infra/postgres.
- Assemble concrete dependencies at the application composition root.
- Avoid circular dependencies between modules.
- Do not access another module's persistence implementation directly.

## Use Cases

- Represent each business operation as an explicitly named use case.
- Examples: RegisterUser, ConfirmEmailVerification, RequestPasswordReset.
- Keep each use case in a dedicated file in app/.
- Name files after their business operations, such as:
  - register_user.go
  - confirm_email_verification.go
  - request_password_reset.go
- Prefer named use-case types with an Execute method.
- Inject required dependencies through constructors.
- Keep operation-specific input and output types in the use-case file.
- Create shared types only when they are genuinely shared.
- Accept context.Context and propagate it to external operations.
- Keep business validation in use cases or domain logic.
- Avoid a single oversized Service type containing every module operation.
- Let handlers call use cases directly.
- Do not add intermediary services without a concrete responsibility.

## Services and Helper Functions

- Create services for reusable capabilities with clear responsibilities,
  such as email delivery or token generation.
- Do not introduce services that only forward calls without adding value.
- Extract functions when they express a meaningful responsibility or
  simplify complex logic.
- Do not split straightforward code merely to reduce line counts.
- Avoid generic utils packages for business logic.
- Prefer names that describe intent rather than implementation details.
- Reuse existing implementations before introducing new abstractions.

## HTTP Handlers

- Limit handlers to adapting HTTP requests to application calls.
- Use named request and response structs in transport/http.
- Map HTTP requests to use-case inputs.
- Map use-case results to HTTP responses when a response body is required.
- Validate request format and required fields at the transport boundary.
- Keep business rules out of handlers.
- Prefer early returns over unnecessary nesting.
- Delegate errors to the centralized error-handling mechanism.
- Do not construct ad hoc error responses with gin.H.
- Do not duplicate error-to-status mappings across handlers.
- Use named HTTP status constants instead of numeric literals.
- Split handler files by operation or cohesive responsibility when needed;
  do not let handler.go become a container for unrelated behavior.


## Persistence and sqlc

- Write application queries in SQL files processed by sqlc.
- Keep module queries in infra/postgres/queries and generated code in
  infra/postgres/sqlc, following the project's sqlc configuration.
- Execute application queries exclusively through sqlc-generated methods.
- Do not write inline SQL in handlers, use cases, or repositories.
- Do not introduce an ORM or query builder.
- Never manually edit generated files.
- Regenerate sqlc code after changing queries or schemas.
- Verify that generated code and its consumers compile.
- Map generated persistence types to domain or application types as needed.
- Translate driver-specific errors at the persistence boundary.
- Distinguish missing records from infrastructure failures.
- Identify the violated constraint before mapping a uniqueness error
  to a specific business error.
- Use transactions when an operation requires atomicity.
- Ensure all queries within a transaction use that transaction.
- Keep transaction implementation details out of HTTP handlers.
- Do not expose database-driver transaction types to domain or app.
- Preserve the business operation's transaction boundary through an
  appropriate repository operation or application-facing transaction contract.

Migrations, connection initialization, and transaction lifecycle management
are infrastructure concerns and do not need to be generated by sqlc.

## Application Errors

- Use a shared, transport-independent application error type containing:
  - A stable machine-readable code.
  - A transport-independent category.
  - A safe public message.
  - An optional internal cause.
- Keep the shared error package independent of Gin and database drivers.
- Define error codes as constants.
- Keep module-specific error definitions within their owning module.
- Keep generic error definitions in the shared error package.
- Give conditions requiring different client behavior distinct codes.
- Reuse a code when its meaning is the same across operations.
- Example codes:
  - EMAIL_NOT_VERIFIED
  - VERIFICATION_TOKEN_INVALID
  - EMAIL_ALREADY_EXISTS
  - VALIDATION_ERROR
  - INTERNAL_ERROR
- Keep HTTP status codes out of domain and app.
- Use categories such as validation, unauthenticated, forbidden,
  not_found, conflict, and internal.
- Preserve internal causes and support errors.Is/errors.As where applicable.
- Never identify errors by comparing err.Error() strings.
- Do not mutate shared error instances with request-specific data.

## Error Response Contract

All API errors must follow this structure:

```json
{
  "error": {
    "code": "EMAIL_NOT_VERIFIED",
    "message": "Verify your email before continuing."
  }
}
```

- Centralize category-to-HTTP-status mapping and JSON serialization.
- Clients must distinguish errors by code, never by message text.
- Return unknown errors as HTTP 500 with INTERNAL_ERROR and a generic message.
- Never expose database errors, stack traces, credentials, or internal causes.
- Log internal causes without logging sensitive tokens or personal data.
- Apply the same response contract to validation, authentication,
  authorization, unmatched routes, and application-controlled panic recovery.
- Write at most one response per request.
- Never attempt to replace a response that has already been written.
- If using Gin error middleware, handlers must attach the error and stop
  processing; the middleware owns error response serialization.
- Keep public message language consistent with the product's locale.

## Code Quality

- Prefer cohesive functions, clear names, and explicit dependencies.
- Avoid unnecessary abstraction layers and generic frameworks.
- Keep reusable business rules outside handlers.
- Avoid deeply nested conditionals and unrelated responsibilities
  in the same function.
- Do not create one-line wrappers unless they establish a useful boundary.
- Follow existing naming and formatting conventions.
- Separate code by responsibility rather than arbitrary file-size limits.
- Do not add interfaces, services, or factories solely to satisfy a pattern.

## Formatting Across All Layers

- Apply these rules to all handwritten Go code, including handlers,
  use cases, repositories, constructors, and helpers.
- Separate functions and methods with one blank line.
- Write function bodies and struct declarations across multiple lines.
- Use named local variables for structured inputs and query parameters.
- Expand struct literals to one field per line.
- Separate parameter construction, execution, and the final response
  or return with blank lines.
- Keep error checks immediately after the operation that produced the error.
- Split long signatures and calls across multiple lines.
- Keep short, simple calls on one line when readable.
- Do not change error semantics or business behavior during style-only edits.
- Do not modify sqlc-generated files.
- Run gofmt after editing.

## Change Discipline

- Inspect similar implementations before adding functionality.
- Follow the established module structure.
- Preserve existing behavior and API contracts unless explicitly asked
  to change them.
- Do not silently change public error codes.
- Apply structural improvements within the requested scope.
- Preserve unrelated user changes.
- Update relevant documentation when an approved architecture or contract
  change makes it outdated.

## Verification

- Test affected use-case rules, error codes, and HTTP contracts.
- Verify persistence behavior and transaction atomicity when relevant.
- Verify that centralized error handling covers the affected failure paths.
- Use existing project commands for formatting, generation, and tests.
- Discover commands from project configuration instead of inventing them.
- Never modify generated files or weaken tests to hide implementation errors.
- Report what changed, what was verified, and any remaining limitations.
- Do not claim a check passed unless it was executed successfully.
