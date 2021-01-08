<template>
  <div class="chart">
    <el-row>
      <el-col :span="24" style="text-align: center">
        <div class="header">
          <el-row>
            <el-col :span="8">
              <!-- <el-radio-group v-model="radio" style="">
                <el-radio :label="3"> 月</el-radio>
                <el-radio :label="6">周</el-radio>
                <el-radio :label="9">日</el-radio>
              </el-radio-group> -->
            </el-col>
            <el-col :span="8">
              <!-- <el-date-picker
                v-model="value1"
                type="date"
                placeholder="选择日期"
                format="yyyy 年 MM 月 dd 日"
              >
              </el-date-picker -->
            </el-col>
            <el-col :span="24" style="text-align: right">
              <el-switch
                @change="test"
                v-model="value"
                active-text="图表"
                inactive-text="词云"
              >
              </el-switch>
            </el-col>
          </el-row>
        </div>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="24" style="text-align: center">
        <v-chart :options="polar" v-if="value == true"
      /></el-col>
    </el-row>
    <el-row>
      <el-col :span="24" style="text-align: center">
        <wordcloud
          :data="defaultWords"
          nameKey="name"
          valueKey="value"
          :color="myColors"
          :showTooltip="false"
          :wordClick="wordClickHandler"
          v-if="value != true"
        >
        </wordcloud>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped>
.echarts {
  width: 800px;
  height: 600px;
  display: inline-block;
}
.header {
  width: 1000px;
  height: 60px;
  display: inline-block;
}

.chart2 {
  float: left;
  padding-top: 20px;
  width: 100%;
  height: 300px;
  margin-top: 10px;
  border: 1px solid #e4e4e4;
  background: #ffffff;
  border-radius: 6px;
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
import wordcloud from "vue-wordcloud";
export default {
  name: "Keywords",
  components: {
    "v-chart": ECharts,
    wordcloud,
  },
  data() {
    return {
      value: true,
      myColors: ["#1f77b4", "#629fc9", "#94bedb", "#c9e0ef"],
      defaultWords: [
        {
          name: "百度",
          value: 38441,
        },
        {
          name: "baidu",
          value: 18312,
        },
        {
          name: "4399小游戏",
          value: 11438,
        },
        {
          name: "qq空间",
          value: 10317,
        },
        {
          name: "优酷",
          value: 10158,
        },
        {
          name: "新亮剑",
          value: 9654,
        },
        {
          name: "馆陶县县长闫宁的父亲",
          value: 9127,
        },
        {
          name: "公安卖萌",
          value: 8192,
        },
        {
          name: "百度一下",
          value: 7104,
        },
        {
          name: "魏特琳",
          value: 6665,
        },
        {
          name: "qq网名",
          value: 6149,
        },
        {
          name: "7k7k小游戏",
          value: 5985,
        },
        {
          name: "黑狐",
          value: 5610,
        },
        {
          name: "新浪微博",
          value: 5369,
        },
        {
          name: "李宇春体",
          value: 5310,
        },
        {
          name: "新疆暴徒被击毙图片",
          value: 4997,
        },
        {
          name: "hao123",
          value: 4834,
        },
        {
          name: "123",
          value: 4829,
        },
        {
          name: "4399洛克王国",
          value: 4112,
        },
        {
          name: "qq头像",
          value: 4085,
        },
        {
          name: "nba",
          value: 4027,
        },
        {
          name: "魏特琳",
          value: 6665,
        },
        {
          name: "龙门飞甲",
          value: 3917,
        },
        {
          name: "qq个性签名",
          value: 3880,
        },
        {
          name: "张去死",
          value: 3848,
        },
        {
          name: "cf官网",
          value: 3729,
        },
        {
          name: "凰图腾",
          value: 3632,
        },
        {
          name: "快播",
          value: 3423,
        },
        {
          name: "金陵十三钗",
          value: 3349,
        },
        {
          name: "吞噬星空",
          value: 3330,
        },
        {
          name: "dnf官网",
          value: 3303,
        },
        {
          name: "武动乾坤",
          value: 3232,
        },
        {
          name: "新亮剑全集",
          value: 3210,
        },
        {
          name: "电影",
          value: 3155,
        },
        {
          name: "优酷网",
          value: 3115,
        },
        {
          name: "两次才处决美女罪犯",
          value: 3028,
        },
        {
          name: "土豆网",
          value: 2969,
        },
        {
          name: "qq分组",
          value: 2940,
        },
        {
          name: "全国各省最低工资标准",
          value: 2872,
        },
        {
          name: "清代姚明",
          value: 2784,
        },
        {
          name: "youku",
          value: 2783,
        },
        {
          name: "争产案",
          value: 2755,
        },
        {
          name: "dnf",
          value: 2686,
        },
        {
          name: "12306",
          value: 2682,
        },
        {
          name: "身份证号码大全",
          value: 2680,
        },
        {
          name: "火影忍者",
          value: 2604,
        },
      ],
      polar: {
        color: ["#3398DB"],
        tooltip: {
          trigger: "axis",
          axisPointer: {
            type: "shadow",
          },
        },
        grid: {
          left: "3%",
          right: "4%",
          bottom: "3%",
          containLabel: true,
        },
        xAxis: [
          {
            type: "category",
            data: [
              "百度",
              "baidu",
              "4399小游戏",
              "qq空间",
              "优酷",
              "新亮剑",
              "馆陶县县长闫宁的父亲",
              "公安卖萌",
              "百度一下",
              "魏特琳",
            ],
            axisTick: {
              alignWithLabel: true,
            },
          },
        ],
        yAxis: [
          {
            type: "value",
          },
        ],
        series: [
          {
            name: "直接访问",
            type: "bar",
            barWidth: "60%",
            data: [
              38441,
              18312,
              11438,
              10317,
              10158,
              9654,
              9127,
              8192,
              7104,
              6665,
            ],
          },
        ],
      },
    };
  },
  methods: {
    test(val) {
      console.log(val);
    },
    wordClickHandler(name, value, vm) {
      console.log("wordClickHandler", name, value, vm);
    },
    initCharts() {
      let myChart2 = echarts.init(this.$refs.chart2);
      myChart2.setOption({
        title: {
          text: "关键词分析",
          x: "center",
        },
        backgroundColor: "#fff",
        // tooltip: {
        //   pointFormat: "{series.name}: <b>{point.percentage:.1f}%</b>"
        // },
        series: [
          {
            type: "wordCloud",
            //用来调整词之间的距离
            gridSize: 10,
            //用来调整字的大小范围
            // Text size range which the value in data will be mapped to.
            // Default to have minimum 12px and maximum 60px size.
            sizeRange: [14, 60],
            // Text rotation range and step in degree. Text will be rotated randomly in range [-90,                                                                             90] by rotationStep 45
            //用来调整词的旋转方向，，[0,0]--代表着没有角度，也就是词为水平方向，需要设置角度参考注释内容
            // rotationRange: [-45, 0, 45, 90],
            // rotationRange: [ 0,90],
            rotationRange: [0, 0],
            //随机生成字体颜色
            textStyle: {
              normal: {
                color: function () {
                  return (
                    "rgb(" +
                    Math.round(Math.random() * 255) +
                    ", " +
                    Math.round(Math.random() * 255) +
                    ", " +
                    Math.round(Math.random() * 255) +
                    ")"
                  );
                },
              },
            },
            //位置相关设置
            left: "center",
            top: "center",
            right: null,
            bottom: null,
            width: "300%",
            height: "300%",
            //数据
            data: this.worddata,
          },
        ],
      });
    },
  },
};
</script>
