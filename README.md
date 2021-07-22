# cf-plugin-apps-manager

The `cf` CLI plugin to open [Apps Manager](https://docs.pivotal.io/application-service/2-11/console/dev-console.html) in browser.

## Installation

Download binary from Github releases and unarchive, then do it:
```
$ cf install-plugin /path/to/cf-plugin-apps-manager
```

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
