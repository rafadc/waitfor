# waitfor

waitfor is a tool to simply waiting for some standard services to be up.

## Usage

```shell
Usage:
  waitfor [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  http        Waits for an HTTP call to be successful
  kafka       Waits for a successful connection to Kafka.
  version     Print the version number of waitfor

```

## Commands

### Global flags

```shell
Global Flags:
  -d, --delay duration   delay in seconds between retries (default 5ns)
  -r, --retries uint     number of retries (default 5)
  -v, --verbose          verbose output
```

### HTTP

Waits for an HTTP call to return a 200 status code.

```shell
Usage:
  waitfor http URL [flags]
```

### Kafka

Waits for a successful connection to Kafka.

```shell
Usage:
  waitfor kafka KAFKA_URL [flags]
```

