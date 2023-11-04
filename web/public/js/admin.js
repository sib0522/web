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
        alert("ログインに失敗しました")
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
        localStorage.setItem("user", data)
        location.replace("/")
    }).fail(function (data){
        alert("ログアウト失敗")
    })
}

function signUp() {
    const obj = document.forms["signup"]
    $.ajax({
        type:'post',
        url:'/account/signup',
        async:false,
        dataType:'html',
        data:$(obj).serialize()
    }).done(function (data){
        alert("アカウント生成成功\nメイン画面に戻ります")
        localStorage.setItem("user", data)
        location.replace("/")
    }).fail(function (data){
        alert("生成に失敗しました")
    })
}

function isLogin(isLogin) {
    if (!isLogin) {
        window.location.href ="/"
        alert("ログインしてください")    
    }
}