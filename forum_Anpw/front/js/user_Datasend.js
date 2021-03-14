//注册按钮触发事件
$(".register-form #register-btn").click(function () {

    var username = $(".register-form #register-username").val()
    var password = $(".register-form #register-password").val()
    var securitycode = $(".register-form #register-securitycode").val()

    $.post({
        url: "http://114.55.107.62:8080/user/register",
        data: {
            username: username,
            password: password,
            securitycode: securitycode,
        },
        success: function (data) {
            alert(data.msg)
        },
        error: function (jqXHR) {
            var responseJSON = JSON.parse(jqXHR.responseText)
            alert(responseJSON.msg)
        }
    })

})


//登录按钮触发事件
$(".login-form #login-btn").click(function () {

    var usernames = $(".login-form #login-username").val()
    var passwords = $(".login-form #login-password").val()

    $.post({
        url: "http://114.55.107.62:8080/user/login",
        data: {
            username: usernames,
            password: passwords,
        },
        dataType:"json",
        xhrFields: {
            //允许带上凭据
            withCredentials: true
        },
        crossDomain: true,
        async: false,
        success: function (data) {
            alert(data.msg)
            var token=data.data.token
            sessionStorage.setItem("token",token)
            $.ajax({
                url: "http://114.55.107.62:8080/user/currentUser",
                async:false,

                beforeSend: function (request) {
                    request.setRequestHeader("Authorization", "Bearer "+token)
                },
                dataType:"json",
                success:function (msg){
                    sessionStorage.setItem("username",msg.data.user.Username)
                    sessionStorage.setItem("sequrityCode",msg.data.user.SequrityCode)
                    sessionStorage.setItem("admin",msg.data.user.Admin)
                    window.location.href="../essay_forum.html"
                },
                error:function (msg){
                    var responseJSON = JSON.parse(jqXHR.responseText)
                    alert(responseJSON.msg)
                }
            })
        },
        error: function (jqXHR) {
            var responseJSON = JSON.parse(jqXHR.responseText)
            alert(responseJSON.msg)
        }
    })

})

//重置密码按钮触发事件
$(".reset-form #reset-btn").click(function () {

    var username = $(".reset-form #reset-username").val()
    var newpassword = $(".reset-form #reset-newpassword").val()
    var newpasswordsure = $(".reset-form #reset-newpassword-sure").val()
    var sequrityCode = $(".reset-form #reset-securitycode").val()


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

        },
        error: function (jqXHR) {

            var responseJSON = JSON.parse(jqXHR.responseText)
            alert(responseJSON.msg)
        }
    })


})