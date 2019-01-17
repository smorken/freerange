
class ConnectScene extends Phaser.Scene {
  constructor () {
    super({ key: 'ConnectScene' })
  }

  preload () {
    this.load.image('play', 'assets/av-symbol/play button.png')
  }
  create () {
    ws.onopen = this.onconnected.bind(this)
    ws.onmessage = this.wsmessage.bind(this)
    this.message = 'connecting'
    this.loadtext = this.add.text(100, 200)
  }

  update (time, delta) {
    this.loadtext.setText(this.message)
  }

  showStartButton () {
    this.sprite = this.add.sprite(400, 300, 'play').setInteractive()
    this.sprite.parent = this
    this.sprite.on('pointerdown', function (pointer) {
      this.setTint(0xff0000)
      this.parent.startLoadScene()
    })
  }

  wsmessage (evt) {
    this.startLoadScene(evt.data)
  }

  onconnected () {
    this.message = 'loading'
    ws.send('request_assets')
  }

  startLoadScene (data) {
    this.scene.start('LoadScene', { assets: data })
  }
}
