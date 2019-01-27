
class LoadScene extends Phaser.Scene {
  constructor () {
    super({ key: 'LoadScene' })
  }

  init (data) {
    // example data stored in "level"
    // {
    //  "images": {
    //    "player": "https://twemoji.maxcdn.com/2/72x72/1f600.png",
    //    "house": "https://twemoji.maxcdn.com/2/72x72/1f3d8.png"
    //  }
    // }
    this.assets = JSON.parse(data.assets)
  }

  preload () {
    var progress = this.add.graphics()

    this.load.on('progress', function (value) {
      progress.clear()
      progress.fillStyle(0xffffff, 1)
      progress.fillRect(0, 270, 800 * value, 60)
    })

    this.load.on('complete', function () {
      progress.destroy()
    })

    for (var key in this.assets['images']) {
      this.load.image(key, this.assets['images'][key])
    }
  }

  create () {
    // this.scene.add('GameScene', GameScene, true)
    ws.send('request_update')
    ws.onmessage = this.wsmessage.bind(this)
  }
  wsmessage (evt) {
    this.startGameScene(evt.data)
  }
  startGameScene (data) {
    this.scene.start('GameScene', { level: data })
  }
}
