class Test{
    constructor(documentParent){
        this.canvas
    }

    
}

function createTest(){
    var b = document.getElementsByTagName("body")[0]
    var testElement = document.createElement("div");
    var canvas = document.createElement("canvas");
    testElement.appendChild(canvas);
    b.appendChild(testElement);
    return testElement;
}