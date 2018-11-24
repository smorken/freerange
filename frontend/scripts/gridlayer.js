function GridLayer(nrows, ncols, tile_atlas_id){

    //fields for data in each cell
    const TileAtlasRow_Field = 0;
    const TileAtlasColumn_Field = 1;

    this.numRows = nrows;
    this.numCols = ncols;
    this.tile_atlas_id = tile_atlas_id;
    this.data = {};

    this.getvalue = function(row, col){
        if(row in this.data && col in this.data[row]){
            return this.data[row][col];
        }
        return undefined;
    }
    // updates this layers data with x,y,value triples
    // ex: value = [[0,0,d(0,0)],[0,1,d(0,1), ... ]
    this.update = function(values){
        for(i = 0; i<values.length; i++){
            value = values[i];
            if(value[0] in this.data){
                this.data[value[0]][value[1]] = value[2];
            }
            else{
                this.data[value[0]] = { };
                this.data[value[0]][value[1]] = value[2];
            }

        }
    }

    //render this grid to the specified context using the specified tile atlas collection
    // @param canvasView an instance of TileCanvasView
    // @world TileBasedGameWorld instance
    // @param tile_atlas_collection a dictionary of tile atlas objects    
    this.render = function(canvasView, world, tile_atlas_collection){
        tile_atlas = tile_atlas_collection[this.tile_atlas_id]
        bounds = canvasView.getDrawBounds();
        for(row = bounds[0]; row<bounds[1]; row++){
            for(col = bounds[2]; col<bounds[3]; col++){

                cell_data = this.getvalue(row,col);
                
                if(cell_data){
                    tile_atlas.draw(
                        canvasView.context,
                        cell_data[TileAtlasRow_Field],
                        cell_data[TileAtlasColumn_Field],
                        row * world.grid_size_x - canvasView.offset_x,
                        col * world.grid_size_y - canvasView.offset_y,
                        world.grid_size_x,
                        world.grid_size_y
                    );
                }
            }
        }
    }
}