<script>
export default {
  name: 'home',
  data() {
    return {
      searchQuery: '',
    }
  },
  methods: {
  },
  created() {

  }
}
</script>

<template>
  <div class="home-page">
    <el-container>
      <el-main>
        <el-row>
          <el-col :span="24" class="search-container">
            <el-input placeholder="请输入企业名、人名、产品名、地址、电话、经营范围等，多关键词用空格隔开" v-model="searchQuery" class="search-input" />
            <el-button type="primary" class="search-button">查一下</el-button>
          </el-col>
        </el-row>
        <el-row class="hot-search">
          <span>热搜：</span>
          <el-link>健身房理发美容整形涉款贪腐</el-link>
          <el-link>合众</el-link>
          <el-link>英伟达总市值突破3万亿美元</el-link>
        </el-row>
        <el-row :gutter="20" class="applications">
          <el-col :span="6">
            <el-card shadow="hover">
              <div slot="header" class="clearfix">
                <span>批量查询</span>
              </div>
              <div class="text item">查企业/股东/集团</div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="hover">
              <div slot="header" class="clearfix">
                <span>风险雷达</span>
              </div>
              <div class="text item">报告企业风险及舆情监测</div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="hover">
              <div slot="header" class="clearfix">
                <span>信用大数据</span>
              </div>
              <div class="text item">查看企业信用信息</div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="hover">
              <div slot="header" class="clearfix">
                <span>司法大数据</span>
              </div>
              <div class="text item">海量司法数据</div>
            </el-card>
          </el-col>
        </el-row>
        <el-row :gutter="20" class="applications">
          <el-col :span="6">
            <el-card shadow="hover">
              <div slot="header" class="clearfix">
                <span>失信被执行人</span>
              </div>
              <div class="text item">查失信、规避合伙风险</div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="hover">
              <div slot="header" class="clearfix">
                <span>查专利</span>
              </div>
              <div class="text item">专利信息查询</div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="hover">
              <div slot="header" class="clearfix">
                <span>招投标查询</span>
              </div>
              <div class="text item">找项目、找工程</div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="hover">
              <div slot="header" class="clearfix">
                <span>全球企业</span>
              </div>
              <div class="text item">60+国家/地区企业数据</div>
            </el-card>
          </el-col>
        </el-row>
      </el-main>
    </el-container>
  </div>
</template>

<style scoped>
.home-page {
  background-color: #f5f5f5;
}

.logo {
  padding: 20px;
}

.search-container {
  margin: 40px 0;
  text-align: center;
}

.search-input {
  width: 600px;
}

.search-button {
  margin-left: 10px;
}

.hot-search {
  margin: 20px 0;
  text-align: center;
}

.applications {
  margin: 20px 0;
}

.item {
  text-align: center;
  font-size: 16px;
  color: #333;
}
</style>
