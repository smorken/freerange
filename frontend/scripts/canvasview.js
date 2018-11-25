//fields for data in each cell
const TileAtlasX_Field = 0;
const TileAtlasY_Field = 1;

class CanvasView {

    constructor(canvas){


        this.canvas = canvas;
        this.context = canvas.getContext("2d")

        this.offset_x = 0;
        this.offset_y = 0;
    }


    clear() {
        this.context.clearRect(0, 0, this.size_x(), this.size_y());
    }

    size_x(){
        return this.canvas.width;
    }

    size_y(){
        return this.canvas.height;
    }

    focusOn(actor){
        this.offset_x = actor.x_position - this.size_x()/2;
        if(this.offset_x < 0) { //enforce minimum offset, so that a gap isnt shown on the left
            this.offset_x = 0;
        }
        this.offset_y = actor.y_position - this.size_y()/2;
        if(this.offset_y < 0){
            this.offset_y = 0;
        }
    }

    setOffset(x,y){
        this.offset_x = x;
        this.offset_y = y;
    }
    
    //gets the draw bounds of this canvas view in world coordinates
    getDrawBounds(cell_size_x, cell_size_y){
        return [
            Math.floor(this.offset_x / cell_size_x),
            Math.ceil((this.offset_x + this.size_x())/cell_size_x),
            Math.floor(this.offset_y / cell_size_y),
            Math.ceil((this.offset_y + this.size_y())/cell_size_y)
        ];
    }

    //render the specified grid to this canvas view's context using the specified tile atlas collection
    // @param gridlayer an instance of gridlayer
    // @param sz_x the rendered size of each column in pixels
    // @param sz_y the rendered size of each row in pixels
    // @param tile_atlas_collection a dictionary of tile atlas objects    
    render(gridlayer, sz_x, sz_y, tile_atlas_collection){
        var tile_atlas = tile_atlas_collection[gridlayer.tile_atlas_id]
        var bounds = this.getDrawBounds(sz_x, sz_y);
        for(var x = bounds[0]; x<bounds[1]; x++){
            for(var y = bounds[2]; y<bounds[3]; y++){

                var cell_data = gridlayer.getvalue(y,x);
                
                if(cell_data){
                    tile_atlas.draw(
                        this.context,
                        cell_data[TileAtlasX_Field],
                        cell_data[TileAtlasY_Field],
                        x * sz_x - this.offset_x,
                        y * sz_y - this.offset_y,
                        sz_x,
                        sz_y
                    );
                }
            }
        }
    }
}