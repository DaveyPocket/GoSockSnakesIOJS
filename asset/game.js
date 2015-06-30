var Point = function(x, y) {
    if (x >= 0 && x <= 60 && y >= 0 && y <= 60)
        return {x: x, y: y};
    else {
        throw 'Invalid point';
    }
}

// Screen Related
function Screen(){
    this.canvas = document.getElementById("canvas");
    this.context = canvas.getContext("2d");
}

Screen.prototype.clearCanvas = function() {
    canvas_height = this.canvas.height;
    canvas_width = this.canvas.width;
    this.context.fillStyle = "black";
    this.context.fillRect(0,0,canvas_height,canvas_width);
}

Screen.prototype.drawPoint = function (point, color) {
    canvas_height = this.canvas.height;
    canvas_width = this.canvas.width;
    var x_size = canvas_width / 60;
    var y_size = canvas_height / 60;
    var x_offset = point.x * x_size;
    var y_offset = point.y * y_size;
    this.context.fillStyle = color;
    this.context.fillRect(x_offset,y_offset,x_size,y_size);
}

// Game Related
function Game(currentStatus){
    this.snakes = currentStatus.snakes;
    this.food = currentStatus.food;
    this.screen = new Screen();
}

Game.prototype.tick = function(){
    this.screen.clearCanvas();

    //Draw Snakes
    for (var color in this.snakes) {
        for (var i=0; i< this.snakes[color].points.length; i++){
            this.screen.drawPoint(this.snakes[color].points[i], color);
        }
    }
    //Draw Food
    for (var i=0; i< this.food.length ; i++){
        this.screen.drawPoint(this.food[i], "red");
    }
}