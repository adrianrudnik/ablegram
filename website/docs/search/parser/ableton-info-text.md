---
description: Parser for Ableton Info-Text elements.
---

# Ableton Info-Text parser

This parser will extract any annotation, also known as Info-Text, within an Ableton Live Set file and create its own dedicated result item per annotation.

This text can be entered by the user in the Live Set and is used to add notes to given tracks, devices, clips or other elements.

The resulting document type is `AbletonInfoText`.

## Fields

### `type`

<IndexProp name="type%" type="string" mapping="exact">

The exact type of the document.

- `+type:AbletonInfoText` will match all documents of this type.
</IndexProp>

<!--@include: ./shared/base.md-->
<!--@include: ./shared/annotation.md-->

### `parent`

<IndexProp name="parent" type="string" mapping="exact" :version="[11,10,9]">

A path-like string that represents the path towards the given Info-Text element.

It is based on a simplified XML structure path to give a hint of the location of the annotation within the Live Set.

- `+parent:"LiveSet/Tracks/AudioTrack/Name/Annotation"` will match all annotations that are at that given location.
</IndexProp>

## Produced tags

<IndexTag base="type:ableton-info-text">

A type tag that matches all documents of this type.

- `+tags:"type:ableton-info-text"` if the result document is of type `AbletonInfoText`. 
</IndexTag>
