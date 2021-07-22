# cf-plugin-apps-manager

The `cf` CLI plugin to open [Apps Manager](https://docs.pivotal.io/application-service/2-11/console/dev-console.html) in browser.

## Installation

```
$ cf install-plugin https://github.com/autopp/cf-plugin-apps-manager/releases/latest/download/cf-plugin-apps-manager-${youros}-${yourarch}
```

Or download binary from [Github releases](https://github.com/autopp/cf-plugin-apps-manager/releases), then do it:
```
$ cf install-plugin /path/to/cf-plugin-apps-manager-${youros}-${yourarch}
```

Supported OS and Arch are:

- OS: `windows`, `darwin`, `linux`
- Arch: `amd64`, `arm64`

(Note: When OS is `windows`, only `amd64` is supported and filename contains `.exe` as extention)

## Usage

```
$ cf apps-manager
```

It opens

- the target space page when both organization and space are targeted
- the target organization page when only organization is targeted
- the top page otherwise

## License

[Apache License 2.0](LICENSE)
