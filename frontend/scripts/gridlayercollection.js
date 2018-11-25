
function GridLayerCollection(worldConfig){
    this.WorldConfig = JSON.parse(worldConfig);
    //the number of cols for the grids in the collection
    this.x_size = this.WorldConfig["x_size"];
    //the number of rows for the grids in the collection
    this.y_size = this.WorldConfig["y_size"];
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
}
