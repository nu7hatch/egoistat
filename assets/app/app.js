/*!
 * Egoistat backbone app.
 * http://egoistat.com/
 *
 * Copyright (c) 2012 by Chris Kowalik a.k.a nu7hatch
 * Released under the AGPLv3 license
 */

var Egoistat = { fn: {}, JST: {} }
  , app

Egoistat.networks = [
    'twitter',
    'facebook',
    'plusone',
    'reddit',
    'hackernews'
]

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
        this.bindSocialButtons()
        $('[role="main"]').show()
    },

    bindSocialButtons: function() {
        var text = "Social popularity statistics for your website:"
          , url = "http://egoistat.com/"
    
        $('a[rel="facebook"]').facebookButton({ text: text, url: url })
        $('a[rel="twitter"]').twitterButton({ text: text, url: url, via: "nu7hatch" })
        $('a[rel="plusone"]').plusoneButton({ url: url })
    }
})

$(function() {
    app = Egoistat.router = new Egoistat.Router()
    Backbone.history.start({ pushState: true })
    app.showPage()
});

