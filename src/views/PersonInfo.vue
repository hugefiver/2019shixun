<template>
    <el-form :model="user">
        <el-form-item>
            <el-image :src="pic" lazy fit="contain" style="height: 200px;"></el-image>
        </el-form-item>
        <el-form-item :hidden="!$parent.isAdmin">
            <el-row>
                <el-col :span="16">
                    <input type="file" id="input-file"></input>
                </el-col>
                <el-col :span="8">
                    <el-button @click="uploadImage">上传图片</el-button>
                </el-col>
            </el-row>
        </el-form-item>
        <el-form-item label="ID">
            <el-input v-model="id" :disabled="!$parent.isAdmin"></el-input>
        </el-form-item>
        <el-form-item label="用户名">
            <el-input v-model="user.name" :disabled="!$parent.isAdmin"></el-input>
        </el-form-item>
        <el-form-item label="用户类型">
            <el-select v-model="user.level" :disabled="!$parent.isAdmin">
                <el-option label="管理员" value="admin"></el-option>
                <el-option label="普通用户" value="default"></el-option>
            </el-select>
        </el-form-item>
        <el-form-item label="密码" :hidden="!$parent.isAdmin">
            <el-input v-model="user.password" ></el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="getPerson" >查询</el-button>
            <el-button @click="updatePerson" :disabled="!$parent.isAdmin">添加更新</el-button>
            <el-button type="danger" @click="deletePerson" :disabled="!$parent.isAdmin">删除用户</el-button>
        </el-form-item>
    </el-form>
</template>

<script>
    export default {
        name: "PersonInfo",
        methods:{
            getPerson: function () {
                let id = this.id;
                this.$http.get('/api/people/' + id, {
                    headers: {
                        token: localStorage.getItem('token')
                    }
                }).then(res => {
                    let d = res.data;
                    this.user.name = d.name;
                    this.user.level = d.level;
                }).catch(err => {
                    alert("查询失败:" + err)
                })
            },
            updatePerson: function () {
                if (this.id === "" | this.user.name === "") {
                    alert("ID/Name 字段不能为空.");
                    return
                }
                this.$http.post('/api/people/' + this.id, {
                    id: this.id,
                    name: this.user.name,
                    level: this.user.level,
                    password: this.user.password,
                }, {
                    headers: {
                        token: localStorage.getItem('token')
                    }}).then(res => {
                        if (res.data.err === 0) {
                            alert(res.data.status)
                        } else {
                            alert("失败:" + res.data.msg)
                        }
                })
            },
            deletePerson: function () {
                this.$confirm('是否删除用户?', '删除用户', {
                    confirmButtonText: '是',
                    cancelButtonText: '否',
                    type: 'warning'
                }).then(() => {
                    this.$http.delete('/api/people/' + this.id, {
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
            uploadImage: function () {
                let files = document.getElementById('input-file').files;

                if (files.length > 0) {
                    if (files[0].name.match(/^.*\.(jpg|jpeg|png|bmp|ico|webp|gif)$/)) {
                        let formData = new FormData();
                        formData.append('file', files[0]);
                        this.$http.post('/api/people/' + this.id + '/image', formData, {
                            headers: {
                                token: localStorage.getItem('token')
                            }
                        }).then(res => {
                            if (res.status === 200) {
                                if (res.data.err > 0) {
                                    alert(res.data.msg);
                                } else {
                                    alert("Upload Succeeded.")
                                }
                            } else {
                                alert('上传失败:' + res.data.status)
                            }
                        })
                    } else {
                        alert("不支持该类型文件。")
                    }
                } else {
                    alert("请选择文件。")
                }
            }
        },
        data: function () {
            return {
                id: "",
                user: {
                    id: "",
                    name: "",
                    level: "",
                    password: ""
                },
            }
        },
        mounted() {
            let u = JSON.parse(localStorage.getItem('user'));
            this.id = u.id;
            this.user.id = u.id;
            this.user.name = u.name;
            this.user.level = u.admin ? "admin" : "default";
            this.user.password = "";
        },
        computed: {
            pic: function () {
                return '/api/people/' + this.id + '/image'
            }
        }
    }
</script>

<style scoped>

</style>