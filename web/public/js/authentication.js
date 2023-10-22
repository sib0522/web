function checkLogin(_url) {
    const url = _url
    const token = getToken()
    const headers = {'Authorization': `Bearer ${token}`}
    if (token === null) {
        alert("ログインしてください")
        location.href = '/account/login';
        return
    }
    fetch(url, {headers})
        .then(response => {
            if (!response.ok && response.statusText === "Unauthorized") {
                localStorage.removeItem('user')
                alert("ログインセッションがタイムアウトしました。\n再度ログインしてください。")
                location.href = '/account/login';
                throw new Error("Unauthorized")
            } else if (!response.ok && response.status === 404) {
                alert("存在しないページです。")
                throw new Error("Not Found")
            }
        })
}

function checkLogin2(_url) {
    const url = _url
    const token = getToken()
    const headers = {'Authorization': `Bearer ${token}`}
    if (token === null) {
        alert("ログインしてください")
        location.href = '/account/login';
        return
    }
    fetch(url, {headers})
        .then(response => {
            if (!response.ok && response.statusText === "Unauthorized") {
                localStorage.removeItem('user')
                alert("ログインセッションがタイムアウトしました。\n再度ログインしてください。")
                location.href = '/account/login';
                throw new Error("Unauthorized")
            }
            else if (!response.ok && response.status === 404) {
                alert("存在しないページです。")
                throw new Error("Not Found")
            }
            return response.json()})
        .then(data => {
            // まずはJSON文字列に変換
            const jsonStr = JSON.stringify(data)

            // そしてJSON文字列をJavaScriptのオブジェクトに変換
            alert("おかえりなさい。" + JSON.parse(jsonStr)['nickname'] + "様")
            //location.href = _url
        }).catch(error => alert(error.message))
}

function getToken() {
    const val = JSON.parse(localStorage.getItem('user'))
    const token = val !== null ? val['token'] : null
    return token;
}