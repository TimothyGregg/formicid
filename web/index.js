document.addEventListener("DOMContentLoaded", function(){
    main();
});

var canvas = document.getElementById("myCanvas");
var ctx = canvas.getContext("2d");

function main() {
    ctx.beginPath();
    ctx.rect(20, 40, 50, 50);
    ctx.fillStyle = "#FF0000";
    ctx.fill();
    ctx.closePath();
}