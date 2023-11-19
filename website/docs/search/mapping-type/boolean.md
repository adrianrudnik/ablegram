---
description: The Boolean mapping type is used for fields that need to be indexed for true/false searching.
---

# Boolean mapping

Boolean mapping is used for fields that need to be indexed for Boolean searching.
This is useful for fields that contain boolean values, such as the status of a document, a property flag or its visibility.

## Examples

- `+field:true` will match all documents that have a value of `field` that is `true`.
- `+field:false` will match all documents that have a value of `field` that is `false`.
