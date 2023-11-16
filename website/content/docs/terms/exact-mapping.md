---
title: "Exact mapping"
description: ""
summary: ""
date: 2023-11-16T16:46:58+01:00
lastmod: 2023-11-16T16:46:58+01:00
menu:
  docs:
    parent: ""
    identifier: "exact-mapping-b94d5b4448313e0d98ca9661e489419c"
weight: 999
toc: true
seo:
  title: "" # custom title (optional)
  description: "" # custom description (recommended)
  canonical: "" # custom canonical URL (optional)
  noindex: false # false (default) or true``
---

The exact mapping is used for fields that should be indexed as-is without any analysis.
It is also known as keyword mapping. 
This is useful for fields that contain identifiers, such as file paths, or tags.

- If the type is of the field is a `string`, an exact match must happen on a single value query.
- If the type is an array, an exact match must happen on one of its values.
