/**
 * Module dependencies.
 */

var express = require('express')
  , http = require('http')
  , path = require('path')
  , uglify = require('uglify-js2')
  , app = express();

/*
 * Express configuration
 */
app.configure(function() {
  app.set('port', process.env.PORT || 9393);
  app.set('views', __dirname + '/views');
  app.set('view engine', 'jade');
  app.use(express.favicon());
  app.use(express.responseTime());
  app.use(express.logger('dev'));
  app.use(express.bodyParser());
  app.use(app.router);
  app.use(express.static(path.join(__dirname, 'public')));
});

app.configure('development', function() {
  app.use(express.errorHandler())
});


/*
 * GET /
 */
app.get('/', function (req, res) {
  res.render('index');
});


/*
 * POST /
 */
app.post('/', function (req, res) {
  var input = req.body.content;
  res.render(
    'index', {
      input: input,
      output: minimize_js(req.body.content)
    }
  );
});


/*
 * Create the server
 */
http.createServer(app).listen(app.get('port'), function(){
  console.log('Express server listening on port ' + app.get('port'));
});


/*
 * Misc
 */
var minimize_js = function(data) {
  return uglify.minify(data, { fromString: true }).code;
};
