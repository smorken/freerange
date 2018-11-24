
function canvasview_test_render(test){
    var displayResult = function(result){
        function mockGridLayer(){
            this.getvalue = function(row, col){
                return [Math.floor(row / 5), Math.floor(row % 5)]
            }
            this.tile_atlas_id = 1
        }
        testCanvas.height = 27*10;
        testCanvas.width = 27*10;
        canvasView = new CanvasView(testCanvas)
        canvasView.render(
            new mockGridLayer(), 25, 25, result);
    }

    createTestTileData(displayResult)

}