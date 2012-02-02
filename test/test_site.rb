# encoding: utf-8
#
#   this is jsmin.de, a sinatra application
#   it is copyright (c) 2012 danilo braband (dbraband @ gmail,
#   then a dot and a 'com')
#

require 'helpers'

#
# These test cases tests the behaviour of the site.
#
class TestSite < TestHelper

  def test_homepage_redirect
    get '/'
    assert last_response.redirect?
    follow_redirect!

    assert_equal 'http://jsmin.de/', last_request.url
  end


  def test_homepage_newrelic_agent
    header 'User-Agent', 'NewRelicPinger/1.0 (39335)'
    get '/'
    assert last_response.ok?
  end


  def test_homepage_curl_agent
    header 'User-Agent', 'curl/7.21.4 (universal-apple-darwin11.0)'
    header 'Accept', 'application/json'
    get '/'
    assert last_response.ok?
  end


  def test_js_get
    get '/js'
    assert !last_response.ok?
  end


  def test_js_post
    post '/js'
    assert !last_response.ok?
  end


  def test_js_post_json
    header 'Accept', 'application/json'
    post '/js'
    assert last_response.ok?
  end

end

