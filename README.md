# gtavd
`gtavd` is a Grand Theft Auto V modification daemon. According to [wikipedia](https://en.wikipedia.org/wiki/Daemon_(computing)), a daemon is a computer program that runs as a background process, rather than being under the direct control of an interactive user.

As just previously stated, this tool allows Grand Theft Auto V modding enthusiasts to run common tasks derived from modding activities with almost no user-interaction. The main benefit of `gtavd` is to ease overall modding operational overhead (_e.g. updating dlclist.xml file everytime when adding an addon mod, keeping files in-sync between multiple devices_).

## How-To

### Install Daemon

`gtavd` is bundled with the **_gtavd.exe_** binary executable and a **_service-installer.bat_** MS-DOS script. To install the daemon, run the _service-installer.bat_ script with _Administrative Privileges/Run as Admin_. Verify the script outputs if there is an error during the executing. If no error was found, verify the daemon is running by opening the _Task Manager_ and look for the _gtavd.exe_ background process.

### Change Configuration

`gtavd` is fully configurable using the [**gtavd.config**](testdata/configs/gtavd.config.yaml) file. It contains a set of definitions the daemon will use to perform tasks. Be aware some configurations such as _module enabling_ require the RESTART of the service to take effect.

The configuration file is located under the user's HOME path (*e.g. C:/Users/YOUR_USERNAME/.gtavd*). Enabling the **Show hidden files** Windows feature MIGHT be required in order to be able to see the _.gtavd_ folder.

### Restart/Stop the Service

`gtavd` runs as a [Windows service](https://docs.microsoft.com/en-us/dotnet/framework/windows-services/introduction-to-windows-service-applications). In order to restart or stop the service: 

- Go to task manager, click on **More details** and look for the **_gtavd.exe_** process. 
- Once found, click on the left arrow to expand the process item and the service will be shown (represented as the only item within the process with a gear icon).
- To stop the service, right-click the service and select **Stop**.
- To restart the service, right-click the service and select **Open Services**.
  - A new window will open with the services list, look for the _Grand Theft Auto V Modification Daemon_ service.
  - Once found, right-click the service and select **Restart**.

### Access to Logs for Troubleshooting

`gtavd` makes use of the [_rotating-file_](https://en.wikipedia.org/wiki/Log_rotation) logging technique. This means it generates logs compressed and partitioned by default along log removal after _n_ days (14 days). Logs are partitioned by size (128 MB).

In addition, `gtavd` streams logs into the _stderr_ too.


The log files are located under the user's HOME path (*e.g. C:/Users/YOUR_USERNAME/.gtavd/logs*). Enabling the **Show hidden files** Windows feature MIGHT be required in order to be able to see the _.gtavd_ folder.

## Features

### Dlclist.xml Automatic Generation

The `dlclist` module is a background-running task detecting changes of both _/mods/x64/dlcpacks_ and _/x64/dlcpacks_ directories of a Grand Theft Auto V game. When a change is detected, the module receives a notification and generates automatically a **dlclist.xml** file ready to be imported into _update.rpf_ file using _OpenIV_. 

## Feature Roadmap

- `Cloud-sync of configuration files`. Useful to keep a global configuration between multiple machines.
- `Cloud-sync of mods files`. Useful to keep a set installed mods between multiple machines. It MIGHT just keep in-sync the _mods_ folder to reduce overall cloud storage consumption.
