---
title: "Fulltext mapping"
description: ""
summary: ""
date: 2023-11-16T18:48:40+01:00
lastmod: 2023-11-16T18:48:40+01:00
draft: true
menu:
  docs:
    parent: ""
    identifier: "fulltext-mapping-468d1e0e6d2ed27e9403a386e92556ad"
weight: 999
toc: true
seo:
  title: "" # custom title (optional)
  description: "" # custom description (recommended)
  canonical: "" # custom canonical URL (optional)
  noindex: false # false (default) or true
---

The fulltext mapping is used for fields that should be indexed for human langauge search.
This is useful for fields that contain text, such as the content of a document, or its title.

## Examples

- `+field:foo` will match all documents that have a value of `field` that contains the word `foo` or something similar.
- `+field:*` will match all documents that have a value of `field` that contains any word.
- `field:"foo bar"` boost matches of documents that have a value of `field` that contains the exact phrase `foo bar` with a higher result score.
