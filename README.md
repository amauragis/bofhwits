bofhwits
========

BOFHwits: A Bad Bot

An IRC chat bot, that, among other things, will provide a quoting service to allow logging of stupid 
and/or amusing things that people say.  Hopefully displaying them in some manner somewhere, also.

Fair warning: it's super brittle and won't work for you.

Configuration
-------------
Everything is configured in `./config/bofhwits.yaml`, an example is provided in the config directory.  
If you wish to use a different configuration file, the `bofhwits -c <path>` option is provided.

To use the twitter integration, app and account api-keys are required.

Dependencies
------------
 - Go
 - Go ircevent (https://github.com/thoj/go-ircevent)
 - Go yaml (https://github.com/go-yaml/yaml)
 - Go anaconda (https://github.com/ChimeraCoder/anaconda)
 
