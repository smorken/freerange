
function canvasview_test_render(test){
    test.appendText("creates a canvas at position 0,0 and renders tiles where tile atlas indices are based on a function");
    var displayResult = function(result){
        function mockGridLayer(){
            this.getvalue = function(x, y){
                return [Math.floor(y % 5), Math.floor(x % 5)];
            }
            this.tile_atlas_id = 1
        }
        var testCanvas = test.getCanvas();
        testCanvas.height = 27*10;
        testCanvas.width = 27*10;
        var canvasView = new CanvasView(testCanvas);
        canvasView.render(
            new mockGridLayer(), 25, 25, result);
    }

    createTestTileData(displayResult);
}

function canvasview_test_render_and_clear(test){
    test.appendText("creates a canvas at position 0,0 and renders tiles where tile atlas indices are based on a function.  Canvas is cleared and re-rendered periodically.");
    var displayResult = function(result){
        function mockGridLayer(){
            this.getvalue = function(x, y){
                return [Math.floor(y % 5), Math.floor(x % 5)];
            }
            this.tile_atlas_id = 1
        }
        var testCanvas = test.getCanvas();
        testCanvas.height = 140;
        testCanvas.width = 110;
        var canvasView = new CanvasView(testCanvas)

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

    createTestTileData(displayResult);
}

function canvasview_test_render_with_undefined_values_in_grid(test){
    test.appendText("tests the effect of undefined grid data, which should be a no-draw in the render function");
    var displayResult = function(result){
        function mockGridLayer(){
            this.data = { 
                0 : { 0: undefined, 1: [0,0], 2: undefined },
                1 : { 0: [1,1], 1: undefined, 2: [0,0] },
                2 : { 0: undefined, 1: [0,0], 2: undefined }
            };
            this.getvalue = function(row, col){
                return this.data[row][col];
            }
            this.tile_atlas_id = 1
        }
        var testCanvas = test.getCanvas();
        testCanvas.height = 75;
        testCanvas.width = 75;
        var canvasView = new CanvasView(testCanvas);
        canvasView.render(
            new mockGridLayer(), 25, 25, result);
    }

    createTestTileData(displayResult);
}

function canvasview_test_focusOn_method(test){
    test.appendText("tests the focus on method by setting it periodically and rendering the result");
    var displayResult = function(result){
        function mockGridLayer(){
            this.getvalue = function(x, y){
                return [Math.floor(y % 5), Math.floor(x % 5)];
            }
            this.tile_atlas_id = 1
        }
        var testCanvas = test.getCanvas();
        testCanvas.height = 100;
        testCanvas.width = 100;
        
        var canvasView = new CanvasView(testCanvas);

        var xpos = 0;
        var ypos = 0;
        
        var render = function(){
            
            canvasView.clear();
            function mockActor(x,y){
                this.x_position = x;
                this.y_position = y;
            }

            xpos = xpos < 100 ? xpos + 1 : 0;
            ypos = ypos < 100 ? ypos + 1 : 0;

            canvasView.focusOn(new mockActor(xpos,ypos))
            canvasView.render(
                new mockGridLayer(), 15, 15, result);
            setTimeout(render, 20);
        };

        render();
    }

    createTestTileData(displayResult);
}
