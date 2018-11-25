function createTestTileData(callback){
    loadTileAtlas([
        {
            "id": 1,
            "src":"testgrid.png",
            "nrows": 5,
            "ncols": 5,
            "xsize": 56,
            "ysize": 56 
        },
        {
            "id": 3,
            "src": "testgrid2.png",
            "nrows": 5,
            "ncols": 5,
            "xsize": 56,
            "ysize": 56 
        }
    ],callback);
}


function tileatlas_test_basic_drawing(test){

    test.appendText("draws four tiles to the canvas");
    var displayResult = function (result){
        testCanvas = test.getCanvas();
        testCanvas.height = 80;
        var context = testCanvas.getContext("2d");
        result[1].draw(context,0,0,0,0,40,40);
        result[1].draw(context,0,0,40,0,40,40);
        result[1].draw(context,1,1,0,40,40,40);
        result[1].draw(context,1,0,40,40,40,40);
    }


    //debugger;
    createTestTileData(displayResult);
}

function tileatlas_test_2_layers(test){

    test.appendText("draws four tiles from the first tile atlas and four from the second tile atlas to the canvas");
    var displayResult = function (result){
        testCanvas = test.getCanvas();
        testCanvas.height = 160;
        var context = testCanvas.getContext("2d");
        result[1].draw(context,0,0,0,0,40,40);
        result[1].draw(context,0,0,40,0,40,40);
        result[1].draw(context,1,1,0,40,40,40);
        result[1].draw(context,1,0,40,40,40,40);
        //result 3 is intentional to check the dictionary is working
        result[3].draw(context,0,0,0,80,40,40); 
        result[3].draw(context,0,0,40,80,40,40);
        result[3].draw(context,1,1,0,120,40,40);
        result[3].draw(context,1,0,40,120,40,40);
    }


    //debugger;
    createTestTileData(displayResult);
}

