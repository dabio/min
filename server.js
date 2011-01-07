(function() {
  var connect, jsp, minimize, pro, server, uglify;
  connect = require('connect');
  uglify = require('uglify-js');
  jsp = uglify.parser;
  pro = uglify.uglify;
  minimize = function(app) {
    return app.post('/post', function(req, res, next) {
      var data;
      data = '';
      req.on('data', function(chunk) {
        return data += chunk.toString();
      });
      return req.on('end', function() {
        var ast;
        ast = jsp.parse(data);
        ast = pro.ast_mangle(ast);
        ast = pro.ast_squeeze(ast);
        ast = pro.gen_code(ast);
        res.writeHead(200, {
          'Content-Type': 'text/plain'
        });
        return res.end(ast);
      });
    });
  };
  server = connect.createServer(connect.router(minimize), connect.conditionalGet(), connect.gzip(), connect.staticProvider("" + __dirname + "/public")).listen(80);
}).call(this);
