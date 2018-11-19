function TileData(id, image, nrows, ncols, xsize, ysize) {
    this.id = id;
    this.image = image;
    this.nrows = nrows;
    this.ncols = ncols;
    this.xsize = xsize;
    this.ysize = ysize;

    this.draw = function(context, tile_row, tile_col, canvas_x_postion,
        canvas_y_position, canvas_x_size, canvas_y_size ){
        context.drawImage(
            this.image, //img - Specifies the image, canvas, or video element to use	 
            tile_col*xsize, //sx - Optional. The x coordinate where to start clipping	
            tile_row*ysize, //sy - Optional. The y coordinate where to start clipping	
            xsize, //swidth - Optional. The width of the clipped image	
            ysize, //sheight - Optional. The height of the clipped image	
            canvas_x_postion, //x - The x coordinate where to place the image on the canvas	
            canvas_y_position,  //y - The y coordinate where to place the image on the canvas	
            canvas_x_size, //width - Optional. The width of the image to use (stretch or reduce the image)	
            canvas_y_size); //height - Optional. The height of the image to use (stretch or reduce the image)
    }
}

//loads tile atlases - based on the json tileData arg
//which contains a sequence of:
// id: identifier for the tile atlas
// image: the url of the image file to load
// nrows: the number of tile rows in the specified image
// ncols: the number of tile columns in the specified image
// xsize: the number of pixls that span the width of each column
// ysize: the number of pixels that span the height of each row
//calls callback function when all images are loaded.  
function loadTileData(tileData, callback) {

    var n,
        result = {},
        count  = tileData.length,
        onload = function() {
             if (--count == 0)
             callback(result);
            };
  
    for(n = 0 ; n < tileData.length ; n++) {
        var id = tileData[n]["id"];
        if(id in result){
            throw `duplicate tile data id: ${id}`
        }
        var image = new Image();
        result[id] = new TileData(
            id,
            image,
            tileData[n]["nrows"],
            tileData[n]["ncols"],
            tileData[n]["xsize"],
            tileData[n]["ysize"]);
        result[id].image.addEventListener("load", onload);
        result[id].image.src = tileData[n]["src"];
    }  
  }

function WorldGridLayer(nrows, ncols, tile_atlas_id){
    this.numRows = nrows
    this.numCols = ncols
    this.tile_atlas_id = tile_atlas_id
    this.data = {}

    this.getvalue = function(row, col){
        if(row in this.data && col in this.data[row]){
            return this.data[row][col]
        }
    }
    // updates this layers data with x,y,value triples
    // ex: value = [[0,0,d(0,0)],[0,1,d(0,1), ... ]
    this.update = function(value){
        for(i = 0; i<value.length; i++){
            this.data[value[0]][value[1]] = value[2]
        }
    }

}

function TileBasedGameWorld(worldConfig){
    this.WorldConfig = JSON.parse(worldConfig);
    //the number of columns in the world
    this.numCols = this.WorldConfig["numcols"];
    //the number of rows in the world
    this.numRows = this.WorldConfig["numrows"];
    //the final x render size of each tile 
    this.grid_size_x = this.WorldConfig["grid_size_x"];
    //the final y render size of each tile
    this.grid_size_y = this.WorldConfig["grid_size_y"];
    //a collection of layers
    this.layers = []

    this.addLayer = function(layer){
        this.layers.push(layer)
    }
    //get the origin x coordinate of the specified column
    this.getCoordinateX = function(col){
        return col*this.grid_size_x;
    }

    //get the origin y coordinate of the specified row
    this.getCoordinateY = function(row){
        return row*this.grid_size_y;
    }

    this.render = function(drawBounds, tileAtasCollection){
        


    }
}

function TileCanvasView(canvas){
    this.canvas = canvas;
    this.context = canvas.getContext("2d")
    this.size_x = canvas.width;
    this.size_y = canvas.height;
    this.offset_x = 0;
    this.offset_y = 0;

    this.clear = function() {
        this.context.clearRect(0, 0, this.size_x, this.size_y);
    }

    //render function
    //renders each item in gameworld grid_data sequentially
    this.render = function(gameworld, tileDataDict){

        var bounds = this.drawBounds(gameworld)
        //render all layers
        for(i = 0; i<gameworld.grid_data.length; i++){
            var griddata = gameworld.grid_data[i];
            var tileData = tileDataDict[ griddata["tile_atlas_id"]];
                        
            for(j = 0; j<griddata.length; j++){
                tileData.draw(
                    this.context,
                    griddata["tile_row"],
                    griddata["tile_col"],
                    griddata["world_row_x"] * world.grid_size_x - this.offset_x,
                    griddata["world_col_y"] * world.grid_size_y - this.offset_y,
                    gameworld.grid_size_x,
                    gameworld.grid_size_y
                );
            }
        }
    }

    this.focusOn = function(actor){
        offset_x = actor.x_position - this.size_x/2
        offset_y = actor.y_position - this.size_y/2
    }
    
    this.drawBounds() = function(world){
        return {
            "x1": Math.floor(this.offset_x/world.grid_size_x),
            "x2": Math.floor((this.offset_x + size_x)/world.grid_size_x),
            "y1": Math.floor(this.offset_y/world.grid_size_y),
            "y2": Math.floor((this.offset_y + size_y)/world.grid_size_y)
        }
    }
}