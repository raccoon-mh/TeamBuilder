
export function isLogin(){
    if (getAuthToken() !== "") {
        return true
    }
    return false
}

export function setSession(data){
    if (process.env.NODE_ENV === 'development') {
        document.cookie = `authToken=${data.token}; path=/;`;
        document.cookie = `user=${data.user}; path=/;`;
    } else {
        document.cookie = `authToken=${data.token}; path=/; Secure; SameSite=Strict`;
        document.cookie = `user=${data.user}; path=/; Secure; SameSite=Strict`;
    }
}

export function detroySession(){
    document.cookie = "authToken=; path=/; expires=Thu, 01 Jan 1970 00:00:00 UTC;";
    document.cookie = "user=; path=/; expires=Thu, 01 Jan 1970 00:00:00 UTC;";
}

export function getAuthToken() {
    const token = document.cookie.split('; ').find(row => row.startsWith('authToken='));
    if (token) {
        return token.split('=')[1];
    }
    return "";
}

export function getUser() {
    const token = document.cookie.split('; ').find(row => row.startsWith('user='));
    if (token) {
        return JSON.parse(token.split('=')[1]);
    }
    return null;
}