$("form").submit( ->
    before = $('#js').val()
    result = $('#result')
    $.post('/post', before, (data, state, obj) ->
        result.show()
        result.children('p').text("#{before.length}/#{data.length}")
        result.children('textarea').val(data).focus().select()
    )
    return false
)
