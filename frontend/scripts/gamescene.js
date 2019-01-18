class GameScene extends Phaser.Scene {
  constructor () {
    super({ key: 'GameScene' })

    this.player = null
    this.bg = null
    this.platforms = null
    this.cursors = null
  }

  init (data) {
    this.level = JSON.parse(data.level)
  }

  preload () {
    this.objectCollection = {}
    var objectList = this.level['objects']
    for (var i = 0; i< objectList.length; i++) {
      var obj = objectList[i]
      var id = obj['id']
      var sprite = this.add.sprite(
        obj['xposition'],
        obj['yposition'],
        obj['img'])
      sprite.displayWidth = obj['xsize']
      sprite.displayHeight = obj['ysize']
      if (obj["clickable"]) {
        sprite.setInteractive()
        sprite.on('pointerdown', function (pointer) {
          ws.send("click")
        })
      }
      this.objectCollection[id] = sprite
    }
    // this.time.addEvent({
    //    delay: 1000,
    //   callback: this.sendWS,
    //   callbackScope: this,
    //   repeat: 100
    // });
  }

  // drawBackgroundObjects () {
  //   //  A simple background for our game
  //   this.bg = this.add.sprite(400, 300, 'bg')
  //   this.bg.displayHeight = 600
  //   this.bg.displayWidth = 800
  //   //  The platforms group contains the ground and the 2 ledges we can jump on
  //   this.platforms = this.physics.add.staticGroup()

  //   var hospital = this.add.sprite(300, 300, 'hospital')
  //   hospital.displayHeight = 300
  //   hospital.displayWidth = 300

  //   for (var i = 0; i < 5; i++) {
  //     var house = this.add.sprite(100 + i * 200, 328, 'house')
  //     house.displayHeight = 200
  //     house.displayWidth = 200
  //   }
  //   //  Here we create the ground.
  //   //  Scale it to fit the width of the game (the original sprite is 400x32 in size)
  //   this.platforms.create(400, 568, 'ground').setScale(2).refreshBody()
  // }

  // destroyNpc (player, npc) {
  //   this.npcgroup.remove(npc)
  //   npc.destroy()
  // }
  wsmessage (evt) {
    this.wsdata = evt.data
  }
  sendWS (message) {
    ws.send(message)
  }
  create () {
    ws.onmessage = this.wsmessage.bind(this)

    // this.drawBackgroundObjects()
    // this.npcgroup = this.physics.add.group({ allowGravity: true })
    // for (var i = 0; i < 5; i++) {
    //   var npc = this.physics.add.sprite(100 + 100 * i, 450, 'npc')
    //   this.npcgroup.add(npc, true)
    // }
    // this.physics.add.collider(this.npcgroup, this.platforms)
    // this.player = this.physics.add.sprite(100, 450, 'player')

    // this.player.setBounce(0.2)
    // this.player.setCollideWorldBounds(true)
    // this.cameras.main.startFollow(this.player, true, 0.08, 0.08)
    // //  Input Events
    // this.cursors = this.input.keyboard.createCursorKeys()

    // this.physics.add.collider(this.player, this.platforms)

    // this.physics.add.overlap(this.player, this.npcgroup, this.destroyNpc, null, this)

    this.timeText = this.add.text(100, 200)
  }

  update (time, delta) {
    this.timeText.setText('Time: ' + time + '\nDelta: ' + delta)
    // if (this.cursors.left.isDown && this.player.body.touching.down) {
    //   this.sendWS('left')
    //   this.player.setVelocityX(-50)
    // } else if (this.cursors.right.isDown && this.player.body.touching.down) {
    //   this.sendWS('right')
    //   this.player.setVelocityX(50)
    // } else {
    //   // this.sendWS('stop');
    //   this.player.setVelocityX(0)
    // }

    // if (this.cursors.up.isDown && this.player.body.touching.down) {
    //   this.sendWS('stop')
    //   this.player.setVelocityY(-150)
    // }
  }
}
