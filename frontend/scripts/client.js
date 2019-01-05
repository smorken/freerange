class Client {
  constructor (url, onOpen, onError) {
    this.websocket = new WebSocket(url)
    this.websocket.onopen = onOpen
    this.websocket.onerror = onError
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
}
