@echo off
rem run this script as admin

if not exist gtavd.exe (
    echo No gtavd executable found
    goto :exit
)

sc create gtavd binpath= "%CD%\gtavd.exe" start= auto DisplayName= "Grand Theft Auto V Modification Daemon"
sc description gtavd "Runs background tasks related to modifications of the Grand Theft Auto V game"
net start gtavd
sc query gtavd

echo Service started

:exit