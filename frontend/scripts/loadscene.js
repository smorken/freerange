
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
    this.data = JSON.parse(data.level)
  }

  preload () {

    var progress = this.add.graphics();
    
    this.load.on('progress', function (value) {

      progress.clear();
      progress.fillStyle(0xffffff, 1);
      progress.fillRect(0, 270, 800 * value, 60);

  });

  this.load.on('complete', function () {

      progress.destroy();

  });

    for (var key in this.data['images']) {
      this.load.image(key, this.data['images'][key])
    }
  }

  create () {
    // this.scene.add('GameScene', GameScene, true)
    this.scene.start('GameScene')
  }
}
