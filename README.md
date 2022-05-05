# Totally Legit Grow Management API

This is the repository for the base API for the totally legit grow system.

## Table of Contents

- [Getting Started](#getting-started)
  - [Running Locally](#running-locally)
- [Notes](#notes)
  - [Patterns](#patterns)
    - [General Route Pattern](#general-route-pattern)
    - [Handler Pattern](#handler-pattern)
    - [Logic Pattern](#logic-pattern)
    - [Persistence Pattern](#persistence-pattern)
  - [Clear DB Quickly](#clear-db-quickly)

## Getting Started

### Running Locally

This project requires the following ENV variables:

- DBHOST
- DBNAME
- DBUSER
- DBPASS
- PORT: This is the port for the service to run on, the service will default DB port to 5432
- ENVIRONMENT

With the ENV variables set, this project can be currently run in isolation by navigating to the run folder and then executing:

    go run .

## Notes

### Patterns

#### General Route Pattern

The general flow of routes should be defined in this order:

1. route is defined under the appropiate section in main.go. Sections have been created by following semi natural divisions in the DB. There will be overlap. When overlap occurs precidence to location should be given to the routes utilization on the frontend. IE if the route is mostly being called from a product page but also requires harvest information, then the route should be in the product section. If the route is called in multiple places an even number of times, then its placement will be subject to the developer with reason.
2. The route will then call the coresponding Handler func, see (handler func section)[#handler-pattern] for more information on this layer.
3. From the Handler layer, the request will move through to the Logic layer, see (logic layer section)[#logic-pattern] for more information on this layer.
4. Once the logic checks have completed, the logic layer will potentially call multiple functions in the persistence layer, see (persistence layer section)[#persistence-pattern] for more information on this layer.

#### Handler Pattern

The basic flow of code through the handler funcs should follow this order with minimal deviation:

1. Start Telemetry
2. Parse form or get variables to create appropiate request struct
3. Call associated Logic func
4. Return appropiate json response

#### Logic Pattern

COMING SOON

#### Persistence Pattern

COMING SOON

### Clear DB Quickly

    DROP SCHEMA public CASCADE;
    CREATE SCHEMA public;
    GRANT ALL ON SCHEMA public TO postgres;
    GRANT ALL ON SCHEMA public TO public;
