source :rubygems

gem 'sinatra', require: 'sinatra/base'
gem 'json'
gem 'uglifier'
gem 'rack-force_domain'
gem 'rack-timeout', require: 'rack/timeout'
gem 'rake', require: false
gem 'thin', require: false

group :development, :test do
  gem 'sinatra-contrib', require: 'sinatra/reloader'
  gem 'heroku', require: false
end

group :test do
  gem 'simplecov'
  gem 'rack-test', require: 'rack/test'
end

group :production do
  gem 'newrelic_rpm'
end

