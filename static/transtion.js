'use strict';
window.onload = function () {
    var date = new Date();
    var year = String(date.getFullYear());
    var month = String(date.getMonth() + 1);
    if (month.length == 1) {
        month = '0' + month;
    }
    var day = String(date.getDate());
    window.location.href = '/home/' + year + month + day;
};
