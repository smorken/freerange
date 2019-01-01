
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