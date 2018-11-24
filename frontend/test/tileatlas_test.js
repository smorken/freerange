function createTestTileData(callback){
    loadTileAtlas([
        {
            "id": 1,
            "src":"testgrid.png",
            "nrows": 5,
            "ncols": 5,
            "xsize": 56,
            "ysize": 56 
        }
    ],callback);
}


function tileatlas_test_basic_drawing(testCanvas){

    var displayResult = function (result){
        var context = testCanvas.getContext("2d");
        result[1].draw(context,0,0,0,0,40,40);
        result[1].draw(context,0,0,40,0,40,40);
        result[1].draw(context,1,1,0,40,40,40);
        result[1].draw(context,1,0,40,40,40,40);
    }


    //debugger;
    createTestTileData(displayResult);
}
