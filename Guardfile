# -*- ruby -*-

notification :off

group :public do
  guard :livereload do
    watch %r{public/.*$}
  end
end

group :assets do
  guard :shell do
    watch %r{^((assets/(.+/)?.+\.(js|jst|css))|(config/assets.yml))$} do |m|
      system 'make precompile_assets DEV=1'
      n "Assets recompiled (changed file: #{m[0]})"
    end
  end
end
