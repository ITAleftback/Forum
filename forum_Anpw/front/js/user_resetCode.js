//重置密码按钮触发事件

$(".layui-input-block #reset-btn").click(function () {
    var username=sessionStorage.getItem("username")
    var newpassword = $(".layui-input-inline #reset-newpassword").val()
    var newpasswordsure = $(".layui-input-inline #reset-newpassword-sure").val()
    var sequrityCode = $(".layui-input-inline #reset-sequrityCode").val()

    $.post({
        url: "http://114.55.107.62:8080/user/resetcode",
        data: {
            username: username,
            newpassword: newpassword,
            newpasswordsure:newpasswordsure,
            sequrityCode: sequrityCode,
        },
        success: function (data) {
            alert(data.msg)
            window.location.href="../essay_forum.html"
        },
        error: function (jqXHR) {
            var responseJSON = JSON.parse(jqXHR.responseText)
            alert(responseJSON.msg)
        }
    })


})

//返回

$(".layui-input-block #return").click(function (){

    window.history.go(-1);
})