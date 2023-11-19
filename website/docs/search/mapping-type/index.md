---
description: The mapping types used to define how information is indexed and searched.
---

# Mapping types

Mapping types are used to define how information is indexed and searched.

A mapping takes a value of a given `type`, i.e. `string`, and is processed by a ruleset that defines how the value is made available for search results. This can be a simple `fulltext` mapping that indexes the value in a human-understandable way, or an `exact` mapping that indexes the value as it is, without any analysis.