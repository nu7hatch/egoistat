# -*- ruby -*-

begin
  require 'bundler/setup'
rescue LoadError
  $stderr.write("ERROR: Bundler not installed, run `gem install bundler` to get it.\n")
end

$LOAD_PATH.unshift("~/Code/Gems/rosey/lib")
require 'rosey/tasks/all'

$GO = ENV['GO'] || "go"
$FOREMAN = ENV['FOREMAN'] || "foreman"
$PUBLIC_DIR ||= "public"
$COMPILED_ASSETS_DIR ||= "#{$PUBLIC_DIR}/assets"
$PRODUCTION_BRANCH ||= "production"

task :default => [:build, :precompile_assets]

desc "Starts application server."
task :server => :install do
  sh "#{$FOREMAN} start"
end

desc "Installs backend application so it can be used by the server."
task :install => :build do
  sh "#{$GO} install ."
end

desc "Rebuilds backend application."
task :build do
  sh "#{$GO} build ."
end

desc "Runs all tests against backend."
task :test do
  sh "#{$GO} test ./..."
end

namespace :deploy do
  task :push do
    sh "#{$GIT} push heroku #{$PRODUCTION_BRANCH}:master"
  end

  task :prepare => :default
end
