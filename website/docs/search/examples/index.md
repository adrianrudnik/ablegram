---
description: A collection of examples of how to use the search.
---

# Search Examples

All examples are in raw query form, as they can be copied and pasted into the free text search field.

There will also be a link to open the query in the demo.
This will be a parsed query, with filters added to the filter list, rather than a raw query version.
For searches containing tags, this is the preferred variant.

## Basics

By default, any piece of information that is not preceded by a `+` or `-` is considered a free text query.

If only one thing is queried, it becomes a MUST (`+`) query.

<QueryExample query="note" />

Requires `note` to be present in the result. Adding another query part makes it a SHOULD query.

<QueryExample query="note midi" />

The result will be hits of documents containing either "note" or "midi" or even both.
The best matching documents will be boosted to the front of the search result.

## Boosting results

Sometimes one part of a query is more important than another.

<QueryExample query="mixtape clip annotation" />

This query will return some very specific results because `CLIP ANNOTATION` already matches parts of the query, so it will be at the top.

But say you were more interested in the `mixtape` part.

<QueryExample query="mixtape^5 clip annotation" />

Adding `^5` to the mixtape part will make it 5 times more important than the other parts, pushing the matching results even higher.

## Must and must not

Sometimes you want to have one thing, but never the other.

<QueryExample query="-mixtape clip annotation" />

Will return hits that never have an association with `mixtape'.

<QueryExample query="+clip -minor" />

Only returns things that have `clip` and never `minor`.

## Find Info-Text content

Searches for the word `clip` in everything that has a specific Ableton info text entered by the user.

<QueryExample query="Clip" :tags="['+user:memo=true']" />

## Tag search

You can perform a tag search. Feel free to open the query in the demo to see tag filtering in action.

<QueryExample :tags="['+ableton-live-set:audio-tracks:available=false', '+ableton-live-set:midi-tracks:available=true']" />

Finds all Ableton Live Sets that have no audio tracks but at least one MIDI track.
The result will be limited to Live Sets as the tag is very specific about the `Live Set` part.
