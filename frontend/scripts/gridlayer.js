function GridLayer(nrows, ncols, tile_atlas_id){



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
}