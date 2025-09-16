# Frontend: Database - PotgreSQL

## Status

Proposed

## Context

**DISCLAIMER**: I forgot to create this ADR before starting the project, so this is a bit of a retrospective ADR.

At this point, the application reads and writes images directly into the filesystem. This is fine for a simple
application, but as the application grows, it might be useful to have a database to store metadata about the images,
users, and other related data. Due to the limited scope of the task at hand and because of time constraints, I will
focus on using something more straightforward with a good integration

- PostgreSQL
- MySQL / MariaDB
- SQLite
- Stay with current solution (filesystem)

## Decision

Since there is no complex relational data to be stored at this point, I could get away with using SQLite. However, I was
not sure if something like user accounts would be added in the future, and I wanted to avoid having to migrate the
database later on. Therefore, I decided to go with PostgreSQL, as it is a powerful and feature-rich database that can
handle a variety of use cases. Since I already have set up Docker / Docker Compose for the project, adding a PostgreSQL
container is straightforward.

## Consequences

**Pros**

- Feature-rich and powerful
- Easy to set up with Docker
- Strong community and ecosystem
- Good support for JSON and other modern data types
- ACID compliant

**Cons**

- Might be overkill for simple use cases
- Requires more setup and maintenance than simpler databases like SQLite


