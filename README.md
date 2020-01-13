# `jmx-cnb`
The Cloud Foundry JMX Buildpack is a Cloud Native Buildpack V3 that enables the JMX in JVM applications.

## Behavior
This buildpack will participate if all of the following conditions are met

* `$BP_JMX` is set

The buildpack will do the following:

* Contribute JMX configuration to `$JAVA_OPTS` 

## Configuration 
| Environment Variable | Description
| -------------------- | -----------
| `BPL_JMX_PORT` | What port the JVM should expose JMX on. Defaults to `5000`. 

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
