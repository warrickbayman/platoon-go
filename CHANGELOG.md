# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## v0.0.1-alpha4 - 2026-04-20

The first version that can actually complete a full Laravel deployment.

- Small change to how commands are run. Added a `Case` member to the `PlatoonCommand` struct
- Added a spinner to the UI
- Added tests for the `spinner.go` file
- Added `shell.FileExists` and `shell.DirectoryExists`
- Updated the logging to write verbose output to the `deploy.log` file.
- `config.Releases.Max` will have a default of 2 set instead of 0 which caused fresh deployments to be then be deleted.

## v0.0.1-alpha.3 - 2026-04-06

Some small changes based on tests that have been added.
