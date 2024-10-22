// store/index.js
import { createStore } from 'vuex';

export default createStore({
    state: {
        user: null, // 유저 정보를 저장할 상태
    },
    mutations: {
    setUser(state, user) {
        state.user = user;  // 유저 정보를 객체로 저장
    },
    logout(state) {
        state.user = null;  // 로그아웃 시 유저 정보를 초기화
    }
    },
    getters: {
        isLoggedIn: state => !!state.user, // 로그인 여부 반환
        getUser: state => state.user,      // 유저 정보 반환
    }
});


