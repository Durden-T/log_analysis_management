<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="app">
          <el-input placeholder="app" v-model="searchInfo.app"></el-input>
        </el-form-item>
        <el-form-item label="报警名">
          <el-input placeholder="报警名" v-model="searchInfo.name"></el-input>
        </el-form-item>
        <el-form-item label="模版id">
          <el-input
            placeholder="模版id"
            v-model="searchInfo.templateId"
          ></el-input>
        </el-form-item>
        <el-form-item label="接收邮箱">
          <el-input
            placeholder="接收邮箱"
            v-model="searchInfo.email"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary"
            >新增模版报警策略</el-button
          >
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

      <el-table-column label="报警名" prop="name" width="120"></el-table-column>

      <el-table-column
        label="模版id"
        prop="templateId"
        width="120"
      ></el-table-column>

      <el-table-column label="创建日期" width="180">
        <template slot-scope="scope">{{
          scope.row.CreatedAt | formatDate
        }}</template>
      </el-table-column>

      <el-table-column
        label="间隔"
        prop="interval"
        width="120"
      ></el-table-column>

      <el-table-column label="阀值" prop="count" width="120"></el-table-column>

      <el-table-column label="使用比例" prop="useRatio" width="120">
        <template slot-scope="scope">{{
          scope.row.useRatio | formatBoolean
        }}</template>
      </el-table-column>

      <el-table-column
        label="比例阀值"
        prop="ratio"
        width="120"
      ></el-table-column>

      <el-table-column
        label="接收邮箱"
        prop="email"
        width="120"
      ></el-table-column>

      <el-table-column label="">
        <template slot-scope="scope">
          <el-button
            class="table-button"
            @click="updateTemplateAlarmStrategy(scope.row)"
            size="small"
            type="primary"
            icon="el-icon-edit"
            >变更</el-button
          >
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

        <el-form-item label="报警名:">
          <el-input
            v-model="formData.name"
            clearable
            placeholder="请输入"
          ></el-input>
        </el-form-item>

        <el-form-item label="接收邮箱:">
          <el-input
            v-model="formData.email"
            clearable
            placeholder="请输入"
          ></el-input>
        </el-form-item>

        <el-form-item label="模版id:">
          <el-input
            v-model.number="formData.templateId"
            clearable
            placeholder="请输入"
          ></el-input>
        </el-form-item>

        <el-form-item label="间隔:">
          <el-input
            placeholder="30s/1m, 精度为s"
            v-model="formData.interval"
            clearable
          ></el-input>
        </el-form-item>

        <el-form-item label="阀值:"
          ><el-input
            v-model.number="formData.count"
            clearable
            placeholder="请输入"
          ></el-input>
        </el-form-item>

        <el-form-item label="使用比例:">
          <el-switch
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            v-model="formData.useRatio"
            clearable
          ></el-switch>
        </el-form-item>

        <el-form-item label="比例阀值:">
          <el-input-number
            v-model="formData.ratio"
            :precision="2"
            clearable
          ></el-input-number>
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
  createTemplateAlarmStrategy,
  deleteTemplateAlarmStrategy,
  deleteTemplateAlarmStrategyByIds,
  updateTemplateAlarmStrategy,
  findTemplateAlarmStrategy,
  getTemplateAlarmStrategyList,
} from "@/api/template_alarm"; //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "TemplateAlarmStrategy",
  mixins: [infoList],
  data() {
    return {
      listApi: getTemplateAlarmStrategyList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        interval: "",
        app: "",
        templateId: 0,
        count: 0,
        useRatio: false,
        ratio: 0,
        name: "",
        email: "",
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
      if (this.searchInfo.useRatio == "") {
        this.searchInfo.useRatio = null;
      }
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
        this.deleteTemplateAlarmStrategy(row);
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

      const res = await deleteTemplateAlarmStrategyByIds({
        ID: ids,
        app: this.searchInfo.app,
      });
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
    async updateTemplateAlarmStrategy(row) {
      const res = await findTemplateAlarmStrategy({ ID: row.ID, app: row.app });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.res;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        interval: "",
        app: "",
        templateId: 0,
        count: 0,
        useRatio: false,
        ratio: 0,
        name: "",
        email: "",
      };
    },
    async deleteTemplateAlarmStrategy(row) {
      const res = await deleteTemplateAlarmStrategy({
        ID: row.ID,
        app: row.app,
      });
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
      let res;
      switch (this.type) {
        case "create":
          res = await createTemplateAlarmStrategy(this.formData);
          break;
        case "update":
          res = await updateTemplateAlarmStrategy(this.formData);
          break;
        default:
          res = await createTemplateAlarmStrategy(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "创建/更改成功",
        });
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    },
  },
};
</script>

<style>
</style>
