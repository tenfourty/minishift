## minishift stop

Stops a running local OpenShift cluster.

### Synopsis


Stops a local OpenShift cluster running in Virtualbox. This command stops the VM
itself, leaving all files intact. The cluster can be started again with the "start" command.

```
minishift stop
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

