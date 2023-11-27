---
description: API endpoints for connected user management.
---

# Users and clients API

The following endpoints are available to work with connected users and their associated users.

## `/api/users`

Returns a list of users.
Requires admin privileges.
Every user is represented by a unique ID and an either given or self-configured display name.
The list only contains users that have at least a single client connected.

```json
[
  {
    "id": "80d8015d-3ec2-4227-a44e-7b6454c927a7",
    "display_name": "Guest",
    "role": "guest"
  },
  {
    "id": "2e70ff7b-0c86-4913-9dee-3d3f4a601acc",
    "display_name": "Admin",
    "role": "admin"
  }
]
```

## `/api/clients`

Returns a list of connected clients.
Requires admin privileges.
Every client is represented by a unique ID and their technical details, as well as the user they are connected to.
The list only contains clients that are currently connected to the server.

```json
[
  {
    "id": "290319dc-2089-4496-a5bf-555b7f23931f",
    "ip": "::1",
    "user_id": "80d8015d-3ec2-4227-a44e-7b6454c927a7"
  },
  {
    "id": "e6207fb4-4f5f-45ad-baad-9b7fd6d1bcb9",
    "ip": "127.0.0.1",
    "user_id": "2e70ff7b-0c86-4913-9dee-3d3f4a601acc"
  },
  {
    "id": "127b1c53-387e-4bcf-b7e1-e7ff3a2fab95",
    "ip": "::1",
    "user_id": "80d8015d-3ec2-4227-a44e-7b6454c927a7"
  }
]
```
