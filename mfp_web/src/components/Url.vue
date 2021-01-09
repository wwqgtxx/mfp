<template>
  <div class="pie-chart" style="text-align: center">
    <v-chart :options="polar" />
  </div>
</template>
<style scoped>
.echarts {
  width: 800px;
  height: 600px;
  display: inline-block;
  top: 60px;
}
</style>
<script>
import ECharts from "vue-echarts";
import "echarts/lib/chart/bar";
import "echarts/lib/chart/line";
import "echarts/lib/chart/pie";
import "echarts/lib/component/tooltip";
import "echarts/lib/component/legend";
import "echarts/lib/component/markPoint";
import "echarts/lib/component/markLine";
import "echarts/lib/component/graphic";

export default {
  name: "Url",
  components: {
    "v-chart": ECharts,
  },
  data() {
    return {
      polar: {
        title: {
          text: "访问量前十的网址分析",
          x: "center",
        },
        tooltip: {
          trigger: "item",
          formatter: "{a} <br/>{b} : {c} ({d}%)",
        },
        legend: {
          show: true,
          orient: "vertical",
          left: "left",
          data: [
            "www.baidu.com",
            "www.4399.com",
            "www.hao123.com",
            "www.youku.com",
            "qzone.qq.com",
            "www.7k7k.com",
            "weibo.com",
            "cf.qq.com",
            "www.xixiwg.com",
            "www.sina.com.cn",
            "其他网址",
          ],
        },
        series: [
          {
            name: "访问来源",
            type: "pie",
            radius: "55%",
            center: ["50%", "60%"],
            data: [
              { value: 73737, name: "www.baidu.com" },
              { value: 19015, name: "www.4399.com" },
              { value: 14338, name: "www.hao123.com" },
              { value: 14086, name: "www.youku.com" },
              { value: 12920, name: "qzone.qq.com" },
              { value: 8326, name: "www.7k7k.com" },
              { value: 7547, name: "weibo.com" },
              { value: 7544, name: "cf.qq.com" },
              { value: 7043, name: "www.xixiwg.com" },
              { value: 4328461, name: "其他网址" },
            ],
            itemStyle: {
              emphasis: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: "rgba(0, 0, 0, 0.5)",
              },
            },
          },
        ],
      },
    };
  },
};
</script>
