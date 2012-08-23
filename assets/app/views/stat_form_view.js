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

    initialize: function(url) {
        this.url = url || "http://"
        this.template = JST["stat_form"]
    },
    
    render: function() {
        this.$el.html(this.template())
        this.$el.find("#stat_url").val(this.url)
        this.submitButton = (new Egoistat.StatFormSubmitButton()).render()
    }
})
