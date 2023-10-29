function loadUserPage() {
    const obj = document.forms["userIdForm"]
    $.ajax({
        type:'post',
        url:'/user',
        async:false,
        dataType:'json',
        data:$(obj).serialize()
    }).done(function (data){
        document.getElementById("level").placeholder = data.level
        document.getElementById("level").value = data.level
        document.getElementById("exp").placeholder = data.exp
        document.getElementById("exp").value = data.exp
        document.getElementById("money").placeholder = data.money
        document.getElementById("money").value = data.money
        document.getElementById("created").value = data.created_at
        document.getElementById("lastLogin").value = data.updated_at
        $("#userDetail").toggle()
    }).fail(function (data){
        alert("検索失敗")
    })
}

function updateUserData(endPoint) {
    const idValue = $('#userId').val();
    const value = $('#'+endPoint).val();
    $.ajax({
        type:'post',
        url:'/user/data/'+endPoint,
        async:false,
        dataType:'html',
        data:{
            id: idValue,
            value: value
        }
    }).done(function (data){
        document.getElementById(endPoint).placeholder = data.endPoint
        alert("更新しました")
    }).fail(function (data){
        alert("更新失敗")
    })
}
