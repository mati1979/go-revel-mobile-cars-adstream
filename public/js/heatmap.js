$(function ($, w) {
        console.log("heatmap")
        var $mapCanvas = $("#map-canvas")[0];
        var adData = [];
        var map;
        var markers = [];

        heatmap = {
            initialize: function() {
                console.log("heatmap init")
                var mapOptions = {
                    zoom: 6,
                    mapTypeId: google.maps.MapTypeId.ROADMAP,
                    zoomControl: true,
                    zoomControlOptions: {
                        style: google.maps.ZoomControlStyle.SMALL,
                        position: google.maps.ControlPosition.TOP_RIGHT
                    },
                    center: new google.maps.LatLng(50.983, 10.317)
                };

                map = new google.maps.Map($mapCanvas, mapOptions);
            },

            updateHeatMap: function(lat, lon) {
                var latlon = new google.maps.LatLng(lat, lon);
                var marker = new google.maps.Marker({
                    position: latlon,
                    map: map
                });
                markers.push(marker);
                //console.log("update lat, lon:" + lat +":" + lon)
                adData.push(latlon);
                if (adData.length % 100 == 0) {
                    heatmap.initialize();
                    var pointArray = new google.maps.MVCArray(adData);
                    heat = new google.maps.visualization.HeatmapLayer({
                        data: pointArray,
                        radius: 20,
                        map: map
                    });
                    heat.setMap(map);
                    //adData = []
                    markers = []
                }
            },

            initSocket: function() {
                var socket = new WebSocket('ws://' + window.location.host + '/websocket/ad/stream')
                socket.onmessage = function(event) {
                    var obj = JSON.parse(event.data)
                    $('#messages').empty();
                    $('#messages').append("<p>" + event.data + "</p>");
                    if (obj.Lat > 0 && obj.Lon > 0) {
                        heatmap.updateHeatMap(obj.Lat, obj.Lon)
                    }
                }
            },

            loadScript: function() {
                heatmap.initSocket();
                $.ajax({
                    url: 'http://maps.google.com/maps/api/js?v=3&sensor=false&libraries=visualization&callback=heatmap.initialize',
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
