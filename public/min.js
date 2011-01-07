(function() {
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
}).call(this);
