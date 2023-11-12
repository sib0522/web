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

function resourceUpload() {
    const input = $("#fileInput")[0]
    const files = input.files

    if (files.length == 0) {
        alert("ファイルをしてください")
        return
    }

    let formData = new FormData()
    for (let i = 0; i < files.length; i++) {
        formData.append('files[]', files[i])
    }

    $.ajax({
        type:'post',
        url:'/resource/upload',
        data:formData,
        processData: false,
        contentType: false,
        async:false,
    }).done(function (data){
        alert("リソースをアップロードしました")
        //location.replace("/")
    }).fail(function (data){
        alert("アップロード失敗")
    })
}