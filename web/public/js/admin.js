function login() {
    const obj = document.forms["loginForm"]
    $.ajax({
        type:'post',
        url:'/account/login',
        async:false,
        dataType:'html',
        data:$(obj).serialize()
    }).done(function (data){
        alert("ログインに成功しました。\nメイン画面に戻ります")
        localStorage.setItem("user", data)
        location.replace("/")
    }).fail(function (data){
        alert("ログインに失敗しました。\n" + data.responseText)
        if (data.status === 400) {
            location.replace("/")
        }
    })
}

function logout() {
    const obj = document.forms["loginForm"]
    $.ajax({
        type:'get',
        url:'/account/logout',
        async:false,
        dataType:'html',
        data:$(obj).serialize()
    }).done(function (data){
        alert("ログアウトしました。\nメイン画面に戻ります")
        location.replace("/")
    }).fail(function (data){
        alert("ログアウト失敗")
    })
}

function signUp() {
    const obj = document.forms["signUpForm"]
    $.ajax({
        type:'post',
        url:'/account/register',
        async:false,
        dataType:'html',
        data:$(obj).serialize()
    }).done(function (data){
        alert("アカウント生成成功\nメイン画面に戻ります")
        location.replace("/")
    }).fail(function (data){
        alert("生成に失敗しました。\n" + data.responseText)
    })
}

function isLogin(isLogin) {
    if (!isLogin) {
        window.location.href ="/"
        alert("ログインしてください")    
    }
}