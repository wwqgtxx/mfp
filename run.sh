#!/bin/bash

#需要先将搜狗数据解压放置在~/mfp/SogouQ
#mfp目录下需要有mfp-linux-amd64和dictionary.txt
cd ~/mfp
chmod +x mfp-linux-amd64

#将搜狗数据从gbk转换为utf8
rm -rf SogouQ_utf8
mkdir -p SogouQ_utf8
./mfp-linux-amd64 gb2utf8 ./SogouQ ./SogouQ_utf8

#数据格式整理（去除损坏数据）
rm -rf format_data
mkdir -p format_data
./mfp-linux-amd64 format_file ./SogouQ_utf8 ./format_data

#将数据进行分词处理
rm -rf db_data
mkdir -p db_data
./mfp-linux-amd64 to_db ./format_data ./db_data

#将分词后数据上传到HDFS
hadoop fs -mkdir /mfp
hadoop fs -mkdir /mfp/data
hadoop fs -put ./db_data/*.txt /mfp/data
hadoop fs -ls /mfp/data

#运行 MR1: 并行计数
yarn jar `find $HADOOP_HOME/share/hadoop/tools/lib/ -name hadoop-streaming*.jar` -file ~/mfp/mfp-linux-amd64 -mapper 'mfp-linux-amd64 mapper1' -reducer 'mfp-linux-amd64 reducer1' -input /mfp/data/*.txt -output /mfp/res1/

#将MR1结果取回
rm -rf res1
rm -rf res1.txt
rm -rf gList.json
mkdir -p res1
hadoop fs -get /mfp/res1/part-* ./res1/

#合并hdfs分段
./mfp-linux-amd64 combine_parts ./res1 res1.txt
#项目分组
./mfp-linux-amd64 sort_kv res1.txt gList.json

#运行 MR2: 并行 FP-Growth
yarn jar `find $HADOOP_HOME/share/hadoop/tools/lib/ -name hadoop-streaming*.jar` -file ~/mfp/mfp-linux-amd64 -file ~/mfp/gList.json -mapper 'mfp-linux-amd64 mapper2' -reducer 'mfp-linux-amd64 reducer2' -input /mfp/data/*.txt -output /mfp/res2/

#运行 MR3: 汇总
yarn jar `find $HADOOP_HOME/share/hadoop/tools/lib/ -name hadoop-streaming*.jar` -file ~/mfp/mfp-linux-amd64 -mapper 'mfp-linux-amd64 mapper3' -reducer 'mfp-linux-amd64 reducer3' -input /mfp/res2/part-* -output /mfp/res3/

#将MR3结果取回
rm -rf result
rm -rf res.txt
mkdir -p result
hadoop fs -get /mfp/res3/part-* ./result/
#合并hdfs分段
./mfp-linux-amd64 combine_parts ./result res.txt

#输入想要查询的关键词，就会返回其频繁二项集
./mfp-linux-amd64 find_pair

#开启API服务器
./mfp-linux-amd64 api_server --listen 0.0.0.0:8000