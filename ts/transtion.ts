'use strict';

window.onload = () => {

  const date:Date = new Date();
  const year = String(date.getFullYear());
  let month = String(date.getMonth() + 1);

  if (month.length == 1) {
    month = '0' + month
  }

  const day = String(date.getDate());

  window.location.href = '/home/' + year + month + day;
}