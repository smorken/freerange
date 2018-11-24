
function CanvasView(canvas){

    //fields for data in each cell
    const TileAtlasRow_Field = 0;
    const TileAtlasColumn_Field = 1;


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
    this.getDrawBounds = function(cell_size_x, cell_size_y){
        return [
            Math.floor(this.offset_y / cell_size_y),
            Math.floor((this.offset_y + this.size_y)/cell_size_y),
            Math.floor(this.offset_x / cell_size_x),
            Math.floor((this.offset_x + this.size_x)/cell_size_x)
        ];
    }

    //render the specified grid to this canvas view's context using the specified tile atlas collection
    // @param gridlayer an instance of gridlayer
    // @param sz_x the rendered size of each column in pixels
    // @param sz_y the rendered size of each row in pixels
    // @param tile_atlas_collection a dictionary of tile atlas objects    
    this.render = function(gridlayer, sz_x, sz_y, tile_atlas_collection){
        tile_atlas = tile_atlas_collection[gridlayer.tile_atlas_id]
        bounds = this.getDrawBounds(sz_x, sz_y);
        for(row = bounds[0]; row<bounds[1]; row++){
            for(col = bounds[2]; col<bounds[3]; col++){
                debugger;
                cell_data = gridlayer.getvalue(row,col);
                
                if(cell_data){
                    tile_atlas.draw(
                        this.context,
                        cell_data[TileAtlasRow_Field],
                        cell_data[TileAtlasColumn_Field],
                        row * sz_y - this.offset_y,
                        col * sz_x - this.offset_x,
                        sz_x,
                        sz_y
                    );
                }
            }
        }
    }
}