var Point = function(x, y) {
    if (x >= 0 && x <= 60 && y >= 0 && y <= 60)
        return {x: x, y: y};
    else {
        throw 'Invalid point';
    }
}

function Game(){
    this.snake = [new Point(30,30),new Point(31,30),new Point(32,30), new Point(33,30) ];
    this.direction = 'LEFT';
    this.food = [new Point(10,30)];
    this.score = 0;
    this.game_running = true;
    this.current_tick = 0;
}

function clearCanvas() {
    canvas = document.getElementById("canvas");
    context = canvas.getContext("2d");
    canvas_height = canvas.height;
    canvas_width = canvas.width;
    context.fillStyle = "black";
    context.fillRect(0,0,canvas_height,canvas_width);
}

function drawPoint(point, color) {
    canvas = document.getElementById("canvas");
    context = canvas.getContext("2d");
    canvas_height = canvas.height;
    canvas_width = canvas.width;
    var x_size = canvas_width / 60;
    var y_size = canvas_height / 60;
    var x_offset = point.x * x_size;
    var y_offset = point.y * y_size;
    context.fillStyle = color;
    context.fillRect(x_offset,y_offset,x_size,y_size);
}

Game.prototype.createFutureHeadPoint = function(head) {
    var possibleMovement = false;
    switch(this.direction){
        case 'LEFT':
            if ((head.x - 1) >= 0)
                return new Point((head.x - 1), head.y)
            break;
        case 'RIGHT':
            if ((head.x + 1) <= 60)
                return new Point((head.x + 1), head.y)
            break;
        case 'UP':
            if ((head.y - 1) >= 0)
                return new Point(head.x, head.y - 1)
            break;
        case 'DOWN':
            if ((head.y + 1) <= 60)
                return new Point(head.x, head.y + 1)
            break;
        };
        this.game_running = false;
};

Game.prototype.createNewFood = function() {
    var newFoodX = Math.ceil(Math.random() * 60);
    var newFoodY = Math.ceil(Math.random() * 60);
    this.food.push(new Point(newFoodX,newFoodY));
}

Game.prototype.collisionsWithBody = function(futureHead) {
    if (futureHead != undefined) {
        for (var i =0; i< this.snake.length; i++){
            if (this.snake[i].x === futureHead.x &&
                this.snake[i].y === futureHead.y){
                return true;
            }
        }
    }
    return false;
}

Game.prototype.updateGame = function() {
    if (this.current_tick == 0) this.createNewFood();
    var newSnake = new Array();
    var eating = false;
    //Check if ete anything
    for (var i=0;i<this.snake.length;i++){
        for(var j=0;j<this.food.length;j++){
            if (this.snake[i].x === this.food[j].x &&
                this.snake[i].y === this.food[j].y ) {
                this.score += 20;
                eating = true;
                this.food.splice(j,1);
            }
        }
    }
    var futureHead = this.createFutureHeadPoint(this.snake[0]);
    if (this.collisionsWithBody(futureHead)) {
        this.game_running = false;
        return;
    }
    newSnake.push(futureHead);
    for(var i=1; i<this.snake.length;i++){
        newSnake.push(this.snake[i-1]);
    }
    if (eating){
        newSnake.push(this.snake[this.snake.length-1]);
    }
    this.snake = newSnake;
};

Game.prototype.tick = function() {
    if (this.game_running){
        //Update the Game
        this.updateGame();
        //Draw stuff
        if (this.game_running == false) {
            return;
        }
        clearCanvas();
        //Draw Snake items
        for (var i=0; i < this.snake.length; i++){
            drawPoint(this.snake[i],"green");
        }
        //Draw the snake food
        for (var i=0; i < this.food.length; i++){
            drawPoint(this.food[i],"orange");
        }

        //Print Score
        var printScore = "<p>Current score:" + parseInt(this.score) + "</p>";
        document.getElementById("score").innerHTML = printScore;

        //update current tick
        this.current_tick += 1;
        this.current_tick = this.current_tick % 20;
    }
};

Game.prototype.setDirection = function(direction) {
    this.direction = direction;
};

window.addEventListener('keydown', function(event) {
  switch (event.keyCode) {
    case 37:
      game.setDirection('LEFT');
    break;
    case 38: 
      game.setDirection('UP');
    break;
    case 39:
      game.setDirection('RIGHT');
    break;
    case 40:
      game.setDirection('DOWN');
    break;
  }
}, false);
