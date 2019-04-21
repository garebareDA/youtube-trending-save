'use strict';
var path = location.pathname;
var pathDay = path.split('/');
var locationDay = parseInt(pathDay[2]);
function previous() {
    var previousDay = String(locationDay - 1);
    location.href = '/home/' + previousDay;
}
function next() {
    var nextDay = String(locationDay + 1);
    location.href = '/home/' + nextDay;
}
window.onload = function () {
    var date = new Date();
    var year = String(date.getFullYear());
    var month = String(date.getMonth() + 1);
    if (month.length == 1) {
        month = '0' + month;
    }
    var day = String(date.getDate());
    console.log(locationDay);
    console.log(year + month + day);
    if (String(locationDay) === year + month + day) {
        console.log('true');
        var button = document.getElementById('next');
        button.setAttribute("disabled", null);
    }
};
