# unofficial Pilw.io CLI

## Local installation
- Install Go 1.12
- Clone this repo
- Run `make install clean` to build `pilw` into `/usr/local/bin`

## Usage
Uses `PILW_API_KEY` from environment. Also attempts to load `.env`

For most listing actions, using `-q` or `--quiet` will restrict output to unique identifiers, like the docker CLI.

Args written as `[key=val]` means they are optional, and you can specify your own value by adding `--key my_val`, defaulting to `val`.
Flags written as `[key=?]` mean they are optional, and won't be used/included unless you specify your own value by adding `--key my_val`.

Supported commands:
- `pilw` root command, shows help
- `pilw user` root for user actions, shows help
  - `pilw user info` Displays authenticated user info
- `pilw token` root for token actions, shows help
  - `pilw token list` lists tokens
  - `pilw token create [description] [restricted=false] [billing_account_id=0]` creates a new token
  - `pilw token delete [token_id]` deletes an existing token
  - `pilw token update [token_id] [restricted=?] [billing_account_id=?]` updates an existing token
- `pilw vm` root for vm actions, shows help
  - `pilw vm list` lists VMs
  - `pilw vm update [uuid] [name=?] [ram=?] [vcpu=?]` updates an existing VM
- `pilw container` root for container actions, shows help
  - `pilw container group` root for container group actions, shows help
    - `pilw container group list` lists container groups
  - `pilw container service` root for container service actions, shows help
    - `pilw container service list [name=?]` lists container services, optionally filtered by name
    - `pilw container service restart [suuid]` stops, then starts a container service

Examples:

Restarting a container service by its name (example-service):
- `pilw container service restart $(pilw container service list -q --name example-service)`

As part of a drone CI docker pipeline step:
```
  - name: deploy
    image: amperl/pilw
    environment:
      PILW_API_KEY:
        from_secret: pilw_key
    commands:
      - pilw container service restart $(pilw container service list -q --name example-service)
```