# encoding: utf-8
#
#   this is jsmin.de, a sinatra application
#   it is copyright (c) 2012 danilo braband (dbraband @ gmail,
#   then a dot and a 'com')
#

RACK_ENV = 'test'
require File.join(File.expand_path(File.dirname(__FILE__)), '..', 'app.rb')

SimpleCov.start

require 'test/unit'

class TestHelper < Test::Unit::TestCase
  include Rack::Test::Methods

  def app
    App
  end

end

