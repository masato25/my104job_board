<template lang="pug">
  .el-row
    el-col
      AppHeader
    el-col
      el-card(class="box-card")
        el-row
          el-col(:span=24)
            el-table(
                :data="jobs"
                v-loading.body="loading"
                height="500"
                style="width: 100%")
              el-table-column(prop="updated_at" label="updated_at")
              el-table-column(prop="job_name" label="job_name")
              el-table-column(prop="company_name" label="company_name")
                template(scope="scope")
                  span
                    a(:href="scope.row.link" target="_blank")
                      | {{ scope.row.company_name }}
              el-table-column(prop="sal_low" label="sal_low")
              el-table-column(prop="sal_high" label="sal_high")
              el-table-column(prop="area" label="area")
        el-row
          el-col(:span=6)
            .grid-conten.bg-purple
              |total: {{total}} / page: {{currentPage + 1}}
          el-col(:span=18)
            .block
              el-pagination(
                  :page-size="limit"
                  :page-sizes="[10,20,30]"
                  @size-change="handleSizeChange"
                  @current-change="handlePageChange"
                  layout="sizes, prev, pager, next"
                  :total="total")
</template>

<script>
import AppHeader from './header.vue';
import { mapState, mapActions } from 'vuex';

export default {
  name: 'AppRow',
  components: {
    'AppHeader': AppHeader
  },
  data:() => {
    return {}
  },
  mounted: function(){
    this.getJobs();
  },
  computed: mapState(['jobs', 'total', 'currentPage', 'limit', 'loading']),
  methods: {
   handleSizeChange(page) {
     this.setCurrentPage({page: page})
   },
   handlePageChange(sizeChange) {
     this.setPageShowLimit({limit: sizeChange})
   },
   ...mapActions(['getJobs', 'setPageShowLimit', 'setCurrentPage'])
  }
}
</script>
