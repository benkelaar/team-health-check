"use strict";

var THC = {};

THC.getParameterByName = function (name) {
  name = name.replace(/[\[\]]/g, "\\$&");
    var url = window.location.href,
        regex = new RegExp('[?&]' + name + '(=([^&#]*)|&|#|$)'),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, " "));
};

THC.populateTable = function (data, team) {
  var table = $('#checks').empty()
    .append($('<thead><tr><th>'+team+'</th></tr></thead><tbody />'));

  function showValue(name, value, column) {
    var id  = '#' + name,
        row = table.find(id),
        img = value.level + '-' + value.direction,
        row = row.length ? row : table.find('tbody')
            .append($('<tr id="'+name+'"><td>'+name+'</td></tr>')).find(id),
        cur = row.find('td').length

    // Add missing columns if previous dates did not have a check
    if (cur != column) {
      for(var i = cur; i < column; i++) {
        row.append($('<td />'));
      }
    }
    row.append($('<td><img src="images/'+img+'.png"></td>'));
  }

  function showCheck(check) {
    var time   = new Date(check.time),
        day    = time.getDate() + '-' + (time.getMonth() + 1),
        column = table.find('thead tr').append($('<th>'+day+'</th>'))
                      .find('th').length - 1;
    for (var key in check.health) {
      showValue(key, check.health[key], column);
    }
  }

  data.forEach(showCheck);
};

THC.loadTable = function (team) {
  $.ajax('http://localhost:3000/v1/checks?team=' + team, {
    dataType: 'json',
    success: function (data) {
      THC.populateTable(data, team);
    }
  });
};

$(document).ready(function () {
  THC.team = THC.getParameterByName('team');
  if (!THC.team) THC.team = 'Team15b';
  THC.loadTable(THC.team);
});
