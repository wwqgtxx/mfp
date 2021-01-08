<template>
  <div class="Test1">
    <div class="history">
      <el-table
        :data="tableData"
        @row-click="rclick"
        border

      >
        <el-table-column prop="time" label="时间" width="200%">
        </el-table-column>
        <el-table-column prop="word" label="关键词" width="100% ">
        </el-table-column>
      </el-table>
    </div>
    <div class="input">
      <div class="input1">
        <el-input
          placeholder="请输入关键词"
          v-model="input"
          class="input-with-select"
        >
          <el-select slot="prepend" placeholder="请选择"> </el-select>
          <el-button
            v-on:click="submit"
            slot="append"
            icon="el-icon-search"
          ></el-button>
        </el-input>
      </div>
      <div class="input2">
        <el-input
          type="textarea"
          :autosize="{ minRows: 16, maxRows: 16 }"
          placeholder="请输入内容"
          v-model="result"
          readonly="true"
        >
        </el-input>
      </div>
    </div>
  </div>
</template>
<style scoped>
.Test1 {
  width: 100%;
  height: 100%;
}
.history {
  width: 25%;
  height: 90%;
  float: left;
  text-align: center;
  margin: 2.5% 2.5%;
}
.input {
  width: 65%;
  height: 90%;
  float: right;
  border: 1px;
  margin: 2.5% 2.5%;
}
.input1 {
    margin: 0% 2.5%;
}
.input2 {
    margin: 2.5% 2.5%;
}
</style>

<script>
import axios from "axios";
export default {
  name: "Queue",
  data() {
    return {
      fit: "scale-down",
      value2: [],
      input: "",
      tableData: [],
      id: 1,
      result: "",
    };
  },
  methods: {
    async submit() {
      let content = {
        id: " ",
        word: " ",
        result: " ",
        time: " ",
      };
      content.id = this.id + "";
      this.id += 1;
      content.word = this.input;
      content.time = new Date().toLocaleString();

      //ajax
      let resp = await axios.post("/api", content);
      let data = resp.data;
      console.log(content);
      this.result = content.result = data.result;

      this.tableData.push(content);
      console.log(this.tableData);
    },
    rclick(row, column, event) {
      console.log([row, column, event]);
      this.input = row.word;
      this.result = row.result;
    },
  },
};
</script>
