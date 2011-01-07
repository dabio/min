connect = require "#{__dirname}/vendor/connect/lib/connect"
jsp = require "#{__dirname}/vendor/uglify-js/lib/parse-js"
pro = require "#{__dirname}/vendor/uglify-js/lib/process"

minimize = (app) ->
    app.post '/post', (req, res, next) ->
        data = ''
        req.on 'data', (chunk) ->
            data +=  chunk.toString()
        req.on 'end', ->
            ast = jsp.parse(data)
            ast = pro.ast_mangle(ast)
            ast = pro.ast_squeeze(ast)
            ast = pro.gen_code(ast)
            res.writeHead(200, {'Content-Type': 'text/plain'})
            res.end(ast)

server = connect.createServer(
    connect.router(minimize),
    connect.conditionalGet(),
    connect.gzip(),
    connect.staticProvider("#{__dirname}/public")
).listen 9393

