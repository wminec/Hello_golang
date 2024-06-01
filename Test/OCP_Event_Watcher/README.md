# OCP Event Watcher
This app collects events occurring in a specific namespace and outputs them to stdout.

The following environment variables must be set:
1. NAMESPACE : namespace what collects events

The following permissions are required:
* "get, watch, list" on events object