# Gymgo

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
O(1) operations instead of O(n). Nonetheless, this brings other issues to the
table, like having to keep both data structures in sync. Currently it's a
non-issue, but definitely something to keep in mind.

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
