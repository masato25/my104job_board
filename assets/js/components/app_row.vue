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
                el-table-column(prop="updated_at" label="updated_at")
                el-table-column(prop="job_name" label="job_name")
                el-table-column(prop="company_name" label="company_name")
                  template(scope="scope")
                    span(v-if="scope.row.link")
                      a(:href="urlwithhttp(scope.row.link)" target="_blank")
                        | {{ scope.row.company_name }}
                    span(v-if="!scope.row.link")
                      | {{ scope.row.company_name }}
                el-table-column(prop="sal_low" label="sal_low")
                el-table-column(prop="sal_high" label="sal_high")
                el-table-column(prop="area" label="area")
          el-row
            el-col(:span=6)
              .grid-conten.bg-purple
                |total: {{totalpage}} / page: {{currentPage}}
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
