
//create a new server connection.
function ServerConnection(websocket){
    this.ws = websocket

    this.onmessage = function(evt) {
        print("RESPONSE: " + evt.data);
    }
    this.send = function(msg){
        this.ws.send(msg)
    }

    this.loadLevel()
}

function ServerToClientMessages(){
    this.messageIds = Object.freeze({
        "login": 0, //prompts the client to log into the server
        "invalid_login": 1, //returned when specified credetial does not work
        "tileConfig": 2, //the list of tile atlas files/dimensions
        "grid": 3, //the level grid data, sent when a player enters a level
        "grid_update": 4,//changes to the level grid data, sent periodically
        "actor_update": 5,//changes to the visible actors (player's self and others) 

    })

    this.parse()
}
function ClientToServerMessages() {
    this.CommandIds = Object.freeze({
        "login": 0,
        "move": 1,
        "action": 2,
        "jump": 3,
        "talk": 4
    })
    this.Directions = Object.freeze({
        "Up": 0, "Down": 1, "Left": 2, "Right": 3
    })

    this.createLoginMessage = function(username, password){
        return [this.CommandIds["login"],username,password]
    }
    this.createActionMesssage = function(direction){
        return [
            this.CommandIds["action"],
            this.Directions[direction]
        ]
    }
    this.createMoveMessage = function(direction){
        return [
            this.CommandIds["move"],
            this.Directions[direction]
        ]
    }

    this.createJumpMessage = function(){
        return [this.CommandIds["jump"]]
    }
}
function Actor(actorConfig) {
    this.x_position = 0;
    this.y_position = 0;
}

const checkImage = path =>
    new Promise(resolve => {
        const img = new Image();
        img.onload = () => resolve({path, status: 'ok'});
        img.onerror = () => resolve({path, status: 'error'});

        img.src = path;
    });

function TileData(id, url, nrows, ncols, xsize, ysize) {
    this.Image = new Image();
    this.img_promise = new Promise((resolve, reject) => {
        image.onload = function(){
            resolve(image);
        }
        image.onerror = function(){
            reject(`error loading image: ${url}`);
        }
        image.url = url;
    }).then();

    this.GetImage(){
        return this.img_promise.
    }
}
function SideScrollGameWorld(worldConfig){
    this.WorldConfig = JSON.parse(worldConfig);
    this.numCols = this.WorldConfig["x_size"];
    this.numRows = this.WorldConfig["y_size"];
    this.grid_size_x = this.WorldConfig["grid_size_x"];
    this.grid_size_y = this.WorldConfig["grid_size_y"];

    this.load

}

function SideScrollCanvasView(canvas){
    this.canvas = canvas;
    this.context = canvas.getContext("2d")
    this.size_x = canvas.width;
    this.size_y = canvas.height;
    this.offset_x = 0;
    this.offset_y = 0;

    this.clear = function() {
        this.context.clearRect(0, 0, this.size_x, this.size_y);
    }

    this.render = function(){

    }

    this.focusOn = function(actor){
        offset_x = actor.x_position - this.size_x/2
        offset_y = actor.y_position - this.size_y/2
    }
    this.minVisibleColumn() = function(world){
        return Math.floor(this.offset_x/world.grid_size_x)
    }
    this.minVisibleRow() = function(world){
        return Math.floor(this.offset_y/world.grid_size_y)
    }

    this.maxVisibleColumn() = function(world){
        return Math.floor((this.offset_x + size_x)/world.grid_size_x)
    }

    this.maxVisibleRow() = function(world){
        return Math.floor((this.offset_y + size_y)/world.grid_size_y)
    }

}