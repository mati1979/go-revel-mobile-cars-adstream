$(function ($, w) {
        console.log("heatmap")
        var $mapCanvas = $("#map-canvas")[0];

        heatmap = {
            initialize: function() {
                console.log("heatmap init")
                var mapOptions = {
                    zoom: 6,
                    mapTypeId: google.maps.MapTypeId.ROADMAP,
                    mapTypeControl: false,
                    panControl: false,
                    zoomControl: true,
                    zoomControlOptions: {
                        style: google.maps.ZoomControlStyle.SMALL,
                        position: google.maps.ControlPosition.TOP_RIGHT
                    },
                    scaleControl: false,
                    scrollwheel: false,
                    streetViewControl: false,
                    center: new google.maps.LatLng(50.983, 10.317)
                };

                var map = new google.maps.Map($mapCanvas, mapOptions);
            },

            loadScript: function() {
                $.ajax({
                    url: 'http://maps.google.com/maps/api/js?v=3&sensor=false&callback=heatmap.initialize',
                    dataType: 'script',
                    cache: true
                });
            }

        }

        console.log("heatmap load")
        $(window).load(function () {
            window.setTimeout(function () {
                heatmap.loadScript();
            }, 0)
        });
    }
);
