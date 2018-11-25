
class CanvasView {

    constructor(canvas){
        //fields for data in each cell
        const TileAtlasRow_Field = 0;
        const TileAtlasColumn_Field = 1;

        this.canvas = canvas;
        this.context = canvas.getContext("2d")

        this.offset_x = 0;
        this.offset_y = 0;
    }

    get size_x(){
        return this.canvas.size_x;
    }

    get size_y(){
        return this.canvas.size_y;
    }

    clear = function() {
        this.context.clearRect(0, 0, this.size_x, this.size_y);
    }

    focusOn = function(actor){
        this.offset_x = actor.x_position - this.size_x/2;
        this.offset_y = actor.y_position - this.size_y/2;
    }

    setOffset(x,y){
        this.offset_x = x;
        this.offset_y = y;
    }
    
    //gets the draw bounds of this canvas view in world coordinates
    getDrawBounds = function(cell_size_x, cell_size_y){
        return [
            Math.floor(this.offset_y / cell_size_y),
            Math.ceil((this.offset_y + this.size_y)/cell_size_y),
            Math.floor(this.offset_x / cell_size_x),
            Math.ceil((this.offset_x + this.size_x)/cell_size_x)
        ];
    }

    //render the specified grid to this canvas view's context using the specified tile atlas collection
    // @param gridlayer an instance of gridlayer
    // @param sz_x the rendered size of each column in pixels
    // @param sz_y the rendered size of each row in pixels
    // @param tile_atlas_collection a dictionary of tile atlas objects    
    render = function(gridlayer, sz_x, sz_y, tile_atlas_collection){
        tile_atlas = tile_atlas_collection[gridlayer.tile_atlas_id]
        bounds = this.getDrawBounds(sz_x, sz_y);
        for(row = bounds[0]; row<bounds[1]; row++){
            for(col = bounds[2]; col<bounds[3]; col++){

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