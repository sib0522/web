const c = getCookie("sideNaviLayout")
if (c != "false") {
    const layoutColor = document.getElementById("sidenavAccordion")
    layoutColor.className = c
}
function changeLayoutColor() {
    const layoutColor = document.getElementById("sidenavAccordion")
    layoutColor.className = layoutColor.className === "sb-sidenav accordion sb-sidenav-dark" ? "sb-sidenav accordion sb-sidenav-light" : "sb-sidenav accordion sb-sidenav-dark"
    setCookie("sideNaviLayout", layoutColor.className, 1)
}

function getCookie(key) {
    let result = null;
    const cookie = document.cookie.split(';');
    cookie.some(function (item) {
        // 공백을 제거
        item = item.replace(' ', '');

        const dic = item.split('=');

        if (key === dic[0]) {
            result = dic[1];
            return true;    // break;
        }
    });
    return result.replaceAll('%20', ' ');
}

function setCookie(key, value, expire) {
    let todayDate = new Date();
    todayDate.setDate(todayDate.getDate() + expire);
    document.cookie = key + "=" + escape(value) + "; path=/; expires=" + todayDate.toUTCString() + ";"
}
