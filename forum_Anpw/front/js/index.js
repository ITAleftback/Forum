var username = sessionStorage.getItem("username")
var token =sessionStorage.getItem("token")
$(".layui-nav-item #username").html(username)
//注意：导航 依赖 element 模块，否则无法进行功能性操作
layui.use('element', function () {
    var element = layui.element;
});
layui.use('table', function () {
    var table = layui.table;
    var token = sessionStorage.getItem("token")
    //第一个实例
    $.post({
        url: "http://114.55.107.62:8080/essay/currentEssay",
        dataType: "json",
        async: false,
        beforeSend: function (request) {
            request.setRequestHeader("Authorization", "Bearer " + token)
        },
        success: function (data) {

            table.render({
                elem: '#demo'
                , height: 500
                , page: true //开启分页
                , data: data.data.essaylist.Value
                , cols: [[ //表头
                    {field: 'Author', title: '用户', width: '15%'}
                    , {field: 'Title', title: '标题', width: '25%', sort: true}
                    , {field: 'CreatedAt', title: '发帖时间', width: '20%'}
                    , {fixed: 'right', width: '20%', align: 'center', toolbar: '#toolbarDemo'}
                ]]


            });
            table.on('tool(test)', function(obj){ //注：tool 是工具条事件名，test 是 table 原始容器的属性 lay-filter="对应的值"
                var data = obj.data; //获得当前行数据
                var layEvent = obj.event; //获得 lay-event 对应的值（也可以是表头的 event 参数对应的值）
                var tr = obj.tr; //获得当前行 tr 的 DOM 对象（如果有的话）
                if(layEvent === 'detail'){ //查看
                    sessionStorage.setItem("title",obj.data.Title)
                    sessionStorage.setItem("essay",obj.data.Essay)
                    sessionStorage.setItem("essayID",obj.data.ID)
                    window.location.href="../essay_View.html"
                } else if(layEvent === 'del'){ //删除
                    layer.confirm('真的删除么', function(index){
                        admindelEssay(obj)
                        layer.close(index);

                        //向服务端发送删除指令
                    });
                } else if(layEvent === 'edit'){ //编辑
                    sessionStorage.setItem("title",obj.data.Title)
                    sessionStorage.setItem("essay",obj.data.Essay)
                    sessionStorage.setItem("id",obj.data.ID)
                    window.location.href="../essay_Edit.html"
                } else if(layEvent === 'LAYTABLE_TIPS'){
                    layer.alert('Hi，头部工具栏扩展的右侧图标。');
                }
            });
        }
    })

});
function admindelEssay(obj){
    $.post({
        url:"http://114.55.107.62:8080/essay/delEssay",
        data: {
            ID:obj.data.ID
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
}