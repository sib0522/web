function resourceDownload() {
    const btn = document.getElementById("submit")
    $.ajax({
        type:'post',
        url:'/resource/download',
        async:false,
        dataType:'html',
    }).done(function (data){
        btn.setAttribute("disabled","true")
        alert("リソースを更新しました")
        //location.replace("/")
    }).fail(function (data){
        alert("更新失敗")
    })
}