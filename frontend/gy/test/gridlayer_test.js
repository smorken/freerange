
function gridlayer_test_update(test){
        
    var gw = new GridLayer(10,10,1)
    gw.update([
        [-10,100,"mock_content"],
    ])        
    test.assertTrue(gw.getvalue(-10,100) == "mock_content", "test update and get value")
}

function gridlayer_test_update_with_overwrite(test){
        
    var gw = new GridLayer(10,10,1)

    test.assertTrue(gw.getvalue(0,0) === undefined, "getvalue(x,y) with no matching value returns undefined.")
    gw.update([
        [-10,100,"mock_content"],
    ])
    test.assertTrue(gw.getvalue(-10,100) == "mock_content", "get result after first update")

    gw.update([
        [-10,100,"mock_content1"],
        [0,0,"mock_content2"],
    ])
    test.assertTrue(
        gw.getvalue(-10,100) == "mock_content1" && gw.getvalue(0,0) == "mock_content2",
         "get result after second update")
}