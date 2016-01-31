# srv

Serve a directory over HTTP. Like `python -m SimpleHTTPServer` except this program allows you to serve any directory (`python -m SimpleHTTPServer` only allows you to serve the current working directory) by passing a file path as the first argument. Defaults to current working directory if no arguments are passed.

Testing prose.io
![]({{site.baseurl}}/autumn-fence.jpg)

Eg:

    srv code/website/public

## Options


    $ srv -h
    Usage of srv
      -p="8000": Port to run server on.


## Install

    go get github.com/nwjlyons/srv
