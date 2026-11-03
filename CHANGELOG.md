# Changelog -- nx-common

## v0.1.0 -- 2026-11-03

- Add `httplog`: one access line per HTTP request on stderr (method, host, path, status, duration), never the query, headers or bodies; the status writer keeps Flusher, Hijacker and Pusher.
