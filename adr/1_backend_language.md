# Backend/API Server: Programming Language and Framework - Go & Gin

## Status

Proposed

## Context

The task at hand is to create a simple API to upload, store and manipulate IPTC metadata. This feature will be built
upon in the coming weeks, so having an extensible foundation that can handle a variety of tasks might be useful.
The options I considered for the backend are:

- Typescript / SvelteKit
- Elixir / Phoenix
- Go / Gin
- Rust / Actix
- Zig / Ziggy

## Decision

The Typescript and SvelteKit combination is a strong contender, as it would allow me to use the same language and
toolkit for both frontend and backend. It is beloved by many and has a large community, albeit not as large as React's
or Vue's.
However, creating this project in TypeScript clashes with my goal of exploring new languages and frameworks a bit.

Elixir and the Phoenix framework are also great options, Elixir being one of the most desired programming languages and
the Phoenix framework
being the most beloved technology in the StackOverflow survey for 3 years in a row. It's strenghts lay in its
concurrency and fault tolerance, but I don't think this is a requirement for the project at hand.
Furthermore, Elixir might be a bit hard for me to get into, as it's syntax is inspired by Ruby which I don't have any
experience with. As much as I'd like to use this technology, it is not the best fit for this project.

Lastly, Rust and Actix are also great options. Rust is a language I have used in a prior project and I'd like to deepen
my knowledge of the language by using it in a context like this. But, due to the time constraints of this project and my
previous experience with Rust as a pretty complicated language to get even simple stuff working, I'm afraid that it
might not be the best fit due to the time constraints of this project.

Zig is a relatively new language that has been gaining traction in the last few years. It is a low-level language with a
strong focus on performance and safety. Ziggy is a web framework for Zig that is still in its early stages and lacks a
large community, just as the language itself. It might take a lot of work to develop new features and find help for
problems I'm facing.

Go and the Gin framework are a great fit for this project. Go is a language I have been wanting to learn for a while,
with a large community and low-level performance. For the task at hand, I can use the goexiv2 wrapper.
Within the Go ecosystem, there are multiple options for creating web servers, such as Echo and Go Chi. My choice is Gin,
as it is the most popular framework and has a large community. It also seems like the most feature-rich out of these
options, although it lacks an out-of-the-box ORM which some web frameworks provide.

## Consequences

**Pros**

- Very fast, single binary, easy to deploy
- Go wrappers for exiv2
- Strong standard library for HTTP, JSON etc.
- Very easy to cross‑compile and containerize
- Great community and ecosystem

**Cons**

- No "batteries included", as in no ORM or templating engine
- A little more time-consuming due to using a new language
