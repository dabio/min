(function() {
  var before_chars, result, _gaq;
  before_chars = $('form .chars');
  result = $('#result');
  $("form").submit(function() {
    var before;
    before = $('#js').val();
    $.post('/post', before, function(data, state, obj) {
      var percent, saved;
      result.show().children('.chars').text(data.length);
      result.children('textarea').val(data).focus().select();
      percent = 100 - Math.round(data.length * 100 / before.length);
      saved = percent > 100 ? 'bloated' : 'saved';
      result.children('.stats').text("" + percent + "% " + saved);
      return before_chars.text(before.length);
    });
    return false;
  });
  $("#error").ajaxError(function() {
    result.hide();
    return $(this).show().text('Failed to process your code. Please check if your code\
    is valid javascript.');
  });
  $("form textarea").keyup(function() {
    return before_chars.text(this.value.length);
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
