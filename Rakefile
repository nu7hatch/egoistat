# -*- ruby -*-

begin
  require 'bundler/setup'
rescue LoadError
  $stderr.write("ERROR: Bundler not installed, run `gem install bundler` to get it.\n")
end

GO                  = ENV['GO'] || "go"
GIT                 = ENV['GIT'] ||"git"
JAMMIT              = ENV['JAMMIT'] || "jammit"
BUNDLE              = ENV['BUNDLE'] || "bundle"
GUARD               = ENV['GUARD'] || "guard"
PUBLIC_DIR          = ENV['PUBLIC_DIR'] || "public"
COMPILED_ASSETS_DIR = ENV['COMPILED_ASSETS_DIR'] || "#{PUBLIC_DIR}/assets"
SERVER              = ENV['SERVER'] || "foreman start"
PRODUCTION_BRANCH   = ENV['PRODUCTION_BRANCH'] || "production"
DEV                 = ENV['DEV'] == 1

task :default => [:build, :precompile_assets]

desc "Starts application server."
task :server => :install do
  sh SERVER
end

desc "Installs backend application so it can be used by the server."
task :install => :build do
  sh "#{GO} install ."
end

desc "Rebuilds backend application."
task :build do
  sh "#{GO} build ."
end

desc "Runs all tests against backend."
task :test do
  sh "#{GO} test ./..."
end

desc "Precompiles static assets."
task :precompile_assets do
  sh "#{JAMMIT} -f -o #{COMPILED_ASSETS_DIR}"
end

desc "Starts watching static assets directory for changes."
task :watch_assets do
  sh "#{GUARD} start -i"
end

desc "Builds backend, compiles static assets and deploys the app."
task :deploy => "deploy:require_clean_tree" do
  puts "# Merging changes to master..."
  Rake::Task["deploy:merge_master"].invoke

  puts "# Building project and precompiling assets..."
  Rake::Task["default"].invoke

  puts "# Committing changes..."
  Rake::Task["deploy:commit_release"].invoke

  puts "# Release v#{$VERSION} ready, deploying..."
  Rake::Task["deploy:push"].invoke
  
  puts "# Cleaning up"
  Rake::Task["deploy:cleanup"].invoke
end

namespace :deploy do
  task :require_clean_tree do
    system "#{GIT} diff --quiet HEAD"
    raise "Uncommited changes detected, commit or stash them before deploy!" if $? != 0
  end
  
  task :merge_master do
    sh "#{GIT} reset HEAD ."
    sh "#{GIT} checkout -f production"
    sh "#{GIT} merge master -q --no-commit -s recursive -Xtheirs"
  end
  
  task :release_bump do
    $VERSION = File.read('VERSION') rescue 0
    $VERSION = $VERSION.to_i + 1
    File.open('VERSION', 'w+') { |f| f.write($VERSION) }
  end

  task :commit_release => "deploy:release_bump" do
    sh "#{GIT} add VERSION #{COMPILED_ASSETS_DIR}"
    sh "#{GIT} commit -qm 'Released v#{$VERSION}'"
  end

  task :push do
    sh "#{GIT} push heroku #{PRODUCTION_BRANCH}:master"
  end

  task :cleanup do
    sh "#{GIT} checkout master"
  end
end
