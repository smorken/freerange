
class LoadScene extends Phaser.Scene {
  constructor () {
    super({ key: 'LoadScene' })
  }

  preload () {
    this.load.image('bg', 'https://twemoji.maxcdn.com/72x72/1f306.png')
    this.load.image('santa', 'https://twemoji.maxcdn.com/36x36/2b1c.png')
    this.load.image('ground', 'assets/platform.png')

    ws = new WebSocket('ws://localhost:8080/client')
    ws.onopen = this.openWS
  }

  create () {

  }

  // update(){
  //
  // }

  // var isOpenWs = false;
  openWS (evt) {
    game.scene.add('GameScene', GameScene, true, { x: 400, y: 300 })
    game.scene.start('GameScene')
    // isOpenWs = true;
  }
}
