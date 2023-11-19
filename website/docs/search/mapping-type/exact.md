---
description: The Exact mapping type is used for fields that should be indexed as they are, without any analysis, for exact matching.
---

# Exact mapping

Exact mapping is used for fields that should be indexed as they are, without any analysis.
It is also known as keyword mapping.
This is useful for fields that contain identifiers, such as file paths or tags.
It will only match if the search value is exactly the same as the indexed value.

- If the field type is a string, an exact match must occur on a single value query.
- If the type is an array, an exact match must occur on one of its values.
