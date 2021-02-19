<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增APP</el-button>
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
      <el-table-column label="接入日期" width="180">
        <template slot-scope="scope">{{
          scope.row.CreatedAt | formatDate
        }}</template>
      </el-table-column>
      <el-table-column label="APP名" prop="name" width="120"></el-table-column>
      <el-table-column
        label="输入Topic"
        prop="kafkaInputTopic"
        width="180"
      ></el-table-column>
      <el-table-column
        label="输出Topic"
        prop="kafkaOutputTopic"
        width="180"
      ></el-table-column>
      <el-table-column label="启用报警" prop="enableAlarm" width="120">
        <template slot-scope="scope">{{
          scope.row.enableAlarm | formatBoolean
        }}</template>
      </el-table-column>
      <el-table-column min-width="160">
        <template slot-scope="scope">
          <el-button @click="updateApp(scope.row)" size="small" type="text">
            {{ scope.row.enableAlarm == true ? "关闭报警" : "开启报警" }}
          </el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button
                size="mini"
                type="text"
                @click="scope.row.visible = false"
                >取消</el-button
              >
              <el-button
                type="primary"
                size="mini"
                @click="deleteApp(scope.row)"
                >确定</el-button
              >
            </div>
            <el-button
              type="danger"
              icon="el-icon-delete"
              size="mini"
              slot="reference"
              >删除</el-button
            >
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{ float: 'right', padding: '20px' }"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog
      :before-close="closeDialog"
      :visible.sync="dialogFormVisible"
      title="APP"
    >
      <el-form :inline="true" :model="form" label-width="80px">
        <el-form-item label="APP名">
          <el-input autocomplete="off" v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="输入Topic">
          <el-input
            autocomplete="off"
            v-model="form.kafkaInputTopic"
          ></el-input>
        </el-form-item>
        <el-form-item label="输出Topic">
          <el-input
            autocomplete="off"
            v-model="form.kafkaOutputTopic"
          ></el-input>
        </el-form-item>
        <el-form-item label="启用报警">
          <el-switch
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            v-model="form.enableAlarm"
            clearable
          ></el-switch>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { createApp, deleteApp, getAppList, updateApp } from "@/api/app_manage";
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "App",
  mixins: [infoList],
  data() {
    return {
      listApi: getAppList,
      dialogFormVisible: false,
      type: "",
      form: {
        name: "",
        kafkaInputTopic: "",
        kafkaOutputTopic: "",
        enableAlarm: false,
      },
    };
  },
  filters: {
    formatDate: function (time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function (bool) {
      if (bool != null) {
        return bool ? "是" : "否";
      } else {
        return "";
      }
    },
  },
  methods: {
    async updateApp(row) {
      row.enableAlarm = !row.enableAlarm;
      const res = await updateApp(row);
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "更新成功",
        });
      }
    },

    closeDialog() {
      this.dialogFormVisible = false;
      this.form = {
        name: "",
        kafkaInputTopic: "",
        kafkaOutputTopic: "",
        enableAlarm: false,
      };
    },
    async deleteApp(row) {
      this.$confirm("确定要删除吗?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(async () => {
        row.visible = false;
        const res = await deleteApp({ ID: row.ID, name: row.name });
        if (res.code == 0) {
          this.$message({
            type: "success",
            message: "删除成功",
          });
          if (this.tableData.length == 1) {
            this.page--;
          }
          this.getTableData();
        }
      });
    },
    async enterDialog() {
      let res = await createApp(this.form);
      if (res.code == 0) {
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.dialogFormVisible = true;
    },
  },
  created() {
    this.getTableData();
  },
};
</script>

<style>
</style>