<template>
  <div class="sticky-top">

    <header class="navbar navbar-expand-md sticky-top d-print-none">
      <div class="container-xl">
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbar-menu"
          aria-controls="navbar-menu" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>

        <h1 class="navbar-brand navbar-brand-autodark d-none-navbar-horizontal pe-0 pe-md-3">
          <router-link class="nav-link" to="/">Team Builder</router-link>
        </h1>

        <div class="navbar-nav flex-row order-md-last">
          <div class="d-none d-md-flex">

            <button @click="setTheme('dark')" class="nav-link px-0 hide-theme-dark" title="Enable dark mode"
              data-bs-toggle="tooltip" data-bs-placement="bottom">
              <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24"
                stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                <path d="M12 3c.132 0 .263 0 .393 0a7.5 7.5 0 0 0 7.92 12.446a9 9 0 1 1 -8.313 -12.454z" />
              </svg>
            </button>

            <button @click="setTheme('light')" class="nav-link px-0 hide-theme-light" title="Enable light mode"
              data-bs-toggle="tooltip" data-bs-placement="bottom">
              <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24"
                stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                <path d="M12 12m-4 0a4 4 0 1 0 8 0a4 4 0 1 0 -8 0" />
                <path d="M3 12h1m8 -9v1m8 8h1m-9 8v1m-6.4 -15.4l.7 .7m12.1 -.7l-.7 .7m0 11.4l.7 .7m-12.1 -.7l-.7 .7" />
              </svg>
            </button>

          </div>

          <div class="nav-item dropdown">

            <!-- Session Exist -->
            <div v-if="user && loginStatus">
              <a href="#" class="nav-link d-flex lh-1 text-reset p-0" data-bs-toggle="dropdown"
                aria-label="Open user menu">
                <span class="avatar avatar-sm"
                  :style="{ backgroundImage: `url(https://cdn.discordapp.com/avatars/${user.id}/${user.avatar}.${user.avatar.startsWith('a_') ? 'gif' : 'png'})`}"></span>
                <div class="d-none d-xl-block ps-2" id="loginBtn">
                  <div>{{ user.global_name }}</div>
                  <div class="mt-1 small text-secondary">{{ user.email }}</div>
                </div>
              </a>
              <div class="dropdown-menu dropdown-menu-end dropdown-menu-arrow">
                <a @click="logout()" class="dropdown-item">Logout</a>
                <!-- <a v-if="user && loginStatus && isadmin" class="dropdown-item">Admin Dashboard</a> -->
              </div>
            </div>

            <!-- Session Not Exist -->
            <div v-if="!(user && loginStatus)">
              <a href="#" class="nav-link d-flex lh-1 text-reset p-0" data-bs-toggle="dropdown"
                aria-label="Open user menu">
                <div class="d-none d-xl-block ps-2" id="loginBtn">
                  <div>Login</div>
                  <div class="mt-1 small text-secondary"></div>
                </div>
              </a>
              <div class="dropdown-menu dropdown-menu-end dropdown-menu-arrow">
                <a @click="openOauth()" class="dropdown-item">Discord Login</a>
              </div>
            </div>
          </div>


        </div>

      </div>
    </header>

    <header class="navbar-expand-md">
      <div class="collapse navbar-collapse" id="navbar-menu">
        <div class="navbar">
          <div class="container-xl">
            <ul class="navbar-nav">
              <li v-if="user && loginStatus && isadmin" class="nav-item">
                <router-link class="nav-link" to="/admin">
                  <span class="nav-link-icon d-md-none d-lg-inline-block">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                      fill="currentColor" class="icon icon-tabler icons-tabler-filled icon-tabler-layout-dashboard">
                      <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                      <path
                        d="M9 3a2 2 0 0 1 2 2v6a2 2 0 0 1 -2 2h-4a2 2 0 0 1 -2 -2v-6a2 2 0 0 1 2 -2zm0 12a2 2 0 0 1 2 2v2a2 2 0 0 1 -2 2h-4a2 2 0 0 1 -2 -2v-2a2 2 0 0 1 2 -2zm10 -4a2 2 0 0 1 2 2v6a2 2 0 0 1 -2 2h-4a2 2 0 0 1 -2 -2v-6a2 2 0 0 1 2 -2zm0 -8a2 2 0 0 1 2 2v2a2 2 0 0 1 -2 2h-4a2 2 0 0 1 -2 -2v-2a2 2 0 0 1 2 -2z" />
                    </svg>
                  </span>
                  <span class="nav-link-title">Admin Dash</span>
                </router-link>
              </li>
              <li class="nav-item">
                <router-link class="nav-link" to="/">
                  <span class="nav-link-icon d-md-none d-lg-inline-block">
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24"
                      stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                      <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                      <path d="M5 12l-2 0l9 -9l9 9l-2 0" />
                      <path d="M5 12v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-7" />
                      <path d="M9 21v-6a2 2 0 0 1 2 -2h2a2 2 0 0 1 2 2v6" />
                    </svg>
                  </span>
                  <span class="nav-link-title">Home</span>
                </router-link>
              </li>
              <li class="nav-item">
                <router-link class="nav-link" to="/about">
                  <span class="nav-link-icon d-md-none d-lg-inline-block">
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24"
                      stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                      <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                      <path d="M5 12l-2 0l9 -9l9 9l-2 0" />
                      <path d="M5 12v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-7" />
                      <path d="M9 21v-6a2 2 0 0 1 2 -2h2a2 2 0 0 1 2 2v6" />
                    </svg>
                  </span>
                  <span class="nav-link-title">About</span>
                </router-link>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </header>

  </div>
</template>

<script>
  import { destroySession, getAuthToken, getUser, getUserFromSession, getUserIsAdmin, setCookie, setUserIsAdmin, setUserToSession } from '../util/auth';

  export default {
    mounted() {
      this.initTheme()
      this.user = getUserFromSession();
    },
    data() {
      return {
        loginStatus: false,
        isadmin: false,
        user: null,
      };
    },
    methods: {
      setTheme(theme) {
        this.currentTheme = theme;
        localStorage.setItem('tablerTheme', theme);
        document.body.className = theme;
        window.location.reload()
      },
      initTheme() {
        const savedTheme = localStorage.getItem('tablerTheme');
        if (savedTheme) {
          this.currentTheme = savedTheme;
        }
      },
      async openOauth() {
        window.open(process.env.VUE_APP_DISCORD_OAUTH_URL, '_blank', 'width=800,height=800');
        return new Promise((resolve) => {
          const handleOAuthMessage = async (event) => {
            // if (event.origin !== window.location.origin) return;
            const data = event.data;
            if (data) {
              setCookie(data);
              this.loginStatus = true;
              window.removeEventListener('message', handleOAuthMessage);
              resolve();
            }
          };
          window.addEventListener('message', handleOAuthMessage);
        });
      },
      async fetchUserInfo() {
        try {
          const token = getAuthToken()
          const userDataResp = await getUser(token);
          if (userDataResp.status == 200) {
            const userInfo = userDataResp.data.userInfo
            const userAdmin = userDataResp.data.admin;
            setUserToSession(userInfo);
            setUserIsAdmin(userAdmin)
            this.isadmin = userAdmin
            this.user = userInfo;
          } else {
            console.error('error fetchUserInfo:', userDataResp.status);
          }
        } catch (error) {
          console.error('error fetchUserInfo:', error);
        }
      },
      logout() {
        destroySession();
        this.loginStatus = false;
        this.user = null;
      },
    },
    watch: {
      loginStatus(newStatus) {
        const sessionUser = getUserFromSession();
        if (newStatus && (sessionUser == undefined)) {
          this.fetchUserInfo();
        }
      }
    },
    created() {
      const sessionUser = getUserFromSession();
      if (sessionUser !== undefined) {
        this.user = sessionUser;
        this.loginStatus = true;
        this.isadmin = getUserIsAdmin();
      } else if (!sessionUser || getAuthToken() !== "") {
        this.fetchUserInfo();
      }
    },
  }


</script>

<style>
  :root {
    --tblr-font-sans-serif: 'Inter Var', -apple-system, BlinkMacSystemFont, San Francisco, Segoe UI, Roboto, Helvetica Neue, sans-serif;
  }

  body {
    font-feature-settings: "cv03", "cv04", "cv11";
  }
</style>