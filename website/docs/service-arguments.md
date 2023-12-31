---
description: Service arguments that can be passed to the Ablegram service binary.
---

# Service arguments

The service can be configured using command line arguments. These arguments can be passed to the service binary when it is started.

##  Log-related flags

### `-enable-logs`

Enable logs for the service. The default is `false`.
If set to `true`, the logs will be written to a file in the to a file in the location of the service binary, i.e. `./Ablegram.runtime.log'.

### `-log-level={level}`

The `log-level` argument allows you to set the log level of the service. The default value is `info`.
The only other allowed value is `debug` which will produce more verbose logs.

### `-enable-scanned-log`

Enable the logging of scanned and processed files. The default is `false`. If set to `true`, the logs are
are written to a file in the location of the service binary, i.e. `./Ablegram.processed.log'.

This will contain additional information about the files, any processing errors and any folders skipped or ignored.

## Behavior related flags

### `-demo-mode`

Runs the service in demo mode. If given, the service will not perform any write operations,
it will not write any files to disk and will not save any changed settings.

### `-no-browser`

When the service is started, it will automatically open the default browser and navigate to the search frontend.
This behaviour can be disabled by passing the `no-browser` flag.

### `-no-gui`

When the service is started, it will automatically open the service GUI.
It will show the current progress and options to open the search frontend in a browser, or to stop the service altogether.
Passing this flag will not open the GUI.

### `-no-webservice`

When the service is started, it will automatically start a web server, providing all API and websocket endpoints.
It will try to bind to the following ports, in the order given

- 10000
- 20000
- 30000
- 40000
- 50000
- 10001

It will then give up and exit with an error.

This flag will disable all web endpoints and only start the file parser and service GUI. This is mainly used for and debugging purposes.

## Indexer flags

The indexer is responsible for moving the generated tags and possible search results into the search index.

### `-indexer-worker-delay={number}`

By default, the indexer threads do not set a delay between payloads.
This means that the indexer will continuously move possible payloads into the search index.
The specified delay is defined in milliseconds.

This introduces a pause between indexing, which can be helpful if you have a slower system.

## Webservice flags

The webservice is responsible for providing the API and websocket endpoints.

### `-master-password={password}`

The master password is used to protect the API endpoints. It is used to authenticate admins.

### `-trusted-platform={header}`

The service will attempt to identify the client's IP address for guest access.

If the service is running behind a reverse proxy, the client's IP address will most likely be in a separate header.

The service should handle most common `X-Forwarded-For` headers, but if you are using a different header, you can specify it using this flag.

For example, if you need to host the service behind [Cloudflare](https://developers.cloudflare.com/fundamentals/reference/http-request-headers/#cf-connecting-ip), you could select their `CF-Connecting-IP` header.

For [Google App Engine](https://cloud.google.com/appengine/docs/flexible/reference/request-headers#app_engine-specific_headers), you could select their `X-Appengine-User-Ip` header.

This will allow the owning guest and all Admins to see the associated IP address of the Client.
Guests will not be able to see the IP address of other guests.
