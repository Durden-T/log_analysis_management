<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="app">
          <el-input placeholder="app" v-model="searchInfo.app"></el-input>
        </el-form-item>
        <el-form-item label="模版id">
          <el-input
            placeholder="模版id"
            v-model="searchInfo.clusterId"
          ></el-input>
        </el-form-item>
        <el-form-item label="tokens">
          <el-input placeholder="tokens" v-model="searchInfo.tokens"></el-input>
        </el-form-item>
        <el-form-item label="级别">
          <el-input placeholder="级别" v-model="searchInfo.level"></el-input>
        </el-form-item>
        <el-form-item label="样例">
          <el-input placeholder="样例" v-model="searchInfo.content"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增日志模版</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button @click="deleteVisible = false" size="mini" type="text"
                >取消</el-button
              >
              <el-button @click="onDelete" size="mini" type="primary"
                >确定</el-button
              >
            </div>
            <el-button
              icon="el-icon-delete"
              size="mini"
              slot="reference"
              type="danger"
              >批量删除</el-button
            >
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
      <el-table-column type="selection" width="55"></el-table-column>
      <el-table-column label="app" prop="app" width="120"></el-table-column>

      <el-table-column
        label="模版id"
        prop="clusterId"
        width="120"
      ></el-table-column>

      <el-table-column
        label="tokens"
        prop="tokens"
        width="120"
      ></el-table-column>

      <el-table-column label="数量" prop="size" width="120"></el-table-column>

      <el-table-column label="级别" prop="level" width="120"></el-table-column>

      <el-table-column
        label="样例"
        prop="content"
        width="120"
      ></el-table-column>

      <el-table-column label="创建日期" width="180">
        <template slot-scope="scope">{{
          scope.row.CreatedAt | formatDate
        }}</template>
      </el-table-column>

      <el-table-column label="更新日期" width="180">
        <template slot-scope="scope">{{
          scope.row.UpdatedAt | formatDate
        }}</template>
      </el-table-column>

      <el-table-column>
        <template slot-scope="scope">
          <el-button
            type="danger"
            icon="el-icon-delete"
            size="mini"
            @click="deleteRow(scope.row)"
            >删除</el-button
          >
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
      title="弹窗操作"
    >
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="app:">
          <el-input
            v-model="formData.app"
            clearable
            placeholder="请输入"
          ></el-input>
        </el-form-item>

        <el-form-item label="级别:">
          <el-input
            v-model="formData.level"
            clearable
            placeholder="请输入"
          ></el-input>
        </el-form-item>

        <el-form-item label="日志内容:">
          <el-input
            v-model="formData.content"
            clearable
            placeholder="请输入"
          ></el-input>
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
import {
  createLogTemplate,
  deleteLogTemplate,
  deleteLogTemplateByIds,
  getLogTemplateList,
} from "@/api/log_template"; //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "LogTemplate",
  mixins: [infoList],
  data() {
    return {
      listApi: getLogTemplateList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        clusterId: 0,
        app: "",
        tokens: "",
        size: 0,
        level: "",
        content: "",
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
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    deleteRow(row) {
      this.$confirm("确定要删除吗?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(() => {
        this.deleteLogTemplate(row);
      });
    },
    async onDelete() {
      const ids = [];
      if (this.multipleSelection.length == 0) {
        this.$message({
          type: "warning",
          message: "请选择要删除的数据",
        });
        return;
      }
      this.multipleSelection &&
        this.multipleSelection.map((item) => {
          ids.push(item.ID);
        });
      const res = await deleteLogTemplateByIds({ ids });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功",
        });
        if (this.tableData.length == ids.length) {
          this.page--;
        }
        this.deleteVisible = false;
        this.getTableData();
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        clusterId: 0,
        app: "",
        tokens: "",
        size: 0,
        level: "",
        content: "",
      };
    },
    async deleteLogTemplate(row) {
      const res = await deleteLogTemplate({ ID: row.ID });
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
    },
    async enterDialog() {
      let res = await createLogTemplate(this.formData);
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "创建成功",
        });
        this.closeDialog();
      }
    },
    openDialog() {
      this.dialogFormVisible = true;
    },
  },
};
</script>

<style>
</style>
