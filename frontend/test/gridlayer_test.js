function test2(testCanvas){
    var displayResult = function(result){
        
        var gw = new GridLayer(10,10,1)
        gw.update([
            [0,0,[0,0]],
            [1,1,[0,1]],
            [2,2,[0,2]],
            [3,3,[0,3]],
            [4,4,[0,4]],
            [5,5,[1,0]],
            [6,6,[1,1]],
            [7,7,[1,2]],
            [8,8,[1,3]],
            [9,9,[1,4]] //draws tile(2,2) at position 1,1
        ])

        function mockCanvasView(){
            this.context = testCanvas.getContext("2d");
            this.getDrawBounds = function(){
                return [0,10,0,10];
            }
            this.offset_x = 0;
            this.offset_y = 0;
        }

        function mockGridCollection(){
            this.grid_size_x = 25;
            this.grid_size_y = 25;
        }
        testCanvas.height = 25*10;
        testCanvas.width = 25*10;
        gw.render(
            new mockCanvasView(),
            new mockGridCollection(),
            result
        )
    }

    createTestTileData(displayResult)
}