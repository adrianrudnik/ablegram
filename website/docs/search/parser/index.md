---
description: Parsers are responsible for extracting information from files and mapping it to a searchable document.
---

# Parsers

Parsers are responsible for extracting information from files and mapping it to a searchable document.

A specific parser has a theme to work against, i.e. `AbletonLiveSet` or `MidiFile`.

A parser worker receives an absolute file path from a source collector and spawns the specific parsers that match the given file extension.

Each spawned parser will receive the absolute path to the source file from the worker and will return a searchable document with mapped information.

The parsers themselves will perform the same three steps:

1) Extract all possible results from the given file found at the given path from the collector.
2) Extract mapped fields that allow Boolean or match-type queries in free text searches.
3) Extract information tags based on custom rules that allow tag-based queries.

The results are returned to the worker thread.

Finally, the parser worker will create an additional document that combines all the extracted tags from all the parsers into one file result document to allow file based (think grouped) queries.

This document and all other documents are then passed to the indexer to be made available for searching.