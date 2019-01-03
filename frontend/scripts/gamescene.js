class GameScene extends Phaser.Scene {
  constructor () {
    super({ key: 'GameScene' })
    this.messages = []
  }

  preload () {
    ws.onmessage = this.getWS
    ws.onerror = this.errorWS

    // this.time.addEvent({
    //    delay: 1000,
    //   callback: this.sendWS,
    //   callbackScope: this,
    //   repeat: 100
    // });
  }
  sendWS (msg) {
    ws.send(msg);
  }
  getWS (evt) {
    this.messages.push(evt.data)
  }
  errorWS (evt) {
    // exit the game scene
  }
  create () {
    //  A simple background for our game
    bg = this.add.sprite(400, 300, 'bg')
    bg.displayHeight = 600
    bg.displayWidth = 800
    //  The platforms group contains the ground and the 2 ledges we can jump on
    platforms = this.physics.add.staticGroup()

    //  Here we create the ground.
    //  Scale it to fit the width of the game (the original sprite is 400x32 in size)
    platforms.create(400, 568, 'ground').setScale(2).refreshBody()

    player = this.physics.add.sprite(100, 450, 'santa')

    player.setBounce(0.2)
    player.setCollideWorldBounds(true)

    //  Input Events
    cursors = this.input.keyboard.createCursorKeys()

    this.physics.add.collider(player, platforms)
    this.timeText = this.add.text(100, 200)
  }

  update (time, delta) {
    this.timeText.setText('Time: ' + time + '\nDelta: ' + delta)
    if (cursors.left.isDown && player.body.touching.down) {
      this.sendWS("left")
      player.setVelocityX(-50)
    }
    else if (cursors.right.isDown && player.body.touching.down) {
      this.sendWS("right")
      player.setVelocityX(50)
    }
    else {
      //this.sendWS("stop");
      player.setVelocityX(0);
    }

    if (cursors.up.isDown && player.body.touching.down) {
      this.sendWS('stop')
      player.setVelocityY(-150)
    }
  }
}
