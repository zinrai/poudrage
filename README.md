# poudrage

poudrage is a CLI tool for declarative package building with FreeBSD [poudriere](https://man.freebsd.org/cgi/man.cgi?poudriere).

## Features

- Declarative package build environment setup using YAML
- Manage poudriere jail creation/update
- Manage poudriere ports tree creation/update
- Configure package build options
- Manage make.conf settings
- Automatic setname configuration based on YAML filename

## Requirements

- poudriere
- git
- sudo ( or doas )

## Installation

To build poudrage from source:

```sh
$ GOOS=freebsd GOARCH=amd64 go build -o poudrage cmd/poudrage/main.go
```

## Usage

See `examples/build.yaml` for settings.

Available commands:

- `setup`: Create jail, ports tree, and configure build options
- `validate`: Validate configuration file
- `build`: Build packages
- `update`: Update jail and ports tree

Example:

Setup build environment:

```sh
$ sudo poudrage setup -c build.yaml
```

Build packages:

```sh
$ sudo poudrage build -c build.yaml
```

Update jail and ports tree:

```sh
$ sudo poudrage update -c build.yaml
```

The setname is automatically derived from the configuration filename:

- `build.yaml` -> setname: "build"
- `build_env.yaml` -> setname: "build"
- `build-env.yaml` -> setname: "build"

## Configuration

### Jail Configuration

The jail section specifies FreeBSD version and architecture:

```yaml
environment:
  jail:
    version: "14.2-RELEASE"  # FreeBSD version
    arch: "amd64"            # Architecture
```

### make.conf Settings

Global make.conf settings can be specified:

```yaml
environment:
  make.conf: |
    OPTIONS_SET+=DOCS
    DEFAULT_VERSIONS+=ssl=openssl
```

### Package Options

Package-specific build options can be configured:

```yaml
packages:
  - name: "www/nginx"
    options: |
      OPTIONS_FILE_UNSET+=HTTP_PERL
      OPTIONS_FILE_SET+=HTTP_RANDOM_INDEX
      OPTIONS_FILE_SET+=HTTP_REALIP
      OPTIONS_FILE_SET+=HTTP_SECURE_LINK
      OPTIONS_FILE_SET+=HTTP_SLICE
      OPTIONS_FILE_SET+=HTTP_SSL
```

## File Locations

- make.conf: `/usr/local/etc/poudriere.d/<jail>-<ports>-<setname>-make.conf`
- Package options: `/usr/local/etc/poudriere.d/<jail>-<ports>-<setname>-options`
- Distfiles: Uses DISTFILES_CACHE directory from poudriere.conf

The options file will be created with comments indicating which package each set of options belongs to:

```
# www/nginx
OPTIONS_FILE_UNSET+=HTTP_PERL
OPTIONS_FILE_SET+=HTTP_RANDOM_INDEX
...

# databases/postgresql15-server
OPTIONS_FILE_SET+=SSL
...
```

## License

This project is licensed under the MIT License - see the [LICENSE](https://opensource.org/license/mit) for details.
