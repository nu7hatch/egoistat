# -*- ruby -*-

notification :off

guard :livereload do
  watch %r{public/.*$}
end

guard :shell do
  watch %r{^((assets/(.+/)?.+\.(js|jst|css))|(config/assets.yml))$} do |m|
    system 'rake assets:precompile DEV=1'
    n "Assets recompiled (changed files: #{m.join(', ')})"
  end
end
