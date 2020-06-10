<template>
    <el-main>
        <el-row type="flex" :gutter="10">
            <el-col :span="16">
                <el-input v-model="kind" placeholder="检索设备型号"></el-input>
            </el-col>
            <el-col :span="8">
                <el-button style="width: 100%" type="primary" @click="search">检索</el-button>
            </el-col>
        </el-row>
        <div style="height: 16px;"></div>
        <el-row type="flex" :gutter="10">
            <el-col :span="16">
                <el-input v-model="id" placeholder="使用设备"></el-input>
            </el-col>
            <el-col :span="8">
                <el-button style="width: 100%" type="primary" @click="use">使用</el-button>
            </el-col>
        </el-row>

        <el-table :data="mochines" stripe>
            <el-table-column prop="Id" label="ID"></el-table-column>
            <el-table-column prop="Kind" label="型号"></el-table-column>
            <el-table-column prop="AdminId" label="管理员ID"></el-table-column>
            <el-table-column prop="LifeYear" label="生命周期"></el-table-column>
        </el-table>
    </el-main>

</template>

<script>
    export default {
        name: "MochineCheck",
        data: function () {
            return {
                mochines: [],
                kind: "",
                id: ""
            }
        },
        methods: {
            search: function () {
                let params = {};
                if (this.kind !== "") {
                    params['kind'] = this.kind;
                }
                this.$http.get('/api/mochines', {
                    headers:{
                        token: localStorage.getItem("token")
                    },
                    params
                }).then(res => {
                    if (res.data.err === 0) {
                        this.mochines = res.data.mochines
                    } else {
                        alert(res.data.err)
                    }
                })
            },
            use: function () {
                this.$http.post('/api/mochines/' + this.id + '/use', {}, {
                    headers: {
                        token: localStorage.getItem('token')
                    }
                }).then(res => {
                    if (res.data.err === 0) {
                        alert('使用成功.')
                    }
                })
            }
        }
    }
</script>

<style scoped>

</style>