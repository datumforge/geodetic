[![Build status](https://badge.buildkite.com/6d265b36442938da91767d8f801e18c1119bea627591f03234.svg)](https://buildkite.com/datum/geodetic?branch=main)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=datumforge_geodetic&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=datumforge_geodetic)
[![Go Reference](https://pkg.go.dev/badge/github.com/datumforge/geodetic.svg)](https://pkg.go.dev/github.com/datumforge/geodetic)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache2.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/datumforge/geodetic)](https://goreportcard.com/report/github.com/datumforge/geodetic)

# geodetic

> A geodetic datum is an abstract coordinate system with a reference surface (such as sea level) that serves to provide known locations to begin surveys and create maps

The Geodetic service is what is typically called a DAL (Data Abstraction Layer) - it acts as an intermediary and pointer between [the datum service](https://github.com/datumforge/datum) and where the actual database is. We've built in initially support for Turso as a third party provider, but plan to use geodetic to facilitate "BYO-database" in the system. This would allow a user / customer of the Datum service to create a database in an external location (e.g. AWS, GCP, Azure), configure their datum organization with that database connection string and credentials which would allow the data being persisted in the Datum organization to reside in the configured data store.

## Developing

Setup [Taskfile](https://taskfile.dev/installation/) by following the instructions and using one of the various convenient package managers or installation scripts. After installation, you can then simply run `task install` to load the associated dependencies. Nearly everything in this repository assumes you already have a local golang environment setup so this is not included. Please see the associated documentation.

## Updating Configuration Settings

See the [README](/config/README.md) in the `config` directory.

