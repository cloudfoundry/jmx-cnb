# `jmx-cnb`
The Cloud Foundry JMX Buildpack is a Cloud Native Buildpack V3 that enables the JMX in JVM applications.

## Detection
The detection phase passes if

* `$BP_JMX` exists and build plan contains `jvm-application`
  * Contributes `jmx` to the build plan

## Build
If the build plan contains

* `jmx`
  * Contributes JMX configuration to `$JAVA_OPTS`
  * If `$BPL_JMX_PORT` is specified, configures the port JMX  will listen on.  Defaults to `5000`.

## Creating SSH Tunnel
After starting an application with JMX enabled, an SSH tunnel must be created to the container.  To create that SSH container, execute the following command:

```bash
$ cf ssh -N -T -L <LOCAL_PORT>:localhost:<REMOTE_PORT> <APPLICATION_NAME>
```

The `REMOTE_PORT` should match the `port` configuration for the application (`5000` by default).  The `LOCAL_PORT` must match the `REMOTE_PORT`.

Once the SSH tunnel has been created, your JConsole should connect to `localhost:<LOCAL_PORT>` for JMX access.

![JConsole Configuration](jconsole.png)

## License
This buildpack is released under version 2.0 of the [Apache License][a].

[a]: https://www.apache.org/licenses/LICENSE-2.0
