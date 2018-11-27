class Test{
    constructor(title){
        var doc = document.getElementsByTagName("body")[0];
        this.canvas = document.createElement("canvas");
        
        this.titleDiv = document.createElement("div");
        this.titleDiv.appendChild(document.createTextNode(title));
        this.titleDiv.style.margin = "5px 5px";

        this.textSectionDiv = document.createElement("div");
        this.textSectionDiv.style.margin = "5px 5px";
        this.textSectionDiv.style.border = "thin solid #000000";

        this.canvasDiv = document.createElement("div");
        this.canvasDiv.appendChild(this.canvas);
        this.canvasDiv.style.border = "thin solid #000000";
        this.canvasDiv.style.margin = "5px 5px";

        this.mainDiv = document.createElement("div");
        this.mainDiv.appendChild(this.titleDiv);
        this.mainDiv.appendChild(this.textSectionDiv);
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
        this.textSectionDiv.appendChild(t);
    }

    setText(index, text){
        var t = this.textSectionDiv.childNodes[index];
        t.childNodes[0].nodeValue = text;
    }

    AssertTrue(value, description){
        var textDiv = document.createElement("div");

        if(value){
            textDiv.style.color = "green";
            description = "PASS: "+ description;
        }
        else{
            textDiv.style.color = "red";
            description = "FAILED: "+ description;
        }
        var descNode = document.createTextNode(description)
        textDiv.appendChild(descNode);
        this.textSectionDiv.appendChild(textDiv);

    }

}

