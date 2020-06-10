<template>
    <div id="app">
        <div id="nav">
            <!-- router-link to="/">Home</router-link>
            <router-link to="/about">About</router-link -->
            <Menu></Menu>
        </div>
        <router-view/>
    </div>
</template>

<style>
    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        /*color: #2c3e50;
        margin-top: 60px;*/
    }
</style>

<script>
    import Menu from "./components/Menu";

    export default {
        components: {
            Menu,
        },
        data() {
            return {
                isLogin: false,
                isAdmin: false,
            }
        },
        methods: {
            checkLogin: function () {
                let token = localStorage.getItem('token');
                if (token) {
                    this.$http.get('/api/check', {
                        headers: {
                            token: token
                        }
                    }).then(res => {
                        if (res.data.err === 0) {
                            this.isLogin = true;
                            this.isAdmin = res.data.admin
                        } else {
                            alert(res.data.msg)
                        }
                    }).then(() => {
                        if (!this.isLogin) {
                            location.href = '/login'
                        } else {
                            //this.$router.push("/me")
                        }
                    })
                } else {
                    location.href = '/login'
                }
            },
        },
        mounted: function () {
            this.checkLogin();
        }
    }
</script>