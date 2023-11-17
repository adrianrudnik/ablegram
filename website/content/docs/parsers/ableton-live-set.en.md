---
title: "Ableton Live Set"
description: ""
summary: ""
date: 2023-11-13T18:39:09+01:00
lastmod: 2023-11-13T18:39:09+01:00
menu:
  docs:
    parent: ""
    identifier: "ableton-live-set-85336a01a42970bd27f4a34a83188c6e"
weight: 999
toc: true
seo:
  title: "" # custom title (optional)
  description: "" # custom description (recommended)
  canonical: "" # custom canonical URL (optional)
  noindex: false # false (default) or true
---

The parser for `.als` files will extract the most common metadata and add it to the search index as `AbletonLiveSet` type.

{{< callout context="info" title="Note" icon="info-circle" >}}
This document is not only about the file itself. The roots are based on the primary XML node within an `.als` file. This means that the document is a combination of file and meta related information.
{{< /callout >}}

## Fields

{{% inc_md "shared/base.mdinc" "AbletonLiveSet" %}}
{{% inc_md "shared/annotation.mdinc" "AbletonLiveSet" %}}
{{% inc_md "shared/scale-info.mdinc" "Live Set" %}}

{{% ds/prop name="bpm" type="number" mapping="numeric" %}}
  The global tempo of the Live Set in BPM.
  - `+bpm:>=120` will match all documents that have a tempo of 120 BPM or higher.
  - `+bpm:<120` will match all documents that have a tempo of less than 120 BPM.
{{%/ ds/prop %}}

{{% ds/prop name="midiTrackCount" type="number" mapping="numeric" %}}
  The count of MIDI tracks in the Live Set.
  - `+midiTrackCount:0` will match all documents that have no MIDI tracks.
  - `+midiTrackCount:>=3 +midiTrackCount:<10` will match all documents that have between 3 and 9 MIDI tracks.
{{%/ ds/prop %}}

{{% ds/prop name="audioTrackCount" type="number" mapping="numeric" %}}
  The count of audio tracks in the Live Set.
  - `+audioTrackCount:0` will match all documents that have no audio tracks.
  - `+audioTrackCount:>=7 +audioTrackCount:<10` will match all documents that have between 7 and 9 audio tracks.
{{%/ ds/prop %}}

{{% ds/prop name="majorVersion" type="string" mapping="exact" %}}
  The major version of the Live Set. This seems to match against an Ableton-internal version number
  maybe representing the used XML schema variant.
  - `+majorVersion:5` will match all documents that have a major version of 5 (Ableton Live 10 and 11).
{{%/ ds/prop %}}

{{% ds/prop name="minorVersion" type="string" mapping="exact" %}}
  The minor version of the Live Set. This is more align towards the  end-user version number used that
  produced the Live Set file.
  - `+minorVersion:"11.0_11300"` will match all documents that were produced with Ableton Live Ableton Live 11.3.13.
{{%/ ds/prop %}}

{{% ds/prop name="creator" type="string" mapping="exact" %}}
  The label of the software used to create the Live Set file. This is a end-user facing product name. 
  - `+creator:"Ableton Live 11.3.13"` will match all documents that were produced with Ableton Live 11.3.13.
{{%/ ds/prop %}}

## Produced tags

{{% ds/tag root="bpm" value="int" %}}
  The global tempo of the given document in `{int}` BPM.
  - `bpm=120` if the document has a tempo of 120 BPM.
{{%/ ds/tag %}}

{{% ds/tag root="ableton:version" value="string" %}}
  The version of Ableton Live that was used to create the given document.
  This tag will be present in multiple variants, due to the SemVer versioning scheme used by Ableton Live.
  
  This means that an Ableton Live Version of `11.3.13` will be split into three available tags: `ableton:version=11`, `ableton:version=11.3` and `ableton:version=11.3.13`.

  - `ableton:version=11` if the document was created with Ableton Live 11 (and any minor/patch version).
{{%/ ds/tag %}}

{{% ds/tag root="ableton-live-set:tracks:count" value="numeric" %}}
  The total count of tracks (MIDI and audio) in the Live Set.
  - `ableton-live-set:tracks:count=8` if the document has no tracks.
{{%/ ds/tag %}}

{{% ds/tag root="ableton-live-set:audio-tracks:available" value="boolean" %}}
  Whether the document has any audio tracks.
  - `ableton-live-set:audio-tracks:available=true` if the document has audio tracks.
  - `ableton-live-set:audio-tracks:available=false` if the document has no audio tracks.
{{%/ ds/tag %}}

{{% ds/tag root="ableton-live-set:midi-tracks:available" value="boolean" %}}
  Whether the document has any MIDI tracks.
  - `ableton-live-set:midi-tracks:available=true` if the document has MIDI tracks.
  - `ableton-live-set:midi-tracks:available=false` if the document has no MIDI tracks.
{{%/ ds/tag %}}

{{% ds/tag root="file:mtime-year" value="int" %}}
  File was last modified in the year `{int}`.
  - `file:mtime-year=2023` if the file was modified in 2023.
{{%/ ds/tag %}}

{{% ds/tag root="file:mtime-weekday" value="int" %}}
  File was last modified on a weekday represented by `{int}`.
  - `file:mtime-weekday=0` if the file was modified on a Sunday.
  - `file:mtime-weekday=6` if the file was modified on a Saturday.
{{%/ ds/tag %}}

{{% ds/tag root="file:mtime-month" value="int" %}}
  File was last modified in the month represented by `{int}`.
  - `file:mtime-month=1` if the file was modified in January.
{{%/ ds/tag %}}

{{% ds/tag root="file:mtime-quarter" value="int" %}}
  File was last modified in the quarter represented by `{int}`.
  - `file:mtime-quarter=3` if the file was modified in the third quarter of the year.
{{%/ ds/tag %}}

{{% ds/tag root="file:mtime-weekno" value="int" %}}
  File was last modified in the week number represented by `{int}`.
  - `file:mtime-weekno=40` if the file was modified in the 40th week of the year.
{{%/ ds/tag %}}

{{% ds/tag root="file:btime-year" value="int" %}}
  File was created in the year `{int}`.
  - `file:btime-year=2023` if the file created in 2023.
{{%/ ds/tag %}}

{{% ds/tag root="file:btime-weekday" value="int" %}}
  File was created on a weekday represented by `{int}`.
  - `file:btime-weekday=0` if the file was created on a Sunday.
  - `file:btime-weekday=6` if the file was created on a Saturday.
{{%/ ds/tag %}}

{{% ds/tag root="file:btime-month" value="int" %}}
  File was created in the month represented by `{int}`.
  - `file:btime-month=1` if the file was created in January.
{{%/ ds/tag %}}

{{% ds/tag root="file:btime-quarter" value="int" %}}
  File was created in the quarter represented by `{int}`.
  - `file:btime-quarter=3` if the file was created in the third quarter of the year.
{{%/ ds/tag %}}

{{% ds/tag root="file:btime-weekno" value="int" %}}
  File was created in the week number represented by `{int}`.
  - `file:btime-weekno=40` if the file was created in the 40th week of the year.
{{%/ ds/tag %}}

{{% ds/tag root="file:zodiac-western" value="string" %}}
  The file was created in the western zodiac sign represented by `{string}`.
  - `file:zodiac-western=scorpio` if the file was created in the Scorpio zodiac sign.
{{%/ ds/tag %}}

{{% ds/tag root="file:zodiac-chinese" value="string" %}}
  The file was created in the chinese zodiac sign represented by `{string}`.
  - `file:zodiac-chinese=rabbit` if the file was created in the Rabbit zodiac sign.
{{%/ ds/tag %}}
