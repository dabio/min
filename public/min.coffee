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

# Google Analytics
_gaq = _gaq || [];
_gaq.push(['_setAccount', 'UA-67221-12'])
_gaq.push(['_trackPageview'])
->
    ga = document.createElement('script')
    ga.type = 'text/javascript'
    ga.async = true
    ga.src = 'http://www.google-analytics.com/ga.js'
    s = document.getElementsByTagName('script')[0]
    s.parentNode.insertBefore(ga, s)

