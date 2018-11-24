
function GridLayerCollection(worldConfig){
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
        this.layers.push(layer);
    }

    this.render = function(canvasView, tileAtlasCollection){
        for(i =0; i<this.layers.length; i++){
            this.layers[i].render(canvasView, this, tileAtlasCollection);
        }
    }
    //get the origin x coordinate of the specified column
    //this.getCoordinateX = function(col){
    //    return col*this.grid_size_x;
    //}

    //get the origin y coordinate of the specified row
    //this.getCoordinateY = function(row){
    //   return row*this.grid_size_y;
    //}
}
