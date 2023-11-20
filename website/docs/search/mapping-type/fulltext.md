---
description: Full text mapping is used for fields that need to be indexed for human search.
---

# Full text mapping

Full text mapping is used for fields that need to be indexed for human search.
This is useful for fields that contain text, such as the content of a document or its title.
It will match if the search value is similar to the indexed value.

## Examples

- `+field:foo` will match all documents that have a value of `field` that contains the word `foo` or something similar.
- `+field:*` will match all documents that have a value of `field` that contains any word.
- `+field:"foo bar"` will give a higher score to matches of documents with a value of `field` that contain the exact phrase `foo bar`.