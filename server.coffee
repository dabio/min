fs = require 'fs'
jsp = require "#{__dirname}/vendor/uglify-js/lib/parse-js"
pro = require "#{__dirname}/vendor/uglify-js/lib/process"
#css = require "#{__dirname}/vendor/cssmin/cssmin"

require('http').createServer (req, res) ->
    switch req.url
        when '/'
            index = fs.readFileSync "#{__dirname}/public/index.html", 'utf8'
            res.writeHead 200, {'Content-Type': 'text/html', 'Content-Length':
                index.length}
            res.end index
        when '/min.js'
            minjs = fs.readFileSync "#{__dirname}/public/min.js", 'utf8'
            res.writeHead 200, {'Content-Type': 'text/javascript', 'Content-Length':
                minjs.length}
            res.end minjs
        when '/post-js'
            data = ''
            req.on 'data', (chunk) ->
                data += chunk.toString()
            req.on 'end', ->
                try
                    data = minimize_js data
                    res.writeHead 200, {'Content-Type': 'text/plain', 'Content-Length':
                        data.length}
                    res.end data
                catch error
                    res.writeHead 404, {'Content-Type': 'text/plain'}
                    res.end "404 Not Found\n"
        else
            res.writeHead 404, {'Content-Type': 'text/plain'}
            res.end "404 Not Found\n"
.listen process.env.PORT or 9393



minimize_js = (data) ->
    ast = jsp.parse data
    ast = pro.ast_mangle ast
    ast = pro.ast_squeeze ast
    pro.gen_code ast

minimize_css = (data) ->
    css.cssmin data


