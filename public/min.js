(function() {
  var _gaq;
  $("form").submit(function() {
    var before, result;
    before = $('#js').val();
    result = $('#result');
    $.post('/post', before, function(data, state, obj) {
      result.show();
      result.children('p').text("" + before.length + "/" + data.length);
      return result.children('textarea').val(data).focus().select();
    });
    return false;
  });
  _gaq = _gaq || [];
  _gaq.push(['_setAccount', 'UA-67221-12']);
  _gaq.push(['_trackPageview']);
  (function() {
    var ga, s;
    ga = document.createElement('script');
    ga.type = 'text/javascript';
    ga.async = true;
    ga.src = 'http://www.google-analytics.com/ga.js';
    s = document.getElementsByTagName('script')[0];
    return s.parentNode.insertBefore(ga, s);
  });
}).call(this);
