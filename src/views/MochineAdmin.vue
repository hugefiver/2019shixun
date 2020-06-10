<template>
    <el-main>
        <el-row :gutter="10">
            <el-col :span="12">
                <el-input v-model="id" placeholder="ID"></el-input>
            </el-col>
            <el-col :span="12">
                <el-input v-model="adminId" placeholder="管理员ID"></el-input>
            </el-col>
        </el-row>
        <el-row :gutter="10">
            <el-col :span="12">
                <el-input v-model="kind" placeholder="型号"></el-input>
            </el-col>
            <el-col :span="12">
                <el-input-number :min="1" :max="10" style="width: 100%"
                                 v-model="lifeYear" placeholder="使用年限">
                </el-input-number>
            </el-col>
        </el-row>

        <el-row>
            <el-col :span="8">
                <el-button @click="updateMochine">添加更新</el-button>
            </el-col>
            <el-col :span="8">
                <el-button @click="deleteMochine">删除设备</el-button>
            </el-col>
            <el-col :span="8">
                <el-button @click="getUseLog">设备使用记录</el-button>
            </el-col>
        </el-row>

        <el-table :data="uses" stripe>
            <el-table-column prop="Id" label="记录ID"></el-table-column>
            <el-table-column prop="UserId" label="使用者"></el-table-column>
            <el-table-column prop="MochineId" label="设备ID"></el-table-column>
            <el-table-column prop="UseTime" label="使用时间"></el-table-column>
          </el-table>
    </el-main>

</template>

<script>
    export default {
        name: "MochineAdmin",
        data: function () {
            return {
                id: '',
                adminId: '',
                kind: '',
                lifeYear: '1',
                uses: []
            }
        },
        methods: {
            updateMochine: function () {
                if (this.id === "" || this.adminId === "" || this.kind === "") {
                    alert("ID/AdminID/Kind 字段不能为空");
                    return
                }
                this.$http.post('/api/mochines/'+this.id, {
                    id: this.id,
                    adminId: this.adminId,
                    kind: this.kind,
                    lifeYear: this.lifeYear
                },{
                    headers:{
                        token: localStorage.getItem('token')
                    }
                }).then(res => {
                    if (res.data.err === 0) {
                        alert(res.data.status)
                    } else {
                        alert(res.data.msg)
                    }
                })
            },
            deleteMochine: function () {
                this.$confirm('是否删除?', '删除', {
                    confirmButtonText: '是',
                    cancelButtonText: '否',
                    type: 'warning'
                }).then(() => {
                    this.$http.delete('/api/mochines/' + this.id, {
                        headers: {
                            token: localStorage.getItem('token')
                        }}).then(res => {
                        if (res.data.err === 0) {
                            this.$alert('删除成功', '', {
                                confirmButtonText: '确定'
                            })
                        } else {
                            alert("失败:" + res.data.msg)
                        }
                    })
                })
            },
            getUseLog: function () {
                if (this.id === "") {
                    return
                }
                this.$http.get('/api/mochines/'+this.id+'/use', {
                    headers: {
                        token: localStorage.getItem('token')
                    }
                }).then(res => {
                    if (res.data.err === 0) {
                        this.uses = res.data.uses
                    } else {
                        alert(res.data.msg)
                    }
                })
            }
        }
    }
</script>

<style scoped>

</style>