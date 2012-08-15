# -*- ruby -*-

notification :off

group :public do
  guard :livereload do
    watch %r{public/.*$}
  end
end

group :assets do
  guard :shell do
    watch %r{^assets/js/.+\.js$} do |m|
      system 'make scripts DEV=1'
      n "Scripts recompiled (changed file: #{m[0]})"
    end

    watch %r{^assets/css/.+\.css} do |m|
      system 'make styles DEV=1'
      n "Styles recompiled (changed file: #{m[0]})"
    end
  end
end
