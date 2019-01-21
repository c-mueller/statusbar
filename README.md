# statusbar

![Statusbar in i3 MultiBlock Mode](docs/img/statusbar-i3mb.png)

My personal implementation of a statusbar for the i3 window manager.

While the application will probably run on other platforms i3 does not
making it only usable on a Linux based system.

## Components

Currently the statusbar application implements the following components:

- Hostname - Displays the current username and hostname.
- Text - Display static text.
- Uptime - Show the current system uptime (in hours and minutes)
- Clock - Shows the current time of the day in HH:MM:SS format
- Date - Shows the current date
- Command - Execute a command, and print the output of the command
- Network - Show current network usage in Bytes per second
- Memory - Show memory usage
- CPUAvgChart - Prints the global average CPu usage as rolling time graph.
- CPULoadBar - Show the current load of every thread.
- Block - Combine multiple child components as one component.
  Using this is mandatory for using multiple components in the `Wheel` component
  Also, when using the i3 Multi Block renderer components grouped using the `Block`
  component will be considered as one block, removing the separator.
- Wheel - Roll the text from a internal component around in a predefined
  string length.


## Configuring

The application is configured using YAML, the following describes the common
definition of a configuration file.

```yaml
refresh_interval: 500 # The Refresh intervall of the statusbar in milliseconds
components: # List of components
  - identifier: hostname # The unique identifier of the component
    type: Hostname
```

A sample config can be seen in the [config.yml](config.yml) file. 

Default configurations for components can be generated using the `statusbar components default-confg <COMPONENT>` command.

## License

This program is Licensed under GPL v3.