
function CanvasView(canvas){
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
        offset_x = actor.x_position - this.size_x/2;
        offset_y = actor.y_position - this.size_y/2;
    }
    
    //gets the draw bounds of this canvas view in world coordinates
    this.getDrawBounds() = function(world){
        return [
            Math.floor(this.offset_y/world.grid_size_y),
            Math.floor((this.offset_y + size_y)/world.grid_size_y),
            Math.floor(this.offset_x/world.grid_size_x),
            Math.floor((this.offset_x + size_x)/world.grid_size_x)
        ];
    }
}