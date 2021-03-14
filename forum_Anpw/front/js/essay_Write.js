
$(".layui-input-block #essay-submit").click(function (){
    var username=sessionStorage.getItem("username")
    var title=$(".layui-input-block #essay-title").val()

    var content=$(".layui-input-block #essay-content").val()

     content = content.replace(/\r\n/g, '<br/>').replace(/\n/g, '<br/>').replace(/\s/g, ' '); //转换格式
    var token=sessionStorage.getItem("token")
    $.post({
        url:"http://114.55.107.62:8080/essay/writeEssay",
        data: {
            username: username,
            essay: content,
            title: title,
        },
        dataType:"json",
        beforeSend: function (request) {
            request.setRequestHeader("Authorization", "Bearer "+token)
        },
        success:function (data){
            alert(data.msg)
            window.location.href="../essay_forum.html"
        },
        error:function (jqXHR){
            var responseJSON = JSON.parse(jqXHR.responseText)
            alert(responseJSON.msg)
        }
    })
})