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
  name: "Time",
  components: {
    "v-chart": ECharts,
  },
  data() {
    return {
      polar: {
        xAxis: {
          type: "category",
          data: [
            "00",
            "01",
            "02",
            "03",
            "04",
            "05",
            "06",
            "07",
            "08",
            "09",
            "10",
            "11",
            "12",
            "13",
            "14",
            "15",
            "16",
            "17",
            "18",
            "19",
            "20",
            "21",
            "22",
            "23",
          ],
        },
        yAxis: {
          type: "value",
        },
        series: [
          {
            data: [
              149968,
              88284,
              57353,
              38323,
              29326,
              31373,
              48437,
              83762,
              162198,
              249605,
              303578,
              284076,
              255557,
              283960,
              293687,
              316266,
              337756,
              303242,
              265825,
              281505,
              321381,
              333715,
              289828,
              190984,
            ],
            type: "line",
            smooth: true,
            itemStyle: {
              normal: {
                color: "#275F82", //改变折线点的颜色
                lineStyle: {
                  color: "#409EFF", //改变折线颜色
                },
              },
            },
          },
        ],
      },
    };
  },
};
</script>
