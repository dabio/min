var uglify_js = require('uglify-js2');

/*
 * GET /
 */
exports.index = function (req, res) {
  res.render('index', { title: 'Compress Your Javascript' });
};

/*
 * POST /
 */
exports.post = function (req, res) {
  minimize_js(req.body.content);
  //res.render('index', { title: 'Compress Your Javascript' });
};


var minimize_js = function(data) {
  console.log(data);
  var result = uglify_js.minify(data, {fromString: true});
  console.log(result.code);
  console.log(result.map);
}

