---
description: Details about the Push 3 Web API and available endpoints.
---

# Push 3 Standalone Web API

This chapter documents the findings in the exploration process against my Push 3 standalone device.

The devices NIC comes with a MAC address associated with the manufacturer `Intel Corporate` with the OUI `B0:A4:60`.

It will try to register itself as `push.local` via DHCP in the local network, but might fail.
For IPv4 it might work, for IPv6 it will not.

If you are unsure what IP address the device has, try to ping through different protocols and DNS zones:

```shell
ping -4 push.local
ping -6 push.local
ping -4 push
ping -6 push
ping -4 push.fritz.box
ping -6 push.fritz.box
```

The following TCP ports are open:

| Port | Protocol | Description                      |
|------|----------|----------------------------------|
| 22   | SSH      | SSH access, ED25519 fingerprint. |
| 80   | HTTP     | HTTP access                      |

## General endpoints

A collection of general endpoints offering more information about the device.

### Device ID

This seems to be the endpoint to identify a specific Push 3 device by an unique ID:

```http
GET /id HTTP/1.1
User-Agent: Live 11.3.13

HTTP/1.1 200 OK
Content-Length: 32
Content-Type: application/octet-stream

68d0163052ab40c4a76681627faef81a
```

It does not require any authentication and returns a 32 byte long string.

### UI web endpoints

http://push.local/legal

Contains a link to legal documents required to be shipped with the standalone device.

http://push.local/pair

General pairing web UI flow entrypoint.

http://push.local/crash

Download endpoint that serves all relevant crash logs in a ZIP bundle.

http://push.local/logs

Download endpoint that services logs in a ZIP bundle.

http://push.local/authorize

Authorize the Ableton Live license against the Push 3 device.

## Access control

The access control is based ona simple challenge-response mechanism. It requires user interaction on the device itself.


### Ableton-Challenge-Response-Token

Referenced as `ACRT` in the following sections.

The token is generally served as a cookie against the API.
It is sent with `credentials: "include"` via
the [Fetch API](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch#sending_a_request_with_credentials_included)
to ensure requests are authenticated.

The validity of the token is longer then the devices up-time and will survive a reboot. It seems to be valid for at least 24 hours.

The cookie settings are bound to `path=/; HttpOnly; SameSite=Strict`.

Example cookie:

```
Ableton-Challenge-Response-Token: 240a3e5f4843b530adc4faa0aaf5f37913435eea3890952e6f6bed5e84a9ba06
```

### Check authorized access

The endpoint itself is an easy way to check if a token is valid against a given Push 3 device, for privileged endpoints.

```http
GET {{push}}/api/v1/access-allowed
Cookie: Ableton-Challenge-Response-Token=ACRT
```

Allowance is indicated by a `200 OK` status code.

Failure is indicated by a `401 Unauthorized` status code.

### Start challenge flow

The challenge flow is the way to obtain a valid `ACRT` token. It will trigger the PIN display on the Push 3 device.

```http
POST /api/v1/challenge

HTTP/1.1 200 OK
```

The shown PIN has to be entered into the [Finish challenge flow](#finish-challenge-flow).

Triggering this flow in succession, while a PIN is shown, has no effect.

### Finish challenge flow

To finish the challenge flow, the displayed PIN from the [Start challenge flow](#start-challenge-flow) has to be sent
back to the device.

```http
POST /api/v1/challenge-response
Content-Type: application/json

{
  "secret": "123456"
}
```

Failure to submit the correct code will result in a `401 Unauthorized` status code and a `{"error": "invalid code"}`
response body.

The response will also contain a `X-Retries-Left:` header transporting the amount of retries left.

Success will result in a `200 OK` status code and a `{"token": "ACRT"}` response body.

The response will also contain a `Set-Cookie: Ableton-Challenge-Response-Token=ACRT` header transporting the `ACRT`
token, more information about the token can be found in
the [Ableton-Challenge-Response-Token](#ableton-challenge-response-token) section.

## Firmware

A collection of endpoints related to firmware updates and required processes.

### Upload and update firmware

The firmware update upload endpoint is used to upload a firmware update to the device.

It requires an update file `.swu` by the [SWUpdate framework](https://sbabic.github.io/swupdate/overview.html), which is
a [CPIO](https://linux.die.net/man/1/cpio) archive.

```http
POST /api/v1/update

Content-Type: multipart/form-data; boundary=boundary

< /. https://cdn-hardware-updates.ableton.com/aks-hardwareupdates-files/update-server/artifacts/push3-release/1.1.13/update.swu
```

Response is not yet known, I'm not keen on debugging this yet.

How to obtain the list of available firmware updates is documented in [List firmware updates](#list-firmware-updates) section.

### List firmware updates

The list of available firmware updates is served by the following website:

https://hardware-updates.ableton.com/api/v1/update/push3-release/

The response contains a list of available firmware updates, including their version, creation date and a list of files
that are part of the update.

```json
[
  {
    "id": 155,
    "version": "1.1.13",
    "created": "2023-10-25T03:52:46.679746",
    "product": "push3-release",
    "mandatory": false,
    "updatefiles": []
  }
]
```

### Check for firmware update

The Push 3 will check its own version against the following endpoint:

https://hardware-updates.ableton.com/api/v1/update/push3-release/1.1.11/

If no updates are available, the response will be a `404 Not Found` status code with a payload of

```json
{
  "detail": "No suitable update found."
}
```

If an update is available, the response will be a `200 OK` status code with a payload of

```json
{
  "id": 155,
  "version": "1.1.13",
  "created": "2023-10-25T03:52:46.679746",
  "product": "push3-release",
  "mandatory": false,
  "updatefiles": [
    {
      "id": 293,
      "path": "push3-release/1.1.13/update.swu",
      "url": "https://cdn-hardware-updates.ableton.com/aks-hardwareupdates-files/update-server/artifacts/push3-release/1.1.13/update.swu",
      "filename": "update.swu",
      "update_rel": {
        "id": 155,
        "version": "1.1.13",
        "created": "2023-10-25T03:52:46.679746",
        "product": "push3-release",
        "mandatory": false
      }
    }
  ]
}
```

### Device reboot

A device reboot can be triggered by using an endpoint exposed by the update process:

```http request
POST /api/v1/update/reboot
Cookie: Ableton-Challenge-Response-Token=ACRT
```

On success the response will be a `200 OK` status code with no payload. The device should reboot after a few seconds.

On failure, the response will be a `401 Unauthorized` status code with a payload similar to

```json
{
  "error": "Unset credentials"
}
```

## Asset file system

The Push 3 offers endpoints to access to all user content available on the device.

These endpoints are primarily used by the Ableton Live browser to access and manage the content.

### Browse content

The endpoint to browse the content is the following:

```http
GET /api/v1/files/
Content-Type: application/json
Cookie: Ableton-Challenge-Response-Token=ACRT

HTTP/1.1 200 OK
Content-Type: application/json

{
  "paths": [
    {
      "path": "Sets",
      "isDirectory": true,
      "lastModifiedDateTime": "2023-10-24T18:33:52Z"
    },
    {
      "path": "Factory%20Packs",
      "isDirectory": true,
      "lastModifiedDateTime": "2023-10-22T20:14:04Z"
    },
    {
      "path": "Live%20Recordings",
      "isDirectory": true,
      "lastModifiedDateTime": "2023-11-20T00:57:23Z"
    },
    {
      "path": "User%20Library",
      "isDirectory": true,
      "lastModifiedDateTime": "2023-10-22T09:29:25Z",
      "isLiveProject": true
    },
    {
      "path": "Crashes",
      "isDirectory": true,
      "lastModifiedDateTime": "2023-11-20T00:57:07Z"
    },
    {
      "path": "AllCrashes",
      "isDirectory": true,
      "lastModifiedDateTime": "2023-10-21T13:54:40Z"
    }
  ]
}
```

::: info
The trailing slash is required to retrieve the content of the root-folder. On sub-folders it is optional.
:::

Sub-folders can be queried by appending the folder name to the endpoint, like `/api/v1/files/User%20Library`.

Files have a slightly different structure, and are not serving the `isDirectory` property (instead of an expected `false` value):

```json
{
  "path": "Juicy%20Tuesday.als",
  "size": 758460,
  "lastModifiedDateTime": "2023-10-22T09:30:21Z"
}
```

### Downloading files

Files can be downloaded by using the same endpoint as for browsing the content, but with a file as target:

```http
GET /api/v1/files/Sets/Juicy%20Tuesday.als
Cookie: Ableton-Challenge-Response-Token=ACRT

HTTP/1.1 200 OK
Content-Disposition: attachment
Content-Length: 758460
Content-Type: application/octet-stream
```

### Uploading files

Files can be uploaded by using the same endpoint as for browsing the content, but with a `POST` method.

```http
POST /api/v1/files/User%20Library/Clips?overwrite=true
Cookie: Ableton-Challenge-Response-Token=ACRT
Content-Length: 184460
Content-Type: multipart/form-data; boundary=

HTTP/1.1 200 OK
Content-Length: 71
Content-Type: application/json
```

The query parameter `overwrite` is optional and will overwrite the file if set to `true`.

::: danger
Further research required:

- How to set the filename (should be Content-Disposition)?
- What is the response?
- What is the possible error when overwrite is `false`?
:::

### Deleting files / folders

Files and folders can be deleted by using the same endpoint as for browsing the content, but with a `DELETE` method.

```http
DELETE /api/v1/files/Sets/Juicy%20Tuesday.als
Cookie: Ableton-Challenge-Response-Token=ACRT

HTTP/1.1 200 OK
```

The response is empty.

This endpoint works for files and folders alike.

### Move files / folders

Files and folders can be moved by using the same endpoint as for browsing the content, but with a `PATCH` method.

The body must be the request payload indicating the new location.

```http
PATCH /api/v1/files/Sets/Halting%20the%20Forest%20Project/Halting%20the%20Forest.als 
Cookie: Ableton-Challenge-Response-Token=ACRT
Content-Length: 65
Content-Type: application/x-www-form-urlencoded

{"path":"Sets/Chrome Orientation Project/Halting the Forest.als"}

HTTP/1.1 200 OK
```

The response is empty on success.

On failure, the response will be a `400 Bad Request` status code with a payload similar to

```json
{
    "error": "Couldn't move file or directory. Path \"Sets/Halting the Forest Project/Halting the Forest.als\" already exists."
}
```
