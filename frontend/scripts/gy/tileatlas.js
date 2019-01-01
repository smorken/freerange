function TileAtlas(id, image, nrows, ncols, xsize, ysize) {
    this.id = id;
    this.image = image;
    this.nrows = nrows;
    this.ncols = ncols;
    this.xsize = xsize;
    this.ysize = ysize;

    this.draw = function(context, tile_col, tile_row, canvas_x_postion,
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
    };
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
function loadTileAtlas(tileAtlas, callback) {

    var n,
        result = {},
        count  = tileAtlas.length,
        onload = function() {
             if (--count == 0)
             callback(result);
            };
  
    for(n = 0 ; n < tileAtlas.length ; n++) {
        var id = tileAtlas[n]["id"];
        if(id in result){
            throw `duplicate tile data id: ${id}`;
        }
        var image = new Image();
        result[id] = new TileAtlas(
            id,
            image,
            tileAtlas[n]["nrows"],
            tileAtlas[n]["ncols"],
            tileAtlas[n]["xsize"],
            tileAtlas[n]["ysize"]);
        result[id].image.addEventListener("load", onload);
        result[id].image.src = tileAtlas[n]["src"];
    }  
}
