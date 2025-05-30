# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

When a new release is proposed:

1. Create a new branch `bump/x.x.x` (this isn't a long-lived branch!!!);
2. The Unreleased section on `CHANGELOG.md` gets a version number and date;
3. Open a Pull Request with the bump version changes targeting the `main` branch;

Releases to productive environments should run from a tagged version.
Exceptions are acceptable depending on the circumstances (critical bug fixes that can be cherry-picked, etc.).

## [Unreleased]

### Added

- added unit tests to increase code coverage
- added information about the project in `README.md`
- added SonarQube analysis to the pipeline
- added `RabbitMQ` support
- added security step in the pipeline
- added unit tests to increase code coverage
- added commands and services to send the email
- added golang project structure
- added basic project structure

### Changed

- changed GitHub pipeline to deploy the application
