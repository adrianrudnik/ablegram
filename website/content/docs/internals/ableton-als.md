---
title: "Ableton .ALS"
description: ""
summary: ""
date: 2023-11-18T14:21:10+01:00
lastmod: 2023-11-18T14:21:10+01:00
menu:
  docs:
    parent: ""
    identifier: "ableton-als-5ea9576331b3cbd250dfb2c99f9357ba"
weight: 999
toc: true
seo:
  title: "" # custom title (optional)
  description: "" # custom description (recommended)
  canonical: "" # custom canonical URL (optional)
  noindex: false # false (default) or true
---

This page details on the findings in analyzing the Ableton .ALS file format.

## Packaging

The `.als` file is a gzip archive containing a single XML file.

We can decompress it via 

```shell
zcat input.als > output.xml
```

This was used to sample Live Sets for the versions 9.7.7, 10.1.43, 11.3.13.

## Samples

Sample for Ableton v9:

```xml
<Ableton MajorVersion="4" MinorVersion="9.5_326" Creator="Ableton Live 9.5b14" Revision="dd97cfbe7b78bb6353aa5bb2e7d0674b8831d11e">
  <LiveSet><!-- ... --></LiveSet>
</Ableton>
```

Sample for Ableton v10:

```xml
<Ableton MajorVersion="5" MinorVersion="10.0_377" SchemaChangeCount="3" Creator="Ableton Live 10.1.3d1" Revision="29bb8f2920bd08036a270fd62ef54610b5b8d4fd">
  <LiveSet><!-- ... --></LiveSet>
</Ableton>
```

Sample for Ableton v11:

```xml
<Ableton MajorVersion="5" MinorVersion="11.0_11300" SchemaChangeCount="3" Creator="Ableton Live 11.3d1" Revision="539c1d1e226a2160a7f80617f3a7a5debcee8bb4">
  <LiveSet><!-- ... --></LiveSet>
</Ableton>
```
