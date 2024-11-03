import { apiGet } from './http';

// cookie
export function setCookie(data){
    if (process.env.NODE_ENV === 'development') {
        document.cookie = `token=${data.token}; path=/;`;
    } else {
        document.cookie = `token=${data.token}; path=/; Secure; SameSite=Strict`;
    }
}

// cookie
export function getAuthToken() {
    const token = document.cookie.split('; ').find(row => row.startsWith('token='));
    if (token) {
        return token.split('=')[1];
    }
    return "";
}

// apiCall
export async function getUser() {
    const user = await apiGet("/user/info");
    return user;
}

//sessionStorage
export function setUserToSession(user) {
    sessionStorage.setItem('user', JSON.stringify(user));
}

//sessionStorage
export function getUserFromSession() {
    const userData = sessionStorage.getItem('user');
    if (userData && userData !== "undefined") {
        const user = JSON.parse(userData);
        return user
    }
    return null;
}

//sessionStorage
export function setUserIsAdmin(isadmin) {
    sessionStorage.setItem('admin', isadmin);
}

//sessionStorage
export function getUserIsAdmin() {
    const isAdmin = sessionStorage.getItem('admin');
    if (isAdmin && isAdmin !== "undefined") {
        return isAdmin
    }
    return false;
}


// sessionStorage and cookie destroy
export function destroySession(){
    document.cookie = "token=; path=/; expires=Thu, 01 Jan 1970 00:00:00 UTC;";
    sessionStorage.clear()
}