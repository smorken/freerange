
function canvasview_test_render(test){
    test.appendText("creates a canvas at position 0,0 and renders tiles where tile atlas indices are based on a function");
    var displayResult = function(result){
        function mockGridLayer(){
            this.getvalue = function(row, col){
                return [Math.floor(row / 5), Math.floor(row % 5)]
            }
            this.tile_atlas_id = 1
        }
        testCanvas = test.getCanvas();
        testCanvas.height = 27*10;
        testCanvas.width = 27*10;
        canvasView = new CanvasView(testCanvas)
        canvasView.render(
            new mockGridLayer(), 25, 25, result);
    }

    createTestTileData(displayResult)
}

function canvasview_test_render_and_clear(test){
    test.appendText("creates a canvas at position 0,0 and renders tiles where tile atlas indices are based on a function.  Canvas is cleared and re-rendered periodically.");
    var displayResult = function(result){
        function mockGridLayer(){
            this.getvalue = function(row, col){
                return [Math.floor(row / 5), Math.floor(row % 5)]
            }
            this.tile_atlas_id = 1
        }
        testCanvas = test.getCanvas();
        testCanvas.height = 27*10;
        testCanvas.width = 27*10;
        canvasView = new CanvasView(testCanvas)

        var clear = function() {
            canvasView.clear();
            setTimeout(render, 1000);
        }
        var render = function() {
            canvasView.render(
            new mockGridLayer(), 25, 25, result);
            setTimeout(clear, 1000);
        }
        render();
    }

    createTestTileData(displayResult)
}