# Launcher
A simple console command Launcher menu using GOCUI (https://github.com/jroimartin/gocui)

## Build
Simply run `go install` to build and install the application.

## Usage
Put the configuration file `.menu.json` (see example below) into your home directory and execute `launch` command in the terminal.
 
```javascript
[
    {
        "title": "Server 1",
        "desc": "Apache Web-Server",
        "cmd": ["ssh", "user@www-server"]
    },
    {
        "title": "Database Server",
        "desc": "MySql",
        "cmd": ["ssh", "user@db-server"]
    }
]
```

