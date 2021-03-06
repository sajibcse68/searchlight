## hyperalert check_influx_query

Check InfluxDB Query Data

### Synopsis


Check InfluxDB Query Data

```
hyperalert check_influx_query [flags]
```

### Options

```
      --A string            InfluxDB query A
      --B string            InfluxDB query B
      --C string            InfluxDB query C
      --D string            InfluxDB query D
      --E string            InfluxDB query E
      --R string            Equation to evaluate result
  -c, --critical string     Critical query which returns [true/false]
  -h, --help                help for check_influx_query
  -H, --host string         Icinga host name
      --influxHost string   URL of InfluxDB host to query
      --kubeconfig string   Path to kubeconfig file with authorization information (the master location is set by the master flag).
      --master string       The address of the Kubernetes API server (overrides any value in kubeconfig)
  -s, --secretName string   Kubernetes secret name
  -w, --warning string      Warning query which returns [true/false]
```

### Options inherited from parent commands

```
      --allow_verification_with_non_compliant_keys   Allow a SignatureVerifier to use keys which are technically non-compliant with RFC6962.
      --alsologtostderr                              log to standard error as well as files
      --log_backtrace_at traceLocation               when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                               If non-empty, write log files in this directory
      --logtostderr                                  log to standard error instead of files (default true)
      --stderrthreshold severity                     logs at or above this threshold go to stderr (default 2)
  -v, --v Level                                      log level for V logs
      --vmodule moduleSpec                           comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO
* [hyperalert](hyperalert.md)	 - AppsCode Icinga2 plugin


