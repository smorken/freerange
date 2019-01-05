class GameScene extends Phaser.Scene {
  constructor () {
    super({ key: 'GameScene' })

    this.player = null
    this.bg = null
    this.platforms = null
    this.cursors = null
  }

  preload () {

    // this.time.addEvent({
    //    delay: 1000,
    //   callback: this.sendWS,
    //   callbackScope: this,
    //   repeat: 100
    // });
  }

  create () {
    //  A simple background for our game
    this.bg = this.add.sprite(400, 300, 'bg')
    this.bg.displayHeight = 600
    this.bg.displayWidth = 800
    //  The platforms group contains the ground and the 2 ledges we can jump on
    this.platforms = this.physics.add.staticGroup()

    //  Here we create the ground.
    //  Scale it to fit the width of the game (the original sprite is 400x32 in size)
    this.platforms.create(400, 568, 'ground').setScale(2).refreshBody()

    this.player = this.physics.add.sprite(100, 450, 'santa')

    this.player.setBounce(0.2)
    this.player.setCollideWorldBounds(true)

    //  Input Events
    this.cursors = this.input.keyboard.createCursorKeys()

    this.physics.add.collider(this.player, this.platforms)
    this.timeText = this.add.text(100, 200)
  }

  update (time, delta) {
    this.timeText.setText('Time: ' + time + '\nDelta: ' + delta)
    if (this.cursors.left.isDown && this.player.body.touching.down) {
      this.client.sendWS('left')
      this.player.setVelocityX(-50)
    } else if (this.cursors.right.isDown && this.player.body.touching.down) {
      this.client.sendWS('right')
      this.player.setVelocityX(50)
    } else {
      // this.sendWS('stop');
      this.player.setVelocityX(0)
    }

    if (this.cursors.up.isDown && this.player.body.touching.down) {
      this.sendWS('stop')
      this.player.setVelocityY(-150)
    }
  }
}