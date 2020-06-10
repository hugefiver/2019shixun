<template>
    <el-main>
        <el-row>
            <el-col :span="9">
                <el-input v-model="mochine" placeholder="设备ID"></el-input>
            </el-col>
            <el-col :span="9">
                <el-input v-model="person" placeholder="使用者"></el-input>
            </el-col>
            <el-col :span="6">
                <el-checkbox-group v-model="returned" :min="1">
                    <el-col>
                        <el-row>
                            <el-checkbox label="false" checked>未归还</el-checkbox>
                        </el-row>
                        <el-row>
                            <el-checkbox label="true" checked>已归还</el-checkbox>
                        </el-row>
                    </el-col>
                </el-checkbox-group>
            </el-col>
        </el-row>

        <el-row>
            <el-col :span="16" :grutter="10">
                <el-date-picker style="width: 100%"
                        v-model="time"
                        type="daterange"
                        range-separator="至"
                        start-placeholder="开始日期"
                        end-placeholder="结束日期">
                </el-date-picker>
            </el-col>
            <el-col :span="8">
                <el-col :span="12">
                    <el-button style="width: 90%" type="primary" @click="search">检索</el-button>
                </el-col>
                <el-col :span="12">
                    <el-button style="width: 90%" @click="clear">清空</el-button>
                </el-col>
            </el-col>
        </el-row>

        <div style="height: 10px"></div>

        <el-row>
            <el-col :span="16">
                <el-input v-model="id" placeholder="设备ID"></el-input>
            </el-col>
            <el-col :span="8">
                <el-col :span="12">
                    <el-button style="width: 90%" @click="borrowMochine">借取</el-button>
                </el-col>
                <el-col :span="12">
                    <el-button style="width: 90%" @click="returnMochine">归还</el-button>
                </el-col>
            </el-col>
        </el-row>

        <el-table :data="mochines" stripe>
            <el-table-column prop="Id" label="记录ID"></el-table-column>
            <el-table-column prop="PersonId" label="用户"></el-table-column>
            <el-table-column prop="MochineId" label="设备ID"></el-table-column>
            <el-table-column prop="BorrowTime" label="借取时间"></el-table-column>
            <el-table-column prop="IsReturn?'Yes':'No'" label="是否归还"></el-table-column>
        </el-table>
    </el-main>

</template>

<script>
    export default {
        name: "MochineBorrow",
        data() {
            return {
                mochine: "",
                person: "",
                returned: [],
                time: [new Date(1), new Date()],
                mochines: [],
                id: ""
            }
        },
        methods: {
            search: function () {
                let params = {};
                if (this.kind !== "") {
                    params['mochine'] = this.mochine
                }
                if (this.person !== "") {
                    params['person'] = this.person
                }
                if (this.returned.length === 1) {
                    params['returned'] = this.returned[0]
                }
                params['before'] = this.time[1].getTime();
                params['after'] = this.time[0].getTime();
                this.$http.get('/api/borrows', {
                    headers: {
                        token: localStorage.getItem('token')
                    },
                    params
                }).then(res => {
                    if (res.data.err === 0) {
                        this.mochines = res.data.list
                    } else {
                        alert(res.data.msg)
                    }
                })
            },
            clear: function () {
                this.mochines = []
            },
            borrowMochine: function () {
                if (this.id === "")
                    return;
                this.$http.post("/api/mochines/"+this.id+"/borrow", {}, {
                    headers: {
                        token: localStorage.getItem('token')
                    }
                }).then(res => {
                    if (res.data.err === 0) {
                        alert("借取成功.")
                    } else {
                        alert(res.data.msg)
                    }
                })
            },
            returnMochine: function () {
                if (this.id === "")
                    return;
                this.$http.post("/api/mochines/"+this.id+"/return", {}, {
                    headers: {
                        token: localStorage.getItem('token')
                    }
                }).then(res => {
                    if (res.data.err === 0) {
                        alert("归还成功.")
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