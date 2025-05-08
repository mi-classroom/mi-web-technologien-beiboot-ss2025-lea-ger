# Frontend: JS Runtime - Bun.js
## Status

Proposed

## Context
A JavaScript runtime is needed to run the frontend code. The options I considered for the frontend are:
- Node.js
- Deno
- Bun.js

## Decision
Bun.js is a modern JavaScript runtime that is designed to be fast and efficient. It has a built-in bundler, transpiler, and package manager, which makes it a great choice for this project. It is also compatible with existing Node.js code, so I can use it with my existing TypeScript code.
It's also beloved by developers and supposed to be more performant than the alternatives.


## Consequences

**Pros**
- Fast
- Built-in bundler, transpiler, and package manager

**Cons**
- Might run into compatibility issues with existing Node.js code


