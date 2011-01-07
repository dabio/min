(function() {
  var connect, jsp, minimize, pro, server;
  connect = require("" + __dirname + "/vendor/connect/lib/connect");
  jsp = require("" + __dirname + "/vendor/uglify-js/lib/parse-js");
  pro = require("" + __dirname + "/vendor/uglify-js/lib/process");
  minimize = function(app) {
    return app.post('/post', function(req, res, next) {
      var data;
      data = '';
      req.on('data', function(chunk) {
        return data += chunk.toString();
      });
      return req.on('end', function() {
        var ast;
        try {
          ast = jsp.parse(data);
          ast = pro.ast_mangle(ast);
          ast = pro.ast_squeeze(ast);
          ast = pro.gen_code(ast);
          res.writeHead(200, {
            'Content-Type': 'text/plain'
          });
          return res.end(ast);
        } catch (error) {
          res.writeHead(404, {
            'Content-Type': 'text/plain'
          });
          return res.end("404 Not Found\n");
        }
      });
    });
  };
  server = connect.createServer(connect.router(minimize), connect.conditionalGet(), connect.gzip(), connect.staticProvider("" + __dirname + "/public")).listen(80);
}).call(this);
