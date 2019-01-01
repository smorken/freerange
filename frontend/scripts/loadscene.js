class LoadScene extends Phaser.Scene {
    constructor ()
    {
        super({ key: 'LoadScene' });
    }

    preload ()
    {
        this.load.image("bg", "https://twemoji.maxcdn.com/72x72/1f306.png");
        this.load.image('santa', 'https://twemoji.maxcdn.com/36x36/2b1c.png');
        this.load.image('ground', 'assets/platform.png');

        ws = new WebSocket("ws://localhost:8080/client");
        ws.onmessage = this.getWS;
        ws.onerror = this.errorWS;
        ws.onopen = this.openWS;
        this.time.addEvent({
            delay: 1000,
            callback: this.sendWS,
            callbackScope: this,
            repeat: 100
        });
    }



    create()
    {        

    }

    //update(){
//
    //}

    //var isOpenWs = false;
    openWS(evt){
        game.scene.start('GameScene');
        // isOpenWs = true;
    }
    sendWS() {
        //   if(isOpenWs){
        ws.send(999);
        // }
    }
    getWS(evt) {
    
    }
    errorWS(evt) {
    
    }
}