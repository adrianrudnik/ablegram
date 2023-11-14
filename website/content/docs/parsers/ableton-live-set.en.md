---
title: "Ableton Live Set"
description: ""
summary: ""
date: 2023-11-13T18:39:09+01:00
lastmod: 2023-11-13T18:39:09+01:00
draft: true
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

## Index mapping

| property        | type     | example | mapping  |
|-----------------|----------|---------|----------|
| type            | string   |         | exact    |
| tags            | []string |         | exact    |
| displayName     | string   |         | fulltext |
| pathAbsolute    | string   |         | exact    |
| pathFolder      | string   |         | exact    |
| filename        | string   |         | exact    |
| annotation      | string   |         | fulltext |
| scaleRootNote   | string   |         | exact    |
| scaleName       | string   |         | exact    |
| majorVersion    | string   |         |          |
| minorVersion    | string   |         |          |
| creator         | string   |         |          |
| revision        | string   |         |          |
| tempo           | int      |         |          |
| midiTrackCount  | int      |         |          |
| audioTrackCount | int      |         |          |

## Produced tags

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

| tag                                            | description |
|------------------------------------------------|-------------|
| file:btime-year={int}                          |             |
| file:btime-weekday={int}                       |             |
| file:btime-month={int}                         |             |
| file:btime-quarter={int}                       |             |
| file:btime-weekno={int}                        |             |
| file:zodiac-western={string}                   |             |
| file:zodiac-chinese={string}                   |             |
| bpm={int}                                      |             |
| ableton-live-set:tracks:count={int}            |             |
| ableton-live-set:audio-tracks:available={bool} |             |
| ableton-live-set:midi-tracks:available={bool}  |             |
| ableton:version={string}                       |             |
