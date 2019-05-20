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
                el-table-column(type="expand")
                  template(slot-scope="props")
                    p 福利: {{props.row.welfare}}
                el-table-column(prop="company_name" label="公司名稱")
                  template(slot-scope="scope")
                    span
                      | {{scope.row.company_name}}
                    span(v-if="scope.row.link")
                      a(:href="urlwithhttp(scope.row.link)" target="_blank")
                        i(class="el-icon-attract")
                el-table-column(prop="job_name" label="職缺名稱")
                  template(slot-scope="scope")
                    span
                      | {{scope.row.job_name}}
                    span
                      a(:href="'https://www.104.com.tw/jobbank/custjob/index.php?j=' + scope.row.j + '&c=' + scope.row.c" target="_blank")
                        i(class="el-icon-thumb")
                el-table-column(prop="sal_low" label="最低薪")
                el-table-column(prop="sal_high" label="最高薪")
                el-table-column(prop="area" label="地區")
                el-table-column(prop="updated_at" label="更新時間")
                  template(slot-scope="scope")
                    div
                      | {{ getDTime(scope.row.updated_at) }}
                    div
                      | [{{ getMtTime(scope.row.updated_at) }}]
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
import moment from 'moment';
moment.locale('zh');

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
   getDTime(ctime) {
     const dtime = moment(ctime).format("YYYY-MM-DD");
     return dtime
   },
   getMtTime(ctime) {
     const mttime = moment(ctime).fromNow();
     return mttime

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
  .el-icon-attract {
    color: #157a56;
  }
}
</style>
