# [min](http://jsmin.de)

This is [jsmin](http://jsmin.de), a javascript compressor. Basically it is an API you can use the minimize your own javascript.

If you don't want to use the API, head over to the [online user interface](http://jsmin.de) to compress your javascript.

## API

All API access is over HTTP and accessed from the `api.jsmin.de` domain. All data is sent and received as JSON.

    $ curl -i http://api.jsmin.de

    HTTP/1.1 200 OK
    X-Frame-Options: sameorigin
    X-XSS-Protection: 1; mode=block
    Content-Type: application/json;charset=utf-8
    Content-Length: 2

    {}

Pass your javascript you want to compress to `/js` and get the minimized result as a response.

###### Request

    POST /js

    {
        "text": "new Array(1, 2, 3, 4);"
    }

###### Response

    {
        "test": "[1,2,3,4]"
    }

You could calculate the compression rate based upon your input if you like.
