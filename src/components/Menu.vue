<template>
    <el-menu router="router" class="el-menu" mode="horizontal">
        <el-menu-item index="/me">
            <router-link to="/me">人员</router-link>
        </el-menu-item>
        <el-menu-item index="/mochines">
            <router-link to="/mochines">设备</router-link>
        </el-menu-item>
        <el-menu-item index="/borrow">
            <router-link to="/borrow">借还</router-link>
        </el-menu-item>
        <!-- el-menu-item index="/admin/people" :hidden="!$parent.isAdmin">
            <router-link to="/admin/people">人员管理</router-link>
        </el-menu-item -->
        <el-menu-item index="/admin/mochines" :hidden="!$parent.isAdmin">
            <router-link to="/admin/mochines">设备管理</router-link>
        </el-menu-item>

        <el-menu-item>
            <el-button v-if="$parent.isLogin" @click="logout">登出</el-button>
            <el-button v-else @click="login">登录</el-button>
        </el-menu-item>
    </el-menu>
</template>

<style>
    a {
        text-decoration-line: none;
    }
</style>

<script>
    export default {
        name: "Menu",
        data() {
            return {
                active: '/me',
            }
        },
        methods: {
            logout: function () {
                localStorage.removeItem('token');
                localStorage.removeItem('user');
                location.href = '/';
                location.reload();
            },
            login: function () {
                location.href = '/login'
            }
        },
    }
</script>

<style scoped>

</style>