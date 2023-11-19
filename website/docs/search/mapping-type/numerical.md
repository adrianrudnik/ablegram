---
description: The numerical mapping type is used for fields that should be indexed as numbers.
---

# Numerical mapping

Numeric mapping is used for fields that are to be indexed as numbers.
This is useful for fields containing numbers, such as the duration of a track or the number of page elements found.

It is possible to use range queries on numeric fields.

## Examples

- field:>=3 will find all documents with a field value greater than or equal to 3.
- field:<10 will match all documents with a field value less than 10.
- field:>=3 field:<10 will match all documents with a value of `field` between `3` and `9`.