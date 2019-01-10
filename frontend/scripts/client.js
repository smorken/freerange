class Client {
  constructor (ws) {
    this.websocket = ws
    ws.onmessage = this.getWS.bind(this)

    this.messages = []
  }

  sendWS (msg) {
    this.websocket.send(msg)
  }

  getWS (evt) {
    this.messages.push(evt.data)
  }

  errorWS (evt) {
    // exit the game scene
  }

  requestLevelAssets(){
    this.sendWS("request_level_assets")
  }
}
