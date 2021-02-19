<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="app">
          <el-input placeholder="app" v-model="searchInfo.app"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">刷新</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
      <el-table-column
        label="模版id"
        prop="clusterId"
        width="100"
      ></el-table-column>

      <el-table-column label="级别" prop="level" width="80"></el-table-column>

      <el-table-column label="数量" prop="size" width="80"></el-table-column>

      <el-table-column label="tokens" prop="tokens"></el-table-column>

      <el-table-column label="样例" prop="content"></el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getRealtimeResult } from "@/api/log_template";
import infoList from "@/mixins/infoList";
export default {
  name: "RealtimeTemplate",
  mixins: [infoList],
  data() {
    return {
      listApi: getRealtimeResult,
    };
  },
  methods: {
    onSubmit() {
      this.refersh();
    },
    refersh() {
      this.getTableData();
      this.$message({
        type: "success",
        message: "刷新成功",
      });
    },
  },
};
</script>

<style>
</style>