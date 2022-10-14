# go-apparmor

Is an initial draft of how apparmor profiles could be managed in golang with a reduced attack surface.
It leverages `libapparmor` to effectively load profiles into the kernel, as well as deleting them too.

`libapparmor` currently does not provide the parsing of plain-text profiles, but rather requires them 
to be in binary format. The `apparmor_parser` (from `apparmor-utils`) is leveraged to make that
conversion, and this operation takes place at lower privilege mode.

## Security Context

When running inside a container, the library will automatically attempt to "privilege escalate" into
the host's mount namespace just for load/delete operations, and then revert back. However, it would
require `HostPID` and run as `privileged` from the get go.

Permissions required:

- Host's PID namespace
- Host's Mount namespace
- run as root
- `CAP_SYS_ADMIN`
- privileged (for containers)

Running directly on the host machine as `root` is enough.

## Next steps

- [] Implement Enforceable() and a func to check whether a profile is already loaded.
- [] Split hostop package from this repo.
- [x] Refactor apparmor package.
- [x] Add some tests.
