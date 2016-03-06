bofhwits  [![Build Status](https://travis-ci.org/amauragis/bofhwits.svg?branch=master)](https://travis-ci.org/amauragis/bofhwits)
========

BOFHwits: A Bad Bot

An IRC chat bot, that, among other things, will provide a quoting service to allow logging of stupid
and/or amusing things that people say.  Hopefully displaying them in some manner somewhere, also.

Fair warning: It's pretty bad and probably won't work for you, but hey, you never know!

Configuration
-------------
Everything is configured in `./config/bofhwits.yaml`, an example is provided in the config directory.  
If you wish to use a different configuration file, the `bofhwits -c <path>` option is provided.

I've also added a -l flag that should direct output to a log file of your choosing.  `bofhwits -l <logfile>`

To use the twitter integration, app and account api-keys are required.

Dependencies
------------
I decided to set up [godep](https://github.com/tools/godep).  I'm not avery good at using it, but it
seemed to work!  Small note though, the [sqlite package](https://github.com/mattn/go-sqlite3) didn't
work right with godep because it has some C code in the "code" subdirectory.  You might need to copy
this in manually if godep doesn't do it for you.

Actual dependency list: [here](https://github.com/amauragis/bofhwits/blob/master/Godeps/Godeps.json)
