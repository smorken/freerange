function GridLayer(n_x, n_y, tile_atlas_id){



    this.N_X = n_x;
    this.N_Y = n_y;
    this.tile_atlas_id = tile_atlas_id;
    this.data = {};

    this.getvalue = function(x, y){
        if(x in this.data && y in this.data[x]){
            return this.data[x][y];
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