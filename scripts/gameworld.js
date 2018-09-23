function ServerConnection(websocket){

}

function Actor(actorConfig) {
    this.x_position = 0;
    this.y_position = 0;
}

function SideScrollGameWorld(worldConfig){
    this.WorldConfig = JSON.parse(worldConfig);
    this.numCols = this.WorldConfig["x_size"];
    this.numRows = this.WorldConfig["y_size"];
    this.grid_size_x = this.WorldConfig["grid_size_x"];
    this.grid_size_y = this.WorldConfig["grid_size_y"];
}

function SideScrollCanvasView(canvas){
    this.canvas = canvas;
    this.context = canvas.getContext("2d")
    this.size_x = canvas.width;
    this.size_y = canvas.height;
    this.offset_x = 0;
    this.offset_y = 0;

    this.clear = function() {
        this.context.clearRect(0, 0, this.size_x, this.size_y);
    }

    this.focusOn = function(actor){
        offset_x = actor.x_position - this.size_x/2
        offset_y = actor.y_position - this.size_y/2
    }
    this.minVisibleColumn() = function(world){
        return Math.floor(this.offset_x/world.grid_size_x)
    }
    this.minVisibleRow() = function(world){
        return Math.floor(this.offset_y/world.grid_size_y)
    }

    this.maxVisibleColumn() = function(world){
        return Math.floor((this.offset_x + size_x)/world.grid_size_x)
    }

    this.maxVisibleRow() = function(world){
        return Math.floor((this.offset_y + size_y)/world.grid_size_y)
    }

}