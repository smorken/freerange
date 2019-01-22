class GameScene extends Phaser.Scene {
  constructor () {
    super({ key: 'GameScene' })
  }

  init (data) {
    this.level = JSON.parse(data.level)
  }

  createObjects (objectList) {
    for (var i = 0; i < objectList.length; i++) {
      var obj = objectList[i]
      var id = obj['id']
      var sprite = this.add.sprite(
        obj['xposition'],
        obj['yposition'],
        obj['img'])
      sprite.displayWidth = obj['xsize']
      sprite.displayHeight = obj['ysize']
      if (obj['clickable']) {
        sprite.setInteractive()
        sprite.on('pointerdown', function (pointer) {
          ws.send('[click, ' + id + ']')
        })
      }
      this.objectCollection[id] = sprite
    }
  }

  destroyObjects (idList) {
    for (var i = 0; i < idList.length; i++) {
      var id = idList[i]
      this.objectCollection[id].destroy()
      delete this.objectCollection[id]
    }
  }

  updatePositions (positions) {
    for (var i = 0; i < positions.length; i++) {
      var data = positions[i]
      var id = data[0]
      var newX = data[1]
      var newY = data[2]
      var obj = this.objectCollection[id]
      obj['xerror'] = newX - obj.x
      obj['yerror'] = newX - obj.x
    }
  }
  preload () {
    this.objectCollection = {}
    var objectList = this.level['objects']
    this.createObjects(objectList)

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
    var wsdata = evt.data
    var data = JSON.parse(wsdata)
    if (data.hasOwnProperty('create')) {
      this.createObjects(data['create'])
    }
    if (data.hasOwnProperty('destroy')) {
      this.destroyObjects(data['destroy'])
    }
    if (data.hasOwnProperty('position')) {
      this.updatePositions(data['position'])
    }
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

    // this sshould be the fraction of server message rate / frame rate
    var smoothFactor = 0.5 // using 30msg/s /60frame/s for now

    for (var i = 0; i < this.objectCollection.length; i++) {
      var obj = this.objectCollection[i]
      if (obj.hasOwnProperty('xerror')) {
        var xcorrection = obj['xerror'] * smoothFactor
        obj.x += xcorrection
        obj['xerror'] -= xcorrection

        var ycorrection = obj['yerror'] * smoothFactor
        obj.y += ycorrection
        obj['yerror'] -= ycorrection
      }
    }
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
