
class ConnectScene extends Phaser.Scene {
    constructor () {
      super({ key: 'ConnectScene' })
    }

    preload () {
        ws.onopen = this.startLoadScene.bind(this)
    }
    
    startLoadScene () {
        this.scene.add('LoadScene', LoadScene, true)
       // this.scene.start('LoadScene')
    }
    
}