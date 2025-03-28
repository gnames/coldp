# Changelog

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v0.3.15] - 2025-03-11 Tue

Add: City to Actor

## [v0.3.14] - 2025-03-10 Mon

Add: update modules.

## [v0.3.13] - 2025-03-10 Mon

Add: more fields from GN.

## [v0.3.12] - 2025-03-05 Wed

Add: better ScientificNameString conversion.

## [v0.3.11] - 2025-03-04 Tue

Add: canonical forms to name and name usage.

## [v0.3.10] - 2025-03-01 Sat

Fix: environment in NameUsage.

## [v0.3.9] - 2025-02-27 Thu

Add: update modules.

## [v0.3.8] - 2025-02-20 Thu

Add: rename arcio to coldpio for clarity.
Add: improve errors types.

## [v0.3.7] - 2025-02-12 Wed

Add: test for ToBool
Add: Header methods
Fix: several typos

## [v0.3.6] - 2025-02-02 Sun

Add: BibTeX references format.

## [v0.3.5] - 2025-01-28 Tue

Add: ID method to all enums. String() now returns low-cased string with
spaces instead of "\_".

## [v0.3.4] - 2025-01-24 Fri

Add: allow NULL for preferred field in vernacular.

## [v0.3.3] - 2025-01-23 Thu

Add: allow NULL for integers and booleans

## [v0.3.2] - 2025-01-23 Thu

Add: rename Name.FullScientificName to Name.ScientificNameString.

## [v0.3.1] - 2025-01-19 Sun

Fix: NameUsage import.

## [v0.3.0] - 2025-01-16 Thu

- Add: sucessufl import of Catalogue of Life from official CoLDP download.
- Add: JSON-CSL parsing for references (did not find their need in production
  yet).

## [v0.2.0] - 2024-11-22 Fri

- Add: fixes in fields names.
- Add: more tests.
- Add: all tables and enums.
- Add: improve synonym and other data.
- Add: authors, taxa, name_relation.
- Add: geotime and oter enums.
- Add: ranks.
- Add: all tables.
- Add: more taxon fields.
- Add: nameusage.
- Fix: taxon data.

## [v0.1.0] - 2024-10-17 Thu

Add [#2]: make tsv/csv files injestable
Add [#1]: parse metadata
Add: extract works
Add: convert to library

Add: makefile
Add: initial commit

## Footnotes

This document follows [changelog guidelines]

[v0.3.14]: https://github.com/gnames/coldp/compare/v0.3.13...v0.3.14
[v0.3.13]: https://github.com/gnames/coldp/compare/v0.3.12...v0.3.13
[v0.3.12]: https://github.com/gnames/coldp/compare/v0.3.11...v0.3.12
[v0.3.11]: https://github.com/gnames/coldp/compare/v0.3.10...v0.3.11
[v0.3.10]: https://github.com/gnames/coldp/compare/v0.3.9...v0.3.10
[v0.3.9]: https://github.com/gnames/coldp/compare/v0.3.8...v0.3.9
[v0.3.8]: https://github.com/gnames/coldp/compare/v0.3.7...v0.3.8
[v0.3.7]: https://github.com/gnames/coldp/compare/v0.3.6...v0.3.7
[v0.3.6]: https://github.com/gnames/coldp/compare/v0.3.5...v0.3.6
[v0.3.5]: https://github.com/gnames/coldp/compare/v0.3.4...v0.3.5
[v0.3.4]: https://github.com/gnames/coldp/compare/v0.3.3...v0.3.4
[v0.3.3]: https://github.com/gnames/coldp/compare/v0.3.2...v0.3.3
[v0.3.2]: https://github.com/gnames/coldp/compare/v0.3.1...v0.3.2
[v0.3.1]: https://github.com/gnames/coldp/compare/v0.3.0...v0.3.1
[v0.3.0]: https://github.com/gnames/coldp/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/gnames/coldp/compare/v0.1.0...v0.2.0
[v0.1.0]: https://github.com/gnames/coldp/compare/v0.0.0...v0.1.0
[v0.0.0]: https://github.com/gnames/coldp/tree/v0.0.0
[#10]: https://github.com/gnames/coldp/issues/10
[#9]: https://github.com/gnames/coldp/issues/9
[#8]: https://github.com/gnames/coldp/issues/8
[#7]: https://github.com/gnames/coldp/issues/7
[#6]: https://github.com/gnames/coldp/issues/6
[#5]: https://github.com/gnames/coldp/issues/5
[#4]: https://github.com/gnames/coldp/issues/4
[#3]: https://github.com/gnames/coldp/issues/3
[#2]: https://github.com/gnames/coldp/issues/2
[#1]: https://github.com/gnames/coldp/issues/1
[changelog guidelines]: https://keepachangelog.com/en/1.0.0/
