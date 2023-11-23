---
description: A working read-only demo of a selection of Ableton projects is available online for you to try out.
---

# Live Demo

A working read-only demo of a selection of Ableton projects is available online.

::: info Access to the live demo
[https://demo.ablegram.app/](https://demo.ablegram.app/)

To authenticate as admin, use the password `admintest`
:::

The demo itself is not a special build of Ablegram, but the same version available for download.

It runs in an isolated container containing some sample project files. It is started with the following arguments

```bash
ablegram \
  -demo-mode \
  -no-gui \
  -no-browser \
  -master-password admintest \
  -log-level debug
```
