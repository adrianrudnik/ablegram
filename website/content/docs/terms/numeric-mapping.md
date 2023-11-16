---
title: "Numeric Mapping"
description: ""
summary: ""
date: 2023-11-16T19:38:14+01:00
lastmod: 2023-11-16T19:38:14+01:00
menu:
  docs:
    parent: ""
    identifier: "numeric-mapping-7410ab9313d491d068d06248669fb41a"
weight: 999
toc: true
seo:
  title: "" # custom title (optional)
  description: "" # custom description (recommended)
  canonical: "" # custom canonical URL (optional)
  noindex: false # false (default) or true
---

The numeric mapping is used for fields that should be indexed as numbers.
This is useful for fields that contain numbers, such as the duration of a track, or the number of pages elements found.

It is possible to use range queries on numeric fields.

## Examples

- `field:>=3` will match all documents that have a value of `field` greater or equal to `3`.
- `field:<10` will match all documents that have a value of `field` less than `10`.
- `field:>=3 field:<10` will match all documents that have a value of `field` between `3` and `9`.
