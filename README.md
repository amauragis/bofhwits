bofhwits  [![Build Status](https://travis-ci.org/amauragis/bofhwits.svg?branch=master)](https://travis-ci.org/amauragis/bofhwits)
========

BOFHwits: A Bad Bot

An IRC chat bot, that, among other things, will provide a quoting service to allow logging of stupid
and/or amusing things that people say.  Hopefully displaying them in some manner somewhere, also.

Fair warning: it's super brittle and won't work for you.

Configuration
-------------
Everything is configured in `./config/bofhwits.yaml`, an example is provided in the config directory.  
If you wish to use a different configuration file, the `bofhwits -c <path>` option is provided.

I've also added a -l flag that should direct output to a log file of your choosing

To use the twitter integration, app and account api-keys are required.

Dependencies
------------
 - Go
 - Go ircevent (https://github.com/thoj/go-ircevent)
 - Go yaml (https://github.com/go-yaml/yaml)
 - Go anaconda (https://github.com/ChimeraCoder/anaconda)
 - This list is incomplete and will be replaced with Godep whenever I figure out how that works
