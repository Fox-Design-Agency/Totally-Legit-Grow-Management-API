# Totally Legit Grow Management API

This is the repository for the base API for the totally legit grow system.

## Table of Contents

- [Getting Started](#getting-started)
  - [Running Locally](#running-locally)
- [Notes](#notes)
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

## Getting Started

### Clear DB Quickly

    DROP SCHEMA public CASCADE;
    CREATE SCHEMA public;
    GRANT ALL ON SCHEMA public TO postgres;
    GRANT ALL ON SCHEMA public TO public;
