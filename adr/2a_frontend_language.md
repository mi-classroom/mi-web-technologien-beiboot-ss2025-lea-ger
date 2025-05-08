# Frontend: Programming Language - TypeScript

## Status

Proposed

## Context

An interactive frontend needs to be created for the application in this task. The web runs on JavaScript, however we do
have options to use other languages that compile to JavaScript. The options I considered for the frontend are:

- TypeScript
- JavaScript
- Elm
- Dart

## Decision

TypeScript is a superset of JavaScript that adds static typing. It is widely used and has a large community. It is
also the language of choice for many popular frameworks, such as Angular and React. It is a good fit for this project
because it allows for better type checking and code quality than vanilla JavaScript.
The other transpiled languages I considered are Elm and Dart. Both have the disadvantage of not being as widely used as
JavaScript and TypeScript.

## Consequences

**Pros**

- Stable frontend code due to type checking
- Large community and ecosystem
- Easy to integrate with existing JavaScript libraries
- Faster development time (I already know TypeScript)

**Cons**

- Slightly more complex build process
- More complexity due to type checking
