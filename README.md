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

Location of the function within the handler files should inherit based on the section of the route within main. Since this function sets up the request to pass to the logic layer, there should be minimal disagreement on proper placement based on the main route placement.

#### Logic Pattern

The basic flow of code through the logic funcs should follow this order with minimal deviation:

1. Run Validation Checks
2. Start Transaction if desired
3. Initiate all persistene layer calls in desired sequence
4. Commit transaction if started
5. Form Response and return back to Handler Func for response

Location of the function within the files should inherit from the location of the function within the handler files. This layer has the chance to call out to many persistence functions and can cause disagreement with proper file location.

#### Persistence Pattern

The basic flow of code through the persistence pattern should follow this order with minimal deviation:

1. State the Stuff to Return. Basically what is going to be scanned from the SQL call (if anything) and returned with the expected response.
2. Define the arguments to be used in the SQL call
3. Execute the SQL call with a SQL statement defined in an associated const file for SQL strings
4. Return the Expected Response

There should only be one SQL call in a single function. If multiple calls are required, create them separately and compose them with a transaction in the logic layer. Try and reuse SQL code when possible, but too much reuse could potentially cause unintended side affects when changing the strings.

### Clear DB Quickly

    DROP SCHEMA public CASCADE;
    CREATE SCHEMA public;
    GRANT ALL ON SCHEMA public TO postgres;
    GRANT ALL ON SCHEMA public TO public;
