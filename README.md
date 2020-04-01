A server for Klaus Teuber's classic board game _Settlers of Catan_.

Please note that this server is still a work-in-progress (see features list below).

# Design overview

A long time ago in a university far, far away, I took a class in which final project was to write a web-based version of _Catan_ as part of a group. Looking back on it, I realize that my group's implementation had some serious drawbacks, namely:

* **Poorly performing**. Despite the fact that _Catan_ naturally lends itself to an event-based model, the front end used a polling mechanism to obtain the latest game state from the server. This model was easy to implement and worked sufficiently well for a school project but would never work in a real microservices product due to its resource intensiveness and performance lag.

* **The server implementation could not scale horizontally**. In order to accommodate the polling model utilized by the front end, all game state was stored in-memory on the server so that it could be read back in a relatively short amount of time. This design violates one of the core principles of microservices, which is that service instances never store state, but should have enough context within a single request to handle it adequately. The consequences of this violation are that the server implementation could never be deployed as a scalable microservice, as in-memory game state could not be transferred from instance to instance.

## A functional approach

My first priority in designing this version of _Catan_ was making the entire system [functional](https://en.wikipedia.org/wiki/Functional_programming). In order to accomplish this, I specifically focused on the following principles:

* **The current state of any game should be determined by applying immutable events to an initial state.** Much in the same way a bank account's current balance can be calculated by taking the most recent balance and applying each transaction one-by-one, I wanted this version of _Catan_ to "calculate" a game's current state by starting with an initial state and then applying each event one-at-a-time. Designing the system this way would force state to be stored outside of the server, essentially making the server a set of pure functions.

* **Event-driven**. Clients publish user actions as events to the server, and the server publishes changes in state as asynchronous messages for all clients.

* **State doesn't limit scaleability**. This implementation was designed as if for a real microservices product. Hence, the server implementation doesn't store state and can scale as needed.

With these principles in mind, I've arrived at the following overall design:

* A game action is performed by the user. The front end sends a corresponding event to the server.

* The server calculates the current state of the game by getting the game's initial state and applying all subsequent events one-at-a-time.

* With the current state of the game calculated, the server verifies that the event can be performed given the current state of the game. If not, it returns an error.

* Having determined that the event can be applied to the current game state, the server applies the event and then saves it.

* After the event has been saved successfully, the server fires off an asynchronous message notifying all clients within the affected game that the particular event occurred.

* Each client receives the event, and then adjusts the user's view accordingly.

## Utilizing the state library

The server functionality doesn't exist as of yet, but the state machine functionality does. This section describes how to use it.

### Prerequisites

In order to use the state machine, you must have go version 1.14 or higher installed on your machine.

### Usage

First download the `game` package:

```
$ go get github.com/jtcotton63/catan/game
```

Once this is finished, you can utilize it like so:

```
TODO
```

## Features list

TODO

## Project layout

This project is based on [the Golang standard project layout](https://github.com/golang-standards/project-layout).