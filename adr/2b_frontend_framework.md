# Frontend: Framework - Svelte

## Status

Proposed

## Context

An interactive frontend needs to be created for the application in this task. To quickly create the frontend, using a
framework that gets rid of development overhead is a good idea. The options I considered for the frontend are:

- Astro.js
- Svelte
- HTMX
- Solid
- Qwik

## Decision
Svelte is a modern JavaScript framework that compiles components into highly optimized vanilla JavaScript at build time.
The other options all have their own strengths, but Svelte's large userbase and closeness to other JavaScript frameworks
like Vue that I'm familiar with make it a good choice for this project.
Astro.js would be the strongest runner-up, but I am already familiar with it and I want to explore new technologies. Also,
my workshop will discuss Astro.js, so I don't want to use it for this project.

## Consequences

**Pros**
- Adds reactivity and interactivity at build time
- Large community and ecosystem
- Integrates well with existing JavaScript libraries
- Familiar syntax for someone with experience in Vue.js
- Great documentation and learning resources
- Quick setup with Vite as build tool

**Cons**
- Slightly more complex build process
- Reuses some concepts from other popular frameworks (boring)


