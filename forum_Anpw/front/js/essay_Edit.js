var title=sessionStorage.getItem("title")
var essay=sessionStorage.getItem("essay")
var ID=sessionStorage.getItem("id")
essay = essay.replace(/<br\s*\/?\s*>/ig,"\n")//转换格式
document.getElementById("newcontent").value=essay
document.getElementById("newtitle").value=title


$(".layui-input-block #essay-update").click(function (){
    var title=$("#newtitle").val()
    var content=$("#newcontent").val()
    content = content.replace(/\r\n/g, '<br/>').replace(/\n/g, '<br/>').replace(/\s/g, ' '); //转换格式
    var token=sessionStorage.getItem("token")
    $.post({
        url:"http://114.55.107.62:8080/essay/editEssay",
        data: {
            ID:ID,
            newessay: content,
            newtitle: title,
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

$(".layui-input-block #return-btn").click(function (){
    window.history.go(-1);
    sessionStorage.removeItem("title")
    sessionStorage.removeItem("essay")
})