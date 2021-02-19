<template>
  <div>
    <div class="clearflex" style="margin: 18px">
      <el-button
        @click="authAppEnter"
        class="fl-right"
        size="small"
        type="primary"
        >确 定</el-button
      >
      <el-button @click="all" class="fl-left" size="small" type="primary"
        >全选</el-button
      >
      <el-button @click="none" class="fl-left" size="small" type="primary"
        >全不选</el-button
      >
    </div>
    <el-checkbox-group v-model="app" @change="selectApp">
      <el-checkbox v-for="(item, key) in allApps" :label="item" :key="key">{{
        item.name
      }}</el-checkbox>
    </el-checkbox-group>
  </div>
</template>
<script>
import { getAppList } from "@/api/app_manage";
import { setAppAuthority } from "@/api/authority";
export default {
  name: "Apps",
  data() {
    return {
      app: [],
      allApps: [],
      needConfirm: false,
    };
  },
  props: {
    row: {
      default: function () {
        return {};
      },
      type: Object,
    },
  },
  methods: {
    // 暴露给外层使用的切换拦截统一方法
    enterAndNext() {
      this.authAppEnter();
    },
    all() {
      this.app = [...this.allApps];
      this.row.app = this.app;
      this.needConfirm = true;
    },
    none() {
      this.app = [];
      this.row.app = this.app;
      this.needConfirm = true;
    },
    // 提交
    async authAppEnter() {
      const res = await setAppAuthority(this.row);
      if (res.code == 0) {
        this.$message({ type: "success", message: "app设置成功" });
      }
    },
    //   选择
    selectApp() {
      this.row.app = this.app;
      this.needConfirm = true;
    },
  },
  async created() {
    const res = await getAppList();
    this.allApps = res.data.list;
    this.row.app &&
      this.allApps.map((app) => {
        this.row.app.map((checkedApp) => {
          if (checkedApp.name === app.name) {
            this.app.push(app);
          }
        });
      });
  },
};
</script>
<style lang="less">
</style>