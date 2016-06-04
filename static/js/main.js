"use strict";

var THC = {};

THC.list = [
  'green-up',
  'green-same',
  'green-down',
  'yellow-up',
  'yellow-same',
  'yellow-down',
  'orange-up',
  'orange-same',
  'orange-down',
  'red-up',
  'red-same',
  'red-down',
];

THC.getParameterByName = function (name) {
  name = name.replace(/[\[\]]/g, "\\$&");
    var url = window.location.href,
        regex = new RegExp('[?&]' + name + '(=([^&#]*)|&|#|$)'),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, " "));
};

THC.activateModifiables = function () {
  function clickEvent(e) {
    var td   = $(e.currentTarget),
        from = td.data().lvl,
        now  = THC.list.indexOf(from),
        down = e.which > 1 || e.altKey || e.ctrlKey || e.metaKey || e.shiftKey,
        last = THC.list.length - 1,
        to   = down ? (now == last ? 0 : now + 1) : (now == 0 ? last : now - 1),
        next = THC.list[to];

    td.data('lvl', next);
    td.find('img').attr('src', 'images/'+next+'.png');
    return false;
  }

  $('.modifiable')
    .click(clickEvent)
    .contextmenu(clickEvent);
};

THC.populateTable = function (data) {
  if (!data || !data.length) {
    console.log('empty data returned from server');
    return;
  }
  var team = THC.team,
      table = $('#checks').empty()
    .append($('<thead><tr><th>'+team+'</th></tr></thead><tbody />'));

  function showValue(name, value, column) {
    var select  = 'tr[id=\'' + name + '\']',
        row     = table.find(select),
        imgName = value.level + '-' + value.direction,
        row     = row.length ? row :
            table.find('tbody')
                .append($('<tr id="'+name+'"><td>'+name+'</td></tr>')).find(select),
        cur     = row.find('td').length,
        img     = $('<td><img src="images/'+imgName+'.png"></td>');

    // Add missing columns if previous dates did not have a check
    if (cur != column) {
      for(var i = cur; i < column; i++) {
        row.append($('<td />'));
      }
    }
    img.data('lvl', imgName);
    row.append(img);
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
  showCheck(data[data.length - 1]);
  table.find('th:last-child').html('New');
  table.find('td:last-child').addClass('modifiable')

  THC.activateModifiables()
};

THC.loadTable = function () {
  $.ajax('http://localhost:3000/v1/checks?team=' + THC.team, {
    success: THC.populateTable
  });
};

THC.storeNewData = function () {
  var checks = $('#checks tbody tr'),
      data   = {name: THC.team, health: {}};
  checks.each(function (i, e) {
    var $e    = $(e),
        label = $e.find('td:first-child').html(),
        value = $e.find('td:last-child').data('lvl').split('-');
    data.health[label] = {level: value[0], direction: value[1]};
  });

  $.ajax({
    url: 'http://localhost:3000/v1/checks',
    method: 'POST',
    data: JSON.stringify(data),
    success: function (data) {
      location.reload();
    }
  });
};

$(document).ready(function () {
  THC.team = THC.getParameterByName('team');
  if (!THC.team) THC.team = 'Team15b';
  THC.loadTable();
  $('#save').click(THC.storeNewData);
});
