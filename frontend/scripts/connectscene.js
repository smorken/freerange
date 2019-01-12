
class ConnectScene extends Phaser.Scene {
  constructor () {
    super({ key: 'ConnectScene' })
  }

  preload () {
    this.load.image('play', 'assets/av-symbol/play button.png')
  }
  create () {
    ws.onopen = this.showStartButton.bind(this)
  }
  showStartButton () {
    this.sprite = this.add.sprite(400, 300, 'play').setInteractive()
    this.sprite.parent = this
    this.sprite.on('pointerdown', function (pointer) {
      this.setTint(0xff0000)
      this.parent.startLoadScene()
    })
  }
  startLoadScene () {
    this.scene.start('LoadScene')
  }
}
