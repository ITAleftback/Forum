
var title=sessionStorage.getItem("title")
var essay=sessionStorage.getItem("essay")
var essayid=sessionStorage.getItem("essayID")
console.log(essayid)
$("#essay-title").html(title)
$(".layui-field-box #essay-content").html(essay)
$(".layui-input-block #return-btn").click(function (){
    window.history.go(-1);
    sessionStorage.removeItem("title")
    sessionStorage.removeItem("essay")
})

$(".layui-input-block #essay-comment-btn").click(function (){
    var token=sessionStorage.getItem("token")
    //第一个实例
    $.post({
        url: "http://114.55.107.62:8080/comment/essaycomment",
        data:{
            username:sessionStorage.getItem("username"),
            comment:$("#essay-comment").val(),
            essayID:essayid
        },
        dataType: "json",
        async:false,
        beforeSend: function (request) {
            request.setRequestHeader("Authorization", "Bearer "+token)
        },
        success: function (data){
            alert(data.msg)
            window.location.href="../essay_View.html"
        },
        error:function (jqXHR){
            var responseJSON = JSON.parse(jqXHR.responseText)
            alert(responseJSON.msg)
        }

    })

})
