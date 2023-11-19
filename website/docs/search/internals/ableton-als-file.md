---
description: Exploring the Ableton .ALS File Format
---

# Ableton .ALS File Format

This document details about the exploration of the `.als` file format.

## Decompression

The `.als` file is a gzip archive containing a single XML file.

We can decompress it using

```shell
zcat input.als > output.xml
```

This was used to sample Live Sets for the versions 9.7.7, 10.1.43, 11.3.13.

## Samples

Example for Ableton v9:

```xml
<Ableton MajorVersion="4" MinorVersion="9.5_326" Creator="Ableton Live 9.5b14" Revision="dd97cfbe7b78bb6353aa5bb2e7d0674b8831d11e">
  <LiveSet><!-- ... --></LiveSet>
</Ableton>
```

Example for Ableton v10:

```xml
<Ableton MajorVersion="5" MinorVersion="10.0_377" SchemaChangeCount="3" Creator="Ableton Live 10.1.3d1" Revision="29bb8f2920bd08036a270fd62ef54610b5b8d4fd">
  <LiveSet><!-- ... --></LiveSet>
</Ableton>
```

Example for Ableton v11:

```xml
<Ableton MajorVersion="5" MinorVersion="11.0_11300" SchemaChangeCount="3" Creator="Ableton Live 11.3d1" Revision="539c1d1e226a2160a7f80617f3a7a5debcee8bb4">
  <LiveSet><!-- ... --></LiveSet>
</Ableton>
```
