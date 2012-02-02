# encoding: utf-8
#
#   this is jsmin.de, a sinatra application
#   it is copyright (c) 2012 danilo braband (dbraband @ gmail,
#   then a dot and a 'com')
#

if defined? Encoding
  Encoding.default_external = Encoding::UTF_8
  Encoding.default_internal = Encoding::UTF_8
end

RACK_ENV = ENV['RACK_ENV'] ||= 'development' unless defined? RACK_ENV
ROOT_DIR = File.dirname(__FILE__) unless defined? ROOT_DIR

require 'bundler/setup'
Bundler.require(:default, RACK_ENV)

# Sinatra::Base. This way, we're not polluting the global namespace with your
# methods and routes and such.
class App < Sinatra::Base; end

class App

  set :app_file, __FILE__
  set :port, ENV['PORT']

  use Rack::ForceDomain, ENV['DOMAIN']
  use Rack::Timeout
  Rack::Timeout.timeout = 10
  #use Rack::Protection, :except => :session_hijacking

  configure :development, :test do
    begin
      require 'ruby-debug'
    rescue LoadError
    end
  end

  configure :development do
    register Sinatra::Reloader
  end


  #
  # GET /
  # This is for the newrelic pinger which visits the app in regular intervalls.
  # We just return an empty site.
  #
  get '/', :agent => /NewRelicPinger.+/ do; end


  #
  # GET /
  # This is for all the folks who pings this site with curl. Return an empty
  # JSON format here
  #
  get '/', :agent => /curl.+/, :provides => 'json' do
    {}.to_json
  end


  #
  # GET /
  # Everyone else gets redirected to the online compression tool.
  #
  get '/' do
    redirect 'http://jsmin.de/', 303
  end


  #
  # POST /js
  # This minimizes the given 'text' given by the JSON container and returns
  # a compressed JSON format.
  #
  post '/js', :provides => 'json' do
    puts params
    {}.to_json
  end


  # must use this for working on heroku cedar
  run! if app_file == $0

end

