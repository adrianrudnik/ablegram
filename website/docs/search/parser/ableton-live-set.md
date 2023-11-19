---
description: Parser for Ableton Live Set files.
---

# Ableton Live Set parser

This parser will extract metadata from Ableton Live Set files.

The resulting document type is `AbletonLiveSet`.

## Fields

### `type`

<IndexProp name="type%" type="string" mapping="exact">

The exact type of the document.

- `+type:AbletonLiveSet` will match all documents of this type.
</IndexProp>

<!--@include: ./shared/base.md-->
<!--@include: ./shared/annotation.md-->
<!--@include: ./shared/scale-info.md-->

### `bpm`

<IndexProp name="bpm" type="number" mapping="numeric" :version="[11,10,9]">

The global tempo of the Live Set in BPM.

- `+bpm:>=120` will match all documents that have a tempo of 120 BPM or higher.
- `+bpm:<120` will match all documents that have a tempo of less than 120 BPM.
- `+bpm:>=120 +bpm:<130` will match all documents that have a tempo between 120 and 129 BPM.
</IndexProp>

### `midiTrackCount`

<IndexProp name="midiTrackCount" type="number" mapping="numeric" :version="[11,10,9]">

The count of MIDI tracks in the Live Set.

- `+midiTrackCount:0` will match all documents that have no MIDI tracks.
- `+midiTrackCount:>=3 +midiTrackCount:<10` will match all documents that have between 3 and 9 MIDI tracks.
</IndexProp>

### `audioTrackCount`
 
<IndexProp name="audioTrackCount" type="number" mapping="numeric" :version="[11,10,9]">

The count of audio tracks in the Live Set.

- `+audioTrackCount:0` will match all documents that have no audio tracks.
- `+audioTrackCount:>=7 +audioTrackCount:<10` will match all documents that have between 7 and 9 audio tracks.
</IndexProp>

### `majorVersion`

<IndexProp name="majorVersion" type="string" mapping="exact" :version="[11,10,9]">

The major version of the Live Set. This seems to match against an Ableton-internal version number
maybe representing the used XML schema variant.

- `+majorVersion:5` will match all documents that have a major version of 5 (Ableton Live 10 and 11).
- `+majorVersion:4` will match all documents that have a major version of 4 (Ableton Live 9).
</IndexProp>

### `minorVersion`

<IndexProp name="minorVersion" type="string" mapping="exact" :version="[11,10,9]">

The minor version of the Live Set. This is more align towards the  end-user version number used that
produced the Live Set file.

- `+minorVersion:"11.0_11300"` will match all documents that were produced with Ableton Live Ableton Live 11.3.13.
</IndexProp>

### `creator`

<IndexProp name="creator" type="string" mapping="exact" :version="[11,10,9]">

The label of the software used to create the Live Set file. This is a end-user facing product name.

- `+creator:"Ableton Live 11.3.13"` will match all documents that were produced with Ableton Live 11.3.13.
</IndexProp>

## Produced tags

### `bpm`

<IndexTag base="bpm" value="int">

The global tempo of the given document in `{int}` BPM.
- `+tags:bpm=120` if the document has a tempo of 120 BPM.
</IndexTag>

### `ableton:version`

<IndexTag base="ableton:version" value="string">

The version of Ableton Live that was used to create the given document.
This tag will be present in multiple variants, due to the SemVer versioning scheme used by Ableton Live.

This means that an Ableton Live Version of `11.3.13` will be split into three available tags: `ableton:version=11`, `ableton:version=11.3` and `ableton:version=11.3.13`.

- `+tags:"ableton:version=11"` if the document was created with Ableton Live 11 (and any minor/patch version).
</IndexTag>

### `ableton-live-set:tracks:count`

<IndexTag base="ableton-live-set:tracks:count" value="int">

The total count of tracks (MIDI and audio) in the Live Set.

- `+tags:"ableton-live-set:tracks:count=8"` if the document has no tracks.
</IndexTag>

### `ableton-live-set:midi-tracks:available`

<IndexTag base="ableton-live-set:midi-tracks:available" value="int">

Whether the document has any audio tracks.

- `+tags:"ableton-live-set:audio-tracks:available=true"` if the document has audio tracks.
- `+tags:"ableton-live-set:audio-tracks:available=false"` if the document has no audio tracks.
</IndexTag>

### `ableton-live-set:midi-tracks:available`

<IndexTag base="ableton-live-set:midi-tracks:available" value="boolean">

Whether the document has any MIDI tracks.

- `+tags:"ableton-live-set:midi-tracks:available=true"` if the document has MIDI tracks.
- `+tags:"ableton-live-set:midi-tracks:available=false"` if the document has no MIDI tracks.
</IndexTag>

### `file:mtime-year`

<IndexTag base="file:mtime-year" value="int">

File was last modified in the year `{int}`.

- `+tags:"file:mtime-year=2023"` if the file was modified in 2023.
</IndexTag>

### `file:mtime-weekday`

<IndexTag base="file:mtime-weekday" value="int">

File was last modified on a weekday represented by `{int}`.

- `+tags:"file:mtime-weekday=0"` if the file was modified on a Sunday.
- `+tags:"file:mtime-weekday=6"` if the file was modified on a Saturday.
</IndexTag>

### `file:mtime-month`

<IndexTag base="file:mtime-month" value="int">

File was last modified in the month represented by `{int}`.

- `+tags:"file:mtime-month=1"` if the file was modified in January.
</IndexTag>

### `file:mtime-quarter`

<IndexTag base="file:mtime-quarter" value="int">

File was last modified in the quarter represented by `{int}`.

- `+tags:"file:mtime-quarter=3"` if the file was modified in the third quarter of the year.
</IndexTag>

### `file:mtime-weekno`

<IndexTag base="file:mtime-weekno" value="int">

File was last modified in the week number represented by `{int}`.

- `+tags:"file:mtime-weekno=40"` if the file was modified in the 40th week of the year.
</IndexTag>

### `file:btime-year`

<IndexTag base="file:btime-year" value="int">

File was created in the year `{int}`.

- `+tags:"file:btime-year=2023"` if the file created in 2023.
</IndexTag>

### `file:btime-weekday`

<IndexTag base="file:btime-weekday" value="int">

File was created on a weekday represented by `{int}`.

- `+tags:"file:btime-weekday=0"` if the file was created on a Sunday.
- `+tags:"file:btime-weekday=6"` if the file was created on a Saturday.
</IndexTag>

### `file:btime-month`

<IndexTag base="file:btime-month" value="int">

File was created in the month represented by `{int}`.

- `+tags:"file:btime-month=1"` if the file was created in January.
</IndexTag>

### `file:btime-quarter`

<IndexTag base="file:btime-quarter" value="int">

File was created in the quarter represented by `{int}`.

- `+tags:"file:btime-quarter=3"` if the file was created in the third quarter of the year.
</IndexTag>

### `file:btime-weekno`

<IndexTag base="file:btime-weekno" value="int">

File was created in the week number represented by `{int}`.

- `+tags:"file:btime-weekno=40"` if the file was created in the 40th week of the year.
</IndexTag>

### `file:zodiac-western`

<IndexTag base="file:zodiac-western" value="string">

The file was created in the western zodiac sign represented by `{string}`.

- `+tags:"file:zodiac-western=scorpio"` if the file was created in the Scorpio zodiac sign.
</IndexTag>

### `file:zodiac-chinese`

<IndexTag base="file:zodiac-chinese" value="string">

The file was created in the chinese zodiac sign represented by `{string}`.

- `+tags:"file:zodiac-chinese=rabbit"` if the file was created in the Rabbit zodiac sign.
</IndexTag>
