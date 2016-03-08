bofhwits  [![Build Status](https://travis-ci.org/amauragis/bofhwits.svg?branch=master)](https://travis-ci.org/amauragis/bofhwits)
========

bofhwits: A Bad Bot

An IRC chat bot, that, among other things, will provide a quoting service to allow logging of stupid
and/or amusing things that people say.  Hopefully displaying them in some manner somewhere, also.

Fair warning: It's pretty bad and probably won't work for you, but hey, you never know!

It's currently running in synirc's #bofh and logging to  [bofh.wtf](https://bofh.wtf).

Installation
------------
I haven't actually tested this, but here is my best guess!

1. [Install Go (>= 1.6)](https://golang.org/doc/install)
  - If you haven't installed go before, do read that page a bit.  $GOPATH and $GOROOT and such are kind of
    non intuitive.
1. Install godep: `go get github.com/tools/godep`
1. Install bofhwits with specified dependencies: `godep get github.com/amauragis/bofhwits`


Configuration
-------------
Everything is configured in `./config/bofhwits.yaml`, an [example](config/bofhwits.yaml.example) is
provided in the config directory.  The bofhwits executable searchs `./config/bofhwits.yaml` specifically
but depending on how you installed it or where you put the config file, you can pass in whatever path
you want.

Running bofhwits
----------------
1. Set up your config file.
1. Use your head to decide if either of these command line options apply to you
 - `-c <path>`: A path to the configuration file you want to use (default is `./config/bofhwits.yaml`)
 - `-l <path>`: A path to a logfile to write/append to (if no option its provided, the bot will print to stdout)
1. You are done.  Make hilarious posts on the internet and let bofhwits record them for you.

bofhwits Commands
-----------------
 - !bofh <msg> : quotes the message
 - !bofhwitsdie : makes the bot quit
 - !buttes : Donges you, what else.

Dependencies
------------
I decided to set up [godep](https://github.com/tools/godep).  I'm not avery good at using it, but it
seemed to work!  Small note though, the [sqlite package](https://github.com/mattn/go-sqlite3) didn't
work right with godep because it has some C code in the "code" subdirectory.  You might need to copy
this in manually if godep doesn't do it for you.

Actual dependency list: [here](Godeps/Godeps.json)
