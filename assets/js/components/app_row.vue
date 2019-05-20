<template lang="pug">
  .el-row
    el-col
      NavHead
    .el-row
      el-col
        el-card(class="box-card")
          el-row
            AppHeader
            el-col(:span=24)
              el-table(
                  :data="jobs"
                  v-loading.body="loading"
                  style="width: 100%")
                el-table-column(prop="updated_at" label="更新時間")
                el-table-column(prop="job_name" label="職缺名稱")
                el-table-column(prop="company_name" label="公司名稱")
                  template(scope="scope")
                    span(v-if="scope.row.link")
                      a(:href="urlwithhttp(scope.row.link)" target="_blank")
                        | {{ scope.row.company_name }}
                    span(v-if="!scope.row.link")
                      | {{ scope.row.company_name }}
                el-table-column(prop="sal_low" label="最低薪")
                el-table-column(prop="sal_high" label="最高薪")
                el-table-column(prop="area" label="地區")
          el-row
            el-col(:span=6)
              .grid-conten.bg-purple
                |筆數: {{totalpage}} / 頁數: {{currentPage}}
            el-col(:span=18)
              .block
                el-pagination(
                    :page-size="limit"
                    :page-sizes="[10,20,30]"
                    @size-change="handleSizeChange"
                    @current-change="handlePageChange"
                    layout="sizes, prev, pager, next"
                    :total="totalpage")
</template>

<script>
import AppHeader from './header.vue';
import NavHead from './navhead.vue';
import { mapState, mapActions } from 'vuex';

export default {
  name: 'AppRow',
  components: {
    'AppHeader': AppHeader,
    'NavHead': NavHead
  },
  data:() => {
    return {
    }
  },
  mounted: function(){
    this.getJobs();
  },
  computed: mapState(['currentPage', 'limit', 'loading', 'totalpage', 'jobs', 'cat']),
  methods: {
   handlePageChange(page) {
     this.setCurrentPage({page: page+1})
   },
   handleSizeChange(sizeChange) {
     this.setPageShowLimit({limit: sizeChange})
   },
   urlwithhttp(uri){
     if(!uri.startsWith("http")){
       uri = "http://" + uri;
     }
     return uri
   },
   ...mapActions(['getJobs', 'setPageShowLimit', 'setCurrentPage'])
  }
}
</script>

<style>
.el-row {
  margin-bottom: 20px;
  &:last-child {
    margin-bottom: 0;
  }
}
.box-card {
  margin-left: 20px;
  margin-right: 20px;
  margin-top: 20px;
}
</style>
