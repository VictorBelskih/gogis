var map = L.map('map').setView([53.7215862, 91.4612620], 13);
    
var googleLayer = L.tileLayer('http://mt0.google.com/vt/lyrs=y&hl=en&x={x}&y={y}&z={z}', {
    attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors',
    maxZoom: 18
});


var openStreetMapLayer = L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors',
    maxZoom: 18
});
var geojson = {{.fields}}   
var geojsonLayer = L.geoJSON(geojson, {              
    style: function(feature) {

    var tluValue = feature.properties.Tlu;
    var color;

    switch (tluValue) {
        case 102:
            color = 'red';
            break;
        case 200:
            color = 'blue';
            break;
        case 300:
            color = 'green';
            break;
        case 500:
            color = 'yellow';
            break;
        default:
            color = 'gray';
    }
    return {
        fillColor: color,
        fillOpacity: 0.1,
        color: 'black',
        weight: 1
    };
},
onEachFeature: function(feature, layer) {
    layer.bindPopup('Номер хозяйства : ' + feature.properties.Farm_name + '<br>Номер поля: ' + feature.properties.Id_eu +
    '<br>TLU: ' + feature.properties.Tlu_name
    + '<br>Площадь: ' + feature.properties.Area_f  + '<br>Гумус: ' + feature.properties.Organic + '<br>Фосфор: ' + feature.properties.El_p
    + '<br>Калий: ' + feature.properties.El_k + '<br>Класс: ' + feature.properties.Humus_class
);
var selectedPolygon = null;
var originalStyle = null;
layer.on('click', function(e) {

    if (e.target.options.weight === 2) {               
        e.target.setStyle({ weight: 1, color:black});
    } else {
        geojsonLayer.eachLayer(function(layer) {
        if (layer.options.weight === 2) {
        layer.setStyle({ weight: 1, color: 'black' });
        }
        });
        e.target.setStyle({ weight: 2, color: 'red' });
    }
    var id_g = feature.properties.Id_eu;
    console.log('ID value:', id_g);
    });
    }
}).addTo(map);
            
var humusLayer = L.geoJSON(geojson, {

    style: function(feature) {

        var HumusClass = feature.properties.Humus_class;
        var color;

        switch (HumusClass) {
            case "Очень низкое":
                color = 'red';
                break;
            case "Низкое":
                color = 'orange';
                break;
            case "Среднее":
                color = 'yellow';
                break;
            case "Повышенное":
                color = 'green';
                break;
            case "Высокое":
                color = 'blue';
                break; 
            case "Очень Высокое":
                color = '#002F55';
                break;     
            default:
                color = '#3f00ff';
            }

            return {
                fillColor: color,
                fillOpacity: 0.7,
                color: 'black',
                weight: 1
            };
    },

onEachFeature: function(feature, layer) {
    layer.bindPopup('Номер хозяйства : ' + feature.properties.Farm_name + '<br>Номер поля: ' + feature.properties.Id_eu +
    '<br>TLU: ' + feature.properties.Tlu_name
    + '<br>Площадь: ' + feature.properties.Area_f  + '<br>Гумус: ' + feature.properties.Organic + '<br>Фосфор: ' + feature.properties.El_p
    + '<br>Калий: ' + feature.properties.El_k + '<br>Класс: ' + feature.properties.Humus_class
);
        
var selectedPolygon = null;
var originalStyle = null;
layer.on('click', function(e) {

    if (e.target.options.weight === 2) {
    // Снятие выделения с выбранного полигона
        e.target.setStyle({ weight: 1, color:black});
    } else {
        humusLayer.eachLayer(function(layer) {
            if (layer.options.weight === 2) {
                layer.setStyle({ weight: 1, color: 'black' });
            }
        });

                // Выделение текущего полигона
        e.target.setStyle({ weight: 2, color: 'red' });
        }

        var id_g = feature.properties.Id_eu;
        console.log('ID value:', id_g);
});
}
}).addTo(map);

var overlayMaps = {
    "GeoJSON": geojsonLayer,
    "Google": googleLayer,
    "OpenStreetMap": openStreetMapLayer,
    "Gumes": humusLayer,
};
        
L.control.layers(null, overlayMaps, { collapsed: true }).addTo(map);
    map.addLayer(googleLayer);

map.on('click', function(e) {
    geojsonLayer.eachLayer(function(layer) {
        if (layer.options.weight === 2 && layer.options.color === 'red') {
            layer.setStyle({ weight: 1, color: 'black' });
        }
    });
        
});
        
var elements = document.getElementsByClassName('leaflet-attribution-flag');
Array.from(elements).forEach(function(element) {
    element.style.height = '0px';
    element.style.width = '0px';
});
var selectControl = L.Control.extend({
    onAdd: function(map) {
        var select = L.DomUtil.create('select');
        select.id = 'selectbox';
        
        var options = {
            "geojson": "GeoJSON",
            "humus": "Humus"
        };

        for (var key in options) {
            var option = document.createElement("option");
            option.value = key;
            option.text = options[key];
            select.appendChild(option);
        }

        L.DomEvent.addListener(select, 'change', function() {
            var selectedValue = select.options[select.selectedIndex].value;
            // Do something with the selected value
            console.log(selectedValue);
        });

        return select;
    }
});

map.addControl(new selectControl());       
        