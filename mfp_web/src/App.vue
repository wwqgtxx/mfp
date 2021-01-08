<template>
  <div id="app">
    <template>
      <el-container style="height: 100% ;weight : 100%">
        <el-header
        >
          <el-menu
            :default-active="activeIndex"
            @select="handleSelect"
            class="el-menu-demo"
            mode="horizontal"
            background-color="#FFFFFF"
            text-color="#409EFF"
            active-text-color="#409EFF"
            style="height: 100%;weight : 100%"
            router="true"
          >
            <el-row>
              <el-col :span="6">
                <div class="title">搜狗搜索日志分析系统</div>
              </el-col>
              <el-col :span="2" style="text-align: center">
                <el-menu-item index="/Queue">查询</el-menu-item>
              </el-col
              >
              <el-col :span="2" style="text-align: center">
                <el-menu-item index="/Keywords">关键词分析</el-menu-item>
              </el-col
              >
              <el-col :span="2" style="text-align: center">
                <el-menu-item index="/Url">网址分析</el-menu-item>
              </el-col
              >
              <el-col :span="2" style="text-align: center">
                <el-menu-item index="/Time">时间分析</el-menu-item>
              </el-col
              >
            </el-row>
          </el-menu>
        </el-header>
        <el-main>
          <router-view></router-view>
        </el-main>
      </el-container>
    </template>
  </div>
</template>
<style>
.app {
  padding: 0;
  margin: 0;
}

.title {
  font-size: 20px;
  height: 60px;
  line-height: 60px;
  color: #409EFF;
}

.el-header {
  padding: 0;
  position: fixed;
  margin: 0;
  left: 10%;
  top: 0;
  width: 80%;
  z-index: 1500;
}

.el-main {
  position: fixed;
  margin-top: 80px;
  left: 10%;
  width: 80%;
  height: calc(100% - 80px);
}
</style>

<script>
export default {
  name: "App",
};
</script>
<style>
/**
 * 默认尺寸为 600px×400px，如果想让图表响应尺寸变化，可以像下面这样
 * 把尺寸设为百分比值（同时请记得为容器设置尺寸）。
 */
.context {
  width: 800px;
  height: 800px;
}
</style>


