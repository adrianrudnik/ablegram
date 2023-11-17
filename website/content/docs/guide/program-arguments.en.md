---
title: "Program arguments"
description: ""
summary: ""
date: 2023-11-17T21:58:51+01:00
lastmod: 2023-11-17T21:58:51+01:00
menu:
  docs:
    parent: ""
    identifier: "program-arguments-91c496698372cb851bf88908b4893a2f"
weight: 60
toc: true
seo:
  title: "" # custom title (optional)
  description: "" # custom description (recommended)
  canonical: "" # custom canonical URL (optional)
  noindex: false # false (default) or true
---

## Log related flags

### `-enable-logs`

Enable the logs of the service. The default value is `false`. If set to `true`, the logs will be written
to a file at the location the service binary is located, i.e. `./Ablegram.runtime.log`.

### `-log-level={level}`

The `log-level` argument allows you to set the log level of the service. The default value is `info`.
The only other allowed value is `debug` which will produce more verbose logs.

### `-enable-scanned-log`

Enable the logs of the scanned and processed files. The default value is `false`. If set to `true`, the logs will
be written to a file at the location the service binary is located, i.e. `./Ablegram.processed.log`.

This will contain additional information about the files, possible processing errors and skipped or ignored folders.

## Behavior related flags

### `-demo-mode`

Will start the service in demo mode. If given, the service will not execute any writing operations,
it will not write any files to disk and not store changed settings.

### `-no-browser` 

When the service is started, it will automatically open the default browser and navigate to the search frontend.
This behavior can be disabled by passing the `no-browser` flag.

### `-no-gui`

When the service is started, it will automatically open the service GUI. It will show the current progress and options
to open the search frontend in a browser or stop the service all together. By passing this flag, the GUI
will not be opened.

### `-no-webservice`

When the service is started, it will automatically start a webserver, offering all API and websocket endpoints.
It will try to bind itself against the following ports, in the given order:

- 10000
- 20000
- 30000
- 40000
- 50000
- 10001

After that, it will give up and exit with an error.

This flag will disable all web endpoints and just start the file parser and service GUI. This is mostly used for
development and debugging purposes.

## Collector flags

The collector is responsible for finding all files that can be processed on your system.

### `-collector-worker-count={number}`

By default, `5` worker threads are spawned per collector. This means that `5` threads will
process possible file locations in parallel. By passing a different number into this flag, the number
of worker threads can be changed. This could be helpful if you have a lot of files and folders to process
on slower hardware.

### `-collector-worker-delay={millis}`

By default, no delay is set between finding files and folders. This means that all worker threads continuously
work through the file system to locate more possible files to send over to the processors.
The given delay is defined in milliseconds.

This will introduce a pause between payloads, which could be helpful if you have a slower system.

### `-parser-worker-count={number}`

By default, `5` worker threads are spawned per collector. This means that `5` threads will
process found files in parallel. By passing a different number into this flag, the number
of worker threads can be changed. This could be helpful if you have a lot of files to process
on slower hardware, that does drown in random access operations.

### `-parser-worker-delay={millis}`

By default, no delay is set between payloads the worker threads process. This means that all worker threads
continuously process files. The given delay is defined in milliseconds.

This will introduce a pause between payload processing, which could be helpful if you have a slower system.

## Indexer flags

The indexer is responsible for moving produced tags and possible search results into the search index.

### `-indexer-worker-delay={number}`

By default, no delay is set between payloads the indexer threads. This means that the indexer will continuously
move possible payloads into the search index. The given delay is defined in milliseconds.

This will introduce a pause between indexing, which could be helpful if you have a slower system.
