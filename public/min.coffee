before_chars = $('form .chars')
result = $('#result')
error = $('#error')
textarea_input = $("form textarea")

$("form").submit( ->
    before = $('#content').val()
    type = 'js' # $('#type').val()
    $.post("/post-#{type}", before, (data, state, obj) ->
        error.hide()
        result.show().children('.chars').text(data.length)
        result.children('textarea').val(data).focus().select()
        percent = 100-Math.round(data.length*100/before.length)
        saved = if percent > 100 then 'bloated' else 'saved'
        result.children('.stats').text("#{percent}% #{saved}")
        before_chars.text(before.length)
    )
    return false
)

error.ajaxError(->
    result.hide()
    $(this).show().text('Failed to process your code. Please check if your code
    is valid javascript.')
)

textarea_input.keyup( ->
    before_chars.text(this.value.length)
)
textarea_input.focus()

# Google Analytics
_gaq = _gaq || []
_gaq.push(['_setAccount', 'UA-67221-12'])
_gaq.push(['_trackPageview'])
ga = document.createElement('script')
ga.type = 'text/javascript'
ga.async = true
ga.src = 'http://www.google-analytics.com/ga.js'
s = document.getElementsByTagName('script')[0]
s.parentNode.insertBefore(ga, s)

