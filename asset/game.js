var Point = function(x, y) {
    if (x >= 0 && x <= 60 && y >= 0 && y <= 60)
        return {x: x, y: y};
    else {
        throw new 'Invalid point';
    }
}

// Screen Related
function Screen(){
    this.canvas = document.getElementById("canvas");
    this.context = canvas.getContext("2d");

    this.clearCanvas = function() {
        canvas_height = this.canvas.height;
        canvas_width = this.canvas.width;
        this.context.fillStyle = "black";
        this.context.fillRect(0,0,canvas_height,canvas_width);
    }

    this.drawPoint = function (point, color) {
        canvas_height = this.canvas.height;
        canvas_width = this.canvas.width;
        var x_size = canvas_width / 60;
        var y_size = canvas_height / 60;
        var x_offset = point.x * x_size;
        var y_offset = point.y * y_size;
        this.context.fillStyle = color;
        this.context.fillRect(x_offset,y_offset,x_size,y_size);
    }
}

// Game Related
function Game(currentStatus){
    this.snakes = currentStatus.snakes;
    this.food = currentStatus.food;
    this.screen = new Screen();

    this.drawOnScreen = function(){
        this.screen.clearCanvas();
        for (var color in this.snakes) {
            for (var i=0; i< this.snakes[color].points.length; i++){
                this.screen.drawPoint(this.snakes[color].points[i], color);
            }
        }
        for (var i=0; i< this.food.length ; i++){
            this.screen.drawPoint(this.food[i], "red");
        }
    }

    this.computeNewBody = function(points, direction){
        var newBody = new Array();
        var newHead = undefined;
        switch (direction) {
            case 'LEFT':
                newHead = new Point(points[0].x -1, points[0].y);
                break;
            case 'RIGHT':
                newHead = new Point(points[0].x +1, points[0].y);
                break;
            case 'UP':
                newHead = new Point(points[0].x, points[0].y - 1);
                break;
            case 'DOWN':
                newHead = new Point(points[0].x, points[0].y + 1);
                break;
        }
        newBody.push(newHead);
        for (var i=0; i < points.length -1; i++) {
            newBody.push(points[i]);
        }
        return newBody;
    }

    this.updateSnakePositions = function() {
        for (var color in this.snakes){
            this.snakes[color].points = this.computeNewBody(this.snakes[color].points,this.snakes[color].direction);
        }
    }

    this.tick = function() {
        this.updateSnakePositions();
        this.drawOnScreen();
    }
}