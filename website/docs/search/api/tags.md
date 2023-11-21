---
description: API endpoints for tags and related results.
---

# Tags API

The following endpoints are available for tags:

## `/api/tags`

Returns a list of tags grouped by their root value, i.e. no `value`.
It returns a count of how many times each tag has been found, without regard to the actual value associated with it.

Example: https://demo.ablegram.app/api/tags

```json
{
  "ableton-device:midi-arpeggiator": 25,
  "bpm": 154,
  "color:ableton": 2371,
  "type:ableton-live-set": 153
}
```

## `/api/tags?verbose=1`

Returns a list of all known tags and their associated `value` and `count`.
This is the most verbose of the tags as it returns the exact tags that can be used in tag based queries.

Example: https://demo.ablegram.app/api/tags?verbose=1

```json
{
  "ableton-device:midi-arpeggiator": 25,
  "bpm=100": 16,
  "bpm=120": 86,
  "color:ableton=0": 1286,
  "color:ableton=30": 2,
  "type:ableton-live-set": 153  
}
```
