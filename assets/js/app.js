/*!
 * Egoistat backbone app.
 * http://egoistat.com/
 *
 * Copyright (c) 2012 by Chris Kowalik a.k.a nu7hatch
 * Released under the AGPLv3 license
 */

/**
 * @depend libs/jquery-1.8.0.js
 * @depend libs/jquery.base64.js
 * @depend libs/jquery.view.ejs.js
 * @depend libs/underscore.js
 * @depend libs/backbone.js
 * @depend social.js
 */

var Egoistat = { fn: {} }
  , app

Egoistat.networks = [
    'twitter',
    'facebook',
    'plusone',
    'reddit',
    'hackernews',
    'pinterest'
]

Egoistat.Stat = Backbone.Model.extend({
    url: function() {
        return "/api/v1/stat.json" + "?url=" + encodeURIComponent(this.address) + "&n=" + this.networks.join(",")
    },
    
    initialize: function(address, networks) {
        this.networks = networks || Egoistat.networks
        this.address = address
    },

    permalink: function() {
        var url = this.address
        return "stat/" + $.base64.encode(url) + "/" + this.networks.join(",")
    },

    parse: function(resp) {
        return resp
    }
})

Egoistat.StatFormSubmitButton = Backbone.View.extend({
    el: "#stat_submit",

    events: {
        'click': 'submit'
    },

    render: function() {
        this.url = $("#stat_url")
    },

    submit: function(e) {
        e.preventDefault()
        this.disable()

        var self = this
          , results = new Egoistat.StatResultsView(this.url.val())
        
        results.render(function() {
            self.enable()
        })
    },

    enable: function() {
        this.$el.removeAttr('disabled')
        return this
    },

    disable: function() {
        this.$el.attr('disabled', 'disabled')
    }
})

Egoistat.StatFormView = Backbone.View.extend({
    el: "#yield",
    template: "stat_form_tpl",

    initialize: function(url) {
        this.url = url || "http://"
    },
    
    render: function() {
        this.$el.html(this.template, {})
        this.$el.find("#stat_url").val(this.url)
        this.submitButton = (new Egoistat.StatFormSubmitButton()).render()
    }
})

Egoistat.StatResultsView = Backbone.View.extend({
    el: "#stat .results",
    
    initialize: function(url) {
        this.url = url
    },
    
    render: function(fn) {
        var self = this
          , stat = new Egoistat.Stat(this.url)
          , networks = this.$el.find('.network')
        
        app.navigate(stat.permalink())

        networks.each(function(_, n) {
            var $points = $(this).find('.points')
            $points.text('...')
        })

        mixpanel.track("Stats fetched", { "URL": this.url })
            
        stat.fetch({
            success: function(model, _) {
                networks.each(function(_, n) {
                    var value = stat.get($(n).attr("tag"))
                      , $points = $(this).find('.points')

                    $points.text(value)
                })
            },
            error: function(model, resp) {
                console.log(resp)
            }
        }).complete(function() {
            if (!!fn) fn()
        })
    }
})

Egoistat.Router = Backbone.Router.extend({
    routes: {
        '': 'home',
        '/': 'home',
        'stat': 'redirectHome',
        'stat/:url/:networks': 'stat'
    },

    home: function() {
        mixpanel.track("Landing page loaded")
        ;(new Egoistat.StatFormView()).render()
    },

    redirectHome: function() {
        this.navigate("/", true)
    },

    stat: function(urlhash, networks) {
        var url = $.base64.decode(urlhash)

        ;(new Egoistat.StatFormView(url)).render()
        ;(new Egoistat.StatResultsView(url)).render()
    },

    showPage: function() {
        $('[role="main"]').show()
    }
})

Egoistat.fn.fillSocialCounters = function() {
    $('a.socialBtn').fillSocialCounter()
}

;(function() {
    var text = "Social popularity statistics for your website:"
      , url = "http://egoistat.com/"
    
    $('a[rel="facebook"]').facebookButton({ text: text, url: url })
    $('a[rel="twitter"]').twitterButton({ text: text, url: url, via: "nu7hatch" })
    $('a[rel="plusone"]').plusoneButton({ url: url })

    app = Egoistat.router = new Egoistat.Router()
    Backbone.history.start({ pushState: true })
    app.showPage()
}())
