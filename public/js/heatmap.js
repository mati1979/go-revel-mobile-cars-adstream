$(function ($, w) {
        var $mapCanvas = $("#map-canvas")[0];
        var map;
        var count = 0;
        var adData = []
        var pointArray;
        var heat;
        var markers = [];

        var styles = [
            {
                "featureType": "water",
                "elementType": "geometry",
                "stylers": [
                    {
                        "color": "#193341"
                    }
                ]
            },
            {
                "featureType": "landscape",
                "elementType": "geometry",
                "stylers": [
                    {
                        "color": "#2c5a71"
                    }
                ]
            },
            {
                "featureType": "road",
                "elementType": "geometry",
                "stylers": [
                    {
                        "color": "#29768a"
                    },
                    {
                        "lightness": -37
                    }
                ]
            },
            {
                "featureType": "poi",
                "elementType": "geometry",
                "stylers": [
                    {
                        "color": "#406d80"
                    }
                ]
            },
            {
                "featureType": "transit",
                "elementType": "geometry",
                "stylers": [
                    {
                        "color": "#406d80"
                    }
                ]
            },
            {
                "elementType": "labels.text.stroke",
                "stylers": [
                    {
                        "visibility": "on"
                    },
                    {
                        "color": "#3e606f"
                    },
                    {
                        "weight": 2
                    },
                    {
                        "gamma": 0.84
                    }
                ]
            },
            {
                "elementType": "labels.text.fill",
                "stylers": [
                    {
                        "color": "#ffffff"
                    }
                ]
            },
            {
                "featureType": "administrative",
                "elementType": "geometry",
                "stylers": [
                    {
                        "weight": 0.6
                    },
                    {
                        "color": "#1a3541"
                    }
                ]
            },
            {
                "elementType": "labels.icon",
                "stylers": [
                    {
                        "visibility": "off"
                    }
                ]
            },
            {
                "featureType": "poi.park",
                "elementType": "geometry",
                "stylers": [
                    {
                        "color": "#2c5a71"
                    }
                ]
            }
        ];

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
                    center: new google.maps.LatLng(50.983, 10.317),
                    styles: styles
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
                    var marker = new google.maps.Marker({
                        position: latlon,
                        map: map
                    });
                    markers.push(marker);
                    if (markers.length > 100) {
                        markers.shift().setMap(null);
                    }
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
