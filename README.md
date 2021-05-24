# Gymgo

[![test](https://github.com/Manzanit0/gymgo/actions/workflows/test.yml/badge.svg)](https://github.com/Manzanit0/gymgo/actions/workflows/test.yml)

Gymgo is an application for boutiques, studios,and gyms which allows business
owners to manage their courses, classes, members, memberships etc,

## 🏗 Getting Started

The easiest way to get started with the application is through the `Makefile`.
There are two targets:

- `run`: starts the web server locally
- `test`: runs all the tests

## 🛠 Design Decisions

There have been 2 major decisions taken to solve the problem:

1. To leverage an in-memory array as well as a dictionary for storing the classes
2. To abstract the domain under a package

### Storing the classes

The simplest possible approach to store the classes is to leverage an in-memory
array. The downside to this approach is that in order to book a class we would
need to iterate the whole array to find the appropriate class, this being O(n)
in the worst scenario.

To address this potential performance issue, a dictionary mapping classes by
date was also added, transforming then the search of a class for booking into an
O(1) operation instead of O(n). However, this brings other issues to the
table, like having to keep both data structures in sync.

### Using a package for the domain

For an application as simple as this one, there's two possible ways of tackling
it: one is to keep the code within the API controller, and another one is to
create a domain layer and have the API consume it.

Generally, for spikes, or until we figure out the domain and relevant
abstractions I'd argue that it's _good enough_ to keep everything in the API
layer, as long as it's the only consumer of the logic. However, in this case
I've decided to differentiate the layers to showcase how I would go about the
solution as it evolves: as the platform grows we want to have our APIs to be as
thin as we can, and keep all the business logic in a domain layer that we can
reuse throughout different services.

## 📃 Regarding the RESTful design

In the API, the approach taken was to make both endpoints use the `PUT` method
because it communicates entity creation. However, they're not the purest of PUTs.
The reason is that ideally a `PUT` endpoint should be idempotent and deal with
being called several times gracefully, replacing the resource at hand. We could
argue it should be a `POST` and not a `PUT`. Nonetheless, I have maintained it and
`PUT` because I believe it keeps the main promise: requests are side-effects-free,
and the verb transmits intent.

If I were to further improve the API I'd say it would be along the following lines:

1. Leverage a purer RESTful approach and allow the consumers of the API to provide 
   an ID for the resource in the request URI, thus allowing the API to discern the
   resource that the request is refering to. This empowers truly idempotent endpoints
   where any `PUT` endpoint can actually replace the previous entity with the new one
   provided. Currently we're making a series of assumptions which don't allow for
   updating resources.

2. Remove the bodies from the responses. Initially I thought that it would be
   useful to respond with the information of the created entity, but when checking
   [MDN's documentation](0), it's not a common practice to do so.

3. Instead of returning a _generic_ `400 Bad Request` every time the user attempts
   to create or book a class but fails due to some precondition not being met, we
   could leverage a `409 Conflict` for some scenarios.

[0]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PUT
