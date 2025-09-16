# Virtualization - Docker / Docker Compose

## Status

Proposed

## Context

**DISCLAIMER**: I forgot to create this ADR before starting the project, so this is a bit of a retrospective ADR.

Building Go applications requires setting up a proper Go environment, with the right version of Go and the necessary
dependencies. This can be a bit of a hassle, especially if you want to ensure that the application runs in the same
environment across different machines. To solve this problem, I considered using virtualization options:

- Docker (with Docker Compose)
- None

## Decision
I decided to use Docker and Docker Compose for this project. Docker allows me to create a containerized environment
that includes all the necessary dependencies and configurations to run the application. This ensures that the application
runs consistently across different machines, and no tooling needs to be installed on the host machine, other than Docker itself.
This way, the application can simply be launched with a single command (theoretically).

## Consequences

**Pros**

- Simple for users to get started with the application
- Ensures consistent environment across different machines

**Cons**

- More effort to set up initially
- Adds a layer of complexity to the development process
- Requires users to have Docker installed