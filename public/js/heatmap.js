$(function ($, w) {
        var $mapCanvas = $("#map-canvas")[0];
        var map;
        var count = 0;
        var adData = []
        var pointArray;
        var heat;

        heatmap = {
            initialize: function() {
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
                pointArray = new google.maps.MVCArray(adData);
                heat = new google.maps.visualization.HeatmapLayer({
                    data: pointArray,
                    radius: 20,
                    map: map
                });
                heat.setMap(map);
                heatmap.initSocket();
            },

            update: function(lat, lon) {
                if (map != null) {
                  var latlon = new google.maps.LatLng(lat, lon);
                  pointArray.push(latlon);
//                var marker = new google.maps.Marker({
//                    position: latlon,
//                    map: map
//                });
//                    if (pointArray.length >= 100000) {
//                        pointArray.shift()
//                    }
                }
            },

            initSocket: function() {
                var socket = new WebSocket('ws://' + window.location.host + '/websocket/ad/stream')
                socket.onmessage = function(event) {
                    count = count + 1;
                    var obj = JSON.parse(event.data)
                    $('#messages').empty();
                    $('#messages').append("<p>" + count + "</p>");
                    if (obj.Lat > 0 && obj.Lon > 0) {
                        heatmap.update(obj.Lat, obj.Lon)
                    }
                }
            },

            loadScript: function() {
                $.ajax({
                    url: 'http://maps.google.com/maps/api/js?v=3&sensor=false&libraries=visualization&callback=heatmap.initialize',
                    dataType: 'script',
                    cache: true
                });
            }

        }

        $(window).load(function () {
            window.setTimeout(function () {
                heatmap.loadScript();
            }, 0)
        });
    }
);
