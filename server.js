(function() {
  var fs, jsp, minimize_css, minimize_js, pro;
  fs = require('fs');
  jsp = require("" + __dirname + "/vendor/uglify-js/lib/parse-js");
  pro = require("" + __dirname + "/vendor/uglify-js/lib/process");
  require('http').createServer(function(req, res) {
    var data, index, minjs;
    switch (req.url) {
      case '/':
        index = fs.readFileSync("" + __dirname + "/public/index.html", 'utf8');
        res.writeHead(200, {
          'Content-Type': 'text/html',
          'Content-Length': index.length
        });
        return res.end(index);
      case '/min.js':
        minjs = fs.readFileSync("" + __dirname + "/public/min.js", 'utf8');
        res.writeHead(200, {
          'Content-Type': 'text/javascript',
          'Content-Length': minjs.length
        });
        return res.end(minjs);
      case '/post-js':
        data = '';
        req.on('data', function(chunk) {
          return data += chunk.toString();
        });
        return req.on('end', function() {
          try {
            data = minimize_js(data);
            res.writeHead(200, {
              'Content-Type': 'text/plain',
              'Content-Length': data.length
            });
            return res.end(data);
          } catch (error) {
            res.writeHead(404, {
              'Content-Type': 'text/plain'
            });
            return res.end("404 Not Found\n");
          }
        });
      default:
        res.writeHead(404, {
          'Content-Type': 'text/plain'
        });
        return res.end("404 Not Found\n");
    }
  }).listen(process.env.PORT || 9393);
  minimize_js = function(data) {
    var ast;
    ast = jsp.parse(data);
    ast = pro.ast_mangle(ast);
    ast = pro.ast_squeeze(ast);
    return pro.gen_code(ast);
  };
  minimize_css = function(data) {
    return css.cssmin(data);
  };
}).call(this);
