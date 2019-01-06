
class LoadScene extends Phaser.Scene {
  constructor () {
    super({ key: 'LoadScene' })
  }

  preload () {
    this.load.image('bg', 'https://twemoji.maxcdn.com/72x72/1f306.png')
    this.load.image('player', 'https://twemoji.maxcdn.com/2/72x72/1f600.png')
    this.load.image('ground', 'assets/platform.png')
    this.load.image('house', 'https://twemoji.maxcdn.com/2/72x72/1f3d8.png')
    this.load.image('hospital', 'https://twemoji.maxcdn.com/2/72x72/1f3e5.png')
    ws.onopen = this.startGame.bind(this)
  }

  create () {

  }

  startGame () {
    this.scene.add('GameScene', GameScene, true)
    this.scene.start('GameScene')
  }
}
