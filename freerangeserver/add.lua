-- add all elements of array `a'
function add (a)
    local sum = 0
    for i,v in ipairs(a) do
        sum = sum + v
    end
    return sum
end

function add2(a, b)
    return a + b
end