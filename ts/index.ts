'use strict';

const path:string = location.pathname;
const pathDay:string[] = path.split('/');
const locationDay:number = parseInt(pathDay[2]);

function previous() {
  const previousDay:string = String(locationDay - 1);
  location.href = '/home/' + previousDay;
}

function next() {
  const nextDay:string = String(locationDay + 1);
  location.href = '/home/' + nextDay;
}

window.onload = () => {

  const date:Date = new Date();
  const year = String(date.getFullYear());
  let month = String(date.getMonth() + 1);

  if (month.length == 1) {
    month = '0' + month;
  }

  const day = String(date.getDate());

  console.log(locationDay)
  console.log(year + month + day)

  if(String(locationDay) === year + month + day){
    console.log('true')
    const button:HTMLElement = document.getElementById('next');
    button.setAttribute("disabled",null);
  }

}