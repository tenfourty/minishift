## minishift docker-env

sets up docker env variables; similar to '$(docker-machine env)'

### Synopsis


sets up docker env variables; similar to '$(docker-machine env)'

```
minishift docker-env
```

### Options

```
      --no-proxy       Add machine IP to NO_PROXY environment variable
      --shell string   Force environment to be configured for a specified shell: [fish, cmd, powershell, tcsh, bash, zsh], default is auto-detect
  -u, --unset          Unset variables instead of setting them
```

### Options inherited from parent commands

```
      --alsologtostderr value          log to standard error as well as files
      --disable-update-notification    Whether to disable VM update check.
      --log-flush-frequency duration   Maximum number of seconds between log flushes (default 5s)
      --log_backtrace_at value         when logging hits line file:N, emit a stack trace (default :0)
      --log_dir value                  If non-empty, write log files in this directory
      --logtostderr value              log to standard error instead of files
      --show-libmachine-logs           Whether or not to show logs from libmachine.
      --stderrthreshold value          logs at or above this threshold go to stderr (default 2)
  -v, --v value                        log level for V logs
      --vmodule value                  comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO
* [minishift](minishift.md)	 - Minishift is a tool for managing local OpenShift clusters.

