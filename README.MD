# Soar Helper

A remote agent for executing actions for soar flow

## Run

Using docker is the most simple way.

```
docker run --rm ghcr.io/yoshino-s/soar-helper:snapshot
```

## Config

You can set all config from:

1. cmdline arguments.

```
Serve runs the HTTP and gRPC server.

Usage:
  soar-helper serve [flags]

Flags:
      --beian.miit_sign string                     sign token for miit
      --beian.werplus_key string                   werplus key
  -h, --help                                       help for serve
      --http.addr string                           http listen address (default ":8080")
      --http.behind_proxy                          http behind proxy
      --http.external_url string                   external url (default "http://127.0.0.1:8080")
      --http.feature uint16                        http feature (default 30)
      --http.log                                   enable http log
      --s3.accelerate_endpoint string              s3 accelerate endpoint
      --s3.access_key_id string                    s3 access key id
      --s3.bucket string                           s3 bucket
      --s3.endpoint string                         s3 endpoint
      --s3.insecure                                s3 insecure
      --s3.presigned_get_object_expires duration   s3 presigned get object expires (default 168h0m0s)
      --s3.public                                  s3 public
      --s3.region string                           s3 region
      --s3.secret_access_key string                s3 secret access key

Global Flags:
      --db.datasource string                 database datasource (default ":memory:?_pragma=foreign_keys(1)")
      --db.driver string                     database driver (default "sqlite3")
      --generate-config.enable               generate config enable
      --generate-config.format string        generate config format, one of json, yaml, env (default "yaml")
      --generate-config.path string          generate config path
      --log.file string                      log file path
      --log.format string                    log format, one of json, console, empty for default (console for dev, json for prod)
      --log.level string                     log level (default "info")
      --log.levels.console string            log level for console, empty for same as log.level
      --log.levels.file string               log level for file, empty for same as log.level
      --log.rotate.enable                    enable log rotation
      --log.rotate.max_age int               max age of log file in days (default 28)
      --log.rotate.max_backups int           max number of log file backups (default 3)
      --log.rotate.max_size int              max size of log file in MB (default 500)
      --proxy.debounce_time duration         order close debounce time (default 1m0s)
      --proxy.secret_id string               kuaidaili account secret id
      --proxy.secret_key string              kuaidaili account secret key
      --telemetry.sentry_dsn string          sentry dsn
      --telemetry.traces_sample_rate float   traces sample rate (default 1)
```

2. By env, map every config by replace `.-` to `_`, and add the prefix `SOAR_HELPER_`. For example, `--telemetry.sentry_dsn` will map to `SOAR_HELPER_TELEMETRY_SENTRY_DSN`

3. Using config file, program will find config at `/etc/soar-helper/` and `$HOME/.soar-helper`, with name `soar-helper.[json|yaml|toml|...]`, supported format can be found at https://github.com/spf13/viper

