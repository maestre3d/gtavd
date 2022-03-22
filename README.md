# gtavd
`gtavd` is a Grand Theft Auto V daemon for third-party modification management. According to 
[wikipedia](https://en.wikipedia.org/wiki/Daemon_(computing)), a daemon is a computer program that runs as a 
background process, rather than being under the direct control of an interactive user.

As just previously stated, this tool allows Grand Theft Auto V modding enthusiasts to run common tasks derived from 
modding activities with almost no user-interaction. 
The main benefit of `gtavd` is to ease overall modding operational overhead (e.g. updating the dlclist.xml file 
everytime a new addon mod is added, enabling/disabling specific scripts, keeping files in-sync between multiple devices).

## Table of Contents

- [gtavd](#gtavd)
  - [Table of Contents](#table-of-contents)
  - [How-To](#how-to)
    - [Install Daemon](#install-daemon)
    - [Run as a Program (Standalone version)](#run-as-a-program-standalone-version)
    - [Change Configuration](#change-configuration)
    - [Restart/Stop the Service](#restartstop-the-service)
    - [Access to Logs for Troubleshooting](#access-to-logs-for-troubleshooting)
  - [Features](#features)
    - [Dlclist.xml Automatic Generation](#dlclistxml-automatic-generation)
  - [Feature Roadmap](#feature-roadmap)

## How-To

### Install Daemon

`gtavd` is bundled with the **_gtavd.exe_** binary executable and a **_service-installer.bat_** Windows script. 
To install the daemon, run the _service-installer.bat_ script with _Administrative Privileges/Run as Admin_. 
Look at the script output to check if an error happened during the execution. If no error was found, 
verify if the daemon is running by opening the _Task Manager_ and looking for the _gtavd.exe_ background process.

### Run as a Program (Standalone version)

`gtavd` is bundled with a **_gtavd_standalone.exe_** binary executable. This executable contains most of the features 
the daemon offers. Nevertheless, the _standalone_ version WILL NOT run as a background process -no Windows Service is 
required-. 

To run the executable, double-click the file. As a matter of fact, this version will prompt a Windows console 
(_CMD or Powershell_).

### Change Configuration

`gtavd` is fully configurable using the [**gtavd.config**](testdata/configs/gtavd.config.yaml) file. It contains a set of definitions the daemon will use to perform tasks. Be aware some configurations such as _module enabling_ require the RESTART of the service to take effect.

The configuration file is located under the user's HOME path (*e.g. C:/Users/YOUR_USERNAME/.gtavd*). Enabling the 
**Show hidden files** _Windows Explorer_ feature MIGHT be required in order to be able to see the _.gtavd_ folder.

### Restart/Stop the Service

`gtavd` runs as a [Windows service](https://docs.microsoft.com/en-us/dotnet/framework/windows-services/introduction-to-windows-service-applications). In order to restart or stop the service: 

- Go to task manager, click on **More details** and look for the **_gtavd.exe_** process. 
- Once found, click on the left arrow to expand the process item and the service will be shown (represented as the only item within the process with a gear icon).
- To stop the service, right-click the service and select **Stop**.
- To restart the service, right-click the service and select **Open Services**.
  - A new window will open with the services list, look for the _Grand Theft Auto V Modification Daemon_ service.
  - Once found, right-click the service and select **Restart**.

### Access to Logs for Troubleshooting

`gtavd` makes use of the [_rotating-file_](https://en.wikipedia.org/wiki/Log_rotation) logging technique. This means it 
generates logs gzip-compressed and partitioned by default along a log removal after _n_ days mechanism 
(_default is 14 days_). Furthermore, the logs are partitioned by size (_default is 128 MB_).
Finally, `gtavd` streams logs into the _stderr_ too.

The log files are located under the user's HOME path (*e.g. C:/Users/YOUR_USERNAME/.gtavd/logs*). Enabling the 
**Show hidden files** _Windows Explorer_ feature MIGHT be required in order to be able to see the _.gtavd_ folder.

## Features

### Dlclist.xml Automatic Generation

The `dlclist` module is a background-running task detecting changes of both _/mods/x64/dlcpacks_ and _/x64/dlcpacks_ directories of a Grand Theft Auto V game. When a change is detected, the module receives a notification and generates automatically a **dlclist.xml** file ready to be imported into _update.rpf_ file using _OpenIV_. 

## Feature Roadmap

- `Script blacklisting`. Enable or disable script mods (_e.g. .rph, .ini, .dll, .asi_) using a blacklist mechanism. 
Blacklisted scripts will be moved to a _temporary_ folder, so they can be later re-enabled. 
- `Cloud-sync of configuration files`. Useful to keep a global configuration between multiple machines.
- `Cloud-sync of mods files`. Useful to keep a set installed mods between multiple machines. It MIGHT just keep in-sync 
the _mods_ folder and lightweight scripts to reduce overall cloud storage consumption.
