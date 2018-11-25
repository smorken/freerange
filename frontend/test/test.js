class Test{
    constructor(title){
        var doc = document.getElementsByTagName("body")[0];
        this.canvas = document.createElement("canvas");
        
        this.titleDiv = document.createElement("div");
        this.titleDiv.appendChild(document.createTextNode(title));
        this.titleDiv.style.margin = "5px 5px";

        this.textDiv = document.createElement("div");
        this.textDiv.style.margin = "5px 5px";
        this.textDiv.style.border = "thin solid #000000";

        this.canvasDiv = document.createElement("div");
        this.canvasDiv.appendChild(this.canvas);
        this.canvasDiv.style.border = "thin solid #000000";
        this.canvasDiv.style.margin = "5px 5px";

        this.mainDiv = document.createElement("div");
        this.mainDiv.appendChild(this.titleDiv);
        this.mainDiv.appendChild(this.textDiv);
        this.mainDiv.appendChild(this.canvasDiv);
        this.mainDiv.style.border = "thin solid #000000";
        

        doc.appendChild(this.mainDiv);
    }
    getCanvas(){
        return this.canvas
    }
    
    appendText(text){
        var t = document.createElement("div");
        t.appendChild(document.createTextNode(text));
        this.textDiv.appendChild(t);
    }

    setText(index, text){
        var t = this.textDiv.childNodes[index];
        t.childNodes[0].nodeValue = text;
    }
}

