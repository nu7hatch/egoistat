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
