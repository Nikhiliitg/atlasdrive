# AtlasDrive

AtlasDrive is a backend-only file storage service inspired by systems like Google Drive and Dropbox.

There is no UI here. That’s intentional.

This project was built to explore correctness, scalability, and clean architecture in a realistic backend system. It is not a feature race, a demo app, or a UI experiment.

If you’re here for flashy screens, this repo probably won’t help you.  
If you care about how backend systems are *actually* put together, keep reading.

---

## Why AtlasDrive?

Most “Drive clones” stop once CRUD works.

AtlasDrive was built with a different mindset:

> Fewer features. Stronger guarantees.

The emphasis is on:
- Explicit ownership and authorization
- Predictable, traceable data flow
- Infrastructure that can be replaced without rewriting the system
- Constraints that resemble production, not tutorials
- Code that explains why it exists

---

## Architecture

AtlasDrive follows Clean Architecture / Hexagonal Architecture ideas.

The core rules are simple:
- Business logic does not depend on frameworks
- Interfaces define contracts
- Infrastructure remains at the edges
- Read paths and write paths are treated differently on purpose


## Project layout:

- domain/ → Core business entities and invariants
- application/ → Use cases and orchestration
- ports/ → Interfaces (repositories, queries)
- adapters/ → External implementations
- cmd/ → Entry points



The domain layer has no knowledge of PostgreSQL, Redis, HTTP, or JWT.  
That separation is deliberate and enforced.

---

## Authentication and Authorization

Authentication is handled using JWT.

Some rules that the system follows:
- Tokens are issued on login
- User identity is extracted in HTTP middleware
- Owner identity flows through context
- Handlers never trust client-supplied owner IDs

This avoids an entire class of authorization bugs that usually show up later.

If a handler needs to know who the user is, it gets that information from context — not from the request body.

---

## What the System Supports

- User registration and login (JWT-based)
- Folder creation with parent–child relationships
- File creation inside folders
- Listing folder contents (folders and files)
- Strict per-user data isolation
- Soft deletes for future recovery and audit paths

Nothing fancy. Just the hard parts done carefully.

---

## Data and Persistence

PostgreSQL is the single source of truth.

- Strong constraints enforce correctness
- Foreign keys and uniqueness rules mirror real storage systems
- Soft deletes make features like trash and restore possible later

The database is trusted to do what databases are good at.

---

## Read Caching with Redis

Read-heavy operations, especially folder listing, are cached using Redis.

Cache design decisions:
- Cache keys are scoped by user ID and folder ID
- Full API responses are cached, not partial results
- PostgreSQL remains authoritative
- Cache entries have TTLs
- Write operations explicitly invalidate affected cache entries

Redis is treated as an optimization, not a requirement for correctness.

---

## Testing

Testing aims to reflect production behavior rather than mock it away.

- Unit tests for application-level use cases
- Authentication enforced via context in tests
- Manual end-to-end testing using Insomnia
- Redis behavior verified using `redis-cli MONITOR`
- Database constraints validated directly through SQL

If a test passes but the system breaks in reality, the test is wrong.

---

## Tech Stack

- Go (Golang)
- PostgreSQL
- Redis
- JWT
- net/http

No heavy frameworks. Everything is wired explicitly.

---

## What’s Next

The architecture leaves room for growth without rewrites:

- Object storage–backed file upload pipeline
- Background workers and event-driven processing
- Sharing and access control (ACLs)
- Search and indexing
- Rate limiting and observability

---

## Final Note

AtlasDrive is not about building fast.

It’s about building correctly first, and earning speed later.

If you’re a senior engineer reading this —  
am I heading in the right direction, or missing something important?
