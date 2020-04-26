function param1(str)
    chainhelper:log("param1 test. str:"..str)
end

function param2(str, flag)
    if (flag)
    then
        chainhelper:log("param2 test. flag true. str:"..str)
    else
        chainhelper:log("param2 test. flag false. str:"..str)
    end
end

function param3(in_str, in_bool, in_num)
    chainhelper:log("test param3, in_str: " .. tostring(in_str) .. ", type: " .. type(in_str))
    chainhelper:log("test param3, in_bool: " .. tostring(in_bool) .. ", type: " .. type(in_bool))
    chainhelper:log("test param3, in_num: " .. tostring(in_num) .. ", type: " .. type(in_num))
end


function test_luatype_function(func)
    chainhelper:log("test luatype_function, func: " .. tostring(func) .. ", type: " .. type(func))
end

function test_luatype(obj)
    chainhelper:log("test luatype, object: " .. tostring(obj) .. ", type: " .. type(obj))
end

