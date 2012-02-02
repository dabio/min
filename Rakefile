# encoding: utf-8
#
#   this is jsmin.de, a sinatra application
#   it is copyright (c) 2012 danilo braband (dbraband @ gmail,
#   then a dot and a 'com')
#


#
# Development
#

task :default => :development
task :development do
    system 'bundle exec thin start -p 9393 -e development'
end


#
# Tests
#

task :test do
  require 'fileutils'
  require 'rake/testtask'

  Rake::TestTask.new do |t|
    t.libs << 'test'
    t.test_files = FileList['test/test_*.rb']
    t.verbose = true
  end

end


#
# Install/Uninstall
#

task :uninstall do
  system 'gem list | cut -d" " -f1 | xargs gem uninstall -aIx'
  File.unlink 'Gemfile.lock'
  system 'rbenv rehash'
end

task :install do
  system 'gem install bundler'
  system 'bundle install --without production'
  system 'rbenv rehash'
end

