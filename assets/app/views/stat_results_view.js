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
